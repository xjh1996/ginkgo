package devops

import (
	"context"

	"github.com/caicloud/nubela/expect"
	"github.com/caicloud/zeus/framework"
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	clientset "k8s.io/client-go/kubernetes"
)

var _ = SIGDescribe("Cargo Smoketest[cargo-smoketest]", func() {
	var k8 clientset.Interface
	f := framework.NewDefaultFramework("deployment-basic")
	ginkgo.BeforeEach(func() {
		k8 = f.ClientSet.K8S
	})
	ginkgo.Context("验证新建自定义镜像仓库并新建项目", func() {
		ginkgo.BeforeEach(func() {
			// TODO 集成外部仓库，校验集成成功
		})
		ginkgo.AfterEach(func() {
			// TODO 删除集成仓库
		})
		ginkgo.It("集成镜像仓库，新建私有/公有项目", func() {

			// TODO 创建公有项目组，创建私有项目组，校验
			// TODO 删除公有项目，私有项目，校验
			// check 1： 业务完成查询操作

			// check 2： 检测 k8s 一致性
			k8sPartition, err := k8.CoreV1().Namespaces().Get(context.TODO(), nameSpace, metav1.GetOptions{})
			expect.NoError(err)
			gomega.Eventually(k8sPartition.Status.Phase, 100).Should(gomega.BeEquivalentTo("Active"), "The status is not Active within 100 seconds")

			// check 3： 业务层面检查
		})
	})

	ginkgo.Context("基础镜像检测", func() {
		ginkgo.BeforeEach(func() {
			// TODO 打开基础镜像检测开关
		})

		ginkgo.It("检查基础镜像的检测结果", func() {
			// TODO 创建私有项目组，校验
			// TODO 上传镜像到私有项目组，校验
			// TODO 获取该镜像信息，为非基础镜像
			// TODO 上传该镜像到baseimages项目组
			// TODO 获取私有项目组内该镜像的信息，为基础镜像
			// TODO 删除私有项目组

		})
	})

	ginkgo.Context("上传/下载/构建镜像", func() {
		ginkgo.BeforeEach(func() {
			// TODO 在租户下新建私有项目组，校验
		})
		ginkgo.AfterEach(func() {
			// TODO 删除私有项目组，校验
		})
		ginkgo.It("上传一个镜像，下载一个镜像，构建一个镜像", func() {

			// TODO 获取私有项目组的信息，校验为空
			// TODO 上传镜像到该私有项目组
			// TODO 构建镜像到该私有项目组
			// TODO 查看构建记录，成功
			// TODO 获取私有项目组的信息，校验有上述上传/构建的镜像
			// TODO 下载镜像，校验成功
			// check 1： 业务完成查询操作

			// check 2： 检测 k8s 一致性
			k8sPartition, err := k8.CoreV1().Namespaces().Get(context.TODO(), nameSpace, metav1.GetOptions{})
			expect.NoError(err)
			gomega.Eventually(k8sPartition.Status.Phase, 100).Should(gomega.BeEquivalentTo("Active"), "The status is not Active within 100 seconds")

			// check 3： 业务层面检查
		})
	})

	ginkgo.Context("镜像仓库同步", func() {
		ginkgo.BeforeEach(func() {
			// TODO 在租户下新建私有项目组，校验
		})
		ginkgo.AfterEach(func() {
			// TODO 删除私有项目组，校验
		})
		ginkgo.It("新建同步策略，进行手动同步", func() {

			// TODO 获取私有项目组的信息，校验为空
			// TODO 上传镜像到该私有项目组
			// TODO 获取私有项目组的信息，校验有上述上传的镜像
			// TODO 新建同步策略，校验成功
			// TODO 触发同步策略，校验成功
			// TODO 查看集成仓库的目标源内含有镜像
			// check 1： 业务完成查询操作

			// check 2： 检测 k8s 一致性
			k8sPartition, err := k8.CoreV1().Namespaces().Get(context.TODO(), nameSpace, metav1.GetOptions{})
			expect.NoError(err)
			gomega.Eventually(k8sPartition.Status.Phase, 100).Should(gomega.BeEquivalentTo("Active"), "The status is not Active within 100 seconds")

			// check 3： 业务层面检查
		})
	})

	ginkgo.Context("使用私有项目上的镜像部署服务", func() {
		var image string
		var namespace string
		ginkgo.BeforeEach(func() {
			// TODO 在租户下新建私有项目组，校验
			// TODO 新建分区，校验
			image = "cargo30.dev.caicloud.xyz/cntest_github-auto/asdasdad:ssnq7"
			namespace = "cnpartition"
		})
		ginkgo.AfterEach(func() {
			// TODO 删除私有项目组，校验
			// TODO 删除分区
		})
		ginkgo.It("上传一个镜像到私有项目组，使用该镜像部署服务", func() {
			depName := "cd"
			label := "app=" + depName
			// TODO 上传镜像到该私有项目组
			// TODO 使用该镜像新建无状态服务，校验
			// check 1： 业务完成查询操作

			// check 2： 检测 k8s 一致性
			gomega.Eventually(podStatusPhase(f, namespace, label), 100, 5).Should(gomega.BeEquivalentTo("Running"), "The status is not Running within 100 seconds")
			gomega.Eventually(depContainerImage(f, namespace, depName), 100, 5).Should(gomega.BeEquivalentTo(image), "image cannot match within 100 seconds")

			// check 3： 业务层面检查
		})
	})
})

