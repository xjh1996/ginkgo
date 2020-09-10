package leafnodes_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/internal/leafnodes"
	. "github.com/onsi/gomega"

	"github.com/onsi/ginkgo/internal/codelocation"
	Failer "github.com/onsi/ginkgo/internal/failer"
	"github.com/onsi/ginkgo/types"
)

type runnable interface {
	Run() (outcome types.SpecState, failure types.SpecFailure)
	CodeLocation() types.CodeLocation
}

func SharedRunnerBehaviors(build func(body func(), failer *Failer.Failer, componentCodeLocation types.CodeLocation) runnable, componentType types.SpecComponentType, componentIndex int) {
	var (
		outcome types.SpecState
		failure types.SpecFailure

		failer *Failer.Failer

		componentCodeLocation types.CodeLocation
		innerCodeLocation     types.CodeLocation

		didRun bool
	)

	BeforeEach(func() {
		failer = Failer.New()
		componentCodeLocation = codelocation.New(0)
		innerCodeLocation = codelocation.New(0)

		didRun = false
	})

	Context("when the function passes", func() {
		BeforeEach(func() {
			outcome, failure = build(func() {
				didRun = true
			}, failer, componentCodeLocation).Run()
		})

		It("should have a successful outcome", func() {
			Ω(didRun).Should(BeTrue())

			Ω(outcome).Should(Equal(types.SpecStatePassed))
			Ω(failure).Should(BeZero())
		})
	})

	Context("when a failure occurs", func() {
		BeforeEach(func() {
			outcome, failure = build(func() {
				didRun = true
				failer.Fail("bam", innerCodeLocation)
				panic("should not matter")
			}, failer, componentCodeLocation).Run()
		})

		It("should return the failure", func() {
			Ω(didRun).Should(BeTrue())

			Ω(outcome).Should(Equal(types.SpecStateFailed))
			Ω(failure).Should(Equal(types.SpecFailure{
				Message:               "bam",
				Location:              innerCodeLocation,
				ForwardedPanic:        "",
				ComponentIndex:        componentIndex,
				ComponentType:         componentType,
				ComponentCodeLocation: componentCodeLocation,
			}))
		})
	})

	Context("when a panic occurs", func() {
		BeforeEach(func() {
			outcome, failure = build(func() {
				didRun = true
				innerCodeLocation = codelocation.New(0)
				panic("ack!")
			}, failer, componentCodeLocation).Run()
		})

		It("should return the panic", func() {
			Ω(didRun).Should(BeTrue())

			Ω(outcome).Should(Equal(types.SpecStatePanicked))
			Ω(failure.ForwardedPanic).Should(Equal("ack!"))
		})
	})

	Context("when a panic occurs with a nil value", func() {
		BeforeEach(func() {
			outcome, failure = build(func() {
				didRun = true
				innerCodeLocation = codelocation.New(0)
				panic(nil)
			}, failer, componentCodeLocation).Run()
		})

		It("should return the nil-valued panic", func() {
			Ω(didRun).Should(BeTrue())

			Ω(outcome).Should(Equal(types.SpecStatePanicked))
			Ω(failure.ForwardedPanic).Should(Equal("<nil>"))
		})
	})
}

var _ = Describe("Shared RunnableNode behavior", func() {
	Describe("It Nodes", func() {
		build := func(body func(), failer *Failer.Failer, componentCodeLocation types.CodeLocation) runnable {
			return NewItNode("", body, types.FlagTypeFocused, componentCodeLocation, failer, 3)
		}

		SharedRunnerBehaviors(build, types.SpecComponentTypeIt, 3)
	})

	Describe("BeforeEach Nodes", func() {
		build := func(body func(), failer *Failer.Failer, componentCodeLocation types.CodeLocation) runnable {
			return NewBeforeEachNode(body, componentCodeLocation, failer, 3)
		}

		SharedRunnerBehaviors(build, types.SpecComponentTypeBeforeEach, 3)
	})

	Describe("AfterEach Nodes", func() {
		build := func(body func(), failer *Failer.Failer, componentCodeLocation types.CodeLocation) runnable {
			return NewAfterEachNode(body, componentCodeLocation, failer, 3)
		}

		SharedRunnerBehaviors(build, types.SpecComponentTypeAfterEach, 3)
	})

	Describe("JustBeforeEach Nodes", func() {
		build := func(body func(), failer *Failer.Failer, componentCodeLocation types.CodeLocation) runnable {
			return NewJustBeforeEachNode(body, componentCodeLocation, failer, 3)
		}

		SharedRunnerBehaviors(build, types.SpecComponentTypeJustBeforeEach, 3)
	})
})
