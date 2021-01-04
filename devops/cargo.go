package devops

import (
	"context"
	cargoclient "github.com/caicloud/cargo-server/pkg/server/client"
	"os"

	v1 "github.com/caicloud/api/meta/v1"
	v20201010 "github.com/caicloud/cargo-server/pkg/server/client/v20201010"
	commonconfig "github.com/caicloud/nubela/config"
	"github.com/caicloud/nubela/expect"
	"github.com/caicloud/nubela/logger"
	"github.com/caicloud/zeus/framework"
	"github.com/onsi/ginkgo"

	k8v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/rand"
)

type cargoPublic struct {
	Username            string `default:"admin" usage:"user to login"`
	Password            string `default:"Pwd123456" usage:"password of this user"`
	CargoDomain         string
	SystemTenant        string
	Description         string
	DefaultRegistryName string
	Basename            string
	Path                string
}

var cargoParam cargoPublic
var _ = commonconfig.AddOptions(&cargoParam, "cargo.registry")

var _ = SIGDescribe("Cargo Smoketest", func() {
	//var k8 clientset.Interface
	//k8 = f.ClientSet.K8S
	var name string
	f := framework.NewDefaultFramework("cargo-test")

	ginkgo.Context("验证新建自定义镜像仓库并新建项目", func() {
		var creReg *v20201010.RegistryResp
		name = cargoParam.Basename + rand.String(5)
		ginkgo.BeforeEach(func() {
			// 集成外部仓库
			cargoc, err := f.APIClient.Cargo()
			expect.NoError(err)
			creReg = createCargoRegistry(cargoc, name)
		})

		ginkgo.AfterEach(func() {
			// 删除集成仓库
			cargoc, err := f.APIClient.Cargo()
			expect.NoError(err)
			err = cargoc.V20201010().DeleteRegistry(context.TODO(), cargoParam.SystemTenant, name)
			expect.NoError(err)
		})

		ginkgo.It("分别在default仓库和集成仓库中新建私有/公有项目", func() {

			// 在default仓库创建公有项目组，创建私有项目组
			cargoc, err := f.APIClient.Cargo()
			expect.NoError(err)
			crePubPro := createPublicProject(cargoc, name, cargoParam.SystemTenant, cargoParam.DefaultRegistryName)
			crePriPro := createPrivateProject(cargoc, name, cargoParam.SystemTenant, cargoParam.DefaultRegistryName)

			// 删除公有项目，私有项目
			err = cargoc.V20201010().DeleteArtifactPublicProject(context.TODO(), cargoParam.SystemTenant, cargoParam.DefaultRegistryName, crePubPro.Name)
			expect.NoError(err)
			err = cargoc.V20201010().DeleteArtifactProject(context.TODO(), cargoParam.SystemTenant, cargoParam.DefaultRegistryName, crePriPro.Name)
			expect.NoError(err)

			// 校验删除成功
			_, err = cargoc.V20201010().GetArtifactPublicProject(context.TODO(), cargoParam.SystemTenant, cargoParam.DefaultRegistryName, crePubPro.Name)
			expect.Error(err, "[default]删除公有项目失败")
			_, err = cargoc.V20201010().GetArtifactProject(context.TODO(), cargoParam.SystemTenant, cargoParam.DefaultRegistryName, crePriPro.Name)
			expect.Error(err, "[default]删除私有项目失败")

			// 在集成仓库创建公有项目组，创建私有项目组
			cargoc, err = f.APIClient.Cargo()
			expect.NoError(err)
			crePubPro2 := createPublicProject(cargoc, name, cargoParam.SystemTenant, creReg.Name)
			crePriPro2 := createPrivateProject(cargoc, name, cargoParam.SystemTenant, creReg.Name)

			// 删除公有项目，私有项目
			err = cargoc.V20201010().DeleteArtifactPublicProject(context.TODO(), cargoParam.SystemTenant, creReg.Name, crePubPro2.Name)
			expect.NoError(err)
			err = cargoc.V20201010().DeleteArtifactProject(context.TODO(), cargoParam.SystemTenant, creReg.Name, crePriPro2.Name)
			expect.NoError(err)

			// 校验删除成功
			_, err = cargoc.V20201010().GetArtifactPublicProject(context.TODO(), cargoParam.SystemTenant, creReg.Name, crePubPro2.Name)
			expect.Error(err, "[集成]删除公有项目失败")
			_, err = cargoc.V20201010().GetArtifactProject(context.TODO(), cargoParam.SystemTenant, creReg.Name, crePriPro2.Name)
			expect.Error(err, "[集成]删除私有项目失败")
		})
	}) //完成

	ginkgo.Context("上传/下载/构建镜像", func() {
		var crePriPro *v20201010.Project
		name = cargoParam.Basename + rand.String(5)
		ginkgo.BeforeEach(func() {
			// 在租户下新建私有项目组
			cargoc, err := f.APIClient.Cargo()
			expect.NoError(err)
			crePriPro = createPrivateProject(cargoc, name, cargoParam.SystemTenant, cargoParam.DefaultRegistryName)
		})

		ginkgo.AfterEach(func() {
			// 删除私有项目组，校验
			cargoc, err := f.APIClient.Cargo()
			expect.NoError(err)
			err = cargoc.V20201010().DeleteArtifactProject(context.TODO(), cargoParam.SystemTenant, cargoParam.DefaultRegistryName, crePriPro.Name)
			expect.NoError(err)

			// 校验删除成功
			_, err = cargoc.V20201010().GetArtifactProject(context.TODO(), cargoParam.SystemTenant, cargoParam.DefaultRegistryName, crePriPro.Name)
			expect.Error(err, "删除私有项目失败")
		})

		ginkgo.It("上传一个镜像，下载一个镜像，构建一个镜像", func() {

			// 获取私有项目组的信息，校验为空
			cargoc, err := f.APIClient.Cargo()
			expect.NoError(err)
			lstRepoOpt := &v20201010.Pagination{
				Start: 0,
				Limit: 10,
			}
			lstRepo, _, err := cargoc.V20201010().ListRepositories(context.TODO(), "", cargoParam.SystemTenant, cargoParam.DefaultRegistryName, crePriPro.Name, "", "", lstRepoOpt)
			expect.NoError(err)
			expect.Equal(lstRepo.Total, 0, "私有项目组不为空")

			// TODO 上传镜像到该私有项目组，校验
			//preUpImage, err := cargoc.V20201010().PrepareImageUpload(context.TODO(), userTenant, registryName, projectName)
			//expect.NoError(err)
			//_, err = cargoc.V20201010().UploadImage(context.TODO(), userTenant, preUpImage.Metadata.ID, registryName, projectName, nil)
			//expect.NoError(err)
			//_, err = cargoc.V20201010().GetRepository(context.TODO(), userTenant, registryName, projectName, imageRepo)
			//expect.NoError(err)

			// TODO 构建镜像到该私有项目组
			dockerfile, err := os.Open(cargoParam.Path)
			expect.NoError(err)
			_, err = cargoc.V20201010().BuildImage(context.TODO(), "admin", cargoParam.SystemTenant, cargoParam.DefaultRegistryName, crePriPro.Name, "cntest", dockerfile)
			expect.NoError(err)
			// TODO 查看构建记录，成功
			// TODO 获取私有项目组的信息，校验有上述上传/构建的镜像
			// TODO 下载镜像，校验成功
			// check 1： 业务完成查询操作

			// check 2： 检测 k8s 一致性
			//k8sPartition, err := k8.CoreV1().Namespaces().Get(context.TODO(), nameSpace, metav1.GetOptions{})
			//framework.ExpectNoError(err)
			//gomega.Eventually(k8sPartition.Status.Phase, 100).Should(gomega.BeEquivalentTo("Active"), "The status is not Active within 100 seconds")

			// check 3： 业务层面检查
		})
	}) //完成


})


