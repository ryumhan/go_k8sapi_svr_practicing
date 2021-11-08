package ServerPropsType

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// user string array
type SupportData struct {
	configmap      struct{}
	configManifest struct{}
	images         struct{}
	schema         struct{}
	deployment     struct{}
	deployManifest struct{} `json:{
        "apiVersion": "apps/v1",
        "kind": "Deployment",
        "metadata": {
            "name": None,
            "namespace": target
        },
        "spec": {
            "replicas": 1,
            "selector": {
                "matchLabels": {
                    "app": None,
                }
            },
            "template": {
                "metadata": {
                    "annotations": {
                        "sidecar.istio.io/inject": "false"
                    },
                    "labels": {
                        "app": None,
                    }
                },
                "spec": {
                    "containers": [
                        {
                            "name": None,
                            "image": None,
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
                                },
                            ]
                        }],
                    "volumes": [
                        {
                            "name": "config-volume",
                            "configMap": {
                                "name": None
                            }
                        },
                        {
                            "name": "script-volume",
                            "configMap": {
                                "name": None
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
}
