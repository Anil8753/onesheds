package ledger

type UserCrpto struct {
	MSP        string `json:"msp"`
	UserId     string `json:"userId"`
	Cert       string `json:"cert"`
	PrivateKey string `json:"privateKey"`
}

type RegisterationData struct {
	DocType string `json:"docType,omitempty"`
	UserId  string `json:"userId"`
	Status  string `json:"status,omitempty"`

	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`
	Phone     string `json:"phone,omitempty"`
	Email     string `json:"email,omitempty"`
	Address   string `json:"address,omitempty"`
	PINCode   string `json:"pincode,omitempty"`
	City      string `json:"city,omitempty"`
	District  string `json:"district,omitempty"`
	State     string `json:"state,omitempty"`

	PANCard string `json:"pancard,omitempty"`
	Aadhar  string `json:"aadharcard,omitempty"`
}

type AssetData struct {
	DocType     string `json:"docType,omitempty"`
	WarehouseId string `json:"warehouseId"`

	Status          string                 `json:"status,omitempty"`
	OwnerId         string                 `json:"ownerId,omitempty"`
	TermsConditions []string               `json:"termsConditions,omitempty"`
	Properties      map[string]interface{} `json:"properties,omitempty"`
	Photos          map[string]interface{} `json:"photos,omitempty"`
	Videos          map[string]interface{} `json:"videos,omitempty"`
}
