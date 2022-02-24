package tfconfig

type Locals struct {
	Name  string      `json:"name"`
	Value interface{} `json:"value"`
}
