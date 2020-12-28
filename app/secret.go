package app

import (
	"context"

	"k8s.io/apimachinery/pkg/util/rand"

	types "github.com/caicloud/app/pkg/server/client/v20201010"
	"github.com/caicloud/nubela/expect"
	"github.com/caicloud/zeus/framework"
	"github.com/caicloud/zeus/framework/app"

	"github.com/onsi/ginkgo"
)

var _ = SIGDescribe("保密字典", func() {

	f := framework.NewDefaultFramework("Secret-basic")

	ginkgo.Context("基础部署", func() {
		ginkgo.It("创建+查询+更新+删除", func() {
			testCRUDSecret(f)
		})
	})
})

func testCRUDSecret(f *framework.Framework) {

	// 随机生成配置名称
	secretName := rand.String(20)
	oldKey := rand.String(20)
	oldValue := rand.String(20)
	key := rand.String(20)
	value := rand.String(20)
	clusterID := f.ClusterID

	a, err := f.APIClient.App()
	expect.NoError(err, "App Client Build Failed")
	secret := a.V20201010()

	//新建Secret 传入配置名称和NameSpace
	secretData := app.NewSecret(secretName, namespace, oldKey, oldValue)
	secretGetOption := app.NewSecretGetOptions(clusterID, namespace, secretName)
	_, err = secret.CreateSecret(context.TODO(), secretGetOption, secretData)
	expect.NoError(err, "Create Secret Failed")

	secretData, err = secret.GetSecret(context.TODO(), secretGetOption)
	expect.NoError(err, "Get NewSecret Failed")

	//后期校验多KV值
	//slice01 := []types.SecretData{
	//	{
	//		Key: oldKey,
	//		Value: oldValue,
	//	},
	//}
	expect.Equal(secretData.Data[0], types.SecretData{Key: oldKey, Value: oldValue}, "kv值下发失败")

	//更新KV值
	secretKVUpdate := app.NewUpdateSecret(secretName, namespace, key, value)
	_, err = secret.UpdateSecret(context.TODO(), secretGetOption, secretKVUpdate)
	expect.NoError(err, "Update Secret Failed")

	// Get secret 更新后的信息
	secretData, err = secret.GetSecret(context.TODO(), secretGetOption)
	expect.NoError(err, "Get UpdateSecret Failed")
	expect.Equal(secretData.Data, secretKVUpdate.Data, "kv值更新失败")

	//删除secret
	secretDeleteOption := app.NewSecretDeleteOptions(clusterID, namespace, secretName)
	err = secret.DeleteSecret(context.TODO(), secretDeleteOption)
	expect.NoError(err, "Del Secret Failed")

	//验证删除成功
	_, err = secret.GetSecret(context.TODO(), secretGetOption)
	expect.Error(err, "Del Secret Failed")

}
