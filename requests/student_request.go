package requests

//CRUD RestAPI Request Data for Create, Update, Delete

type ClassroomIDRequest struct {
	ClassroomID int `json:"classroom_id" validate:"required"`
}

type SigUpRequest struct {
	Phone    string `json:"phone" validate:"required,min=9,max=10"`
	Password string `json:"password" validate:"required"`
	UserType string `json:"user_type" validate:"required"`
	//Token    string `json:"token"`
}

type SignInRequest struct {
	Phone    string `json:"phone" validate:"required,min=9,max=10"`
	Password string `json:"password" validate:"required"`
	UserType string `json:"user_type" validate:"required"`
	//Token    string `json:"token"`
}

type CustomerRequest struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required"`
}

type UpdateCustomerRequest struct {
	Id    int    `json:"id" validate:"required"`
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required"`
}

type DeleteCustomerRequest struct {
	Id    int    `json:"id" validate:"required"`
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required"`
}

type StudentIdRequest struct {
	StudentID string `json:"student_id"`
}

type StudentRequest struct {
	StudentID string `json:"student_id"`
	Firstname string `json:"firstname" `
	Lastname  string `json:"lastname"`
	Phone     string `json:"phone" validate:"required,min=9,max=10"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Birthday  string `json:"birthday"`
	Gender    string `json:"gender"`
	Status    int    `json:"status"`
}

type StudentImageRequest struct {
	StudentID string `json:"student_id" validate:"required"`
	Image     []byte `json:"image" validate:"required"`
}
