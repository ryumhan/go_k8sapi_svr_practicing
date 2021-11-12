package ServerPropsType

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
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
	Metadata   interface{} `json:"metadata"`
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
                "containers": [
                    {
                        "name": "undefined",
                        "image": "undefined",
                        "volumeMounts": [
                            {
                                "name": "config-volume",
                                "mountPath": "/var/lib/oncue/config"
                            },
                            {
                                "name": "schema-volume",
                                "mountPath": "/var/lib/oncue/schema"
                            },
                            {
                                "name": "schema-volume",
                                "mountPath": "/var/lib/oncue/schema"
                            }
                                        ]
                }          
                            ],
                "volumes": [
                    {
                        "name": "config-volume",
                        "configMap": {
                            "name": "undefined"
                        }
                    },
                    {
                        "name": "script-volume",
                        "configMap": {
                            "name": "undefined"
                        }
                    },
                    {
                        "name": "schema-volume",
                        "hostPath": {
                            "path": "/var/lib/oncue/schema",
                            "type": "Directory"
                        }
                    },
                    {
                        "name": "actcode",
                        "secret": {"secretName": "actcode",
                                   "defaultMode": 400
                                   }
                    }
                ]
            }
        }
    }
}`

func GetManifestConfigMap() {

}

func GetManifestDeployment() {

}
