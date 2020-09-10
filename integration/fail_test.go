package integration_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("Failing Specs", func() {
	var pathToTest string

	BeforeEach(func() {
		pathToTest = tmpPath("failing")
		copyIn(fixturePath("fail_fixture"), pathToTest, false)
	})

	It("should fail in all the possible ways", func() {
		session := startGinkgo(pathToTest, "--noColor")
		Eventually(session).Should(gexec.Exit(1))
		output := string(session.Out.Contents())

		Ω(output).ShouldNot(ContainSubstring("NEVER SEE THIS"))

		Ω(output).Should(ContainSubstring("a top level failure on line 10"))
		Ω(output).Should(ContainSubstring("fail_fixture_test.go:10"))

		Ω(output).Should(ContainSubstring("a sync failure"))
		Ω(output).Should(MatchRegexp(`Test Panicked\n\s+a sync panic`))
		Ω(output).Should(ContainSubstring("a sync FAIL failure"))

		Ω(output).Should(ContainSubstring("a top level specify"))
		Ω(output).ShouldNot(ContainSubstring("ginkgo_dsl.go"))
		Ω(output).Should(ContainSubstring("fail_fixture_test.go:31"))

		Ω(output).ShouldNot(ContainSubstring("table.go"))
		Ω(output).Should(MatchRegexp(`a top level DescribeTable\n.*fail_fixture_test\.go:35`),
			"the output of a failing DescribeTable should include its file path and line number")
		Ω(output).ShouldNot(ContainSubstring("table_entry.go"))
		Ω(output).Should(MatchRegexp(`a TableEntry constructed by Entry \[It\]\n.*fail_fixture_test\.go:39`),
			"the output of a failing Entry should include its file path and line number")

		Ω(output).Should(ContainSubstring("0 Passed | 7 Failed"))
	})
})
