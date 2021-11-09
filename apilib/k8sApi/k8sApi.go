package k8sApi

import (
	"context"

	"flag"
	"log"
	"path/filepath"

	ServerPropsType "oncue/apiserver/apilib"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	restclient "k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

// package variable
var (
	kubeconfig *string
	config     *restclient.Config
	clientset  *kubernetes.Clientset
)

func init() {
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}

	flag.Parse()

	var err error
	config, err = clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err)
	}

	// creates the clientset
	clientset, err = kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	var pods *v1.PodList
	pods, err = clientset.CoreV1().Pods("oncue").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	log.Print("k8sApi Imported - got Pods ", len(pods.Items))
}

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

func PutApi(category string) ServerPropsType.Response {
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
	// get configmaps in "oncue" namespace.
	configmaps, err := clientset.CoreV1().ConfigMaps("oncue").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	configs := make(map[string]interface{})
	// parsing the config data.
	for _, v := range configmaps.Items {
		if len(v.Data["config.toml"]) > 0 {
			configs[v.ObjectMeta.Name] = v.Data["config.toml"]
		} else if v.ObjectMeta.Name == "oncue-script" {
			configs["oncue-script"] = v.Data
		}
	}

	log.Printf("There are %d configmaps in the cluster\n", len(configs))
	// var str, _ = json.Marshal(configs)
	return ServerPropsType.Response{400, "Config & Script", configs}
}

func getCustomImageMap() ServerPropsType.Response {

	// // get configmaps in "oncue" namespace.
	// configmaps, err := clientset.
	// if err != nil {
	// 	panic(err.Error())
	// }

	// configs := make(map[string]interface{})
	// // parsing the config data.
	// for _, v := range configmaps.Items {
	// 	if len(v.Data["config.toml"]) > 0 {
	// 		configs[v.ObjectMeta.Name] = v.Data["config.toml"]
	// 	} else if v.ObjectMeta.Name == "oncue-script" {
	// 		configs["oncue-script"] = v.Data
	// 	}
	// }

	// log.Printf("There are %d configmaps in the cluster\n", len(configs))
	// // var str, _ = json.Marshal(configs)
	return ServerPropsType.Response{400, "Config & Script", "configs"}
}

func getSchema() ServerPropsType.Response {
	// var str, _ = json.Marshal(configs)
	return ServerPropsType.Response{400, "Config & Script", "configs"}
}
