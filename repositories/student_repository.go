package repositories

import (
	"github.com/pkg/errors"
	"go_starter/logs"
	"go_starter/models"
	"gorm.io/gorm"
	"strings"
)

type StudentRepository interface {
	GetStudentClassroomByClassroomIDRepository(classroomID uint) ([]models.StudentClassroom, error)

	//
	GetTeacherByPhoneRepository(phone string) (*models.Teacher, error)
	GetStudentByPhoneRepository(phone string) (*models.Student, error)

	//
	SignUpForTeacherRepository(request models.Teacher) (*models.Teacher, error)
	SignUpForStudentRepository(request models.Student) (*models.Student, error)

	//
	GetStudentsRepository() ([]models.Student, error)
	GetStudentByIdRepository(id int) (*models.Student, error)
	GetStudentByStudentIdRepository(studentID string) (*models.Student, error)
	CreateStudentRepository(request *models.Student) error
	UpdateStudentRepository(request *models.Student) error
	DeleteStudentByStudentIDRepository(studentID string) error

	//
	CheckTeacherPhoneAlreadyHas(phone string) (bool, error)
	CheckStudentPhoneAlreadyHas(phone string) (bool, error)
	CheckStudentIDAlreadyHas(studentID string) (bool, error)

	//image
	GetStudentImageRepository(studentID string) (string, error)
	UpdateStudentImageRepository(request *models.Student) error
	DeleteStudentImageRepository(studentID string) error
}

type studentRepository struct{ db *gorm.DB }

func (s studentRepository) GetStudentClassroomByClassroomIDRepository(classroomID uint) ([]models.StudentClassroom, error) {
	//var model []models.StudentClassroom
	//err := s.db.Where("id", classroomID).Find(&model).Error
	//if err != nil {
	//	return nil, err
	//}
	//return model, err

	var studentClassrooms []models.StudentClassroom

	// Join query to fetch StudentClassroom with related Student and Classroom
	err := s.db.Preload("Student").Preload("Classroom").
		Where("classroom_id = ?", classroomID).
		Find(&studentClassrooms).Error

	if err != nil {
		return nil, err
	}
	return studentClassrooms, nil
}

//-----------------------------------------new---------------------------------------------------//

func (s studentRepository) GetTeacherByPhoneRepository(phone string) (*models.Teacher, error) {
	var model models.Teacher
	query := s.db.First(&model, "phone = ?", phone)
	if query.Error != nil {
		return nil, nil
	}
	return &model, nil
}

func (s studentRepository) GetStudentByPhoneRepository(phone string) (*models.Student, error) {
	var model models.Student
	query := s.db.First(&model, "phone = ?", phone)
	if query.Error != nil {
		return nil, nil
	}
	return &model, nil
}

func (s studentRepository) SignUpForTeacherRepository(request models.Teacher) (*models.Teacher, error) {
	request = models.Teacher{
		ID:       request.ID,
		Phone:    request.Phone,
		Password: request.Password,
		Token:    request.Token,
	}
	create := s.db.Create(&request)
	if create.Error != nil {
		logs.Error(create.Error)
		return nil, create.Error
	}
	return &request, nil
}

func (s studentRepository) SignUpForStudentRepository(request models.Student) (*models.Student, error) {
	//select specific that you want to insert
	request = models.Student{
		ID:       request.ID,
		Phone:    request.Phone,
		Password: request.Password,
		Token:    request.Token,
	}
	create := s.db.Create(&request)
	if create.Error != nil {
		logs.Error(create.Error)
		return nil, create.Error
	}
	return &request, nil
}

//---------------------------------------------------------------------------------------//

func (s studentRepository) GetStudentImageRepository(studentID string) (string, error) {
	var image string
	query := s.db.Raw("SELECT image FROM students WHERE student_id = ?", studentID).Scan(&image)
	if query.Error != nil {
		return "", query.Error
	}
	return image, nil
}

