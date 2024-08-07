package main

import (
	"errors"
	"fmt"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please provide a filename")
		return
	}
	filename := strings.ToLower(os.Args[1])
	CreateRequests(filename)
	CreateResponses(filename)
	CreateModels(filename)
	CreateRepositories(filename)
	CreateServices(filename)
	CreateControllers(filename)

}

func CreateRequests(filename string) {
	pathFolder := "requests"
	if _, err := os.Stat(pathFolder); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(pathFolder, os.ModePerm)
		if err != nil {
			fmt.Println(err)
		}
	}

	path := "requests/"
	file := path + filename + "_request" + ".go"
	var _, err = os.Stat(file)

	if os.IsNotExist(err) {

		destination, err := os.Create(file)

		if err != nil {
			fmt.Println(err)
			return
		}
		defer destination.Close()
		fmt.Fprintf(destination, "package requests")
		//fmt.Fprintf(destination, " %s\n", filename)

	} else {
		fmt.Println("File already exists!", file)
		return
	}

	fmt.Println("Created Request successfully", file)

}

func CreateResponses(filename string) {
	pathFolder := "responses"
	if _, err := os.Stat(pathFolder); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(pathFolder, os.ModePerm)
		if err != nil {
			fmt.Println(err)
		}
	}
	path := "responses/"
	file := path + filename + "_response" + ".go"
	var _, err = os.Stat(file)

	if os.IsNotExist(err) {

		destination, err := os.Create(file)

		if err != nil {
			fmt.Println(err)
			return
		}
		defer destination.Close()
		fmt.Fprintf(destination, "package responses")

	} else {
		fmt.Println("File already exists!", file)
		return
	}

	fmt.Println("Created Response successfully", file)

}

func CreateModels(filename string) {
	pathFolder := "models"
	if _, err := os.Stat(pathFolder); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(pathFolder, os.ModePerm)
		if err != nil {
			fmt.Println(err)
		}
	}

	path := "models/"
	file := path + filename + ".go"
	var _, err = os.Stat(file)

	if os.IsNotExist(err) {
		destination, err := os.Create(file)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer destination.Close()
		upperString := strings.Replace(cases.Title(language.Und, cases.NoLower).String(strings.Replace(filename, "_", " ", -1)), " ", "", -1)

		fmt.Fprintf(destination, "package models")
		fmt.Fprintf(destination, "\n\n")
		fmt.Fprintf(destination, `import "gorm.io/gorm"`)
		fmt.Fprintf(destination, "\n\n")
		fmt.Fprintf(destination, `type %s struct {`, upperString)
		fmt.Fprintf(destination, "\n")
		fmt.Fprintf(destination, `gorm.Model`)
		fmt.Fprintf(destination, "\n")
		fmt.Fprintf(destination, "}")
	} else {
		fmt.Println("File already exists!", file)
		return
	}

	fmt.Println("Created Model successfully", file)
}

