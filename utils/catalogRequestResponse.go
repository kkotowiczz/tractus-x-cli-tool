package utils

type CatalogRequestResponse struct {
	ID            string `json:"@id"`
	Type          string `json:"@type"`
	OdrlHasPolicy struct {
		ID             string `json:"@id"`
		Type           string `json:"@type"`
		OdrlPermission struct {
			OdrlAction struct {
				ID string `json:"@id"`
			} `json:"odrl:action"`
			OdrlConstraint struct {
				OdrlOr struct {
					OdrlLeftOperand struct {
						ID string `json:"@id"`
					} `json:"odrl:leftOperand"`
					OdrlOperator struct {
						ID string `json:"@id"`
					} `json:"odrl:operator"`
					OdrlRightOperand string `json:"odrl:rightOperand"`
				} `json:"odrl:or"`
			} `json:"odrl:constraint"`
		} `json:"odrl:permission"`
	} `json:"odrl:hasPolicy"`
}
