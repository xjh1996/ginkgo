package devops

import "github.com/onsi/ginkgo"

// SIGDescribe annotates the test with the SIG(Special Interest Groups) label.
func SIGDescribe(text string, body func()) bool {
	return ginkgo.Describe("[cps-devops] "+text, body)
}