var _ = SIGDescribe("镜像仓库管理", func() {
	//var k8 clientset.Interface
	//f := framework.NewDefaultFramework("deployment-basic")
	ginkgo.BeforeEach(func() {
		//k8 = f.ClientSet.K8S
	})

	ginkgo.Context("镜像仓库集成和删除", func() {
		ginkgo.It("集成一个镜像仓库", func() {
			//TODO 集成镜像仓库，错误的用户名/密码
			//TODO 集成镜像仓库，正确的用户名/密码
			//TODO 删除集成的镜像仓库
		})
	})
})

var _ = SIGDescribe("项目管理", func() {
	//var k8 clientset.Interface
	//f := framework.NewDefaultFramework("deployment-basic")
	ginkgo.BeforeEach(func() {
		//k8 = f.ClientSet.K8S
	})

	ginkgo.Context("公有项目管理", func() {
		ginkgo.It("新增公有项目，修改该公有项目，在此项目上传镜像，删除该公有项目", func() {
			//TODO 新增公有项目
			//TODO 修改公有项目
			//TODO 上传镜像到该公有项目
			//TODO 删除该公有项目
		})
	})
	ginkgo.Context("私有项目管理", func() {
		ginkgo.It("新增私有项目，修改该私有项目，在此项目上传镜像，删除该私有项目", func() {
			//TODO 新增私有项目
			//TODO 修改私有项目
			//TODO 上传镜像到该私有项目
			//TODO 删除该私有项目
		})
	})
})

var _ = SIGDescribe("镜像管理", func() {
	//var k8 clientset.Interface
	//f := framework.NewDefaultFramework("deployment-basic")
	ginkgo.BeforeEach(func() {
		//k8 = f.ClientSet.K8S
		//TODO 预置公有/私有项目组
	})

	ginkgo.Context("镜像构建", func() {
		ginkgo.It("在公有项目组中构建镜像，删除镜像", func() {
			//TODO 构建镜像
			//TODO 删除该镜像
		})
		ginkgo.It("在私有项目组中构建镜像，删除镜像", func() {
			//TODO 构建镜像
			//TODO 删除该镜像
		})
	})

	ginkgo.Context("镜像上传/下载", func() {
		ginkgo.It("在公有项目组中上传镜像，下载该镜像", func() {
			//TODO 在公有项目组中上传镜像
			//TODO 下载该镜像
		})

		ginkgo.It("在私有项目组中上传镜像，下载该镜像", func() {
			//TODO 在私有项目组中上传镜像
			//TODO 下载该镜像
		})
	})
	ginkgo.Context("部署服务", func() {
		ginkgo.It("在default公有项目组中上传镜像，使用该镜像部署服务", func() {
			//TODO 在default公有项目组中上传镜像
			//TODO 使用该镜像部署服务
		})
		ginkgo.It("在default私有项目组中上传镜像，使用该镜像部署服务", func() {
			//TODO 在default私有项目组中上传镜像
			//TODO 使用该镜像部署服务
		})
		ginkgo.It("在集成仓库公有项目组中上传镜像，使用该镜像部署服务", func() {
			//TODO 在集成仓库公有项目组中上传镜像
			//TODO 使用该镜像部署服务
		})
		ginkgo.It("在集成仓库私有项目组中上传镜像，使用该镜像部署服务", func() {
			//TODO 在集成仓库私有项目组中上传镜像
			//TODO 使用该镜像部署服务
		})
	})
})

//获取 pod 状态
func podStatusPhase(f *framework.Framework, x string, y string) v1.PodPhase {
	var k8 = f.ClientSet.K8S
	res, err := k8.CoreV1().Pods(x).List(context.TODO(), metav1.ListOptions{LabelSelector: y})
	expect.NoError(err)
	return res.Items[0].Status.Phase
}

//获取无状态服务所使用的镜像
func depContainerImage(f *framework.Framework, x string, y string) string {
	var k8 = f.ClientSet.K8S
	res, err := k8.AppsV1().Deployments(x).Get(context.TODO(), y, metav1.GetOptions{})
	expect.NoError(err)
	return res.Spec.Template.Spec.Containers[0].Image
}
