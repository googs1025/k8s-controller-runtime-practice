package main

import (
	"context"
	"fmt"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/test/test/common"
	"time"
)

func main() {
	// 创建新的manager对象
	mgr, err := manager.New(common.K8sRestConfig(),
		manager.Options{
		Logger: logf.Log.WithName("test"),
		})
	common.Check(err)

	// 需要放到子goroutine中
	go func() {
		time.Sleep(time.Second * 5)
		p := &v1.Pod{}
		mgr.GetClient().Get(context.Background(),
			types.NamespacedName{
			Namespace: "default",
			Name: "hello-world-68fdbf5747-w789w",
			}, p)
		fmt.Println(p.Name, p.Namespace)
	}()
	err = mgr.Start(context.Background())
	common.Check(err)
}
