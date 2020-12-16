Zeus 古希腊神话中的众神之王，此 repo 是面向 Compass 的自动化测试用例集。

自动化框架脱胎于 kubernetes 社区 [e2e 项目](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-testing/e2e-tests.md#what-is-ci)，旨在借鉴社区优秀经验在执行用例，在线监控，持续集成取得较好的效果。同时该框架也易于与公司 devops 及 Tesla 等平台集成对接。

# TLDR
用例编写采用 Behavior-Driven Development("BDD") 风格，基于 go 测试框架 [Ginkgo](https://github.com/onsi/ginkgo) 以及匹配断言库 [Gomega](https://github.com/onsi/gomega). 此外我们也对框架进行了一定封装以满足自己的测试需求。
```
var _ = SIGDescribe("无状态服务基础部署", func() {
   Context("使用不同镜像", func() {
        It("自定义镜像，容器数量和存储数量为1，部署能够成功", func() { 
        })
        It("选择集群内镜像，容器数量和存储数量为1，部署能够成功", func() { 
        })
    })
   Context("会话保持", func() {
        It("开关状态修改，配置生效", func() { 
        })
    })
})

var _ = SIGDescribe("无状态服务高级配置", func() {}
```

# 核心特性
* 继承 ginkgo 和 gomega 全部特性，实现 BDD 用例编写风格
* 用例独立随机运行，避免用例按序执行所存在的隐性问题
* 配置管理能力，参数可通过命令行或配置文件进行配置
* 轻松对接研发侧 API Client，用例无需再维护 API 层，提升敏捷效率
* 多种形式的测试报告产出，支持主流 Junit 报告
* 提供丰富的运行日志，包括用例执行前，执行中，及执行后系统信息收集（TODO）
* 提供预置资源能力，框架可对接第三方服务亦或框架自动预置相关资源（TODO）
* 细粒度执行用例，按执行复杂度，模块及项目运行指定用例
* 多参数用例组合生成，探索不同参数组合对结果影响（TODO）


# 安装说明
本框架基于 Golang 开发，并使用 go mod 进行包管理。版本配套关系依赖被测产品版本及 k8s 版本。

* 本地安装
下载用例仓库代码
```
github.com/caicloud/zeus
```
依赖包下载
```
cd zeus/
go mod download
```
* 容器化安装
TODO

# 用例执行
执行 compass 所有用例
```
$ ginkgo 
```
在 192.168.129.30:6060 集群运行 compass app 组用例的命令如下：
```
ginkgo -- --ginkgo.foucus=[cps-app] --baseurl=192.168.129.30:6060
```

## 配置管理
作为 E2E 测试，被测系统会对接第三方服务，例如外部存储，github 等。诸如此类的对接信息可以通过如下方式传入
1. CLI 方式
```
$ ginkgo -- --baseurl=192.168.129.30:6060
```
注意参数配置于 `--` 之后，才能正确生效。所有可用参数及解释可以通过如下方式查看
```
ginkgo -- --help
```
2. 配置文件

e2e 测试配置在未来会显著增加，单纯使用 CLI 方式输入将有很大的局限，这里也提供配置文件对参数进行管理。默认在用例入口路径存在 `e2e-config.yaml` 文件，并在其中定义相关参数。 例如 
```
$ cat e2e-config.yaml
baseurl: 192.168.129.30:6060
```

3. 混合传入

基于 CLI 和 配置文件的的输入也混合使用，例如
```
$ cat /tmp/test.yaml
baseurl: 192.168.129.30:6060

$ ginkgo -- -config=/tmp/test.yaml
```
以上 CLI 指定了配置文件的地址 `/tmp/test.yaml`，框架会从 `/tmp/test.yaml` 中读取其他参数。若 CLI 和 配置文件中的参数重复，CLI 的输入优先级将更高。全量的参数说明可见 [RegisterFlags](https://github.com/caicloud/zeus/blob/master/framework/config/context.go#L93)

## 执行粒度
每个用例有一组标签决定了该用例所属模块，使用场景等。标签作为用例的属性信息，和用例本身没有直接关系，仅作为用例执行范围的依据，通过 `--ginkgo.focus=""` 或 `--ginkgo.skip=""` 运行或跳过指定的带有某类标签的用例。
我们约定有如下标签类型(持续补充中)

类型|名称|使用场景
---|:---| :---
时间复杂度|| 默认与其他用例可并行执行，运行时长在 5min 内
-|[Slow]| 用例运行时长超过 5 min
-|[Serial]| 串型执行，与可并行用例将在不同的测试套中，例如占用大量资源，使用有限资源 (GPU)
-|[Disruptive]| 是 Serial 的子集，标识影响 wokload 服务创建的用例，例如重启组件，taint 节点
-|[Flaky]| 用例含有 bug 并短时间内难以修复，标识该标签，此类用例默认不连跑
外部资源|| 默认不需要特殊环境
-|[WindowOnly]| 仅在 window 环境下运行
-|[Internet]|需要外网环境，比如 develop 连通 github
项目类型| |默认没有模块版本约束
-|[cps-*]| 业务组 label
-|[Smoketest]| 冒烟测试用例
-|[Unionpay]|银联 OEM 用例

> 参考社区 [kind of tests](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-testing/e2e-tests.md#kinds-of-tests)

## 报告管理
用例执行后默认会在 reports 目录下生成对应的 Junit 格式的报告及日志文件。Junit 为报表展示统计信息，日志文件为详细的用例记录信息。

(可选) 本地可视化报表方法
这个一种本地生成可视化报表的方法，后期集成至 Jenkins 或者 Tesla 这一步可以使用平台能力完成，只需要提供上文生成的 Junit 格式的报告就可以。
可视化 Junit 格式的报表可以通过 allure 工具实现。
1. (Mac)下载安装 allure
```
brew install allure
```
2. 生成 HTML 测试报告界面
```
$ allure generate reports/ -o allure-report --clean 
```
测试报告将生成到 allure-report 目录中