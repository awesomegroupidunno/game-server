package cmd

type Command struct {
	Type     string `json:"Type"`
	Subtype  string `json:"Subtype"`
	UniqueId string `json:"UniqueId"`
}
