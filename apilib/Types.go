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
	deployManifest struct{}
}