//获取 pod 状态
func podStatusPhase(f *framework.Framework, x string, y string) k8v1.PodPhase {
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

//集成外部仓库
func createCargoRegistry(cargoc cargoclient.Interface, name string) (creReg *v20201010.RegistryResp) {
	creRegOpt := &v20201010.CreateRegistryReq{
		ObjectMeta: &v1.ObjectMeta{
			Alias: name,
		},
		Spec: &v20201010.RegistrySpec{
			Host:     "https://" + cargoParam.CargoDomain,
			Domain:   cargoParam.CargoDomain,
			Username: cargoParam.Username,
			Password: cargoParam.Password,
		},
	}
	creReg, err := cargoc.V20201010().CreateRegistry(context.TODO(), cargoParam.SystemTenant, "admin", creRegOpt)
	expect.NoError(err)
	expect.Equal(creReg.Alias, name, "create Registry failed")
	return creReg
}

//创建公有项目组
func createPublicProject(cargoc cargoclient.Interface, name, tenant, registry string) (crePubPro *v20201010.PublicProject) {
	crePubProOpt := &v20201010.CreatePublicProjectReq{
		ObjectMeta: &v1.ObjectMeta{
			Name:        name,
			Description: cargoParam.Description,
		},
		Spec: &v20201010.ProjectSpec{
			Registry: registry,
		},
	}
	crePubPro, err := cargoc.V20201010().CreateArtifactPublicProject(context.TODO(), tenant, registry, crePubProOpt)
	expect.NoError(err)
	expect.Equal(crePubPro.Status.Synced, true, "[default]Synced not match")
	expect.Equal(crePubPro.Name, name, "[default]public projectname not match")
	expect.Equal(crePubPro.Description, cargoParam.Description, "[default]public project description not match")
	expect.Equal(crePubPro.Spec.Registry, registry, "[default]public project based registry not match")

	// 校验创建成功
	_, err = cargoc.V20201010().GetArtifactPublicProject(context.TODO(), tenant, registry, name)
	expect.NoError(err)
	return crePubPro
}

//创建私有项目组
func createPrivateProject(cargoc cargoclient.Interface, name, tenant, registry string) (crePriPro *v20201010.Project) {
	crePriProOpt := &v20201010.CreateProjectReq{
		ObjectMeta: &v1.ObjectMeta{
			Name:        name,
			Description: cargoParam.Description,
		},
		Spec: &v20201010.ProjectSpec{
			Registry: registry,
		},
	}
	crePriPro, err := cargoc.V20201010().CreateArtifactProject(context.TODO(), tenant, registry, crePriProOpt)
	expect.NoError(err)
	logger.Infof("%v", crePriPro.Status)
	expect.Equal(crePriPro.Status.Synced, true, "[default]Synced not match")
	expect.Equal(crePriPro.Name, tenant+"_"+name, "[default]private projectname not match")
	expect.Equal(crePriPro.Description, cargoParam.Description, "[default]private project description not match")
	expect.Equal(crePriPro.Spec.Registry, registry, "[default]private project based registry not match")

	// 校验创建成功
	_, err = cargoc.V20201010().GetArtifactProject(context.TODO(), cargoParam.SystemTenant, registry, name)
	expect.NoError(err)
	return crePriPro
}
