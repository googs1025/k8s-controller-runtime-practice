package main

import (
	"context"
	v1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	cc "sigs.k8s.io/controller-runtime/pkg/internal/controller"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/source"
	"sigs.k8s.io/controller-runtime/pkg/test/test/common"
	"sigs.k8s.io/controller-runtime/pkg/test/test/src"
)

func main() {

	mgr, err := manager.New(common.K8sRestConfig(),
		manager.Options{
		Logger: logf.Log.WithName("test"),
		Namespace: "default",
		})
	common.Check(err)
	controllerDemo, err := controller.New("test-controller", mgr, controller.Options{
		Reconciler: &src.ControllerDemo{},
	})
	common.Check(err)
	// 资源对象
	resources := &source.Kind{
		Type: &v1.Pod{},
	}
	handlerFunc := &handler.EnqueueRequestForObject{}
	err = controllerDemo.Watch(resources, handlerFunc)
	common.Check(err)
	web := src.NewWeb(handlerFunc, controllerDemo.(*cc.Controller))

	err = mgr.Add(web) // 入参：Runnable接口对象，实现Start方法。
	common.Check(err)
	err = mgr.Start(context.Background())
	common.Check(err)

}
