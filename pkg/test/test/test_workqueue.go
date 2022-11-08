package main

import (
	"fmt"
	"golang.org/x/time/rate"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/util/workqueue"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"strconv"
	"time"
)

func main() {

}


func newItem(name, ns string) reconcile.Request {
	return reconcile.Request{NamespacedName: types.NamespacedName{
		Name:      name,
		Namespace: ns,
	}}
}

func queuePractice() {
	queue := workqueue.New()
	go func() {
		for {
			item, _ := queue.Get()
			obj := item.(reconcile.Request).NamespacedName
			fmt.Println(obj)
			time.Sleep(time.Millisecond * 20)
		}
	}()

	for {
		queue.Add(newItem("abc", "default"))
		queue.Add(newItem("abc", "default"))
		queue.Add(newItem("abc2", "default"))
		time.Sleep(time.Second * 1)
	}
}

func queuePractice1() {
	queue := workqueue.NewRateLimitingQueue(&workqueue.BucketRateLimiter{
		Limiter: rate.NewLimiter(1, 1),
	})
	go func() {
		for {
			item, _ := queue.Get()
			fmt.Println(item.(reconcile.Request).NamespacedName)

			//手动模拟处理 数据
			queue.Done(item)
		}
	}()

	for i := 0; i < 100; i++ {
		queue.AddRateLimited(newItem("abc"+strconv.Itoa(i), "default"))
		//time.Sleep(time.Millisecond * 200)

	}
	select {}
}

func queuePractice2() {
	limter := rate.NewLimiter(1, 3)
	for i := 0; ; i++ {
		r := limter.Reserve()
		if r.Delay() > 0 {
			fmt.Println("令牌不够了，需要等:", r.Delay())
		}

		fmt.Println(i)
		time.Sleep(time.Millisecond * 200)

	}

	que := workqueue.NewRateLimitingQueue(&workqueue.BucketRateLimiter{
		Limiter: rate.NewLimiter(1, 1),
	})
	go func() {
		for {
			item, _ := que.Get()
			fmt.Println(item.(reconcile.Request).NamespacedName)

			//手动模拟处理 数据
			que.Done(item)
		}
	}()

	for {
		que.AddRateLimited(newItem("abc", "default"))
		time.Sleep(time.Millisecond * 200)

	}

}