func CreateRepositories(filename string) {
	pathFolder := "repositories"
	if _, err := os.Stat(pathFolder); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(pathFolder, os.ModePerm)
		if err != nil {
			fmt.Println(err)
		}
	}

	path := "repositories/"
	file := path + filename + "_repository" + ".go"
	var _, err = os.Stat(file)

	if os.IsNotExist(err) {
		destination, err := os.Create(file)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer destination.Close()
		upperString := strings.Replace(cases.Title(language.Und, cases.NoLower).String(strings.Replace(filename, "_", " ", -1)), " ", "", -1)
		lowerString := strings.ToLower(string(upperString[0])) + string(upperString[1:len(upperString)])
		//pwd, err := os.Getwd()
		//if err != nil {
		//	fmt.Println(err)
		//	os.Exit(1)
		//}
		//arr := strings.Split(pwd, "/")
		//projectName := arr[len(arr)-1]

		projectName := "go_starter"

		fmt.Fprintf(destination, "package repositories")
		fmt.Fprintf(destination, "\n\n")
		fmt.Fprintf(destination, `import (`)
		fmt.Fprintf(destination, "\n")
		fmt.Fprintf(destination, `"gorm.io/gorm"`)
		fmt.Fprintf(destination, "\n")
		fmt.Fprintf(destination, `"%s/models"`, projectName)
		fmt.Fprintf(destination, "\n")
		fmt.Fprintf(destination, ")")
		fmt.Fprintf(destination, "\n\n")
		fmt.Fprintf(destination, `type %sRepository interface{`, upperString)
		fmt.Fprintf(destination, "\n")
		fmt.Fprintf(destination, `//Insert your function interface`)
		fmt.Fprintf(destination, "\n")
		fmt.Fprintf(destination, `}`)
		fmt.Fprintf(destination, "\n\n")
		fmt.Fprintf(destination, `type %sRepository struct {db *gorm.DB}`, lowerString)
		fmt.Fprintf(destination, "\n\n")
		fmt.Fprintf(destination, `func New%sRepository(db *gorm.DB) %sRepository {`, upperString, upperString)
		fmt.Fprintf(destination, "\n")
		fmt.Fprintf(destination, `// db.Migrator().DropTable(models.%s{})`, upperString)
		fmt.Fprintf(destination, "\n")
		fmt.Fprintf(destination, `db.AutoMigrate(models.%s{})`, upperString)
		fmt.Fprintf(destination, "\n")
		fmt.Fprintf(destination, `	return &%sRepository{db: db}`, lowerString)
		fmt.Fprintf(destination, "\n")
		fmt.Fprintf(destination, `}`)
	} else {
		fmt.Println("File already exists!", file)
		return
	}

	fmt.Println("Created Repository successfully", file)
}

func CreateServices(filename string) {
	pathFolder := "services"
	if _, err := os.Stat(pathFolder); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(pathFolder, os.ModePerm)
		if err != nil {
			fmt.Println(err)
		}
	}

	path := "services/"
	file := path + filename + "_service" + ".go"
	var _, err = os.Stat(file)

	if os.IsNotExist(err) {
		destination, err := os.Create(file)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer destination.Close()
		upperString := strings.Replace(cases.Title(language.Und, cases.NoLower).String(strings.Replace(filename, "_", " ", -1)), " ", "", -1)
		lowerString := strings.ToLower(string(upperString[0])) + string(upperString[1:len(upperString)])
		//pwd, err := os.Getwd()
		//if err != nil {
		//	fmt.Println(err)
		//	os.Exit(1)
		//}
		//arr := strings.Split(pwd, "/")
		//projectName := arr[len(arr)-1]

		projectName := "go_starter"

		fmt.Fprintf(destination, "package services")
		fmt.Fprintf(destination, "\n\n")
		fmt.Fprintf(destination, `import (`)

		fmt.Fprintf(destination, "\n")
		fmt.Fprintf(destination, `"%s/repositories"`, projectName)
		fmt.Fprintf(destination, "\n")
		fmt.Fprintf(destination, `)`)

		fmt.Fprintf(destination, "\n\n")
		fmt.Fprintf(destination, `type %sService interface{`, upperString)
		fmt.Fprintf(destination, "\n")
		fmt.Fprintf(destination, `//Insert your function interface`)
		fmt.Fprintf(destination, "\n")
		fmt.Fprintf(destination, `}`)
		fmt.Fprintf(destination, "\n\n")
		fmt.Fprintf(destination, `type %sService struct {`, lowerString)
		fmt.Fprintf(destination, "\n")
		fmt.Fprintf(destination, `repository%s repositories.%sRepository`, upperString, upperString)
		fmt.Fprintf(destination, "\n")
		fmt.Fprintf(destination, `}`)

		fmt.Fprintf(destination, "\n\n")
		fmt.Fprintf(destination, `func New%sService(`, upperString)
		fmt.Fprintf(destination, "\n")
		fmt.Fprintf(destination, `repository%s repositories.%sRepository,`, upperString, upperString)
		fmt.Fprintf(destination, "\n")
		fmt.Fprintf(destination, "//repo")
		fmt.Fprintf(destination, "\n")
		fmt.Fprintf(destination, ") %sService {", upperString)
		fmt.Fprintf(destination, "\n")
		fmt.Fprintf(destination, `	return &%sService{`, lowerString)
		fmt.Fprintf(destination, "\n")
		fmt.Fprintf(destination, `repository%s :repository%s,`, upperString, upperString)
		fmt.Fprintf(destination, "\n")
		fmt.Fprintf(destination, "//repo")
		fmt.Fprintf(destination, "\n")
		fmt.Fprintf(destination, `}`)

		fmt.Fprintf(destination, "\n")
		fmt.Fprintf(destination, `}`)
	} else {
		fmt.Println("File already exists!", file)
		return
	}

	fmt.Println("Created Service successfully", file)
}

