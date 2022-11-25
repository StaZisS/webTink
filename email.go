package listing

type Email struct {
	EmailValue string `json:"email" form:"email" validate:"required,email"`
	Name       string `json:"name" form:"name" validate:"required"`
	Message    string `json:"message" form:"message" validate:"required"`
}
