package devops

import (
	"context"

	commonconfig "github.com/caicloud/nubela/config"
	"github.com/caicloud/nubela/expect"
	"github.com/caicloud/zeus/framework"
	"github.com/caicloud/zeus/framework/app"
	"github.com/caicloud/zeus/framework/devops"

	appclient "github.com/caicloud/app/pkg/server/client"
	types "github.com/caicloud/app/pkg/server/client/v20201010"
	cargoclient "github.com/caicloud/cargo-server/pkg/server/client"
	v20201010 "github.com/caicloud/cargo-server/pkg/server/client/v20201010"

	"github.com/onsi/ginkgo"
	"k8s.io/apimachinery/pkg/util/rand"
)

type cargoPublic struct {
	Username            string `default:"admin" usage:"user to login"`
	Password            string `default:"Pwd123456" usage:"password of this user"`
	CargoDomain         string `usage:"domain for harbor"`
	Description         string `usage:"description"`
	DefaultRegistryName string `usage:"name from default Cargo"`
	Basename            string `usage:"name for all resource"`
	Path                string `usage:"path for Dockerfile"`
	ImageRepo           string `usage:"Image Repository for deployment"`
}

var cargoParam cargoPublic
var _ = commonconfig.AddOptions(&cargoParam, "cargo.registry")