func CreateControllers(filename string) {
	pathFolder := "controllers"
	if _, err := os.Stat(pathFolder); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(pathFolder, os.ModePerm)
		if err != nil {
			fmt.Println(err)
		}
	}

	path := "controllers/"
	file := path + filename + "_controller" + ".go"
	var _, err = os.Stat(file)

	if os.IsNotExist(err) {
		destination, err := os.Create(file)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer destination.Close()
		upperString := strings.Replace(cases.Title(language.Und, cases.NoLower).String(strings.Replace(filename, "_", " ", -1)), " ", "", -1)
		lowerString := strings.ToLower(string(upperString[0])) + string(upperString[1:len(upperString)])
		//pwd, err := os.Getwd()
		//if err != nil {
		//	fmt.Println(err)
		//	os.Exit(1)
		//}
		//arr := strings.Split(pwd, "/")
		//projectName := arr[len(arr)-1]

		projectName := "go_starter"

		fmt.Fprintf(destination, "package controllers")
		fmt.Fprintf(destination, "\n\n")
		fmt.Fprintf(destination, `import (`)

		fmt.Fprintf(destination, "\n")
		fmt.Fprintf(destination, `"%s/services"`, projectName)
		fmt.Fprintf(destination, "\n")
		fmt.Fprintf(destination, `)`)

		fmt.Fprintf(destination, "\n\n")
		fmt.Fprintf(destination, `type %sController interface{`, upperString)
		fmt.Fprintf(destination, "\n")
		fmt.Fprintf(destination, `//Insert your function interface`)
		fmt.Fprintf(destination, "\n")
		fmt.Fprintf(destination, `}`)
		fmt.Fprintf(destination, "\n\n")
		fmt.Fprintf(destination, `type %sController struct {`, lowerString)
		fmt.Fprintf(destination, "\n")
		fmt.Fprintf(destination, `service%s services.%sService`, upperString, upperString)
		fmt.Fprintf(destination, "\n")
		fmt.Fprintf(destination, `}`)

		fmt.Fprintf(destination, "\n\n")
		fmt.Fprintf(destination, `func New%sController(`, upperString)
		fmt.Fprintf(destination, "\n")
		fmt.Fprintf(destination, `service%s services.%sService,`, upperString, upperString)
		fmt.Fprintf(destination, "\n")
		fmt.Fprintf(destination, "//services")
		fmt.Fprintf(destination, "\n")
		fmt.Fprintf(destination, ") %sController {", upperString)
		fmt.Fprintf(destination, "\n")
		fmt.Fprintf(destination, `	return &%sController{`, lowerString)
		fmt.Fprintf(destination, "\n")
		fmt.Fprintf(destination, `service%s :service%s,`, upperString, upperString)
		fmt.Fprintf(destination, "\n")
		fmt.Fprintf(destination, "//services")
		fmt.Fprintf(destination, "\n")
		fmt.Fprintf(destination, `}`)

		fmt.Fprintf(destination, "\n")
		fmt.Fprintf(destination, `}`)
	} else {
		fmt.Println("File already exists!", file)
		return
	}

	fmt.Println("Created Controller successfully", file)
}
