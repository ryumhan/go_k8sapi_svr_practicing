package K8sApi

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"strings"

	ServerPropsType "oncue/apiserver/apilib"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

func applyConfigMap(body io.ReadCloser) ServerPropsType.Response {
	var names []string
	var decoded = make(map[string]interface{})

	err := json.NewDecoder(body).Decode(&decoded)
	if err != nil {
		log.Print("Json Decode ERROR, Body - ", body)
		panic(err.Error())
	}

	// get configmaps in "oncue" namespace.
	for name := range decoded {
		var configManifest ServerPropsType.ApiMeta2

		jsonString := []byte(ServerPropsType.ConfigManifest)

		err := json.Unmarshal(jsonString, &configManifest)
		if err != nil {
			log.Print("Json.Unmarshal ERROR, jsonString, ", jsonString)
			panic(err.Error())
		}

		data, err := json.Marshal(decoded[name])
		if err != nil {
			log.Print("Json.Marshal ERROR, decoded[name], ", decoded[name])
			panic(err.Error())
		}

		configManifest.Metadata.Name = name
		configManifest.Data.Config = string(data[:])

		editedConfig, err := json.Marshal(configManifest)
		if err != nil {
			log.Print("Json.Marshal ERROR, configManifest, ", configManifest)
			panic(err.Error())
		}

		result, err := clientset.CoreV1().ConfigMaps("oncue").Patch(context.TODO(), name, types.MergePatchType, editedConfig, metav1.PatchOptions{})
		if err != nil {
			// normal error
			if !strings.HasSuffix(err.Error(), "not found") {
				log.Print("Patch ERROR - ", err.Error())
				panic(err.Error())
			}

			// exist error.
			var newConifg v1.ConfigMap

			err := json.Unmarshal(editedConfig, &newConifg)
			if err != nil {
				log.Print("Json.Unmarshal ERROR, editedConfig, ", editedConfig)
				panic(err.Error())
			}

			result, err := clientset.CoreV1().ConfigMaps("oncue").Create(context.TODO(), &newConifg, metav1.CreateOptions{})
			if err != nil {
				log.Print("ConfigMap Create ERROR")
				panic(err.Error())
			}

			names = append(names, result.Name)
			continue
		}

		names = append(names, result.Name)
	}

	log.Print("PUT - ConfigMap, ", names)
	return ServerPropsType.Response{200, "Applied Configmap - ", names}
}

func applyDeployment(body io.ReadCloser) ServerPropsType.Response {
	var decoded = make(map[string]interface{})

	json.NewDecoder(body).Decode(&decoded)

	return ServerPropsType.Response{200, "Applied ", nil}
}

func applyScript(body io.ReadCloser) ServerPropsType.Response {
	var names []string
	var decoded = make(map[string]interface{})

	err := json.NewDecoder(body).Decode(&decoded)
	if err != nil {
		log.Print("Json Decode ERROR, Body - ", body)
		panic(err.Error())
	}

	// Make script configmap from decoded.
	var scriptManifest ServerPropsType.ApiMeta3
	jsonString := []byte(ServerPropsType.ConfigManifest)

	err = json.Unmarshal(jsonString, &scriptManifest)
	if err != nil {
		log.Print("Json.Unmarshal ERROR, jsonString, ", jsonString)
		panic(err.Error())
	}

	scriptManifest.Metadata.Name = "oncue-script"
	scriptManifest.Data = make(map[string]string)

	script := make(map[string]string)

	for name := range decoded {
		data, err := json.Marshal(decoded[name])
		if err != nil {
			log.Print("Json.Marshal ERROR, decoded[name], ", decoded[name])
			panic(err.Error())
		}

		script[name] = string(data[:])
		names = append(names, name)
	}

	scriptManifest.Data = script
	//Marshal for Patch new configmap
	editedConfig, err := json.Marshal(scriptManifest)
	if err != nil {
		log.Print("Json.Marshal ERROR, script scriptManifest, ", scriptManifest)
		panic(err.Error())
	}

	result, err := clientset.CoreV1().ConfigMaps("oncue").Patch(context.TODO(), "oncue-script", types.MergePatchType, editedConfig, metav1.PatchOptions{})
	if err == nil {
		log.Print("PUT - Patch Script Configmap, ", result.Name, names)
		return ServerPropsType.Response{200, "Applied Script Configmap - ", result.Name}
	}

	// normal error
	if !strings.HasSuffix(err.Error(), "not found") {
		log.Print("Patch ERROR - ", err.Error())
		panic(err.Error())
	}

	// Not Exist Error, Create new configmap.
	var newConifg v1.ConfigMap

	err = json.Unmarshal(editedConfig, &newConifg)
	if err != nil {
		log.Print("Json.Unmarshal ERROR, editedConfig, ", newConifg)
		panic(err.Error())
	}

	// Create new script config map
	result, err = clientset.CoreV1().ConfigMaps("oncue").Create(context.TODO(), &newConifg, metav1.CreateOptions{})
	if err != nil {
		log.Print("Script Configmap Create ERROR")
		panic(err.Error())
	}

	log.Print("PUT - Create Script Configmap, ", result.Name, names)
	return ServerPropsType.Response{200, "Applied Script Configmap - ", result.Name}
}
