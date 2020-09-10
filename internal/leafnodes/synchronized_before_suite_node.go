package leafnodes

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/onsi/ginkgo/internal/failer"
	"github.com/onsi/ginkgo/types"
)

type synchronizedBeforeSuiteNode struct {
	node1Runner    *runner
	allNodesRunner *runner

	data []byte

	outcome types.SpecState
	failure types.SpecFailure
	runTime time.Duration
}

func NewSynchronizedBeforeSuiteNode(node1Body func() []byte, allNodesBody func([]byte), codeLocation types.CodeLocation, failer *failer.Failer) SuiteNode {
	node := &synchronizedBeforeSuiteNode{}

	node.node1Runner = newRunner(node.wrapNode1Body(node1Body), codeLocation, failer, types.SpecComponentTypeBeforeSuite, 0)
	node.allNodesRunner = newRunner(node.wrapAllNodesBody(allNodesBody), codeLocation, failer, types.SpecComponentTypeBeforeSuite, 0)

	return node
}

func (node *synchronizedBeforeSuiteNode) Run(parallelNode int, parallelTotal int, syncHost string) bool {
	t := time.Now()
	defer func() {
		node.runTime = time.Since(t)
	}()

	if parallelNode == 1 {
		node.outcome, node.failure = node.runNode1Portion(parallelTotal, syncHost)
	} else {
		node.outcome, node.failure = node.waitForNode1(syncHost)
	}

	if node.outcome != types.SpecStatePassed {
		return false
	}
	node.outcome, node.failure = node.allNodesRunner.run()

	return node.outcome == types.SpecStatePassed
}

func (node *synchronizedBeforeSuiteNode) runNode1Portion(parallelTotal int, syncHost string) (types.SpecState, types.SpecFailure) {
	outcome, failure := node.node1Runner.run()

	if parallelTotal > 1 {
		state := types.RemoteBeforeSuiteStatePassed
		if outcome != types.SpecStatePassed {
			state = types.RemoteBeforeSuiteStateFailed
		}
		json := (types.RemoteBeforeSuiteData{
			Data:  node.data,
			State: state,
		}).ToJSON()
		http.Post(syncHost+"/BeforeSuiteState", "application/json", bytes.NewBuffer(json))
	}

	return outcome, failure
}

func (node *synchronizedBeforeSuiteNode) waitForNode1(syncHost string) (types.SpecState, types.SpecFailure) {
	failure := func(message string) types.SpecFailure {
		return types.SpecFailure{
			Message:               message,
			Location:              node.node1Runner.codeLocation,
			ComponentType:         node.node1Runner.nodeType,
			ComponentIndex:        node.node1Runner.componentIndex,
			ComponentCodeLocation: node.node1Runner.codeLocation,
		}
	}
	for {
		resp, err := http.Get(syncHost + "/BeforeSuiteState")
		if err != nil || resp.StatusCode != http.StatusOK {
			return types.SpecStateFailed, failure("Failed to fetch BeforeSuite state")
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return types.SpecStateFailed, failure("Failed to read BeforeSuite state")
		}
		resp.Body.Close()

		beforeSuiteData := types.RemoteBeforeSuiteData{}
		err = json.Unmarshal(body, &beforeSuiteData)
		if err != nil {
			return types.SpecStateFailed, failure("Failed to decode BeforeSuite state")
		}

		switch beforeSuiteData.State {
		case types.RemoteBeforeSuiteStatePassed:
			node.data = beforeSuiteData.Data
			return types.SpecStatePassed, types.SpecFailure{}
		case types.RemoteBeforeSuiteStateFailed:
			return types.SpecStateFailed, failure("BeforeSuite on Node 1 failed")
		case types.RemoteBeforeSuiteStateDisappeared:
			return types.SpecStateFailed, failure("Node 1 disappeared before completing BeforeSuite")
		}

		time.Sleep(50 * time.Millisecond)
	}
}

func (node *synchronizedBeforeSuiteNode) Passed() bool {
	return node.outcome == types.SpecStatePassed
}

func (node *synchronizedBeforeSuiteNode) Summary() *types.SetupSummary {
	return &types.SetupSummary{
		ComponentType: node.node1Runner.nodeType,
		CodeLocation:  node.node1Runner.codeLocation,
		State:         node.outcome,
		RunTime:       node.runTime,
		Failure:       node.failure,
	}
}

func (node *synchronizedBeforeSuiteNode) wrapNode1Body(node1Body func() []byte) func() {
	return func() {
		node.data = node1Body()
	}
}

func (node *synchronizedBeforeSuiteNode) wrapAllNodesBody(allNodesBody func([]byte)) func() {
	return func() {
		allNodesBody(node.data)
	}
}
