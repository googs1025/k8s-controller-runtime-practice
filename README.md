# 调用controller-runtime练习
### 本项目是在controller-runtime的包下进行练习的项目，会把一些test文件与代码注释在上面。

```
1. download controller-runtime包文件
2. 加入resources文件，里面放入.kube/config文件
3. 主要测试代码都在pkg/test/test文件中
```

### 项目目录
(主要是在test文件中，其馀都是源码)
 controller_runtime/pkg/test
```
test
├── common  // 读取配置
│   └── common.go
├── src // 代码
│   ├── controllerDemo.go   // 控制器对象，重要
│   ├── watches.go  // 监听不同资源 
│   └── webForReconciler.go // 回调
├ // 测试文件
├── test_controller.go  
├── test_reconciler.go
├── test_scheme.go
├── test_watches.go
└── test_workqueue.go

```

### 不同组件

#### Schema: 
定义k8s中资源序列化和反序列化的方法以及资源类型和版本的对应关系(本身是个map对应关系)。

可以根据GVK找到Go Type, 也可以通过Go Type找到GVK。

#### Informer机制
客户端均通过client-go(K8s系统使用client-go作为Go语言的官方编程式交互客户端库,提供对 K8s API Server服务的交互访问)的Informer机制与Kubernetes API Server进行通信的。

#### Clients 
提供访问API对象的客户端。

#### Caches 
默认情况下客户端从本地缓存读取对象。缓存将自动缓存需要Watch的对象，同时也会缓存其他被请求的结构化对象。
Cache内部是通过Informer负责监听对应 GVK 的 GVR 的创建/删除/更新操作,然后通知所有 Watch 该 GVK 的 Controller, Controller 将对应的资源名称添加到 Queue 里面,最终触发 Reconciler 的调协。

#### Managers
Controller runtime抽象的最外层的管理对象，负责管理 Controller、Caches、Client。

#### Controllers
控制器响应事件(Create/Update/Delete)来触发调协(reconcile)请求,与要实现的调协逻辑一一对应，会创建限速Queue, 一个Controller 可以关注很多**GVK**,然后根据**GVK**到Cache里面找到对应的Share Informer去Watch资源,Watch到的事件会加入到 Queue里面, Queue 最终触发开发者的Reconciler 的调和。

#### Reconcilers
开发者主要实现的逻辑，用来接收Controller的GVK事件，然后获取GVR 进行协调并决定是否更新或者重新入队。

### controller-runtime流程
```bigquery
1. 初始化Schema注册表, 注册内置资源以及自定义资源;
2. 创建并初始化manager对象，将注册表schema传入，并在内部初始化cache和client等资源;
3. 初始化 Reconciler, 传入 client 和 schema
4. 将Reconciler方法加入manager对象，并创建controller对象与Reconciler方法绑定;
5. 用Controller Watch 资源(可以是CR)，controller会从 Cache 里面去拿到Share Informer,如果没有则创建,
6. 然后对ShareInformer 进行 Watch,将得到的资源的名字和 Namespace存入到Queue中；
7. Controller 不断获取 Queue 中的数据并调用 Reconciler 进行调协；
```

