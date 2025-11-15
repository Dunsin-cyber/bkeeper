package requests

type CreateCatagoryRequest struct {
	Name     string `json:"name" validate:"required,min=2,max=100"`
	IsCustom bool   `json:"is_custom"`
}
