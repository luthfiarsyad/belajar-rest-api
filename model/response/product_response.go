package response

type ResponseProduct struct {
	Id         int    `json:"id"`
	CategoryId int    `json:"category_id"`
	Name       string `json:"name"`
}
