package k8sApi

import (
	"context"
	"encoding/json"
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
	case category == "configmanifest":
		return getConfigManifest()
	case category == "images":
		return getCustomImageMap()
	case category == "schema":
		return getSchema()
	case category == "deployment":
		return getConfigMap()
	case category == "deployManifest":
		return getConfigMap()
	default:
		return ServerPropsType.Response{200, "Not Supported", nil}
	}
}

func PutApi(category string) ServerPropsType.Response {
	switch {
	case category == "configmap":
		return getConfigMap()
	case category == "configmanifest":
		return getConfigManifest()
	case category == "images":
		return getCustomImageMap()
	case category == "schema":
		return getConfigMap()
	case category == "deployment":
		return getConfigMap()
	case category == "deploymanifest":
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
			log.Println("Load Config - ", v.ObjectMeta.Name)
		} else if v.ObjectMeta.Name == "oncue-script" {
			configs["oncue-script"] = v.Data
			log.Println("Load Config - ", v.ObjectMeta.Name)
		}
	}

	log.Printf("There are %d configmaps in the cluster\n", len(configs))
	return ServerPropsType.Response{200, "ConfigMaps, OnCue-Script", configs}
}

func getCustomImageMap() ServerPropsType.Response {
	data, err := clientset.RESTClient().Get().AbsPath("apis/oncue.sdplex.com/v1").Namespace("oncue").Resource("images").DoRaw(context.TODO())
	if err != nil {
		panic(err.Error())
	}

	// encode binary to struct type.
	var encoded *v1.ComponentStatusList
	err = json.Unmarshal(data, &encoded)
	if err != nil {
		log.Print("Json.Unmarshal ERROR, *v1.ComponentStatusList", encoded)
		panic(err.Error())
	}

	// parsing the config data.
	images := make(map[string]interface{})
	for _, v := range encoded.Items {
		name := v.ObjectMeta.Name
		annotation := v.GetAnnotations()["kubectl.kubernetes.io/last-applied-configuration"]
		//make jsonstring
		jsonString := []byte(annotation)

		var imagemeta ServerPropsType.ImageSpec
		err = json.Unmarshal(jsonString, &imagemeta)
		if err != nil {
			log.Print("Json.Unmarshal ERROR, imagemeta", imagemeta)
			panic(err.Error())
		}

		images[name] = imagemeta.Spec
		log.Println("Load Image - ", name)
	}

	log.Printf("There are %d images in the cluster\n", len(images))
	return ServerPropsType.Response{400, "Images", images}
}

func getSchema() ServerPropsType.Response {
	// var str, _ = json.Marshal(configs)
	return ServerPropsType.Response{200, "Config & Script", "configs"}
}

func getConfigManifest() ServerPropsType.Response {
	jsonString := []byte(ServerPropsType.ConfigManifest)

	var configManifest ServerPropsType.ApiMeta

	err := json.Unmarshal(jsonString, &configManifest)
	if err != nil {
		log.Print("Json.Unmarshal ERROR, configManifest", configManifest)
		panic(err.Error())
	}

	log.Print("getConfigManifest - ", configManifest)
	return ServerPropsType.Response{200, "ConfigManifestFile", configManifest}
}
