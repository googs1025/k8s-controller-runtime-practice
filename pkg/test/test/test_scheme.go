package main

import (
	"context"
	"fmt"
	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/rest"
	_ "k8s.io/client-go/tools/clientcmd"
	"sigs.k8s.io/controller-runtime/pkg/cache"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/cluster"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/test/test/common"
	"time"
)

/*
	使用controller-runtime打印GVK
 */

func main() {
	// 创建manager
	mgr, err := manager.New(common.K8sRestConfig(),
		manager.Options{
			Logger: logf.Log.WithName("test"),
			NewClient: func(cache cache.Cache, config *rest.Config, options client.Options, uncachedObjects ...client.Object) (client.Client, error) {
				return cluster.DefaultNewClient(cache, config, options, &v1.Pod{})
			},
		})

	common.Check(err)
	// 打印gvk对象
	ret, _, err := mgr.GetScheme().ObjectKinds(&v1.Pod{})
	ret1, _, err := mgr.GetScheme().ObjectKinds(&appsv1.Deployment{})
	fmt.Println(ret)
	fmt.Println(ret1)
	for k, v := range mgr.GetScheme().AllKnownTypes() {
		fmt.Println("key:", k.String())
		fmt.Println("value:", v.String())
	}

	go func() {

		time.Sleep(time.Second * 3)
		pod := &v1.Pod{}
		fmt.Printf("%T\n", mgr.GetClient())
		err = mgr.GetClient().Get(context.TODO(),
			types.NamespacedName{Namespace: "default", Name: "hello-world-68fdbf5747-679dk"}, pod)
		common.Check(err)
		fmt.Println(pod.Name)
	}()

	err = mgr.Start(context.Background())
	common.Check(err)

}
