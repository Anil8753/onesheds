package ledger

type UserCrpto struct {
	MSP        string
	UserId     string
	Cert       string
	PrivateKey string
}

type RegisterationData struct {
	DocType  string `json:"docType,omitempty"`
	UniqueId string `json:"uniqueId"`
	Status   string `json:"status,omitempty"`

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

	Status     string                 `json:"status,omitempty"`
	OwnerId    string                 `json:"ownerId,omitempty"`
	Properties map[string]interface{} `json:"properties,omitempty"`
	Photos     map[string]interface{} `json:"photos,omitempty"`
	Videos     map[string]interface{} `json:"videos,omitempty"`
}
