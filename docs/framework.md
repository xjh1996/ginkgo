# 介绍
默认框架为每个用例提供一个 framework 作为联通框架资源和用例的桥梁。目前作用有
* 提供系统预置资源信息
* 初始化 kubernetes client 以及业务层 client
* 为用例预置独立的测试分区
* 提供默认的系统资源

# 使用
framework 一般在用例执行前定义，内部处理机制利用了 `ginkgo.BeforeEach`，通过 framework 创建的资源会在用例执行后自动进行删除。

## 初始化
1. NewDefaultFramework
默认使用的 framework 初始化函数，同时创建一个大小为 RequestCPU = "10m" RequestMem = "10M" LimitCPU = "100m" LimitMem = "100M" 的分区，分区名称为 "basename + 随机字符" (basename 是函数入参)
```
func NewDefaultFramework(baseName string) *Framework {}
```
如果对默认预置的分区资源大小不满意，也可以自定义配额，形式如下
```
f := framework.NewDefaultFramework("deployment-basic")
f.NamespceMetadate = &auth.NamespceMetadate{"10", "20m", "100", "200m"}
```

2. NewFramework
如果不希望初始化函数预置分区及初始化 k8s client，可以根据需要显式关闭。
```
func NewFramework(baseName string, skipK8sClientsetCreation, skipNamespaceCreation bool) *Framework {}
```
其中 `skipK8sClientsetCreation` 跳过初始化 k8s client 的过程，`skipNamespaceCreation` 跳过预置分区的过程
## 资源调用
framework 可供外部使用的[参数](https://github.com/caicloud/zeus/blob/master/framework/framework.go#L23)如下更多

```
type Framework struct {
	// Guaranteed to be unique in the cluster even when running the same
	// test multiple times in parallel.
	UniqueName string

	ClientSet         *client.BaseClientType          // 后端所有可用 client，含有 kubernetes 和 containeros clientset
	RestClient        Client.User                     // 平台测试账户 api client
	AdminRestClient   Client.User                     // 平台管理员账户 api client
	ClusterID         string                          // 测试集群 ID
	PresetResource    e2econfig.PresetCompassResource // 平台预置资源
	Namespace         *v1.Namespace                   // 框架预置的分区资源，不创建时为 nil
    ... ...
}
```

可被外部调用的函数
```
func (f *Framework) CreateNamespace(metadate *auth.NamespceMetadate) (*v1.Namespace, error) {}
```
通过该函数创建分区，该分区不需要再手动写删除函数，框架会在用例执行完成后统一进行删除
局限性：使用默认 tenant，cluster，如需要在其他租户或者集群下创建，需要自行进行创建并在用例结束后删除