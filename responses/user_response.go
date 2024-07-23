package responses

type UserResponse struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
type MessageUserResponse struct {
	Message string `json:"message"`
}

type SignUpUserResponse struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	AccessToken string `json:"access_token"`
	Message     string `json:"message"`
}

type SignInUserResponse struct {
	Email       string `json:"email"`
	AccessToken string `json:"access_token"`
	Message     string `json:"message"`
}

type Login struct {
	Token string `json:"token"`
}

type ResponseLogin struct {
	//Name  string `json:"name"`
	Email       string `json:"email"`
	AccessToken string `json:"access_token"`
}
