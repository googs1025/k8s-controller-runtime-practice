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

func AddCmWatch(controller controller.Controller) error {
	resources := &source.Kind{
		Type: &v1.ConfigMap{},
	}
	e := handler.Funcs{
		CreateFunc: func(event event.CreateEvent, limitingInterface workqueue.RateLimitingInterface) {
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
