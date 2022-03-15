package review

type Reply struct {
	Id      string      `json:"id"`
	UserId  string      `json:"userId"`
	Text    string      `json:"text"`
	Extra   interface{} `json:"extra"`
	Replies []Reply     `json:"replies"`
}

type UserReview struct {
	// Review  Review `json:"review"`
	Id      string      `json:"id"`
	UserId  string      `json:"userId"`
	Rating  float32     `json:"userRating"`
	Text    string      `json:"text"`
	Extra   interface{} `json:"extra"`
	Replies []Reply     `json:"replies"`
}

type Entry struct {
	Id          string     `json:"id"`
	WarehouseId string     `json:"warehouseId"`
	AdminRating float32    `json:"adminRating"`
	UserReview  UserReview `json:"userReview"`
}
