package user

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
