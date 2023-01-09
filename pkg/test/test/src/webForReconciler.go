package src

import (
	"context"
	"github.com/gin-gonic/gin"
	v1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/internal/controller"
	"sigs.k8s.io/controller-runtime/pkg/manager"
)

type Web struct {
	h handler.EventHandler
	controllerDemo *controller.Controller // 控制器

}


func NewWeb(h handler.EventHandler, controllerDemo *controller.Controller) *Web {
	return &Web{
		h: h,
		controllerDemo: controllerDemo,
	}
}

var _ manager.Runnable = &Web{}	// 是否实现此接口

func (w *Web) Start(ctx context.Context) error {
	r := gin.New()
	r.GET("/add", func(c *gin.Context) {
		p := &v1.Pod{}
		p.Name = "test-controller"
		p.Namespace = "default"

		w.h.Create(event.CreateEvent{Object: p}, w.controllerDemo.Queue)
	})
	r.Run(":8081")
	return nil
}


