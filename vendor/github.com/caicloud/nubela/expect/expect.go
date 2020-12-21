package expect

import (
	"github.com/onsi/gomega"
)

// Equal expects the specified two are the same, otherwise an exception raises
func Equal(actual interface{}, extra interface{}, explain ...interface{}) {
	gomega.ExpectWithOffset(1, actual).To(gomega.Equal(extra), explain...)
}

// NotEqual expects the specified two are not the same, otherwise an exception raises
func NotEqual(actual interface{}, extra interface{}, explain ...interface{}) {
	gomega.ExpectWithOffset(1, actual).NotTo(gomega.Equal(extra), explain...)
}

// Error expects an error happens, otherwise an exception raises
func Error(err error, explain ...interface{}) {
	gomega.ExpectWithOffset(1, err).To(gomega.HaveOccurred(), explain...)
}

// NoError checks if "err" is set, and if so, fails assertion while logging the error.
func NoError(err error, explain ...interface{}) {
	NoErrorWithOffset(1, err, explain...)
}

// NoErrorWithOffset checks if "err" is set, and if so, fails assertion while logging the error at "offset" levels above its caller
// (for example, for call chain f -> g -> NoErrorWithOffset(1, ...) error would be logged for "f").
func NoErrorWithOffset(offset int, err error, explain ...interface{}) {
	gomega.ExpectWithOffset(1+offset, err).NotTo(gomega.HaveOccurred(), explain...)
}

// ConsistOf expects actual contains precisely the extra elements.  The ordering of the elements does not matter.
func ConsistOf(actual interface{}, extra interface{}, explain ...interface{}) {
	gomega.ExpectWithOffset(1, actual).To(gomega.ConsistOf(extra), explain...)
}

// HaveKey expects the actual map has the key in the keyset
func HaveKey(actual interface{}, key interface{}, explain ...interface{}) {
	gomega.ExpectWithOffset(1, actual).To(gomega.HaveKey(key), explain...)
}

// Empty expects actual is empty
func Empty(actual interface{}, explain ...interface{}) {
	gomega.ExpectWithOffset(1, actual).To(gomega.BeEmpty(), explain...)
}
