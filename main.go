package main

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"go_starter/config"
	"go_starter/controllers"
	"go_starter/controllers/web"
	web2 "go_starter/routes/web"

	//"go_starter/controllers/web"
	"go_starter/database"
	"go_starter/logs"
	"go_starter/partners"
	"go_starter/repositories"
	//web2 "go_starter/routes/web"
	"go_starter/services"
	"go_starter/trails"
	"log"
	"net/http"
)

func main() {

	//connect database
	//postgres
	postgresConnection, err := database.PostgresConnection()
	if err != nil {
		logs.Error(err)
		return
	}

	////mysql
	//mySQLConnection, err := database.MysqlConnection()
	//if err != nil {
	//	logs.Error(err)
	//	return
	//}

	//call api client interface
	httpClient := http.Client{}
	newHttpClientTrail := trails.NewHttpClientTrail(httpClient)
	partners.NewPartner(newHttpClientTrail)

	//basic structure
	newRepository := repositories.NewRepository(postgresConnection)
	newService := services.NewService(newRepository)
	//newControllerApi := api.NewControllerApi(newService)

	//student
	studentRepository := repositories.NewStudentRepository(postgresConnection)
	studentService := services.NewStudentServices(studentRepository)
	studentController := controllers.NewCustomerController(studentService)

	// User
	userRepository := repositories.NewUserRepository(postgresConnection)
	userService := services.NewUserService(userRepository)
	userController := controllers.NewUserController(userService)

	//connect route
	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
		BodyLimit:   16 * 1024 * 1024,
	})
	app.Use(logger.New())
	app.Use(cors.New())

	// Serve static files from the "assets/ceit/2024/images" directory
	app.Static("/ceit/2024/images", "./assets/ceit/2024/images")

	//Web routes
	newController := web.NewController(newService)
	newWebRoute := web2.NewWebRoutes(
		newController,
		studentController,
		userController,
		//new web controller
	)
	newWebRoute.Install(app)

	//Api routes
	// newControllerApi = api.NewControllerApi(newService)
	// newApiRoute := api2.NewApiRoutes(
	// 	newControllerApi,
	// 	studentController,
	// 	userController,
	// )
	// newApiRoute.Install(app)

	log.Fatal(app.Listen(fmt.Sprintf(":%s", config.Env("app.port"))))
}
