package utils

type CatalogRequestResponse struct {
	ID        string `json:"@id"`
	Type      string `json:"@type"`
	HasPolicy []struct {
		ID         string `json:"@id"`
		Type       string `json:"@type"`
		Permission []struct {
			Action string `json:"action"`
		} `json:"permission"`
	} `json:"hasPolicy"`
}
