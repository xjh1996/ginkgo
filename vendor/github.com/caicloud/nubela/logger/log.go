package logger

import (
	"bytes"
	"fmt"
	"k8s.io/kubernetes/test/e2e/framework/ginkgowrapper"
	"regexp"
	"runtime"
	"runtime/debug"
	"time"

	"github.com/onsi/ginkgo"
)

// This is a modified copy of log in https://github.com/kubernetes/kubernetes/blob/v1.19.2/test/e2e/framework/log.go

func nowStamp() string {
	return time.Now().Format(time.StampMicro)
}

func log(level string, format string, args ...interface{}) {
	fmt.Fprintf(ginkgo.GinkgoWriter, level+" "+nowStamp()+": "+format+"\n", args...)
	_, file, line, _ := runtime.Caller(2)
	fmt.Fprintf(ginkgo.GinkgoWriter, "%v:%v\n", file, line)
}

// Logf logs the info.
func Infof(format string, args ...interface{}) {
	log("[INFO]", format, args...)
}

// Warningf logs the Warning.
func Warningf(format string, args ...interface{}) {
	log("[WARN]", format, args...)
}

// Errorf logs the Error and will not make spec exits
func Errorf(format string, args ...interface{}) {
	log("[Error]", format, args...)
}

// Failf logs the fail info, including a stack trace. This will make spec fails and exits.
func Failf(format string, args ...interface{}) {
	FailfWithOffset(1, format, args...)
}

// FailfWithOffset calls "Fail" and logs the error with a stack trace that starts at "offset" levels above its caller
// (for example, for call chain f -> g -> FailfWithOffset(1, ...) error would be logged for "f").
func FailfWithOffset(offset int, format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	skip := offset + 1
	log("[FAIL]", "%s\n\nFull Stack Trace\n%s", msg, PrunedStack(skip))
	ginkgowrapper.Fail(nowStamp()+": "+msg, skip)
}

// Fail is a replacement for ginkgo.Fail which logs the problem as it occurs
// together with a stack trace and then calls ginkgowrapper.Fail.
func Fail(msg string, callerSkip ...int) {
	skip := 1
	if len(callerSkip) > 0 {
		skip += callerSkip[0]
	}
	log("[FAIL]", "%s\n\nFull Stack Trace\n%s", msg, PrunedStack(skip))
	ginkgowrapper.Fail(nowStamp()+": "+msg, skip)
}

var codeFilterRE = regexp.MustCompile(`/github.com/onsi/ginkgo/`)

// PrunedStack is a wrapper around debug.Stack() that removes information
// about the current goroutine and optionally skips some of the initial stack entries.
// With skip == 0, the returned stack will start with the caller of PruneStack.
// From the remaining entries it automatically filters out useless ones like
// entries coming from Ginkgo.
//
// This is a modified copy of PruneStack in https://github.com/onsi/ginkgo/blob/f90f37d87fa6b1dd9625e2b1e83c23ffae3de228/internal/codelocation/code_location.go#L25:
// - simplified API and thus renamed (calls debug.Stack() instead of taking a parameter)
// - source code filtering updated to be specific to Kubernetes
// - optimized to use bytes and in-place slice filtering from
//   https://github.com/golang/go/wiki/SliceTricks#filter-in-place
func PrunedStack(skip int) []byte {
	fullStackTrace := debug.Stack()
	stack := bytes.Split(fullStackTrace, []byte("\n"))
	// Ensure that the even entries are the method names and the
	// the odd entries the source code information.
	if len(stack) > 0 && bytes.HasPrefix(stack[0], []byte("goroutine ")) {
		// Ignore "goroutine 29 [running]:" line.
		stack = stack[1:]
	}
	// The "+2" is for skipping over:
	// - runtime/debug.Stack()
	// - PrunedStack()
	skip += 2
	if len(stack) > 2*skip {
		stack = stack[2*skip:]
	}
	n := 0
	for i := 0; i < len(stack)/2; i++ {
		// We filter out based on the source code file name.
		if !codeFilterRE.Match([]byte(stack[i*2+1])) {
			stack[n] = stack[i*2]
			stack[n+1] = stack[i*2+1]
			n += 2
		}
	}
	stack = stack[:n]

	return bytes.Join(stack, []byte("\n"))
}
