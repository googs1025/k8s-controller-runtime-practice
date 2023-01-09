package src


import (
	"context"
	"fmt"
	controllerruntime "sigs.k8s.io/controller-runtime"
)

// 实现Reconcile方法的对象
type ControllerDemo struct {
}

// Reconcile 协调函数
func (c *ControllerDemo) Reconcile(ctx context.Context, req controllerruntime.Request) (controllerruntime.Result, error) {

	fmt.Println(req.NamespacedName)
	return controllerruntime.Result{}, nil
}

