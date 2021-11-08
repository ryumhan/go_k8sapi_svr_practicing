package k8sApi

import (
	"context"
	"log"
	ServerPropsType "oncue/apiserver/apilib"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func GetApi(category string) ServerPropsType.Response {
	switch {
	case category == "configmap":
		return getConfigMap()
	case category == "configManifest":
		return getConfigMap()
	case category == "images":
		return getConfigMap()
	case category == "schema":
		return getConfigMap()
	case category == "deployment":
		return getConfigMap()
	case category == "deployManifest":
		return getConfigMap()
	default:
		return ServerPropsType.Response{400, "", nil}
	}
}

func PostApi(category string) ServerPropsType.Response {
	switch {
	case category == "configmap":
		return getConfigMap()
	case category == "configManifest":
		return getConfigMap()
	case category == "images":
		return getConfigMap()
	case category == "schema":
		return getConfigMap()
	case category == "deployment":
		return getConfigMap()
	case category == "deployManifest":
		return getConfigMap()
	default:
		return ServerPropsType.Response{400, "", nil}
	}
}

func getConfigMap() ServerPropsType.Response {
	// creates the in-cluster config
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	// get pods in all the namespaces by omitting namespace
	// Or specify namespace to get pods in particular namespace
	configmaps, err := clientset.CoreV1().ConfigMaps("oncue").List(context.TODO(),metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	log.Printf("There are %d pods in the cluster\n", len(configmaps.Items))

	return ServerPropsType.Response{400, "", configmaps.Items}
}
