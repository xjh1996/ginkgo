package auth

const (
	defaultRequestCPU = "10m"
	defaultRequestMem = "10M"
	defaultLimitCPU   = "100m"
	defaultLimitMem   = "100M"
)

// Describe resource metadate for a namespace
type NamespceMetadate struct {
	LimitCPU   string
	LimitMem   string
	RequestCPU string
	RequestMem string
}

// DefaultNM returns default namespace metadata
func DefaultNamespaceMeta() *NamespceMetadate {
	return &NamespceMetadate{
		LimitCPU:   defaultLimitCPU,
		LimitMem:   defaultLimitMem,
		RequestCPU: defaultRequestCPU,
		RequestMem: defaultRequestMem,
	}
}
