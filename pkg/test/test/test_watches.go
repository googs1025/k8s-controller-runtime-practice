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
	"sigs.k8s.io/controller-runtime/pkg/test/test/pkg"
)

func main() {
	// 新增 manager管理器
	mgr, err := manager.New(common.K8sRestConfig(),
		manager.Options{
			Logger:    logf.Log.WithName("test"),
			Namespace: "default",
		})
	common.Check(err)
	// 新增 controller
	controllerDemo, err := controller.New("test-controller", mgr, controller.Options{
		Reconciler: &pkg.ControllerDemo{},
	})
	common.Check(err)

	// source：事件流，用于监听资源对象
	resources := &source.Kind{
		Type: &v1.Pod{},
	}
	// 框架定义好的回调方法
	handlerFunc := &handler.EnqueueRequestForObject{}
	err = controllerDemo.Watch(resources, handlerFunc)
	common.Check(err)
	// 注入自定义对象
	err = mgr.Add(pkg.NewMyManager(handlerFunc, controllerDemo.(*cc.Controller)))
	common.Check(err)

	err = pkg.AddConfigmapWatch(controllerDemo)
	common.Check(err)
	// 启动manager
	err = mgr.Start(context.Background())
	common.Check(err)

}
