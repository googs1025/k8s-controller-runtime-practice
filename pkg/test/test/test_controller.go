package main

import (
	"context"
	v1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/source"
	"sigs.k8s.io/controller-runtime/pkg/test/test/common"
	"sigs.k8s.io/controller-runtime/pkg/test/test/src"
)

func main() {
	// 建立管理器 manager
	mgr, err := manager.New(common.K8sRestConfig(),
		manager.Options{
		Logger: logf.Log.WithName("test"),
		Namespace: "default",
		})
	common.Check(err)
	// 创建一个controller控制器
	controllerDemo, err := controller.New("test-controller", mgr, controller.Options{
		Reconciler: &src.ControllerDemo{},
	})
	common.Check(err)
	// 资源
	resources := &source.Kind{
		Type: &v1.Pod{},// 监听资源
	}
	// 监听时 当有不同事件触发时，应该如何回调，使用内置对象(也可以自己实现)
	handlerFunc := &handler.EnqueueRequestForObject{}
	err = controllerDemo.Watch(resources, handlerFunc)
	common.Check(err)

	// 启动管理器
	err = mgr.Start(context.Background())
	common.Check(err)


}

