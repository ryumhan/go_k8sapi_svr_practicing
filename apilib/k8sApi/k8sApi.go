package K8sApi

import (
	"context"
	"flag"
	"io"
	"log"
	"path/filepath"

	ServerPropsType "oncue/apiserver/apilib"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
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
	var err error
	var incluster bool

	// parsing flag
	flag.BoolVar(&incluster, "incluster", false, "Default : false")
	if incluster {
		config = getClusterConfig()
	} else {
		config = getKubeConfig()
	}

	flag.Parse()
	// error check.
	if err != nil {
		panic(err.Error())
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

func getClusterConfig() *restclient.Config {
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}

	log.Print("k8sApi getClusterConfig, Incluster : true")
	return config
}

func getKubeConfig() *restclient.Config {
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	log.Print("k8sApi getKubeConfig, Incluster : false")
	return config
}

func GetApi(category string) ServerPropsType.Response {
	switch {
	case category == "confimap":
		return getConfigMap()
	case category == "configmanifes":
		return getConfigManifest()
	case category == "images":
		return getCustomImageMap()
	case category == "schema":
		return getSchema()
	case category == "deployment":
		return getDeployment()
	case category == "deploymanifst":
		return getDeploymentManifest()
	default:
		return ServerPropsType.Response{400, "Not Supported", nil}
	}
}

func PutApi(category string, body io.ReadCloser) ServerPropsType.Response {
	switch {
	case category == "configmap":
		return applyConfigMap(body)
	case category == "deployment":
		return applyDeployment(body)
	case category == "script":
		return applyScript(body)
	default:
		return ServerPropsType.Response{400, "Not Supported", nil}
	}
}
