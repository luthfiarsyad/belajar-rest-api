package request

type RequestCreateProduct struct {
	CategoryId int    `json:"category_id" validate:"required,numeric"`
	Name       string `json:"name" validate:"required,min=1,max=200"`
}

type RequestUpdateProduct struct {
	Id         int    `json:"id" validate:"required,numeric"`
	CategoryId int    `json:"category_id" validate:"required,numeric"`
	Name       string `json:"name" validate:"required,min=1,max=200"`
}

type RequestDeleteProduct struct {
	Id int `json:"id" validate:"required,numeric"`
}
