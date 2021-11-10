package ServerPropsType

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ImageSpec struct {
	ApiVersion string      `json:"apiVersion"`
	Kind       string      `json:"kind"`
	Metadata   interface{} `json:"metadata"`
	Spec       interface{} `json:"spec"`
}

type ApiMeta struct {
	ApiVersion string      `json:"apiVersion"`
	Kind       string      `json:"kind"`
	Metadata   interface{} `json:"metadata"`
	Data       interface{} `json:"data"`
}

type Data map[string]interface{}

// user string array
type SupportData struct {
	configmap      struct{}
	configManifest struct{}
	images         struct{}
	schema         struct{}
	deployment     struct{}
	deployManifest struct{}
}

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
