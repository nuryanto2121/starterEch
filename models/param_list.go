package models

// ParamList :
type ParamList struct {
	Page       int    `json:"page" valid:"Required"`
	PerPage    int    `json:"perpage" valid:"Required"`
	Search     string `json:"search,omitempty"`
	InitSearch string `json:"initsearch,omitempty"`
}
