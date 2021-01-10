package devops

import (
	"context"
	"os"
	"time"

	v1 "github.com/caicloud/api/meta/v1"
	cargoclient "github.com/caicloud/cargo-server/pkg/server/client"
	v20201010 "github.com/caicloud/cargo-server/pkg/server/client/v20201010"

	"k8s.io/apimachinery/pkg/util/wait"
)

//集成外部仓库
func CreateCargoRegistry(cargoc cargoclient.Interface, name, tenant, domain, user, passwd string) (registry *v20201010.RegistryResp, err error) {
	registryOpt := &v20201010.CreateRegistryReq{
		ObjectMeta: &v1.ObjectMeta{
			Alias: name,
		},
		Spec: &v20201010.RegistrySpec{
			Host:     "https://" + domain,
			Domain:   domain,
			Username: user,
			Password: passwd,
		},
	}
	return cargoc.V20201010().CreateRegistry(context.TODO(), tenant, "admin", registryOpt)
}

//创建公有项目组
func CreatePublicProject(cargoc cargoclient.Interface, name, tenant, registry, describe string) (publicProject *v20201010.PublicProject, err error) {
	publicProjectOpt := &v20201010.CreatePublicProjectReq{
		ObjectMeta: &v1.ObjectMeta{
			Name:        name,
			Description: describe,
		},
		Spec: &v20201010.ProjectSpec{
			Registry: registry,
		},
	}
	publicProject, err = cargoc.V20201010().CreateArtifactPublicProject(context.TODO(), tenant, registry, publicProjectOpt)
	if err != nil {
		return nil, err
	}
	// 校验创建成功
	_, err = cargoc.V20201010().GetArtifactPublicProject(context.TODO(), tenant, registry, name)
	if err != nil {
		return nil, err
	}
	return publicProject, nil
}

//创建私有项目组
func CreatePrivateProject(cargoc cargoclient.Interface, name, tenant, registry, describe string) (privateProject *v20201010.Project, err error) {
	privateProjectOpt := &v20201010.CreateProjectReq{
		ObjectMeta: &v1.ObjectMeta{
			Name:        name,
			Description: describe,
		},
		Spec: &v20201010.ProjectSpec{
			Registry: registry,
		},
	}
	privateProject, err = cargoc.V20201010().CreateArtifactProject(context.TODO(), tenant, registry, privateProjectOpt)
	if err != nil {
		return nil, err
	}

	// 校验创建成功
	_, err = cargoc.V20201010().GetArtifactProject(context.TODO(), tenant, registry, tenant+"_"+name)
	if err != nil {
		return nil, err
	}
	return privateProject, nil
}

//创建同步策略
func CreateReplication(cargoc cargoclient.Interface, name, tenant, project, sourceReg, targetReg, trigger string) (replication *v20201010.Replication, err error) {
	replicationOpt := &v20201010.CreateReplicationReq{
		ObjectMeta: v1.ObjectMeta{
			Alias: name,
		},
		Spec: v20201010.ReplicationSpec{
			Project: project,
			Source: &v20201010.ReplicationObject{
				Name: sourceReg,
			},
			Target: &v20201010.ReplicationObject{
				Name: targetReg,
			},
			ReplicateDeletion: false,
			Trigger: &v20201010.ReplicationTrigger{
				Kind:     trigger,
				Settings: &v20201010.TriggerSettings{},
			},
		},
	}
	return cargoc.V20201010().CreateReplication(context.TODO(), tenant, replicationOpt)
}

func ImageBuild(cargoc cargoclient.Interface, path, name, tenant, registry, project string) (imageBuildRecord *v20201010.ImageBuildRecordResp, err error) {
	dockerfile, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	//构建镜像
	imageBuild, err := cargoc.V20201010().BuildImage(context.TODO(), "admin", tenant, registry, project, name+":test", dockerfile)
	if err != nil {
		return nil, err
	}

	//查看构建记录，成功
	err = wait.PollImmediate(time.Second*2, time.Second*100, func() (done bool, err error) {
		imageBuildRecord, err = cargoc.V20201010().GetImageBuildRecord(context.TODO(), registry, project, imageBuild.UID)
		if err != nil {
			return false, err
		}
		if imageBuildRecord.Status.Status == "Waiting" { //TODO 等开发代码条后，修改状态为success
			return true, nil
		} else {
			return false, nil
		}
	})
	return imageBuildRecord, err
}
