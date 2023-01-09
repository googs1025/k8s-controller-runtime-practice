package src

import (
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/util/workqueue"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

// AddConfigmapWatch 监听configmap资源
func AddConfigmapWatch(controller controller.Controller) error {
	resources := &source.Kind{
		Type: &v1.ConfigMap{},
	}
	// 创建时，仍然会调用controller的Reconcile方法
	e := handler.Funcs{
		CreateFunc: func(event event.CreateEvent, limitingInterface workqueue.RateLimitingInterface) {
			// 把add事件的Request加入工作队列
			limitingInterface.Add(reconcile.Request{
				types.NamespacedName{
					Name: "test",
					Namespace: "default",
				},
			})
		},
	}
	return controller.Watch(resources, e)
}
