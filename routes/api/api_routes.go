package api

import (
	"go_starter/controllers"
	"go_starter/controllers/api"
	"go_starter/routes"

	"github.com/gofiber/fiber/v2"
)

type apiRoutes struct {
	controllerApi     api.ControllerApi
	studentController controllers.StudentController
	userController    controllers.UserController
}

func (a apiRoutes) Install(app *fiber.App) {
	route := app.Group("api/", func(ctx *fiber.Ctx) error {
		return ctx.Next()
	})
	route.Post("hello", a.controllerApi.StartController)
	//route.Post("customer", a.studentController.CreateCustomer)
	//route.Get("customer/:id", a.studentController.GetCustomer)
	route.Get("student", a.studentController.GetStudentController)
	//route.Put("/customer/:id", a.studentController.UpdateCustomer)
	//route.Delete("/customer/:id", a.studentController.DeleteCustomer)
	route.Post("get-all-user", a.userController.GetAllUserController)
}

func NewApiRoutes(controllerApi api.ControllerApi, customerApi controllers.StudentController, userApi controllers.UserController) routes.Routes {
	return &apiRoutes{
		controllerApi:     controllerApi,
		studentController: customerApi,
		userController:    userApi,
		//controller
	}
}
