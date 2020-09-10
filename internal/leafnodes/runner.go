package leafnodes

import (
	"github.com/onsi/ginkgo/internal/codelocation"
	"github.com/onsi/ginkgo/internal/failer"
	"github.com/onsi/ginkgo/types"
)

type runner struct {
	body           func()
	codeLocation   types.CodeLocation
	nodeType       types.SpecComponentType
	componentIndex int
	failer         *failer.Failer
}

func newRunner(body func(), codeLocation types.CodeLocation, failer *failer.Failer, nodeType types.SpecComponentType, componentIndex int) *runner {
	return &runner{
		body:           body,
		codeLocation:   codeLocation,
		failer:         failer,
		nodeType:       nodeType,
		componentIndex: componentIndex,
	}
}

func (r *runner) run() (outcome types.SpecState, failure types.SpecFailure) {
	finished := false

	defer func() {
		if e := recover(); e != nil || !finished {
			r.failer.Panic(codelocation.New(2), e)
		}

		failure, outcome = r.failer.Drain(r.nodeType, r.componentIndex, r.codeLocation)
	}()

	r.body()
	finished = true

	return
}
