package common

import (
	"fmt"
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
func K8sRestConfig() *rest.Config {
	myDir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	// 读取配置文件，体外
	config, err := clientcmd.BuildConfigFromFlags("", myDir + "/resources/config")
	if err != nil {
		log.Fatal(err)
	}
	//config.Insecure=true
	return config
}
