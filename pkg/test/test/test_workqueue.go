package main

import (
	"context"
	"fmt"
	"golang.org/x/time/rate"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/util/workqueue"
	"log"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"strconv"
	"time"
)

func main() {
	//queuePractice()
	//queuePractice1()
	queuePractice2()
}


func newItem(name, ns string) reconcile.Request {
	return reconcile.Request{NamespacedName: types.NamespacedName{
		Name:      name,
		Namespace: ns,
	}}
}

func queuePractice() {
	// 对列queue
	queue := workqueue.New()
	go func() {
		// 取出
		// 多次取出也只能取出一次
		for {
			item, _ := queue.Get()
			obj := item.(reconcile.Request).NamespacedName
			fmt.Println(obj)
			time.Sleep(time.Millisecond * 20)
			// queue.Done(item) 不使用Done方法，会无法再次调用

		}
	}()

	// 放入queue
	for {
		// 多次加入只能处理一次
		queue.Add(newItem("test", "default"))
		queue.Add(newItem("test", "default"))
		queue.Add(newItem("test2", "default"))
		time.Sleep(time.Second * 1)
	}
}

func queuePractice1() {
	// 限速队列，防止加入队列过快
	queue := workqueue.NewRateLimitingQueue(&workqueue.BucketRateLimiter{	// 令排桶算法
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
		queue.AddRateLimited(newItem("abc"+strconv.Itoa(i), "default")) // 限流
		//time.Sleep(time.Millisecond * 200)

	}
	select {}
}

func queuePractice2() {
	limter := rate.NewLimiter(1, 3)	// 桶里令排有三个，一秒放入一个
	for i := 0; ; i++ {

		// 法一：wait 消费token，如果token不足，会阻塞等待
		err := limter.Wait(context.Background())
		if err != nil {
			log.Fatal(err)
		}

		// 法二：
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

