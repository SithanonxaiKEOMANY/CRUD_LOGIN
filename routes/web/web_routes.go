package web

import (
	"go_starter/controllers"
	"go_starter/controllers/web"
	"go_starter/routes"

	"github.com/gofiber/fiber/v2"
)

type webRoutes struct {
	controller        web.Controller
	studentController controllers.StudentController
	userController    controllers.UserController
}

func (w webRoutes) Install(app *fiber.App) {
	route := app.Group("web/", func(ctx *fiber.Ctx) error {
		return ctx.Next()
	})
	route.Post("hello", w.controller.StartController)
	route.Get("students", w.studentController.GetStudentController)
	route.Get("student/:id", w.studentController.GetStudentByIDController)
	route.Get("student", w.studentController.GetStudentByStudentIDControllerV2)
	route.Post("create-student", w.studentController.CreateStudentController)
	route.Put("update-student", w.studentController.UpdateStudentController)
	route.Delete("delete-student", w.studentController.DeleteStudentByIDController)

	//image
	route.Post("update-image", w.studentController.UploadStudentImageController)

	route.Post("signup", w.studentController.SignUpController)
	route.Post("signin", w.studentController.SignInController)
	route.Post("student-classroom", w.studentController.GetStudentClassroomByClassroomIDController)

	// User LogIn and User CRUD

	//LogIn
	route.Post("login", w.userController.LoginController)
	//route.Post("sign-up", w.userController.SignUpUserController)
	route.Post("sign-in", w.userController.SignInUserController)

	//CRUD
	route.Post("get-all-user", w.userController.GetAllUserController)
	route.Post("get-by-id/:id", w.userController.GetUserByIdController)
	//route.Post("create-user", w.userController.CreateUserController)
	route.Post("update-user", w.userController.UpdateUserController)
	route.Post("delete-user", w.userController.DeleteUserController)

}

func NewWebRoutes(
	controller web.Controller,
	studentController controllers.StudentController,
	userController controllers.UserController,
	// controller
) routes.Routes {
	return &webRoutes{
		controller:        controller,
		studentController: studentController,
		userController:    userController,
		//controller
	}
}
