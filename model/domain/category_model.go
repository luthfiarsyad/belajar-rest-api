package domain

import "belajar-rest-api/model/response"

type Category struct {
	id   int
	name string
}

func (category *Category) ToResponseCategory() response.ResponseCategory {
	return response.ResponseCategory{
		Id:   category.id,
		Name: category.name,
	}
}

func (c *Category) GetId() *int {
	return &c.id
}

func (c *Category) GetName() *string {
	return &c.name
}

func (c *Category) SetId(id *int) {
	c.id = *id
}
func (c *Category) SetName(name *string) {
	c.name = *name
}
