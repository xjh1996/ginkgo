module github.com/caicloud/zeus

go 1.15

require (
	github.com/caicloud/app v0.0.0-20201214102315-9cf375c41ea5
	github.com/caicloud/cargo-server v0.0.0-20201214095028-bcb03a7a8fe2
	github.com/caicloud/containeros v1.0.0-alpha.1.0.20201214055653-3fd25271b910
	github.com/caicloud/nirvana v0.3.0-rc.3
	github.com/caicloud/nubela v0.0.0-20201214094516-041195f0fe70
	github.com/caicloud/pipeline v0.0.0-20201214105434-1cfdcc4ed9be
	github.com/go-test/deep v1.0.7 // indirect
	github.com/google/uuid v1.1.2 // indirect
	github.com/imdario/mergo v0.3.11 // indirect
	github.com/kr/pretty v0.2.1 // indirect
	github.com/onsi/ginkgo v1.14.0
	github.com/onsi/gomega v1.10.1
	github.com/xanzy/go-gitlab v0.28.0
	k8s.io/api v0.19.2
	k8s.io/apimachinery v0.19.2
	k8s.io/client-go v12.0.0+incompatible
	k8s.io/component-base v0.19.2
	k8s.io/klog/v2 v2.4.0
)

replace (
	github.com/xanzy/go-gitlab => github.com/xanzy/go-gitlab v0.39.0
	k8s.io/api => k8s.io/api v0.19.2
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.19.2
	k8s.io/apimachinery => k8s.io/apimachinery v0.19.2
	k8s.io/apiserver => k8s.io/apiserver v0.19.2
	k8s.io/cli-runtime => k8s.io/cli-runtime v0.19.2
	k8s.io/client-go => k8s.io/client-go v0.19.2
	k8s.io/cloud-provider => k8s.io/cloud-provider v0.19.2
	k8s.io/cluster-bootstrap => k8s.io/cluster-bootstrap v0.19.2
	k8s.io/code-generator => k8s.io/code-generator v0.19.2
	k8s.io/component-base => k8s.io/component-base v0.19.2
	k8s.io/cri-api => k8s.io/cri-api v0.19.2
	k8s.io/csi-translation-lib => k8s.io/csi-translation-lib v0.19.2
	k8s.io/kube-aggregator => k8s.io/kube-aggregator v0.19.2
	k8s.io/kube-controller-manager => k8s.io/kube-controller-manager v0.19.2
	k8s.io/kube-proxy => k8s.io/kube-proxy v0.19.2
	k8s.io/kube-scheduler => k8s.io/kube-scheduler v0.19.2
	k8s.io/kubectl => k8s.io/kubectl v0.19.2
	k8s.io/kubelet => k8s.io/kubelet v0.19.2
	k8s.io/legacy-cloud-providers => k8s.io/legacy-cloud-providers v0.19.2
	k8s.io/metrics => k8s.io/metrics v0.19.2
	k8s.io/sample-apiserver => k8s.io/sample-apiserver v0.19.2
	qiniupkg.com/x => github.com/qiniu/x v1.11.5
)
