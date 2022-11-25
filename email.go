package listing

type Email struct {
	EmailValue string `json:"email" validate:"required,email"`
	Name       string `json:"name" validate:"required"`
	Message    string `json:"message" validate:"required"`
}
