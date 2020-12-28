package app

import (
	"context"

	"k8s.io/apimachinery/pkg/util/rand"

	types "github.com/caicloud/app/pkg/server/client/v20201010"
	"github.com/caicloud/zeus/framework"
	"github.com/caicloud/zeus/framework/app"

	"github.com/caicloud/nubela/expect"

	"github.com/onsi/ginkgo"
)

// TODO 框架实现Namespace的创建
const namespace = "app"

var _ = SIGDescribe("配置项", func() {
	f := framework.NewDefaultFramework("Configmap-basic")
	ginkgo.Context("基础部署", func() {
		ginkgo.It("创建+查询+更新+删除", func() {
			testCRUDConfigMap(f)
		})
	})

	ginkgo.Context("管理", func() {
		ginkgo.It("罗列配置数据", func() {
			testlistConfigmap(f)
		})
	})

})

func testCRUDConfigMap(f *framework.Framework) {

	ConfigName := rand.String(20)
	oldKey := rand.String(20)
	oldValue := rand.String(20)
	key := rand.String(20)
	value := rand.String(20)
	clusterID := f.ClusterID

	a, err := f.APIClient.App()
	expect.NoError(err, "App Client Build Failed")
	configmap := a.V20201010()

	configmapData := app.NewConfigMap(ConfigName, namespace, oldKey, oldValue)
	configmapGetOption := app.NewConfigGetOptions(clusterID, namespace, ConfigName)
	_, err = configmap.CreateConfigMap(context.TODO(), configmapGetOption, configmapData)
	expect.NoError(err, "Create Configmap Failed")

	configmapData, err = a.V20201010().GetConfigMap(context.TODO(), configmapGetOption)
	expect.NoError(err, "Get Configmap Failed")
	expect.Equal(configmapData.Data[0], types.ConfigMapData{Key: oldKey, Value: oldValue}, "kv值下发失败")

	configmapUpdate := app.NewUpdateConfigMap(ConfigName, namespace, key, value)
	_, err = configmap.UpdateConfigMap(context.TODO(), configmapGetOption, configmapUpdate)
	expect.NoError(err, "Update Configmap Failed")

	configmapData, err = a.V20201010().GetConfigMap(context.TODO(), configmapGetOption)
	expect.NoError(err, "Get Configmap Failed")
	expect.Equal(configmapData.Data, configmapUpdate.Data, "kv值更新失败")

	configmapDeleteOption := app.NewConfigDeleteOptions(clusterID, namespace, ConfigName)
	err = configmap.DeleteConfigMap(context.TODO(), configmapDeleteOption)
	expect.NoError(err, "Del Configmap Failed")

	_, err = configmap.GetConfigMap(context.TODO(), configmapGetOption)
	expect.Error(err, "Configmap Deleted")
}

func testlistConfigmap(f *framework.Framework) {

	clusterID := f.ClusterID
	key := rand.String(20)
	value := rand.String(20)

	var ConfigName [10]string
	a, err := f.APIClient.App()
	expect.NoError(err, "App Client Build Failed")

	//创建10个Configmap
	for i := 0; i < 10; i++ {
		ConfigName[i] = rand.String(20)
		configNameData := app.NewConfigMap(ConfigName[i], namespace, key, value)
		configmapGetOption := app.NewConfigGetOptions(clusterID, namespace, ConfigName[i])
		_, err = a.V20201010().CreateConfigMap(context.TODO(), configmapGetOption, configNameData)
		expect.NoError(err, "Create Configmap Failed")
	}

	listConfigmap := app.NewListOptions(clusterID, namespace)
	_, err = a.V20201010().ListConfigMaps(context.TODO(), listConfigmap)
	expect.NoError(err, "List Configmap Failed")

	//删除所有创建的Configmap
	for i := 0; i < 10; i++ {
		configmapDeleteOption := app.NewConfigDeleteOptions(clusterID, namespace, ConfigName[i])
		err = a.V20201010().DeleteConfigMap(context.TODO(), configmapDeleteOption)
		expect.NoError(err, "Del Configmap Failed")
	}

}
