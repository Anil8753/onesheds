package knowledgebase

type Answer struct {
	Text  string      `json:"text"`
	Extra interface{} `json:"extra"`
}

type Entry struct {
	Questioner string `json:"questioner"`
	Question   string `json:"question"`
	Answer     Answer `json:"answer"`
	Answerer   string `json:"answerer"`
}

type Document struct {
	Id            string  `json:"id"`
	WarehouseId   string  `json:"warehouseId"`
	KnowledgeBase []Entry `json:"knowledgebase"`
}

const IDPrefix = "knowledgebase"
