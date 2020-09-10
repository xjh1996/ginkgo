package leafnodes

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/onsi/ginkgo/internal/failer"
	"github.com/onsi/ginkgo/types"
)

type synchronizedAfterSuiteNode struct {
	allNodesRunner *runner
	node1Runner    *runner

	outcome types.SpecState
	failure types.SpecFailure
	runTime time.Duration
}

func NewSynchronizedAfterSuiteNode(allNodesBody func(), node1Body func(), codeLocation types.CodeLocation, failer *failer.Failer) SuiteNode {
	return &synchronizedAfterSuiteNode{
		allNodesRunner: newRunner(allNodesBody, codeLocation, failer, types.SpecComponentTypeAfterSuite, 0),
		node1Runner:    newRunner(node1Body, codeLocation, failer, types.SpecComponentTypeAfterSuite, 0),
	}
}

func (node *synchronizedAfterSuiteNode) Run(parallelNode int, parallelTotal int, syncHost string) bool {
	node.outcome, node.failure = node.allNodesRunner.run()

	if parallelNode == 1 {
		if parallelTotal > 1 {
			node.waitUntilOtherNodesAreDone(syncHost)
		}

		outcome, failure := node.node1Runner.run()

		if node.outcome == types.SpecStatePassed {
			node.outcome, node.failure = outcome, failure
		}
	}

	return node.outcome == types.SpecStatePassed
}

func (node *synchronizedAfterSuiteNode) Passed() bool {
	return node.outcome == types.SpecStatePassed
}

func (node *synchronizedAfterSuiteNode) Summary() *types.SetupSummary {
	return &types.SetupSummary{
		ComponentType: node.allNodesRunner.nodeType,
		CodeLocation:  node.allNodesRunner.codeLocation,
		State:         node.outcome,
		RunTime:       node.runTime,
		Failure:       node.failure,
	}
}

func (node *synchronizedAfterSuiteNode) waitUntilOtherNodesAreDone(syncHost string) {
	for {
		if node.canRun(syncHost) {
			return
		}

		time.Sleep(50 * time.Millisecond)
	}
}

func (node *synchronizedAfterSuiteNode) canRun(syncHost string) bool {
	resp, err := http.Get(syncHost + "/RemoteAfterSuiteData")
	if err != nil || resp.StatusCode != http.StatusOK {
		return false
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false
	}
	resp.Body.Close()

	afterSuiteData := types.RemoteAfterSuiteData{}
	err = json.Unmarshal(body, &afterSuiteData)
	if err != nil {
		return false
	}

	return afterSuiteData.CanRun
}
