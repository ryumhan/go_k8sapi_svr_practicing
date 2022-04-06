package ServerPropsType

import (
	"encoding/json"
	"log"

	v1 "k8s.io/api/apps/v1"
	coreV1 "k8s.io/api/core/v1"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type Deployment struct {
	Image     string `json:"image"`
	Category  string `json:"category"`
	Configmap string `json:"configmap"`
}

type metadata struct {
	Name string `json:"name"`
}

type config struct {
	Config interface{} `json:"config.toml"`
}

type ApiMeta1 struct {
	ApiVersion string      `json:"apiVersion"`
	Kind       string      `json:"kind"`
	Metadata   metadata    `json:"metadata"`
	Spec       interface{} `json:"spec"`
}

type ApiMeta2 struct {
	ApiVersion string   `json:"apiVersion"`
	Kind       string   `json:"kind"`
	Metadata   metadata `json:"metadata"`
	Data       config   `json:"data"`
}

type ApiMeta3 struct {
	ApiVersion string      `json:"apiVersion"`
	Kind       string      `json:"kind"`
	Metadata   metadata    `json:"metadata"`
	Data       interface{} `json:"data"`
}

type DeploymentMeta struct {
	Status int32  `json:"status"`
	Config string `json:"config"`
}

type Data map[string]interface{}

var ConfigManifest = `{
    "kind": "ConfigMap",
    "apiVersion": "v1",
    "metadata": {
        "name": "undefined",
        "namespace": "oncue"
    },
    "data": {
        "config.toml": "undefined"
    }
}`

var DeploymentManifest = `{
    "apiVersion": "apps/v1",
    "kind": "Deployment",
    "metadata": {
        "name": "undefined",
        "namespace": "oncue"
    },
    "spec": {
        "replicas": 1,
        "selector": {
            "matchLabels": {
                "app": "undefined"
            }
        },
        "template": {
            "metadata": {
                "annotations": {
                    "sidecar.istio.io/inject": "false"
                },
                "labels": {
                    "app": "undefined"
                }
            },
            "spec": {
                "containers": [],
                "volumes": []
            }
        }
    }
}`

func MakeDeploymentManifest(name string, image string, configmap string, category string) *v1.Deployment {
	var deploymentManifest *v1.Deployment
	jsonString := []byte(DeploymentManifest)

	err := json.Unmarshal(jsonString, &deploymentManifest)
	if err != nil {
		log.Print("Json.Unmarshal ERROR, jsonString, ", jsonString)
		panic(err.Error())
	}

	deploymentManifest.ObjectMeta.Name = name
	deploymentManifest.Spec.Selector.MatchLabels = map[string]string{"app": name}
	deploymentManifest.Spec.Template.Labels = map[string]string{"app": name}

	var container coreV1.Container
	container.Name = name
	container.Image = image

	//Set config-mount
	var configMount coreV1.VolumeMount
	configMount.Name = "config-volume"
	configMount.MountPath = "/var/lib/oncue/config"

	//Set schema-mount
	var schemaMount coreV1.VolumeMount
	schemaMount.Name = "schema-volume"
	schemaMount.MountPath = "/var/lib/oncue/schema"

	//Set script-mount
	var scriptMount coreV1.VolumeMount
	scriptMount.Name = "script-volume"
	scriptMount.MountPath = "/var/lib/oncue/script"

	var configVolumes coreV1.Volume = coreV1.Volume{
		Name: "config-volume",
		VolumeSource: coreV1.VolumeSource{
			ConfigMap: &coreV1.ConfigMapVolumeSource{
				LocalObjectReference: coreV1.LocalObjectReference{Name: configmap},
			},
		},
	}

	//Set script-Volumes
	var scriptVolumes coreV1.Volume = coreV1.Volume{
		Name: "script-volume",
		VolumeSource: coreV1.VolumeSource{
			ConfigMap: &coreV1.ConfigMapVolumeSource{
				LocalObjectReference: coreV1.LocalObjectReference{Name: "oncue-script"},
			},
		},
	}

	//Set Schema-Volumes
	var dir = coreV1.HostPathDirectory
	var schemaVolumes coreV1.Volume = coreV1.Volume{
		Name: "schema-volume",
		VolumeSource: coreV1.VolumeSource{
			HostPath: &coreV1.HostPathVolumeSource{
				Path: "/var/lib/oncue/schema",
				Type: &dir,
			},
		},
	}

	if category == "protocol" {
		//Set actcode-Volumes
		var actcodeVolumes coreV1.Volume
		actcodeVolumes.Name = "actcode"

		var mode int32 = 400
		actcodeVolumes.Secret = &coreV1.SecretVolumeSource{
			SecretName:  "actcode",
			DefaultMode: &mode,
		}

		container.VolumeMounts = append(container.VolumeMounts, configMount, schemaMount, scriptMount)

		deploymentManifest.Spec.Template.Spec.Containers = append(deploymentManifest.Spec.Template.Spec.Containers, container)
		deploymentManifest.Spec.Template.Spec.Volumes = append(deploymentManifest.Spec.Template.Spec.Volumes, configVolumes, schemaVolumes, scriptVolumes, actcodeVolumes)
	} else if category == "processor" {
		container.VolumeMounts = append(container.VolumeMounts, configMount, schemaMount, scriptMount)

		deploymentManifest.Spec.Template.Spec.Containers = append(deploymentManifest.Spec.Template.Spec.Containers, container)
		deploymentManifest.Spec.Template.Spec.Volumes = append(deploymentManifest.Spec.Template.Spec.Volumes, configVolumes, schemaVolumes, scriptVolumes)
	} else {
		container.VolumeMounts = append(container.VolumeMounts, configMount, schemaMount)

		deploymentManifest.Spec.Template.Spec.Containers = append(deploymentManifest.Spec.Template.Spec.Containers, container)
		deploymentManifest.Spec.Template.Spec.Volumes = append(deploymentManifest.Spec.Template.Spec.Volumes, configVolumes, schemaVolumes)
	}

	return deploymentManifest
}
