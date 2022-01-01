package user

type RegisterationData struct {
	UniqueId string `json:"uniqueId"`
	Status   string `json:"status,omitempty"`

	Basic struct {
		FirstName string `json:"firstName,omitempty"`
		LastName  string `json:"lastName,omitempty"`
		Phone     string `json:"phone,omitempty"`
		Email     string `json:"email,omitempty"`
	} `json:"basic,omitempty"`

	Documents struct {
		PAN    string `json:"pan,omitempty"`
		Aadhar string `json:"aadhar,omitempty"`
	} `json:"documents,omitempty"`
}