var _ = SIGDescribe("Cargo Smoketest", func() {
	var name string
	var err error
	var registry *v20201010.RegistryResp
	var repoOpt = &v20201010.Pagination{Start: 0, Limit: 10}
	var cargoc cargoclient.Interface
	var appc appclient.Interface

	f := framework.NewDefaultFramework("cargo-test")

	ginkgo.BeforeEach(func() {
		cargoc, err = f.AdminAPIClient.Cargo()
		expect.NoError(err)
		appc, err = f.AdminAPIClient.App()
		expect.NoError(err)
	})

	ginkgo.Context("验证新建自定义镜像仓库并", func() {

		ginkgo.BeforeEach(func() {
			name = cargoParam.Basename + rand.String(5)
			registry, err = devops.CreateCargoRegistry(cargoc, name, f.AdminAPIClient.Tenant, cargoParam.CargoDomain, cargoParam.Username, cargoParam.Password)
			expect.NoError(err)
		})

		ginkgo.AfterEach(func() {
			// 删除集成仓库
			err = cargoc.V20201010().DeleteRegistry(context.TODO(), f.AdminAPIClient.Tenant, registry.Name)
			expect.NoError(err)
		})

		ginkgo.It("在系统租户下，在default仓库中新建私有/公有项目", func() {

			// 在default仓库创建公有项目组，创建私有项目组
			publicProject, err := devops.CreatePublicProject(cargoc, name, f.AdminAPIClient.Tenant, cargoParam.DefaultRegistryName, cargoParam.Description)
			expect.NoError(err)
			privateProject, err := devops.CreatePrivateProject(cargoc, name, f.AdminAPIClient.Tenant, cargoParam.DefaultRegistryName, cargoParam.Description)
			expect.NoError(err)

			// 删除公有项目，私有项目
			err = cargoc.V20201010().DeleteArtifactPublicProject(context.TODO(), f.AdminAPIClient.Tenant, cargoParam.DefaultRegistryName, publicProject.Name)
			expect.NoError(err)
			err = cargoc.V20201010().DeleteArtifactProject(context.TODO(), f.AdminAPIClient.Tenant, cargoParam.DefaultRegistryName, privateProject.Name)
			expect.NoError(err)

			// 校验删除成功
			_, err = cargoc.V20201010().GetArtifactPublicProject(context.TODO(), f.AdminAPIClient.Tenant, cargoParam.DefaultRegistryName, publicProject.Name)
			expect.Error(err, "系统租户<default>删除公有项目失败") // FIXME 校验返回码 404
			_, err = cargoc.V20201010().GetArtifactProject(context.TODO(), f.AdminAPIClient.Tenant, cargoParam.DefaultRegistryName, privateProject.Name)
			expect.Error(err, "系统租户<default>删除私有项目失败") // FIXME 校验返回码 404
		})

		ginkgo.It("在系统租户下，在集成仓库中新建私有/公有项目", func() {

			// 在集成仓库创建公有项目组，创建私有项目组
			publicProject2, err := devops.CreatePublicProject(cargoc, name, f.AdminAPIClient.Tenant, registry.Name, cargoParam.Description)
			expect.NoError(err)
			privateProject2, err := devops.CreatePrivateProject(cargoc, name, f.AdminAPIClient.Tenant, registry.Name, cargoParam.Description)
			expect.NoError(err)

			// 删除公有项目，私有项目
			err = cargoc.V20201010().DeleteArtifactPublicProject(context.TODO(), f.AdminAPIClient.Tenant, registry.Name, publicProject2.Name)
			expect.NoError(err)
			err = cargoc.V20201010().DeleteArtifactProject(context.TODO(), f.AdminAPIClient.Tenant, registry.Name, privateProject2.Name)
			expect.NoError(err)

			// 校验删除成功
			_, err = cargoc.V20201010().GetArtifactPublicProject(context.TODO(), f.AdminAPIClient.Tenant, registry.Name, publicProject2.Name)
			expect.Error(err, "系统租户<集成>删除公有项目失败") // FIXME 校验返回码 404
			_, err = cargoc.V20201010().GetArtifactProject(context.TODO(), f.AdminAPIClient.Tenant, registry.Name, privateProject2.Name)
			expect.Error(err, "系统租户<集成>删除私有项目失败") // FIXME 校验返回码 404
		})

		ginkgo.It("在普通租户下，在default仓库中新建私有项目", func() {

			// 在default仓库创建私有项目组
			privateProject, err := devops.CreatePrivateProject(cargoc, name, f.APIClient.Tenant, cargoParam.DefaultRegistryName, cargoParam.Description)
			expect.NoError(err)

			// 删除私有项目
			err = cargoc.V20201010().DeleteArtifactProject(context.TODO(), f.APIClient.Tenant, cargoParam.DefaultRegistryName, privateProject.Name)
			expect.NoError(err)

			// 校验删除成功
			_, err = cargoc.V20201010().GetArtifactProject(context.TODO(), f.APIClient.Tenant, cargoParam.DefaultRegistryName, privateProject.Name)
			expect.Error(err, "普通租户<default>删除私有项目失败") // FIXME 校验返回码 404
		})

		ginkgo.It("在普通租户下，在集成仓库中新建私有项目", func() {

			// 在集成仓库创建私有项目组
			privateProject2, err := devops.CreatePrivateProject(cargoc, name, f.APIClient.Tenant, registry.Name, cargoParam.Description)
			expect.NoError(err)

			// 删除私有项目
			err = cargoc.V20201010().DeleteArtifactProject(context.TODO(), f.APIClient.Tenant, registry.Name, privateProject2.Name)
			expect.NoError(err)

			// 校验删除成功
			_, err = cargoc.V20201010().GetArtifactProject(context.TODO(), f.APIClient.Tenant, registry.Name, privateProject2.Name)
			expect.Error(err, "普通租户<集成>删除私有项目失败") // FIXME 校验返回码 404
		})

	}) //完成

	ginkgo.Context("验证上传/下载/构建镜像在", func() {
		var privateProject *v20201010.Project
		var publicProject *v20201010.PublicProject

		ginkgo.BeforeEach(func() {
			name = cargoParam.Basename + rand.String(5)
		})

		ginkgo.It("系统租户下，default仓库公有项目组中", func() {

			// 在系统租户下<default仓库>新建公有项目组
			publicProject, err = devops.CreatePublicProject(cargoc, name, f.AdminAPIClient.Tenant, cargoParam.DefaultRegistryName, cargoParam.Description)
			expect.NoError(err)

			// 获取公有项目组的信息，校验为空
			repositoryList, _, err := cargoc.V20201010().ListRepositories(context.TODO(), "", f.AdminAPIClient.Tenant, cargoParam.DefaultRegistryName, publicProject.Name, "", "", repoOpt)
			expect.NoError(err)
			expect.Equal(repositoryList.Total, 0, "公有项目组不为空")

			// TODO 上传镜像到该项目组，校验

			// 构建镜像到项目组
			_, err = devops.ImageBuild(cargoc, cargoParam.Path, name, f.AdminAPIClient.Tenant, cargoParam.DefaultRegistryName, publicProject.Name)
			expect.NoError(err)

			// TODO 获取项目组的信息，校验有上述上传/构建的镜像
			// TODO 下载镜像，校验成功

			//删除公有项目组
			err = cargoc.V20201010().DeleteArtifactPublicProject(context.TODO(), f.AdminAPIClient.Tenant, cargoParam.DefaultRegistryName, publicProject.Name)
			expect.NoError(err)

			//校验删除成功
			_, err = cargoc.V20201010().GetArtifactPublicProject(context.TODO(), f.AdminAPIClient.Tenant, cargoParam.DefaultRegistryName, publicProject.Name)
			expect.Error(err, "系统租户<default仓库>删除公有项目失败") // FIXME 校验返回码 404
		})

		ginkgo.It("系统租户下，default仓库私有项目组中", func() {

			// 在系统租户下<default仓库>新建私有项目组
			privateProject, err = devops.CreatePrivateProject(cargoc, name, f.AdminAPIClient.Tenant, cargoParam.DefaultRegistryName, cargoParam.Description)
			expect.NoError(err)

			// 获取私有项目组的信息，校验为空
			repositoryList, _, err := cargoc.V20201010().ListRepositories(context.TODO(), "", f.AdminAPIClient.Tenant, cargoParam.DefaultRegistryName, privateProject.Name, "", "", repoOpt)
			expect.NoError(err)
			expect.Equal(repositoryList.Total, 0, "私有项目组不为空")

			// TODO 上传镜像到该项目组，校验

			// 构建镜像到项目组
			_, err = devops.ImageBuild(cargoc, cargoParam.Path, name, f.AdminAPIClient.Tenant, cargoParam.DefaultRegistryName, privateProject.Name)
			expect.NoError(err)

			// TODO 获取项目组的信息，校验有上述上传/构建的镜像
			// TODO 下载镜像，校验成功

			// 删除系统租户<default仓库>私有项目组
			err = cargoc.V20201010().DeleteArtifactProject(context.TODO(), f.AdminAPIClient.Tenant, cargoParam.DefaultRegistryName, privateProject.Name)
			expect.NoError(err)

			//校验删除成功
			_, err = cargoc.V20201010().GetArtifactProject(context.TODO(), f.AdminAPIClient.Tenant, cargoParam.DefaultRegistryName, privateProject.Name)
			expect.Error(err, "系统租户<default仓库>删除私有项目失败") // FIXME 校验返回码 404
		})

		ginkgo.It("系统租户下，集成仓库公有项目组中", func() {

			// 在系统租户下<集成仓库>新建公有项目组
			publicProject, err = devops.CreatePublicProject(cargoc, name, f.AdminAPIClient.Tenant, f.PresetResource.Cargo.ID, cargoParam.Description)
			expect.NoError(err)

			// 获取项目组的信息，校验为空
			repositoryList, _, err := cargoc.V20201010().ListRepositories(context.TODO(), "", f.AdminAPIClient.Tenant, f.PresetResource.Cargo.ID, publicProject.Name, "", "", repoOpt)
			expect.NoError(err)
			expect.Equal(repositoryList.Total, 0, "公有项目组不为空")

			// TODO 上传镜像到该项目组，校验

			// 构建镜像到项目组
			_, err = devops.ImageBuild(cargoc, cargoParam.Path, name, f.AdminAPIClient.Tenant, f.PresetResource.Cargo.ID, publicProject.Name)
			expect.NoError(err)

			// TODO 获取项目组的信息，校验有上述上传/构建的镜像
			// TODO 下载镜像，校验成功

			// 删除系统租户<集成仓库>公有项目组，校验
			err = cargoc.V20201010().DeleteArtifactPublicProject(context.TODO(), f.AdminAPIClient.Tenant, f.PresetResource.Cargo.ID, publicProject.Name)
			expect.NoError(err)

			//校验删除成功
			_, err = cargoc.V20201010().GetArtifactPublicProject(context.TODO(), f.AdminAPIClient.Tenant, f.PresetResource.Cargo.ID, publicProject.Name)
			expect.Error(err, "系统租户<集成仓库>删除公有项目失败") // FIXME 校验返回码 404
		})

		ginkgo.It("系统租户下，集成仓库私有项目组中", func() {

			// 在系统租户下<集成仓库>新建私有项目组
			privateProject, err = devops.CreatePrivateProject(cargoc, name, f.AdminAPIClient.Tenant, f.PresetResource.Cargo.ID, cargoParam.Description)
			expect.NoError(err)

			// 获取项目组的信息，校验为空
			repositoryList, _, err := cargoc.V20201010().ListRepositories(context.TODO(), "", f.AdminAPIClient.Tenant, f.PresetResource.Cargo.ID, privateProject.Name, "", "", repoOpt)
			expect.NoError(err)
			expect.Equal(repositoryList.Total, 0, "私有项目组不为空")

			// TODO 上传镜像到该项目组，校验

			// 构建镜像到项目组
			_, err = devops.ImageBuild(cargoc, cargoParam.Path, name, f.AdminAPIClient.Tenant, f.PresetResource.Cargo.ID, privateProject.Name)
			expect.NoError(err)

			// TODO 获取项目组的信息，校验有上述上传/构建的镜像
			// TODO 下载镜像，校验成功

			// 删除系统租户<集成仓库>私有项目组，校验
			err = cargoc.V20201010().DeleteArtifactProject(context.TODO(), f.AdminAPIClient.Tenant, f.PresetResource.Cargo.ID, privateProject.Name)
			expect.NoError(err)

			// 校验删除成功
			_, err = cargoc.V20201010().GetArtifactProject(context.TODO(), f.AdminAPIClient.Tenant, f.PresetResource.Cargo.ID, privateProject.Name)
			expect.Error(err, "系统租户<集成仓库>删除私有项目失败") // FIXME 校验返回码 404

		})

		ginkgo.It("普通租户下，default仓库私有项目组中", func() {

			// 在普通租户下<default仓库>新建私有项目组
			privateProject, err = devops.CreatePrivateProject(cargoc, name, f.APIClient.Tenant, cargoParam.DefaultRegistryName, cargoParam.Description)
			expect.NoError(err)

			// 获取私有项目组的信息，校验为空
			repositoryList, _, err := cargoc.V20201010().ListRepositories(context.TODO(), "", f.APIClient.Tenant, cargoParam.DefaultRegistryName, privateProject.Name, "", "", repoOpt)
			expect.NoError(err)
			expect.Equal(repositoryList.Total, 0, "私有项目组不为空")

			// TODO 上传镜像到该项目组，校验

			// 构建镜像到项目组
			_, err = devops.ImageBuild(cargoc, cargoParam.Path, name, f.APIClient.Tenant, cargoParam.DefaultRegistryName, privateProject.Name)
			expect.NoError(err)

			// TODO 获取项目组的信息，校验有上述上传/构建的镜像
			// TODO 下载镜像，校验成功

			// 删除普通租户<default仓库>私有项目组，校验
			err = cargoc.V20201010().DeleteArtifactProject(context.TODO(), f.APIClient.Tenant, cargoParam.DefaultRegistryName, privateProject.Name)
			expect.NoError(err)

			// 校验删除成功
			_, err = cargoc.V20201010().GetArtifactProject(context.TODO(), f.APIClient.Tenant, cargoParam.DefaultRegistryName, privateProject.Name)
			expect.Error(err, "普通租户<default仓库>删除私有项目失败") // FIXME 校验返回码 404
		})

		ginkgo.It("普通租户下，集成仓库私有项目组中", func() {

			// 在普通租户下<集成仓库>新建私有项目组
			privateProject, err = devops.CreatePrivateProject(cargoc, name, f.APIClient.Tenant, f.PresetResource.Cargo.ID, cargoParam.Description)
			expect.NoError(err)

			// 获取私有项目组的信息，校验为空
			repositoryList, _, err := cargoc.V20201010().ListRepositories(context.TODO(), "", f.APIClient.Tenant, f.PresetResource.Cargo.ID, privateProject.Name, "", "", repoOpt)
			expect.NoError(err)
			expect.Equal(repositoryList.Total, 0, "私有项目组不为空")

			// TODO 上传镜像到该项目组，校验

			// 构建镜像到项目组
			_, err = devops.ImageBuild(cargoc, cargoParam.Path, name, f.APIClient.Tenant, f.PresetResource.Cargo.ID, privateProject.Name)
			expect.NoError(err)

			// TODO 获取项目组的信息，校验有上述上传/构建的镜像
			// TODO 下载镜像，校验成功

			// 删除普通租户<集成仓库>私有/公有项目组，校验
			err = cargoc.V20201010().DeleteArtifactProject(context.TODO(), f.APIClient.Tenant, f.PresetResource.Cargo.ID, privateProject.Name)
			expect.NoError(err)

			// 校验删除成功
			_, err = cargoc.V20201010().GetArtifactProject(context.TODO(), f.APIClient.Tenant, f.PresetResource.Cargo.ID, privateProject.Name)
			expect.Error(err, "普通租户<集成仓库>删除私有项目失败") // FIXME 校验返回码 404
		})

	}) // 3.0仅上车镜像构建 // FIXME 之后优化封装重复步骤

	ginkgo.Context("验证镜像仓库同步", func() {
		var privateProject *v20201010.Project
		var publicProject *v20201010.PublicProject
		var trigger = []string{"Manual", "OnPush", "Scheduled"}
		var startReplicationOpt = &v20201010.TriggerReplicationReq{Action: "Start"}

		ginkgo.BeforeEach(func() {
			name = cargoParam.Basename + rand.String(5)
		})

		ginkgo.It("在系统租户新建公有项目同步策略，进行同步", func() {

			// 在系统租户下<default仓库>新建公有项目组
			publicProject, err = devops.CreatePublicProject(cargoc, name, f.AdminAPIClient.Tenant, cargoParam.DefaultRegistryName, cargoParam.Description)
			expect.NoError(err)

			// 获取系统租户下<default仓库>公有项目组，校验为空
			publicRepositoryList, _, err := cargoc.V20201010().ListRepositories(context.TODO(), "", f.AdminAPIClient.Tenant, cargoParam.DefaultRegistryName, publicProject.Name, "", "", repoOpt)
			expect.NoError(err)
			expect.Equal(publicRepositoryList.Total, 0, "系统租户<default仓库>公有项目组不为空")

			// 获取系统租户下<集成仓库>的同名公有项目组，校验不存在(同步策略创建后才会创建)
			publicProjectList, _, err := cargoc.V20201010().ListArtifactPublicProjects(context.TODO(), "", f.AdminAPIClient.Tenant, f.PresetResource.Cargo.ID, repoOpt)
			expect.NoError(err)
			for _, v := range publicProjectList.Items {
				expect.NotEqual(publicProject.Name, v.Name, "系统租户下<集成仓库>存在同名公有项目组")
			}

			// TODO 上传镜像到 系统租户<default仓库>的公有/私有项目组 普通租户<default仓库>的私有项目组，并校验上传成功
			// 暂时通过构建来实现
			_, err = devops.ImageBuild(cargoc, cargoParam.Path, name, f.AdminAPIClient.Tenant, cargoParam.DefaultRegistryName, publicProject.Name)
			expect.NoError(err)

			// 在系统租户下新建公有项目组的同步策略
			replication, err := devops.CreateReplication(cargoc, name, f.AdminAPIClient.Tenant, publicProject.Name, cargoParam.DefaultRegistryName, f.PresetResource.Cargo.ID, trigger[0])
			expect.NoError(err)

			//校验同步策略创建后，target仓库<集成仓库>中项目组已创建
			_, err = cargoc.V20201010().GetArtifactPublicProject(context.TODO(), f.AdminAPIClient.Tenant, f.PresetResource.Cargo.ID, publicProject.Name)
			expect.NoError(err)

			//触发同步策略
			err = cargoc.V20201010().TriggerReplication(context.TODO(), f.AdminAPIClient.Tenant, replication.Name, startReplicationOpt)
			expect.NoError(err)

			// TODO 查看record成功 (暂时通过校验 target 项目组有没有镜像来确定是否同步成功)

			//查看系统/普通租户集成仓库的目标源内含有项目组和镜像
			publicRepositoryList, _, err = cargoc.V20201010().ListRepositories(context.TODO(), "", f.AdminAPIClient.Tenant, f.PresetResource.Cargo.ID, publicProject.Name, "", "", repoOpt)
			expect.NoError(err)
			expect.Equal(publicRepositoryList.Total, 1, "同步失败") //FIXME 准确校验镜像

			//删除同步策略
			err = cargoc.V20201010().DeleteReplication(context.TODO(), f.AdminAPIClient.Tenant, replication.Name)
			expect.NoError(err)

			//校验删除成功
			_, err = cargoc.V20201010().GetReplication(context.TODO(), f.AdminAPIClient.Tenant, replication.Name)
			expect.Error(err) // FIXME 校验返回码404

			// 删除系统租户<default仓库>公有项目组，校验
			err = cargoc.V20201010().DeleteArtifactPublicProject(context.TODO(), f.AdminAPIClient.Tenant, cargoParam.DefaultRegistryName, publicProject.Name)
			expect.NoError(err)
		})

		ginkgo.It("在系统租户新建私有项目同步策略，进行同步", func() {

			// 在系统租户下<default仓库>新建私有项目组
			privateProject, err = devops.CreatePrivateProject(cargoc, name, f.AdminAPIClient.Tenant, cargoParam.DefaultRegistryName, cargoParam.Description)
			expect.NoError(err)

			// 获取系统租户下<default仓库>私有项目组，校验为空
			privateRepositoryList, _, err := cargoc.V20201010().ListRepositories(context.TODO(), "", f.AdminAPIClient.Tenant, cargoParam.DefaultRegistryName, privateProject.Name, "", "", repoOpt)
			expect.NoError(err)
			expect.Equal(privateRepositoryList.Total, 0, "系统租户<default仓库>私有项目组不为空")

			// 获取系统租户下<集成仓库>的同名私有项目组，校验不存在(同步策略创建后才会创建)
			privateProjectList, _, err := cargoc.V20201010().ListArtifactProjects(context.TODO(), "", f.AdminAPIClient.Tenant, "admin", f.PresetResource.Cargo.ID, false, "", repoOpt)
			expect.NoError(err)
			for _, v := range privateProjectList.Items {
				expect.NotEqual(privateProject.Name, v.Name, "系统租户下<集成仓库>存在同名私有项目组")
			}

			// TODO 上传镜像到 系统租户<default仓库>的公有/私有项目组 普通租户<default仓库>的私有项目组，并校验上传成功
			// 暂时通过构建来实现
			_, err = devops.ImageBuild(cargoc, cargoParam.Path, name, f.AdminAPIClient.Tenant, cargoParam.DefaultRegistryName, privateProject.Name)
			expect.NoError(err)

			// 在系统租户下新建私有项目组的同步策略
			replication, err := devops.CreateReplication(cargoc, name, f.AdminAPIClient.Tenant, privateProject.Name, cargoParam.DefaultRegistryName, f.PresetResource.Cargo.ID, trigger[0])
			expect.NoError(err)

			//校验同步策略创建后，target仓库<集成仓库>中项目组已创建
			_, err = cargoc.V20201010().GetArtifactProject(context.TODO(), f.AdminAPIClient.Tenant, f.PresetResource.Cargo.ID, privateProject.Name)
			expect.NoError(err)

			// 触发同步策略
			err = cargoc.V20201010().TriggerReplication(context.TODO(), f.AdminAPIClient.Tenant, replication.Name, startReplicationOpt)
			expect.NoError(err)

			// TODO 查看record成功 (暂时通过校验 target 项目组有没有镜像来确定是否同步成功)

			//查看系统/普通租户集成仓库的目标源内含有项目组和镜像
			privateRepositoryList, _, err = cargoc.V20201010().ListRepositories(context.TODO(), "", f.AdminAPIClient.Tenant, f.PresetResource.Cargo.ID, privateProject.Name, "", "", repoOpt)
			expect.NoError(err)
			expect.Equal(privateRepositoryList.Total, 1, "同步失败") //FIXME 准确校验镜像

			//删除同步策略
			err = cargoc.V20201010().DeleteReplication(context.TODO(), f.AdminAPIClient.Tenant, replication.Name)
			expect.NoError(err)

			//校验删除成功
			_, err = cargoc.V20201010().GetReplication(context.TODO(), f.AdminAPIClient.Tenant, replication.Name)
			expect.Error(err) // FIXME 校验返回码404

			// 删除系统租户<default仓库>私有/公有项目组，校验
			err = cargoc.V20201010().DeleteArtifactProject(context.TODO(), f.AdminAPIClient.Tenant, cargoParam.DefaultRegistryName, privateProject.Name)
			expect.NoError(err)
		})

		ginkgo.It("在普通租户新建私有项目同步策略，进行同步", func() {

			// 在普通租户下<default仓库>新建私有项目组
			privateProject, err = devops.CreatePrivateProject(cargoc, name, f.APIClient.Tenant, cargoParam.DefaultRegistryName, cargoParam.Description)
			expect.NoError(err)

			// 获取普通租户下<default仓库>私有项目组，校验为空
			privateRepositoryList, _, err := cargoc.V20201010().ListRepositories(context.TODO(), "", f.APIClient.Tenant, cargoParam.DefaultRegistryName, privateProject.Name, "", "", repoOpt)
			expect.NoError(err)
			expect.Equal(privateRepositoryList.Total, 0, "普通租户<default仓库>私有项目组不为空")

			// 获取普通租户下<集成仓库>的同名私有项目组，校验不存在(同步策略创建后才会创建)
			privateProjectList, _, err := cargoc.V20201010().ListArtifactProjects(context.TODO(), "", f.APIClient.Tenant, "admin", f.PresetResource.Cargo.ID, false, "", repoOpt)
			expect.NoError(err)
			for _, v := range privateProjectList.Items {
				expect.NotEqual(privateProject.Name, v.Name, "普通租户下<集成仓库>存在同名私有项目组")
			}

			// TODO 上传镜像到 系统租户<default仓库>的公有/私有项目组 普通租户<default仓库>的私有项目组，并校验上传成功
			// 暂时通过构建来实现
			_, err = devops.ImageBuild(cargoc, cargoParam.Path, name, f.AdminAPIClient.Tenant, cargoParam.DefaultRegistryName, privateProject.Name)
			expect.NoError(err)

			//在普通租户下新建私有项目组的同步策略
			replication, err := devops.CreateReplication(cargoc, name, f.APIClient.Tenant, privateProject.Name, cargoParam.DefaultRegistryName, f.PresetResource.Cargo.ID, trigger[0])
			expect.NoError(err)

			//校验同步策略创建后，target仓库<集成仓库>中项目组已创建
			_, err = cargoc.V20201010().GetArtifactProject(context.TODO(), f.APIClient.Tenant, f.PresetResource.Cargo.ID, privateProject.Name)
			expect.NoError(err)

			// 触发同步策略
			err = cargoc.V20201010().TriggerReplication(context.TODO(), f.APIClient.Tenant, replication.Name, startReplicationOpt)
			expect.NoError(err)

			// TODO 查看record成功 (暂时通过校验 target 项目组有没有镜像来确定是否同步成功)

			// TODO 查看系统/普通租户集成仓库的目标源内含有项目组和镜像
			privateRepositoryList, _, err = cargoc.V20201010().ListRepositories(context.TODO(), "", f.APIClient.Tenant, f.PresetResource.Cargo.ID, privateProject.Name, "", "", repoOpt)
			expect.NoError(err)
			expect.Equal(privateRepositoryList.Total, 1, "普通租户<default仓库>私有项目组不为空")

			//删除同步策略
			err = cargoc.V20201010().DeleteReplication(context.TODO(), f.APIClient.Tenant, replication.Name)
			expect.NoError(err)

			//校验删除成功
			_, err = cargoc.V20201010().GetReplication(context.TODO(), f.APIClient.Tenant, replication.Name)
			expect.Error(err) // FIXME 校验返回码404

			// 删除普通租户<default仓库>私有项目组，校验
			err = cargoc.V20201010().DeleteArtifactProject(context.TODO(), f.APIClient.Tenant, cargoParam.DefaultRegistryName, privateProject.Name)
			expect.NoError(err)
		})

	}) // 3.0不上镜像上传/复制，所以通过构建来新增镜像 // FIXME 之后优化封装重复步骤

	ginkgo.Context("镜像扫描", func() {

		ginkgo.BeforeEach(func() {
			name = cargoParam.Basename + rand.String(5)
		})

		ginkgo.It("在default仓库，扫描公有项目library中的一个repo", func() {

			// TODO 上传镜像到该私有项目组

			// 扫描default仓库下library项目组的镜像
			publicRepositoryList, _, err := cargoc.V20201010().ListRepositories(context.TODO(), "", f.AdminAPIClient.Tenant, cargoParam.DefaultRegistryName, "library", "", "", repoOpt)
			expect.NoError(err)
			tagList, err := cargoc.V20201010().ListArtifactPublicTags(context.TODO(), f.AdminAPIClient.Tenant, cargoParam.DefaultRegistryName, "library", publicRepositoryList.Items[0].Name, "", "", repoOpt)
			expect.NoError(err)
			err = cargoc.V20201010().ScanPublicImage(context.TODO(), f.AdminAPIClient.Tenant, cargoParam.DefaultRegistryName, "library", publicRepositoryList.Items[0].Name, tagList.Items[0].Name, tagList.Items[0].Spec.ArtifactDigest)
			expect.NoError(err)

		})

	}) // 3.0不上镜像上传/复制，所以先实现扫描公有项目

	ginkgo.Context("验证镜像部署负载", func() {

		ginkgo.BeforeEach(func() {
			name = cargoParam.Basename + rand.String(5)
		})

		ginkgo.It("在普通租户私有项目组，上传一个镜像到私有项目组，使用该镜像部署服务", func() {

			privateProject, err := devops.CreatePrivateProject(cargoc, name, f.APIClient.Tenant, cargoParam.DefaultRegistryName, cargoParam.Description)
			expect.NoError(err)
			// TODO 上传镜像到该私有项目组
			// TODO 使用该镜像新建无状态服务，校验

			// 构建镜像到项目组
			_, err = devops.ImageBuild(cargoc, cargoParam.Path, name, f.APIClient.Tenant, cargoParam.DefaultRegistryName, privateProject.Name)
			expect.NoError(err)

			deployment := app.NewDeployment(name, "app", int32(1), func(deployment1 *types.Deployment) {
				deployment1.Spec.Template.Spec.Containers[0].Image = cargoParam.ImageRepo
			})

			_, err = app.CreateDeployment(appc, deployment, f.ClusterID, "app", name)
			expect.NoError(err)
		})
	})
})
