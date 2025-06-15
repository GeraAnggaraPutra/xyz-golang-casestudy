package payload

type UserPayload struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password"`
}
