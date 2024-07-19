package services

import (
	"fmt"
	"go_starter/errs"
	"go_starter/models"
	"go_starter/repositories"
	"go_starter/requests"
	"go_starter/responses"
	"go_starter/security"
	"strings"

	"github.com/pkg/errors"
)

type UserService interface {

	// Login
	SignUpUserService(request requests.SignUpUserRequest) (*responses.SignUpUserResponse, error)
	SignInUserService(request requests.SignInUserRequest) (*responses.SignInUserResponse, error)

	//CRUD
	GetAllUserService() ([]responses.UserResponse, error)
	GetByIdUserService(id uint) (*responses.UserResponse, error)
	GetByPhoneService(phone string) (*responses.UserResponse, error)
	CreateUserService(request requests.CreateUserRequest) (*responses.MessageUserResponse, error)
	UpdateUserService(request requests.UpdateUserRequest) (*responses.MessageUserResponse, error)
	DeleteUserService(request requests.DeleteUserRequest) (*responses.MessageUserResponse, error)
}

type userService struct {
	repositoryUserRepository repositories.UserRepository
}

// CreateUserService implements UserService.
func (u *userService) CreateUserService(request requests.CreateUserRequest) (*responses.MessageUserResponse, error) {

	email := strings.ToUpper(request.Email)

	if checkEmail, err := u.repositoryUserRepository.CheckEmailAlreadyHas(email); err != nil {

		return nil, err
	} else if checkEmail {
		return nil, errors.New("Email already in User")
	}

	model := models.User{

		Name:  request.Name,
		Email: request.Email,
		Phone: request.Phone,
	}

	if err := u.repositoryUserRepository.CreateUserRepository(&model); err != nil {

		return nil, err
	}
	response := &responses.MessageUserResponse{Message: "Success"}
	return response, nil
}

// DeleteUserService implements UserService.
func (u *userService) DeleteUserService(request requests.DeleteUserRequest) (*responses.MessageUserResponse, error) {

	if request.ID == 0 {
		return nil, errors.New("ID can't be empty")
	}
	err := u.repositoryUserRepository.DeleteUserRepository(request.ID)
	if err != nil {
		return nil, err
	}
	response := &responses.MessageUserResponse{Message: "Success"}

	return response, nil
}

// GetAllUserService implements UserService.
func (u *userService) GetAllUserService() ([]responses.UserResponse, error) {

	getAllUser, err := u.repositoryUserRepository.GetAllUserRepository()

	if err != nil {
		return nil, err
	}

	if getAllUser == nil {
		return nil, errors.New("getAllUser slice is nil")
	}

	var response []responses.UserResponse
	for _, data := range getAllUser {
		userResponse := responses.UserResponse{

			ID:        data.ID,
			Name:      data.Name,
			Email:     data.Email,
			Phone:     data.Phone,
			CreatedAt: data.CreatedAt.Format("02-01-2006 15:01:05"),
			UpdatedAt: data.UpdatedAt.Format("02-01-2006 15:01:05"),
		}
		response = append(response, userResponse)
	}
	return response, err
}

// GetByIdUserService implements UserService.
func (u *userService) GetByIdUserService(id uint) (*responses.UserResponse, error) {

	data, err := u.repositoryUserRepository.GetByIdUserRepository(uint(id))

	if err != nil {
		return nil, err
	}

	response := &responses.UserResponse{

		ID:        data.ID,
		Name:      data.Name,
		Email:     data.Email,
		Phone:     data.Phone,
		CreatedAt: data.CreatedAt.Format("02-01-2006 15:01:05"),
		UpdatedAt: data.UpdatedAt.Format("02-01-2006 15:01:05"),
	}
	return response, nil
}

// GetByUserNameService implements UserService.
func (u *userService) GetByPhoneService(phone string) (*responses.UserResponse, error) {

	data, err := u.repositoryUserRepository.GetByPhoneRepository(phone)

	if err != nil {
		return nil, err
	}

	response := &responses.UserResponse{

		ID:        data.ID,
		Name:      data.Name,
		Email:     data.Email,
		Phone:     data.Phone,
		CreatedAt: data.CreatedAt.Format("02-01-2006 15:01:05"),
		UpdatedAt: data.UpdatedAt.Format("02-01-2006 15:01:05"),
	}
	return response, nil
}

// SignUpUserService implements UserService.
func (u *userService) SignUpUserService(request requests.SignUpUserRequest) (*responses.SignUpUserResponse, error) {

	if request.Email == "" {
		return nil, errs.ErrorBadRequest("Email Cant Be Empty")
	}
	if checkUserName, err := u.repositoryUserRepository.CheckEmailAlreadyHas(request.Email); err != nil {
		return nil, err
		
	} else if checkUserName {
		return nil, errors.New("UserName already in User")
	}
	trimSpaceUser := strings.TrimSpace(request.Password)
	if trimSpaceUser == "" {
		return nil, errs.ErrorBadRequest("Password Cant Be Empty")
	}
	encryptPassword, err := security.EncryptPassword(request.Password)
	if err != nil {
		return nil, err
	}
	newAccessToken, err := security.NewAccessToken(request.Email)

	if err != nil {
		return nil, err
	}

	data := models.User{
		Email:    request.Email,
		Password: encryptPassword,
		Name:        request.Name,
		Phone:       request.Phone,
		Token:    newAccessToken,
	}
	signUpUser, err := u.repositoryUserRepository.SignUpUserRepository(data)

	if err != nil {
		return nil, err
	}
	response := responses.SignUpUserResponse{
		Email:       signUpUser.Email,
		Name:        request.Name,
		Phone:       request.Phone,
		AccessToken: signUpUser.Token,
	}

	return &response, nil
}

// SignInUserService implements UserService.
func (u *userService) SignInUserService(request requests.SignInUserRequest) (*responses.SignInUserResponse, error) {

	if request.Email == "" {
		return nil, errs.ErrorBadRequest("Email Cant Be Empty")
	}

	trimSpaceUser := strings.TrimSpace(request.Password)
	if trimSpaceUser == "" {
		return nil, errs.ErrorBadRequest("Password Cant Be Empty")
	}

	getUserData, err := u.repositoryUserRepository.GetByEmailRepository(request.Email)

	if err != nil {
		return nil, err
	}
	err = security.VerifyPassword(getUserData.Password, request.Password)

	if err != nil {
		return nil, fmt.Errorf("password does not match")
	}
	response := responses.SignInUserResponse{
		Email:       getUserData.Email,
		AccessToken: "",
	}
	return &response, err
}

// UpdateUserService implements UserService.
func (u *userService) UpdateUserService(request requests.UpdateUserRequest) (*responses.MessageUserResponse, error) {

	data := models.User{
		ID:    request.ID,
		Name:  request.Name,
		Email: request.Email,
	}
	if err := u.repositoryUserRepository.UpdateUserRepository(&data); err != nil {
		return nil, err
	}
	response := &responses.MessageUserResponse{Message: "Success"}

	return response, nil
}

func NewUserService(repositoryUserRepository repositories.UserRepository) UserService {
	return &userService{
		repositoryUserRepository: repositoryUserRepository,
	}
}
