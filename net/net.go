package net

import (
	"context"
	"encoding/json"

	"github.com/caicloud/nubela/expect"

	"github.com/caicloud/zeus/framework"
	"github.com/caicloud/zeus/framework/config"
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/tools/clientcmd"
)

type NetworkIPv4Data struct {
	Subnet []SubnetIPv4Data
	PhyDev string `json:"phyDev"`
	VlanID int64  `json:"vlanID"`
}

type SubnetIPv4Data struct {
	Cidr       string `json:"cidr"`
	RangeStart string `json:"rangeStart"`
	RangeEnd   string `json:"rangeEnd"`
	Gateway    string `json:"gateway"`
	Id         string `json:"id,omitempty"`
}

// 初始化 IPv4 参数
var networkIPv4Data = NetworkIPv4Data{
	[]SubnetIPv4Data{
		{
			"137.101.235.0/24",
			"137.101.235.10",
			"137.101.235.20",
			"137.101.235.1",
			"",
		},
	},
	"eth0",
	2418,
}

var _ = SIGDescribe("Bridge-Vlan IPv4 网络管理", func() {
	f := framework.NewFramework("Bridge-Vlan IPv4 MGMT", true, true)

	ginkgo.BeforeEach(func() {
		// do something
	})

	ginkgo.Context("创建 Bridge Vlan 网络", func() {
		var bridgeVlanID string
		ginkgo.BeforeEach(func() {
			//业务 Create Bridge Vlan
			bridgeVlanID = "bridge-vlan-612f14-20201201132840-mjh" // need modify
		})

		ginkgo.It("IPv4 固定", func() {
			// TODO 业务校验
			// check 1： 业务完成查询操作
			// check 2： 检测 k8s 一致性
			testBridgeVlanCheck(f, bridgeVlanID)
		})
		ginkgo.AfterEach(func() {
			// do something
		})
	})
})

func testBridgeVlanCheck(f *framework.Framework, bridgeVlanID string) {
	config, err := clientcmd.BuildConfigFromFlags("", config.Context.KubeConfig)
	expect.NoError(err, "Clientcmd build failed")
	dynamicClient, err := dynamic.NewForConfig(config)
	expect.NoError(err, "dynamic Client init error")

	networksResource := schema.GroupVersionResource{
		Group:    "resource.caicloud.io",
		Version:  "v1beta1",
		Resource: "networks",
	}
	networkCR, err := dynamicClient.Resource(networksResource).Get(context.TODO(), bridgeVlanID, metav1.GetOptions{})
	expect.NoError(err, "dynamicClient get network cr failed")

	name, _, err := unstructured.NestedSlice(networkCR.Object, "spec", "subnets")
	expect.NoError(err, "unstructured.NestedSlice network get failed")

	data, err := json.Marshal(name)
	expect.NoError(err, "Marshal unstructured.NestedSlice failed")

	var subnet []SubnetIPv4Data
	err = json.Unmarshal([]byte(data), &subnet)
	expect.NoError(err, "Unmarshal data failed")

	phyDev, _, err := unstructured.NestedString(networkCR.Object, "spec", "cni", "bridge-vlan", "phyDev")
	expect.NoError(err, "unstructured.NestedSlice phyDev get failed")

	vlanID, _, err := unstructured.NestedInt64(networkCR.Object, "spec", "cni", "bridge-vlan", "vlanID")
	expect.NoError(err, "unstructured.NestedSlice vlanID get failed")
	for i, _ := range subnet {
		gomega.Expect(subnet[i].Cidr).To(gomega.BeEquivalentTo(networkIPv4Data.Subnet[i].Cidr), "cidr 匹配错误")
		gomega.Expect(subnet[i].RangeStart).To(gomega.BeEquivalentTo(networkIPv4Data.Subnet[i].RangeStart), "rangeStart 匹配错误")
		gomega.Expect(subnet[i].RangeEnd).To(gomega.BeEquivalentTo(networkIPv4Data.Subnet[i].RangeEnd), "rangeEnd 匹配错误")
		gomega.Expect(subnet[i].Gateway).To(gomega.BeEquivalentTo(networkIPv4Data.Subnet[i].Gateway), "gateway 匹配错误")
	}
	gomega.Expect(vlanID).To(gomega.BeEquivalentTo(networkIPv4Data.VlanID), "vlanID 匹配错误")
	gomega.Expect(phyDev).To(gomega.BeEquivalentTo(networkIPv4Data.PhyDev), "phyDev 匹配错误")

}
