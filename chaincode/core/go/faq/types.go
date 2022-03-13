package faq

type Answer struct {
	Text  string      `json:"text"`
	Extra interface{} `json:"extra"`
}

type FAQ struct {
	Question string `json:"question"`
	Answer   Answer `json:"answer"`
}

type Entry struct {
	Id          string `json:"id"`
	WarehouseId string `json:"warehouseId"`
	Question    string `json:"question"`
	FAQs        []FAQ  `json:"faqs"`
}

const IDPrefix = "faq"
