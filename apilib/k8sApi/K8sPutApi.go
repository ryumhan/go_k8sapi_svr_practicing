package K8sApi

import (
	"context"
	"encoding/json"
	"io"
	"log"

	ServerPropsType "oncue/apiserver/apilib"

	v1 "k8s.io/api/apps/v1"
	coreV1 "k8s.io/api/core/v1"
	k8sError "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

func applyConfigMap(body io.ReadCloser) ServerPropsType.Response {
	var names []string
	var decoded = make(map[string]interface{})

	err := json.NewDecoder(body).Decode(&decoded)
	if err != nil {
		log.Print("Json Decode ERROR, Body Has invalid Character - - ", body)
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
		if err == nil {
			// success patch
			names = append(names, result.Name)
			log.Print("Patch Success- ", result.Name)
			continue
		}

		// another error
		if !k8sError.IsNotFound(err) {
			log.Print("Patch ERROR - ", err.Error())
			panic(err.Error())
		}

		// exist error.
		var newConifg coreV1.ConfigMap

		err = json.Unmarshal(editedConfig, &newConifg)
		if err != nil {
			log.Print("Json.Unmarshal ERROR, editedConfig, ", editedConfig)
			panic(err.Error())
		}

		result, err = clientset.CoreV1().ConfigMaps("oncue").Create(context.TODO(), &newConifg, metav1.CreateOptions{})
		if err != nil {
			log.Print("ConfigMap Create ERROR")
			panic(err.Error())
		}

		names = append(names, result.Name)
		log.Print("Create Success- ", result.Name)
	}

	log.Print("PUT - ConfigMap, ", names)
	return ServerPropsType.Response{200, "Applied Configmap - ", names}
}

func applyDeployment(body io.ReadCloser) ServerPropsType.Response {
	var names []string
	var decoded = make(map[string]ServerPropsType.Deployment)

	err := json.NewDecoder(body).Decode(&decoded)
	if err != nil {
		log.Print("Json Decode ServerPropsType.Deployment Type ERROR, Body Has invalid Character - ", body)
		panic(err.Error())
	}

	// get configmaps in "oncue" namespace.
	for name := range decoded {
		var image = decoded[name].Image
		if image == "" {
			log.Print("Invalid data request, Image value does not exist - ", decoded[name])
			panic(err.Error())
		}

		var category = decoded[name].Category
		if category == "" {
			log.Print("Invalid data request, category value does not exist - ", decoded[name])
			panic(err.Error())
		}

		var configmap = decoded[name].Configmap
		if configmap == "" {
			log.Print("Invalid data request, configmap value does not exist - ", decoded[name])
			panic(err.Error())
		}

		deploymentManifest := ServerPropsType.MakeDeploymentManifest(name, image, configmap, category)
		editedDeployment, err := json.Marshal(deploymentManifest)

		log.Print("MadeDeploymentManifest, ", string(editedDeployment[:]))
		if err != nil {
			log.Print("Json.Marshal ERROR, deploymentManifest, ", deploymentManifest)
			panic(err.Error())
		}

		result, err := clientset.AppsV1().Deployments("oncue").Patch(context.TODO(), name, types.MergePatchType, editedDeployment, metav1.PatchOptions{})
		if err == nil {
			// success patch
			names = append(names, result.Name)
			log.Print("Patch Success- ", result.Name)
			continue
		}

		// another error
		if !k8sError.IsNotFound(err) {
			log.Print("Patch ERROR - ", err.Error())
			panic(err.Error())
		}

		// exist error.
		var newDeployment v1.Deployment
		err = json.Unmarshal(editedDeployment, &newDeployment)
		if err != nil {
			log.Print("Json.Unmarshal ERROR, editedDeployment, ", editedDeployment)
			panic(err.Error())
		}

		result, err = clientset.AppsV1().Deployments("oncue").Create(context.TODO(), &newDeployment, metav1.CreateOptions{})
		if err != nil {
			log.Print("Deployment Create ERROR")
			panic(err.Error())
		}

		names = append(names, result.Name)
		log.Print("Create Success- ", result.Name)
	}

	log.Print("PUT - Deployment, ", names)
	return ServerPropsType.Response{200, "Applied ", nil}
}

func applyScript(body io.ReadCloser) ServerPropsType.Response {
	var names []string
	var decoded = make(map[string]interface{})

	err := json.NewDecoder(body).Decode(&decoded)
	if err != nil {
		log.Print("Json Decode ERROR, Body Has invalid Character - ", body)
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
	if !k8sError.IsNotFound(err) {
		log.Print("Patch ERROR - ", err.Error())
		panic(err.Error())
	}

	// Not Exist Error, Create new configmap.
	var newConifg coreV1.ConfigMap
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
