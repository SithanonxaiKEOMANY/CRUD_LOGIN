package requests

type UserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignUpUserRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type SignInUserRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
	Token    string `json:"token"`
}
type CreateUserRequest struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required"`
}
type UpdateUserRequest struct {
	ID    uint   `json:"id" validate:"required"`
	Email string `json:"email" validate:"required"`
	Name  string `json:"name" `
}
type DeleteUserRequest struct {
	ID uint `json:"id" validate:"required"`
}
type UserIdRequest struct {
	ID uint `json:"id" validate:"required"`
}

type LoginRequest struct {
	Name     string `json:"name" `
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}
