package services

import (
	"fmt"
	"github.com/pkg/errors"
	"go_starter/errs"
	"go_starter/models"
	"go_starter/repositories"
	"go_starter/requests"
	"go_starter/responses"
	"go_starter/security"
	"go_starter/trails"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type StudentService interface {
	GetStudentClassroomByClassroomIDService(request requests.ClassroomIDRequest) (*responses.StudentClassroomResponse, error)

	SignInService(request requests.SignInRequest) (*responses.SignInResponse, error)
	SignUpService(request requests.SigUpRequest) (*responses.SignUpResponse, error)

	GetStudentService() ([]responses.StudentResponse, error)
	GetStudentByIdService(id uint) (*responses.StudentResponse, error)
	GetStudentByStudentIdServiceV2(request requests.StudentIdRequest) (*responses.StudentResponse, error)
	CreateStudentService(request requests.StudentRequest) (*responses.MessageResponse, error)
	UpdateStudentService(request requests.StudentRequest) (*responses.MessageResponse, error)
	DeleteStudentByIDService(request requests.StudentIdRequest) (*responses.MessageResponse, error)

	//image

	UploadStudentImageService(request requests.StudentImageRequest) (*responses.MessageResponse, error)
}

type studentService struct {
	repositoryStudent repositories.StudentRepository
}

func (s studentService) GetStudentClassroomByClassroomIDService(request requests.ClassroomIDRequest) (*responses.StudentClassroomResponse, error) {
	studentClassrooms, err := s.repositoryStudent.GetStudentClassroomByClassroomIDRepository(uint(request.ClassroomID))
	if err != nil {
		return nil, err
	}
	//fmt.Println("data:", studentClassrooms)

	if len(studentClassrooms) == 0 {
		// Return an empty response or handle the case where no data is found
		return &responses.StudentClassroomResponse{}, nil
	}

	// Assuming the first record to populate classroom details
	classroom := studentClassrooms[0].Classroom
	response := &responses.StudentClassroomResponse{
		ID:          studentClassrooms[0].ID,
		ClassroomID: classroom.ID,
		ClassName:   classroom.ClassName,
		SubjectName: classroom.SubjectName,
		Student:     []responses.Student{},
	}

	for _, sc := range studentClassrooms {
		student := sc.Student
		response.Student = append(response.Student, responses.Student{
			StudentID: student.StudentID,
			Firstname: student.Firstname,
			Lastname:  student.Lastname,
		})
	}

	return response, nil
}

func (s studentService) SignInService(request requests.SignInRequest) (*responses.SignInResponse, error) {
	// Validate phone number
	if request.Phone == "" {
		return nil, errs.ErrorBadRequest("PHONE_CANT_BE_EMPTY")
	}
	if len(request.Phone) > 10 || len(request.Phone) < 9 {
		return nil, errs.ErrorBadRequest("PHONE_INVALID")
	}
	trimSpacePassword := strings.TrimSpace(request.Password)
	if trimSpacePassword == "" {
		return nil, errs.ErrorBadRequest("PASSWORD_CANT_BE_EMPTY")
	}
	switch request.UserType {
	case "teacher":
		getTeacherData, err := s.repositoryStudent.GetTeacherByPhoneRepository(request.Phone)
		if err != nil {
			return nil, errs.ErrorBadRequest("not found teacher id")
		}
		err = security.VerifyPassword(getTeacherData.Password, request.Password)
		if err != nil {
			return nil, fmt.Errorf("password doesn't match")
		}
		response := responses.SignInResponse{
			Phone:       getTeacherData.Phone,
			UserType:    "student",
			AccessToken: "",
		}
		return &response, err

	case "student":
		getStudentData, err := s.repositoryStudent.GetStudentByPhoneRepository(request.Phone)
		if err != nil {
			return nil, err
		}
		err = security.VerifyPassword(getStudentData.Password, request.Password)
		if err != nil {
			return nil, fmt.Errorf("password doesn't match")
		}
		response := responses.SignInResponse{
			Phone:       getStudentData.Phone,
			UserType:    "student",
			AccessToken: "",
		}
		return &response, err
	default:
		return nil, fmt.Errorf("invalid user type")
	}
}

func (s studentService) SignUpService(request requests.SigUpRequest) (*responses.SignUpResponse, error) {
	// Validate phone number
	if request.Phone == "" {
		return nil, errs.ErrorBadRequest("PHONE_CANT_BE_EMPTY")
	}
	if len(request.Phone) > 10 || len(request.Phone) < 9 {
		return nil, errs.ErrorBadRequest("PHONE_INVALID")
	}

	// Handle user type specific logic
	switch request.UserType {
	case "teacher":
		// Check if student's phone number is already in use
		if checkTeacherPhone, err := s.repositoryStudent.CheckTeacherPhoneAlreadyHas(request.Phone); err != nil {
			return nil, err
		} else if checkTeacherPhone {
			return nil, errors.New("phone number already in use")
		}
		trimSpacePassword := strings.TrimSpace(request.Password)
		if trimSpacePassword == "" {
			return nil, errs.ErrorBadRequest("PASSWORD_CANT_BE_EMPTY")
		}
		encryptPassword, err := security.EncryptPassword(request.Password)
		if err != nil {
			return nil, err
		}
		newAccessToken, err := security.NewAccessToken(request.Phone)
		if err != nil {
			return nil, err
		}

		student := models.Teacher{
			Phone:    request.Phone,
			Password: encryptPassword,
			Token:    newAccessToken,
		}
		signUpTeacher, err := s.repositoryStudent.SignUpForTeacherRepository(student)
		if err != nil {
			return nil, err
		}
		responseStudent := responses.SignUpResponse{
			Phone:       signUpTeacher.Phone,
			UserType:    "Teacher",
			AccessToken: signUpTeacher.Token,
		}
		return &responseStudent, nil

	case "student":
		// Check if student's phone number is already in use
		if checkStudentPhone, err := s.repositoryStudent.CheckStudentPhoneAlreadyHas(request.Phone); err != nil {
			return nil, err
		} else if checkStudentPhone {
			return nil, errors.New("phone number already in use")
		}
		trimSpacePassword := strings.TrimSpace(request.Password)
		if trimSpacePassword == "" {
			return nil, errs.ErrorBadRequest("PASSWORD_CANT_BE_EMPTY")
		}
		encryptPassword, err := security.EncryptPassword(request.Password)
		if err != nil {
			return nil, err
		}
		newAccessToken, err := security.NewAccessToken(request.Phone)
		if err != nil {
			return nil, err
		}

		student := models.Student{
			Phone:    request.Phone,
			Password: encryptPassword,
			Token:    newAccessToken,
		}
		signUpStudent, err := s.repositoryStudent.SignUpForStudentRepository(student)
		if err != nil {
			return nil, err
		}
		responseStudent := responses.SignUpResponse{
			Phone:       signUpStudent.Phone,
			UserType:    "Student",
			AccessToken: signUpStudent.Token,
		}
		return &responseStudent, nil

	default:
		return nil, fmt.Errorf("invalid user type")
	}

}

func (s studentService) GetStudentService() ([]responses.StudentResponse, error) {
	// fetch getStudent data from repository(database)
	getStudent, err := s.repositoryStudent.GetStudentsRepository()
	if err != nil {
		return nil, err
	}

	// Check if getStudent is nil
	if getStudent == nil {
		return nil, errors.New("getStudent slice is nil")
	}

	// Business logic
	var response []responses.StudentResponse
	for _, studentData := range getStudent {
		studentResponse := responses.StudentResponse{
			ID:        studentData.ID,
			StudentID: studentData.StudentID,
			Firstname: studentData.Firstname,
			Lastname:  studentData.Lastname,
			Phone:     studentData.Phone,
			Email:     studentData.Email,
			Birthday:  studentData.Birthday.Format("02-01-2006"),
			Gender:    studentData.Gender,
			Status:    studentData.Status,
			Image:     studentData.Image,
			CreatedAt: studentData.CreatedAt.Format("02-01-2006 15:01:05"),
			UpdatedAt: studentData.UpdatedAt.Format("02-01-2006 15:01:05"),
		}

		response = append(response, studentResponse)
	}
	return response, err
}

func (s studentService) GetStudentByIdService(id uint) (*responses.StudentResponse, error) {
	studentData, err := s.repositoryStudent.GetStudentByIdRepository(int(id))
	if err != nil {
		return nil, err
	}
	response := &responses.StudentResponse{
		ID:        studentData.ID,
		StudentID: studentData.StudentID,
		Firstname: studentData.Firstname,
		Lastname:  studentData.Lastname,
		Phone:     studentData.Phone,
		Email:     studentData.Email,
		Birthday:  studentData.Birthday.Format("02-01-2006"),
		Gender:    studentData.Gender,
		Status:    studentData.Status,
		Image:     studentData.Image,
		CreatedAt: studentData.CreatedAt.Format("02-01-2006 15:01:05"),
		UpdatedAt: studentData.UpdatedAt.Format("02-01-2006 15:01:05"),
	}
	return response, err
}

func (s studentService) GetStudentByStudentIdServiceV2(request requests.StudentIdRequest) (*responses.StudentResponse, error) {
	studentData, err := s.repositoryStudent.GetStudentByStudentIdRepository(request.StudentID)
	if err != nil {
		return nil, err
	}
	response := &responses.StudentResponse{
		ID:        studentData.ID,
		StudentID: studentData.StudentID,
		Firstname: studentData.Firstname,
		Lastname:  studentData.Lastname,
		Phone:     studentData.Phone,
		Email:     studentData.Email,
		Birthday:  studentData.Birthday.Format("02-01-2006"),
		Gender:    studentData.Gender,
		Status:    studentData.Status,
		Image:     studentData.Image,
		CreatedAt: studentData.CreatedAt.Format("02-01-2006 15:01:05"),
		UpdatedAt: studentData.UpdatedAt.Format("02-01-2006 15:01:05"),
	}
	return response, err
}

func (s studentService) CreateStudentService(request requests.StudentRequest) (*responses.MessageResponse, error) {
	// Convert the student ID to uppercase
	studentID := strings.ToUpper(request.StudentID)

	// Check if the student ID or phone number is already in use
	if checkStudentID, err := s.repositoryStudent.CheckStudentIDAlreadyHas(studentID); err != nil {
		return nil, err
	} else if checkStudentID {
		return nil, errors.New("student ID already in use")
	}

	if checkPhone, err := s.repositoryStudent.CheckStudentPhoneAlreadyHas(request.Phone); err != nil {
		return nil, err
	} else if checkPhone {
		return nil, errors.New("phone number already in use")
	}

	// Initialize the birthday variable
	var birth time.Time
	// Check if the birthday is provided
	if request.Birthday != "" {
		// Parse the birthday string
		parsedBirth, err := time.Parse("02-01-2006", request.Birthday)
		if err != nil {
			return nil, fmt.Errorf("invalid birthday format: %v", err)
		}
		birth = parsedBirth
	}

	// Create the student model
	model := models.Student{
		StudentID: studentID,
		Firstname: request.Firstname,
		Lastname:  request.Lastname,
		Phone:     request.Phone,
		Email:     request.Email,
		Password:  request.Password,
		Birthday:  birth, // Assign the *time.Time object or nil
		Gender:    request.Gender,
	}

	// Call the repository method to create the student record
	if err := s.repositoryStudent.CreateStudentRepository(&model); err != nil {
		return nil, err
	}

	// If successful, return a success message response
	response := &responses.MessageResponse{Message: "success"}
	return response, nil
}

func (s studentService) UpdateStudentService(request requests.StudentRequest) (*responses.MessageResponse, error) {
	// Convert the student ID to uppercase
	studentID := strings.ToUpper(request.StudentID)

	// Initialize the birthday variable
	var birth time.Time
	// Check if the birthday is provided
	if request.Birthday != "" {
		// Parse the birthday string
		parsedBirth, err := time.Parse("02-01-2006", request.Birthday)
		if err != nil {
			return nil, fmt.Errorf("invalid birthday format: %v", err)
		}
		birth = parsedBirth
	}
	// Create the student model
	model := models.Student{
		StudentID: studentID,
		Firstname: request.Firstname,
		Lastname:  request.Lastname,
		Phone:     request.Phone,
		Email:     request.Email,
		Password:  request.Password,
		Birthday:  birth,
		Gender:    request.Gender}

	// Call the repository method to update the student record
	if err := s.repositoryStudent.UpdateStudentRepository(&model); err != nil {
		return nil, err
	}

	// If successful, return a success message response
	response := &responses.MessageResponse{Message: "success"}
	return response, nil
}

func (s studentService) DeleteStudentByIDService(request requests.StudentIdRequest) (*responses.MessageResponse, error) {
	// Check if the student ID is empty
	if request.StudentID == "" {
		return nil, errors.New("student ID cannot be empty")
	}

	// Call the repository method to delete the student record by ID
	err := s.repositoryStudent.DeleteStudentByStudentIDRepository(request.StudentID)
	if err != nil {
		return nil, err
	}

	// If successful, return a success message response
	response := &responses.MessageResponse{Message: "success"}
	return response, nil
}

func (s studentService) UploadStudentImageService(request requests.StudentImageRequest) (*responses.MessageResponse, error) {
	// Check student id
	if checkStudentID, err := s.repositoryStudent.CheckStudentIDAlreadyHas(request.StudentID); err != nil {
		return nil, err
	} else if !checkStudentID {
		return nil, errors.New("student ID not found")
	}

	// Check if the image exists for the student
	checkData, err := s.repositoryStudent.GetStudentImageRepository(request.StudentID)
	if err != nil {
		return nil, err
	}

	// If an image already exists, delete it
	if checkData != "" {
		err = trails.DeleteImageFile(checkData)
		if err != nil {
			return nil, err
		}
		err = s.repositoryStudent.DeleteStudentImageRepository(request.StudentID)
		if err != nil {
			return nil, err
		}
	}

	// Define the directory structure where the images will be stored
	//directory := filepath.Join("go-starter", "assets", "ceit", "2024", "images")
	//
	directory := "assets/ceit/2024/images"

	// Create the directory if it doesn't exist
	err = os.MkdirAll(directory, os.ModePerm)
	if err != nil {
		return nil, fmt.Errorf("failed to create directory: %v", err)
	}

	// Generate a random number
	randomNumber := trails.GenerateRandomNumber()

	// Generate the file path for storing the image
	imagePath := filepath.Join(directory, request.StudentID+""+randomNumber+".png")

	// Write the image data to a file
	err = ioutil.WriteFile(imagePath, request.Image, 0644)
	if err != nil {
		return nil, fmt.Errorf("failed to write image to file: %v", err)
	}

	// Update the student image path in the database
	err = s.repositoryStudent.UpdateStudentImageRepository(&models.Student{
		StudentID: request.StudentID,
		Image:     imagePath,
	})
	if err != nil {
		return nil, err
	}

	// Return success message
	response := &responses.MessageResponse{Message: "uploaded success"}
	return response, nil
}

// ------ handle with pointer

//func (s studentService) UploadStudentImageService(request requests.StudentImageRequest) (*responses.MessageResponse, error) {
//	// Check student id
//	if checkStudentID, err := s.repositoryStudent.CheckStudentIDAlreadyHas(request.StudentID); err != nil {
//		return nil, err
//	} else if !checkStudentID {
//		return nil, errors.New("student ID not found")
//	}
//
//	// Check if the image exists for the student
//	checkData, err := s.repositoryStudent.GetStudentImageRepository(request.StudentID)
//	if err != nil {
//		return nil, err
//	}
//	fmt.Printf("%v\n", checkData)
//
//	// If an image already exists, delete it
//	if checkData != "" {
//		err = trails.DeleteImageFile(checkData)
//		if err != nil {
//			return nil, err
//		}
//		err = s.repositoryStudent.DeleteStudentImageRepository(request.StudentID)
//		if err != nil {
//			return nil, err
//		}
//	}
//
//	// Define the directory structure where the images will be stored
//	directory := "/Users/kidtikonemks/Documents/Projects/go/Starter/go-starter/assets/ceit/2024/images"
//
//	// Create the directory if it doesn't exist
//	err = os.MkdirAll(directory, os.ModePerm)
//	if err != nil {
//		return nil, fmt.Errorf("failed to create directory: %v", err)
//	}
//
//	//random number
//	studentIDWithRandom := request.StudentID + trails.GenerateRandomNumber()
//
//	// Generate the file path for storing the image
//	imagePath := filepath.Join(directory, studentIDWithRandom+".png")
//
//	// Write the image data to a file
//	err = ioutil.WriteFile(imagePath, request.Image, 0644)
//	if err != nil {
//		return nil, fmt.Errorf("failed to write image to file: %v", err)
//	}
//
//	// Update the student image path in the database
//	err = s.repositoryStudent.UpdateStudentImageRepository(&models.Student{
//		StudentID: request.StudentID,
//		Image:     imagePath,
//	})
//	if err != nil {
//		return nil, err
//	}
//
//	// Return success message
//	response := &responses.MessageResponse{Message: "uploaded successfully"}
//	return response, nil
//}

//
//	func getStringPointerValue(s *string) *string {
//		if s == nil {
//			return nil // Return nil if the string pointer is nil
//		}
//		return s // Otherwise, return the pointer itself
//	}
//
// // getStringPointer returns a pointer to string if the input is not empty, otherwise returns nil
//
//	func getStringPointer(s string) *string {
//		if s != "" {
//			return &s
//		}
//		return nil
//	}
//
//	func parseBirthday(birthday string) (*time.Time, error) {
//		if birthday == "" {
//			return nil, nil
//		}
//		parsedBirth, err := time.Parse("2006-01-02", birthday)
//		if err != nil {
//			return nil, errors.New("invalid birthday format")
//		}
//		return &parsedBirth, nil
//	}
//
// func (s studentService) CreateStudentService(request requests.StudentRequest) (*responses.MessageResponse, error) {
//
//	// Convert the student ID to uppercase
//	studentID := strings.ToUpper(request.StudentID)
//
//	// Check if the student ID or phone number is already in use
//	if checkStudentID, err := s.repositoryStudent.CheckStudentIDAlreadyHas(studentID); err != nil {
//		return nil, err
//	} else if checkStudentID {
//		return nil, errors.New("student ID already in use")
//	}
//
//	if checkPhone, err := s.repositoryStudent.CheckStudentPhoneAlreadyHas(request.Phone); err != nil {
//		return nil, err
//	} else if checkPhone {
//		return nil, errors.New("phone number already in use")
//	}
//
//	// Parse birthday
//	birth, err := parseBirthday(request.Birthday)
//	if err != nil {
//		return nil, err
//	}
//
//	// Create the student model
//	model := models.Student{
//		StudentID: getStringPointer(request.StudentID),
//		Firstname: getStringPointer(request.Firstname),
//		Lastname:  getStringPointer(request.Lastname),
//		Phone:     getStringPointer(request.Phone),
//		Email:     getStringPointer(request.Email),
//		Password:  request.Password,
//		Birthday:  birth, // Assign the time.Time object or nil
//		Gender:    getStringPointer(request.Gender),
//		//Image:     getStringPointer(request.Image),
//	}
//
//	// Call the repository method to create the student record
//	if err = s.repositoryStudent.CreateStudentRepository(&model); err != nil {
//		return nil, err
//	}
//
//	// If successful, return a success message response
//	response := &responses.MessageResponse{Message: "success"}
//	return response, nil
//
// }
//
//	func (s studentService) UpdateStudentService(request requests.StudentRequest) (*responses.MessageResponse, error) {
//		// Convert the student ID to uppercase
//		studentID := strings.ToUpper(request.StudentID)
//
//		// Check if the student ID exists
//		if checkStudentID, err := s.repositoryStudent.CheckStudentIDAlreadyHas(studentID); err != nil {
//			return nil, err
//		} else if !checkStudentID {
//			return nil, errors.New("student ID does not exist")
//		}
//
//		// Check if the phone number is already in use by another student
//		if checkPhone, err := s.repositoryStudent.CheckStudentPhoneAlreadyHas(request.Phone); err != nil {
//			return nil, err
//		} else if !checkPhone {
//			return nil, errors.New("phone number already in use")
//		}
//
//		// Parse birthday
//		birth, err := parseBirthday(request.Birthday)
//		if err != nil {
//			return nil, err
//		}
//
//		// Create the student model
//		model := models.Student{
//			StudentID: getStringPointer(request.StudentID),
//			Firstname: getStringPointer(request.Firstname),
//			Lastname:  getStringPointer(request.Lastname),
//			Phone:     getStringPointer(request.Phone),
//			Email:     getStringPointer(request.Email),
//			Password:  request.Password,
//			Birthday:  birth, // Assign the time.Time object or nil
//			Gender:    getStringPointer(request.Gender),
//			//Image:     getStringPointer(request.Image),
//		}
//
//		// Call the repository method to update the student record
//		if err = s.repositoryStudent.UpdateStudentRepository(&model); err != nil {
//			return nil, err
//		}
//
//		// If successful, return a success message response
//		response := &responses.MessageResponse{Message: "success"}
//		return response, nil
//	}

//
//func (s studentService) GetStudentByStudentIdServiceV2(request requests.StudentIdRequest) (*responses.StudentResponse, error) {
//	getStudent, err := s.repositoryStudent.GetStudentByStudentIdRepository(request.StudentID)
//	if err != nil {
//		return nil, err
//	}
//
//	response := &responses.StudentResponse{
//		ID:        getStudent.ID,
//		StudentID: getStringPointerValue(getStudent.StudentID),
//		Firstname: getStringPointerValue(getStudent.Firstname),
//		Lastname:  getStringPointerValue(getStudent.Lastname),
//		Phone:     getStringPointerValue(getStudent.Phone),
//		Email:     getStringPointerValue(getStudent.Email),
//		Gender:    getStringPointerValue(getStudent.Gender),
//		Status:    getStudent.Status,
//		//Image:     getStringPointerValue(getStudent.Image),
//		CreatedAt: getStudent.CreatedAt.Format(`02-01-2006 15:04:05`),
//		UpdatedAt: getStudent.UpdatedAt.Format(`02-01-2006 15:04:05`),
//	}
//
//	// Check if Birthday is not nil before formatting
//	if getStudent.Birthday != nil {
//		birthdayStr := getStudent.Birthday.Format(`02-01-2006`)
//		response.Birthday = &birthdayStr // Assign the address of the string
//	}
//
//	return response, nil
//}
//
//func (s studentService) GetStudentByIdService(id uint) (*responses.StudentResponse, error) {
//	getStudent, err := s.repositoryStudent.GetStudentByIdRepository(int(id))
//	if err != nil {
//		return nil, err
//	}
//	response := &responses.StudentResponse{
//		ID:        getStudent.ID,
//		StudentID: getStringPointerValue(getStudent.StudentID),
//		Firstname: getStringPointerValue(getStudent.Firstname),
//		Lastname:  getStringPointerValue(getStudent.Lastname),
//		Phone:     getStringPointerValue(getStudent.Phone),
//		Email:     getStringPointerValue(getStudent.Email),
//		Gender:    getStringPointerValue(getStudent.Gender),
//		Status:    getStudent.Status,
//		//Image:     getStringPointerValue(getStudent.Image),
//		CreatedAt: getStudent.CreatedAt.Format(`02-01-2006 15:04:05`),
//		UpdatedAt: getStudent.UpdatedAt.Format(`02-01-2006 15:04:05`),
//	}
//
//	// Check if Birthday is not nil before formatting
//	if getStudent.Birthday != nil {
//		birthdayStr := getStudent.Birthday.Format(`02-01-2006`)
//		response.Birthday = &birthdayStr // Assign the address of the string
//	}
//
//	return response, nil
//}
//
//func (s studentService) GetStudentService() ([]responses.StudentResponse, error) {
//	// fetch getStudent data from repository(database)
//	getStudent, err := s.repositoryStudent.GetStudentsRepository()
//	if err != nil {
//		return nil, err
//	}
//
//	// Check if getStudent is nil
//	if getStudent == nil {
//		return nil, errors.New("getStudent slice is nil")
//	}
//
//	// Business logic
//	// Response studentData
//	var response []responses.StudentResponse
//	for _, studentData := range getStudent {
//		// Convert birthday to a string or use nil if it's nil
//		var birthdayStr *string // Change to a pointer to string
//		if studentData.Birthday != nil {
//			bStr := studentData.Birthday.Format(`02-01-2006`)
//			birthdayStr = &bStr // Assign the address of the string
//		}
//
//		// Create pointers to strings for other fields
//		studentID := getStringPointerValue(studentData.StudentID)
//		firstname := getStringPointerValue(studentData.Firstname)
//		lastname := getStringPointerValue(studentData.Lastname)
//		phone := getStringPointerValue(studentData.Phone)
//		email := getStringPointerValue(studentData.Email)
//		gender := getStringPointerValue(studentData.Gender)
//		image := getStringPointerValue(studentData.Image)
//
//		response = append(response, responses.StudentResponse{
//			ID:        studentData.ID,
//			StudentID: studentID,
//			Firstname: firstname,
//			Lastname:  lastname,
//			Phone:     phone,
//			Email:     email,
//			Birthday:  birthdayStr, // Pass the pointer to the string
//			Gender:    gender,
//			Status:    studentData.Status,
//			Image:     image, // Pass the base64 encoded image data
//			CreatedAt: studentData.CreatedAt.Format(`02-01-2006 15:04:05`),
//			UpdatedAt: studentData.UpdatedAt.Format(`02-01-2006 15:04:05`),
//			//DeletedAt: studentData.DeletedAt.Format(`02/01/2006`),
//		})
//	}
//	return response, nil
//}

func NewStudentServices(repositoryStudent repositories.StudentRepository) StudentService {
	return &studentService{
		repositoryStudent: repositoryStudent,
	}
}
