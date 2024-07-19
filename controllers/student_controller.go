package controllers

import (
	"github.com/gofiber/fiber/v2"
	"go_starter/logs"
	"go_starter/requests"
	"go_starter/services"
	"go_starter/trails"
	"go_starter/validation"
)

type StudentController interface {
	GetStudentClassroomByClassroomIDController(ctx *fiber.Ctx) error

	SignInController(ctx *fiber.Ctx) error
	SignUpController(ctx *fiber.Ctx) error

	GetStudentController(ctx *fiber.Ctx) error
	GetStudentByIDController(ctx *fiber.Ctx) error
	GetStudentByStudentIDControllerV2(ctx *fiber.Ctx) error
	CreateStudentController(ctx *fiber.Ctx) error
	UpdateStudentController(ctx *fiber.Ctx) error
	DeleteStudentByIDController(ctx *fiber.Ctx) error

	//
	UploadStudentImageController(ctx *fiber.Ctx) error
}
type studentController struct {
	serviceStudent services.StudentService
}

func (c *studentController) GetStudentClassroomByClassroomIDController(ctx *fiber.Ctx) error {
	req := new(requests.ClassroomIDRequest)
	if err := ctx.BodyParser(req); err != nil {
		logs.Error(err)
		return NewErrorResponses(ctx, err)
	}
	errValidate := validation.Validate(req)
	if errValidate != nil {
		return NewErrorValidate(ctx, errValidate[0].Error)
	}
	response, err := c.serviceStudent.GetStudentClassroomByClassroomIDService(*req)
	if err != nil {
		return NewErrorResponses(ctx, err)
	}
	return NewSuccessResponse(ctx, response)
}

func (c *studentController) SignInController(ctx *fiber.Ctx) error {
	req := new(requests.SignInRequest)
	if err := ctx.BodyParser(req); err != nil {
		logs.Error(err)
		return NewErrorResponses(ctx, err)
	}
	errValidate := validation.Validate(req)
	if errValidate != nil {
		return NewErrorValidate(ctx, errValidate[0].Error)
	}
	response, err := c.serviceStudent.SignInService(*req)
	if err != nil {
		return NewErrorResponses(ctx, err)
	}
	return NewSuccessResponse(ctx, response)
}

func (c *studentController) SignUpController(ctx *fiber.Ctx) error {
	req := new(requests.SigUpRequest)
	if err := ctx.BodyParser(req); err != nil {
		logs.Error(err)
		return NewErrorResponses(ctx, err)
	}
	errValidate := validation.Validate(req)
	if errValidate != nil {
		return NewErrorValidate(ctx, errValidate[0].Error)
	}
	response, err := c.serviceStudent.SignUpService(*req)
	if err != nil {
		return NewErrorResponses(ctx, err)
	}
	return NewSuccessResponse(ctx, response)
}

func (c *studentController) UploadStudentImageController(ctx *fiber.Ctx) error {
	// Multipart form data handle
	imageData, err := trails.HandleMultipartFormData(ctx)
	if err != nil {
		return err
	}

	// Create a student image request instance
	request := requests.StudentImageRequest{
		StudentID: ctx.FormValue("student_id"),
		Image:     imageData,
	}
	//fmt.Printf("%v\n", request)
	errValidate := validation.Validate(request)
	if errValidate != nil {
		return NewErrorValidate(ctx, errValidate[0].Error)
	}
	// Call the service
	response, err := c.serviceStudent.UploadStudentImageService(request)
	if err != nil {
		return NewErrorResponses(ctx, err)
	}
	return NewSuccessMsg(ctx, response.Message)
}

func (c *studentController) CreateStudentController(ctx *fiber.Ctx) error {
	request := new(requests.StudentRequest)
	if err := ctx.BodyParser(request); err != nil {
		return NewErrorResponses(ctx, err)
	}
	errValidate := validation.Validate(request)
	if errValidate != nil {
		return NewErrorValidate(ctx, errValidate[0].Error)
	}
	response, err := c.serviceStudent.CreateStudentService(*request)
	if err != nil {
		return NewErrorResponses(ctx, err)
	}
	return NewSuccessMsg(ctx, response.Message)
}

func (c *studentController) UpdateStudentController(ctx *fiber.Ctx) error {
	request := new(requests.StudentRequest)
	if err := ctx.BodyParser(request); err != nil {
		return NewErrorResponses(ctx, err)
	}
	errValidate := validation.Validate(request)
	if errValidate != nil {
		return NewErrorValidate(ctx, errValidate[0].Error)
	}
	response, err := c.serviceStudent.UpdateStudentService(*request)
	if err != nil {
		return NewErrorResponses(ctx, err)
	}
	return NewSuccessMsg(ctx, response.Message)
}

func (c *studentController) DeleteStudentByIDController(ctx *fiber.Ctx) error {
	request := new(requests.StudentIdRequest)
	if err := ctx.BodyParser(request); err != nil {
		return NewErrorResponses(ctx, err)
	}
	errValidate := validation.Validate(request)
	if errValidate != nil {
		return NewErrorValidate(ctx, errValidate[0].Error)
	}
	response, err := c.serviceStudent.DeleteStudentByIDService(*request)
	if err != nil {
		return NewErrorResponses(ctx, err)
	}
	return NewSuccessMsg(ctx, response.Message)
}

func (c *studentController) GetStudentByStudentIDControllerV2(ctx *fiber.Ctx) error {
	req := new(requests.StudentIdRequest)
	if err := ctx.BodyParser(req); err != nil {
		logs.Error(err)
		return NewErrorResponses(ctx, err)
	}
	errValidate := validation.Validate(req)
	if errValidate != nil {
		return NewErrorValidate(ctx, errValidate[0].Error)
	}
	response, err := c.serviceStudent.GetStudentByStudentIdServiceV2(*req)
	if err != nil {
		return NewErrorResponses(ctx, err)
	}
	return NewSuccessResponse(ctx, response)
}

func (c *studentController) GetStudentByIDController(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id", 1)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Failed to retrieve customer data",
			"error":   err.Error(),
		})
	}
	response, err := c.serviceStudent.GetStudentByIdService(uint(id))
	if err != nil {
		return NewErrorResponses(ctx, err)
	}
	return NewSuccessResponse(ctx, response)
}

func (c *studentController) GetStudentController(ctx *fiber.Ctx) error {

	//fetch customer data from service folder
	customers, err := c.serviceStudent.GetStudentService()
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
		"data":    customers,
	})
}

func NewCustomerController(serviceService services.StudentService) StudentController {
	return &studentController{serviceStudent: serviceService}
}