func (s studentRepository) UpdateStudentImageRepository(request *models.Student) error {
	query := s.db.Model(&models.Student{}).Where("student_id = ?", request.StudentID).Update("image", request.Image)
	if query.Error != nil {
		return nil
	}
	return nil
}

func (s studentRepository) DeleteStudentImageRepository(studentID string) error {
	query := s.db.Model(&models.Student{}).Where("student_id = ?", studentID).Update("image", nil)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func (s studentRepository) GetStudentByIdRepository(id int) (*models.Student, error) {
	var model models.Student

	// Execute raw SQL query
	query := s.db.Raw("SELECT * FROM students WHERE id = ?", id).Scan(&model).Error
	if query != nil {
		return nil, query
	}
	return &model, nil
}

func (s studentRepository) CheckStudentIDAlreadyHas(studentID string) (bool, error) {
	var count int64
	// Convert studentID to uppercase for comparison
	upperStudentID := strings.ToUpper(studentID)
	// Perform a case-insensitive comparison
	query := s.db.Model(&models.Student{}).Where("UPPER(student_id) = ?", upperStudentID).Count(&count)
	if query.Error != nil {
		return false, query.Error
	}
	return count > 0, nil // Return true if count is greater than 0, indicating student ID exists
}

func (s studentRepository) CheckTeacherPhoneAlreadyHas(phone string) (bool, error) {
	var count int64
	query := s.db.Model(&models.Teacher{}).Where("phone = ?", phone).Count(&count)
	if query.Error != nil {
		return false, query.Error
	}
	return count > 0, nil
}

func (s studentRepository) CheckStudentPhoneAlreadyHas(phone string) (bool, error) {
	var count int64
	query := s.db.Model(&models.Student{}).Where("phone = ?", phone).Count(&count)
	if query.Error != nil {
		return false, query.Error
	}
	return count > 0, nil
}

func (s studentRepository) GetStudentsRepository() ([]models.Student, error) {
	var model []models.Student
	query := s.db.Find(&model).Error
	if query != nil {
		return nil, query
	}
	return model, nil
}

func (s studentRepository) GetStudentByStudentIdRepository(studentID string) (*models.Student, error) {
	var model models.Student

	// Execute raw SQL query
	query := s.db.Raw("SELECT * FROM students WHERE student_id = ?", studentID).Scan(&model).Error
	if query != nil {
		return nil, query
	}
	return &model, nil
}

func (s studentRepository) CreateStudentRepository(model *models.Student) error {
	if err := s.db.Create(model).Error; err != nil {
		return err
	}
	return nil
}
func (s studentRepository) UpdateStudentRepository(request *models.Student) error {
	// raw function no check data
	//query := s.db.Model(&models.Student{}).Where("student_id =?", request.StudentID).Updates(request)
	//if query.Error != nil {
	//	return query.Error
	//}
	//return nil

	// add check student_id on database
	query := s.db.Model(&models.Student{}).Where("student_id = ?", request.StudentID).Updates(request)
	if query.Error != nil {
		return query.Error
	}
	if query.RowsAffected == 0 {
		return errors.New("no student_id found")
	}
	return nil
}

func (s studentRepository) DeleteStudentByStudentIDRepository(studentID string) error {
	// raw function no check data
	//	query := models.Student{StudentID: studentID}
	//	if err := s.db.Where("student_id = ?", studentID).Delete(&query).Error; err != nil {
	//		return err
	//	}
	//	return nil

	// add check student_id on database
	var count int64
	if err := s.db.Model(&models.Student{}).Where("student_id = ?", studentID).Count(&count).Error; err != nil {
		return err
	}
	if count == 0 {
		return errors.New("student not found")
	}

	// Delete the student
	if err := s.db.Where("student_id = ?", studentID).Delete(&models.Student{}).Error; err != nil {
		return err
	}
	return nil
}

func NewStudentRepository(db *gorm.DB) StudentRepository {
	//db.Migrator().DropTable(models.Student{})
	//db.AutoMigrate(models.Classroom{})
	//db.AutoMigrate(models.StudentClassroom{})
	return &studentRepository{db: db}
}
