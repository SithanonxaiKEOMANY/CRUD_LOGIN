package responses

// Read API using Response Data

//	type StudentResponse struct {
//		ID        uint    `json:"id" gorm:"primaryKey"`
//		StudentID *string `json:"student_id"`
//		Firstname *string `json:"firstname"`
//		Lastname  *string `json:"lastname"`
//		Phone     *string `json:"phone"`
//		Email     *string `json:"email"`
//		Birthday  *string `json:"birthday"`
//		Gender    *string `json:"gender"`
//		Status    int     `json:"status"`
//		Image     *string `json:"image"`
//		CreatedAt string  `json:"created_at"`
//		UpdatedAt string  `json:"updated_at"`
//		DeletedAt *string `json:"deleted_at"`
//	}

type StudentClassroomResponse struct {
	ID          uint      `json:"id"`
	ClassroomID uint      `json:"classroom_id"`
	ClassName   string    `json:"className"`
	SubjectName string    `json:"subject_name"`
	Student     []Student `json:"student"`
}

type Student struct {
	StudentID string `json:"student_id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

type StudentResponse struct {
	ID        uint   `json:"id"`
	StudentID string `json:"student_id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	Birthday  string `json:"birthday"`
	Gender    string `json:"gender"`
	Status    int    `json:"status"`
	Image     string `json:"image"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type MessageResponse struct {
	Message string `json:"message"`
}

type SignUpResponse struct {
	Phone       string `json:"phone"`
	UserType    string `json:"user_type"`
	AccessToken string `json:"access_token"`
}

type SignInResponse struct {
	Phone       string `json:"phone"`
	UserType    string `json:"user_type"`
	AccessToken string `json:"access_token"`
}
