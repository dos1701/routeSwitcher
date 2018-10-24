package model

type Route struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Destination string `json:"destination"`
	Mask        string `json:"mask"`
	Metric      int    `json:"metric"`
	Interface   string `json:"interface"`
}
