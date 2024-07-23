package controllers

import (
	"go_starter/logs"
	"go_starter/requests"
	"go_starter/services"
	"go_starter/validation"

	"github.com/gofiber/fiber/v2"
)

type UserController interface {
	SignInUserController(ctx *fiber.Ctx) error
	//SignUpUserController(ctx *fiber.Ctx) error
	LoginController(ctx *fiber.Ctx) error

	GetAllUserController(ctx *fiber.Ctx) error
	GetUserByIdController(ctx *fiber.Ctx) error
	GetUserByPhoneController(ctx *fiber.Ctx) error
	UpdateUserController(ctx *fiber.Ctx) error
	DeleteUserController(ctx *fiber.Ctx) error
}

type userController struct {
	serviceUser services.UserService
}

// LoginController implements UserController.
func (u *userController) LoginController(ctx *fiber.Ctx) error {
	request := new(requests.LoginRequest)

	// Parse request body
	if err := ctx.BodyParser(request); err != nil {
		return NewErrorResponses(ctx, err)
	}

	// Validate request data
	errValidate := validation.Validate(request)
	if errValidate != nil {
		return NewErrorValidate(ctx, errValidate[0].Error)
	}

	// Call the login service
	response, token, err := u.serviceUser.LoginService(*request)
	if err != nil {
		return NewErrorResponses(ctx, err)
	}

	// Call NewSuccessResponseSignIn with the correct number of arguments
	return NewSuccessResponseSignIn(ctx, response, token)
}

// CreateUserController implements UserController.

// DeleteUserController implements UserController.
func (u *userController) DeleteUserController(ctx *fiber.Ctx) error {

	request := new(requests.DeleteUserRequest)

	if err := ctx.BodyParser(request); err != nil {
		return NewErrorResponses(ctx, err)
	}
	errValidate := validation.Validate(request)
	if errValidate != nil {

		return NewErrorValidate(ctx, errValidate[0].Error)
	}
	response, err := u.serviceUser.DeleteUserService(*request)

	if err != nil {
		return NewErrorResponses(ctx, err)
	}

	return NewSuccessMessage(ctx, response.Message)

}

// GetAllUserController implements UserController.
func (u *userController) GetAllUserController(ctx *fiber.Ctx) error {

	//fetch User data from service folder
	data, err := u.serviceUser.GetAllUserService()
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Failed to retrieve customer data",
			"error":   err.Error(),
		})
	}

	//return http response
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    data,
	})
}

// GetUserByIdController implements UserController.
func (u *userController) GetUserByIdController(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Failed to retrieve customer data",
			"error":   err.Error(),
		})
	}
	response, err := u.serviceUser.GetByIdUserService(uint(id))
	if err != nil {
		return NewErrorResponses(ctx, err)
	}
	return NewSuccessResponse(ctx, response)
}

// GetUserByUserNameControllerV2 implements UserController.
func (u *userController) GetUserByPhoneController(ctx *fiber.Ctx) error {

	panic("unimplemented")
}

// SignInUserController implements UserController.
func (u *userController) SignInUserController(ctx *fiber.Ctx) error {

	request := new(requests.SignInUserRequest)
	if err := ctx.BodyParser(request); err != nil {
		logs.Error(err)
		return NewErrorResponses(ctx, err)
	}
	errValidate := validation.Validate(request)
	if errValidate != nil {
		return NewErrorValidate(ctx, errValidate[0].Error)
	}
	response, err := u.serviceUser.SignInUserService(*request)

	if err != nil {
		return NewErrorResponses(ctx, err)
	}
	return NewSuccessResponse(ctx, response)
}

// SignUpUserController implements UserController.
// func (u *userController) SignUpUserController(ctx *fiber.Ctx) error {

// 	request := new(requests.SignUpUserRequest)
// 	if err := ctx.BodyParser(request); err != nil {
// 		logs.Error(err)
// 		return NewErrorResponses(ctx, err)
// 	}
// 	errValidate := validation.Validate(request)
// 	if errValidate != nil {
// 		return NewErrorValidate(ctx, errValidate[0].Error)
// 	}
// 	response, err := u.serviceUser.SignUpUserService(*request)

// 	if err != nil {
// 		return NewErrorResponses(ctx, err)
// 	}
// 	return NewSuccessResponse(ctx, response)
// }

// UpdateUserController implements UserController.
func (u *userController) UpdateUserController(ctx *fiber.Ctx) error {

	request := new(requests.UpdateUserRequest)
	if err := ctx.BodyParser(request); err != nil {
		return NewErrorResponses(ctx, err)
	}
	errValidate := validation.Validate(request)
	if errValidate != nil {
		return NewErrorValidate(ctx, errValidate[0].Error)
	}
	response, err := u.serviceUser.UpdateUserService(*request)
	if err != nil {
		return NewErrorResponses(ctx, err)
	}
	return NewSuccessMsg(ctx, response.Message)
}

func NewUserController(serviceUser services.UserService) UserController {
	return &userController{serviceUser: serviceUser}
}
