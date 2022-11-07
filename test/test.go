package main

import (
	"context"
	"fmt"
	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"log"
	"os"
	"sigs.k8s.io/controller-runtime/pkg/cache"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/cluster"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"time"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
func K8sRestConfig() *rest.Config {
	myDir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	config, err := clientcmd.BuildConfigFromFlags("", myDir + "/resources/config")
	if err != nil {
		log.Fatal(err)
	}
	//config.Insecure=true
	return config
}

func main() {

	mgr, err := manager.New(K8sRestConfig(),
		manager.Options{
			Logger: logf.Log.WithName("test"),
			NewClient: func(cache cache.Cache, config *rest.Config, options client.Options, uncachedObjects ...client.Object) (client.Client, error) {
				return cluster.DefaultNewClient(cache, config, options, &v1.Pod{})
			},
		})

	check(err)
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
		check(err)
		fmt.Println(pod.Name)
	}()

	//本课程来自 程序员在囧途(www.jtthink.com) 咨询群：98514334
	err = mgr.Start(context.Background())
	check(err)

}
