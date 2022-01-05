package user

type RegisterationData struct {
	DocType  string `json:"docType,omitempty"`
	UniqueId string `json:"uniqueId"`
	Status   string `json:"status,omitempty"`

	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`
	Phone     string `json:"phone,omitempty"`
	Email     string `json:"email,omitempty"`

	PANCard string `json:"pancard,omitempty"`
	Aadhar  string `json:"aadharcard,omitempty"`
}
