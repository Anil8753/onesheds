package enquiry

const (
	EnquiryDocType = "Enquiry"
)

type EnquiryData struct {
	DocType    string                 `json:"docType,omitempty"`
	EnquiryId  string                 `json:"enquiryId"`
	Depositor  string                 `json:"depositor"`
	Warehouse  string                 `json:"warehouse"`
	Attributes map[string]interface{} `json:"attributes"`
	Response   map[string]interface{} `json:"response"`
}
