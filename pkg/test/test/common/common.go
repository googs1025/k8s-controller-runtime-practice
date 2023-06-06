package common

import (
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"log"
	"os"
)

func Check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// K8sRestConfig 测试使用
func K8sRestConfig() *rest.Config {
	myDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	// 读取配置文件，体外
	config, err := clientcmd.BuildConfigFromFlags("", myDir+"/resources/config")
	if err != nil {
		log.Fatal(err)
	}
	//config.Insecure=true
	return config
}
