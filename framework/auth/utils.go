package auth

import "time"

const (
	defaultRequestCPU = "10m"
	defaultRequestMem = "10M"
	defaultLimitCPU   = "100m"
	defaultLimitMem   = "100M"
	interval          = time.Second * 2
	timeout           = time.Second * 10
)
