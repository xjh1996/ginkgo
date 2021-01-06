module github.com/caicloud/zeus

go 1.15

require (
	github.com/caicloud/api v0.0.0-20201214034828-9e5d698f165c
	github.com/caicloud/app v1.0.0-alpha.2.0.20201229132246-88cca4ad919f
	github.com/caicloud/auth v0.0.0-20201230062221-41d054766aef
	github.com/caicloud/cargo-server v0.1.0-alpha.2
	github.com/caicloud/containeros v1.0.0-alpha.2.0.20201229105221-6d983eb5abd9
	github.com/caicloud/nirvana v0.3.0-rc.4.0.20201230040208-d9bc298813a9
	github.com/caicloud/nubela v0.0.0-20201230131338-89d6de7a5bc2
	github.com/caicloud/pipeline v0.1.0-alpha.1
	github.com/caicloud/resource v0.0.0-20201228065446-072d17bb1451
	github.com/google/go-github v17.0.0+incompatible
	github.com/onsi/ginkgo v1.14.2
	github.com/onsi/gomega v1.10.1
	github.com/xanzy/go-gitlab v0.28.0
	golang.org/x/mod v0.4.0 // indirect
	golang.org/x/oauth2 v0.0.0-20200107190931-bf48bf16ab8d
	golang.org/x/tools v0.0.0-20201229221835-b8413747bbd4 // indirect
	k8s.io/api v0.19.3
	k8s.io/apimachinery v0.19.3
	k8s.io/client-go v12.0.0+incompatible
	k8s.io/component-base v0.19.3
	k8s.io/klog/v2 v2.4.0
	k8s.io/utils v0.0.0-20201015054608-420da100c033 // indirect
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
