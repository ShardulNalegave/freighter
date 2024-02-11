package compose

type ComposeConfig struct {
	Address  string `json:"address"`
	Backends []struct {
		Address  string                 `json:"address"`
		Metadata map[string]interface{} `json:"metadata"`
	} `json:"backends"`
}
