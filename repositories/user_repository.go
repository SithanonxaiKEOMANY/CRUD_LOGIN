package repositories

import (
	"go_starter/logs"
	"go_starter/models"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type UserRepository interface {
	//Login
	SignUpUserRepository(request models.User) (*models.User, error)

	//CRUD
	CreateUserRepository(request *models.User) error
	GetAllUserRepository() ([]models.User, error)
	GetByIdUserRepository(id uint) (*models.User, error)
	GetByPhoneRepository(phone string) (*models.User, error)
	GetByEmailRepository(email string) (*models.User, error)
	UpdateUserRepository(request *models.User) error
	DeleteUserRepository(id uint) error

	//Check UserName and Check Phone
	CheckEmailAlreadyHas(email string) (bool, error)
	CheckPhoneAlreadyHas(phone string) (bool, error)
}

type userRepository struct{ db *gorm.DB }

// CheckPhoneAlreadyHas implements UserRepository.
func (u *userRepository) CheckPhoneAlreadyHas(phone string) (bool, error) {

	var model models.User
	result := u.db.Where("phone = ?", phone).First(&model)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return false, result.Error
	}
	return result.RowsAffected > 0, nil
}

// CheckUserNameAlreadyHas implements UserRepository.
func (u *userRepository) CheckEmailAlreadyHas(email string) (bool, error) {

	var model models.User
	result := u.db.Where("email = ?", email).First(&model)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return false, result.Error
	}
	return result.RowsAffected > 0, nil
}

// SignUpUserRepository implements UserRepository.
func (u *userRepository) SignUpUserRepository(request models.User) (*models.User, error) {

	create := u.db.Create(&request)
	if create.Error != nil {
		logs.Error(create.Error)
		return nil, create.Error
	}
	return &request, nil
}

// CreateUserRepository implements UserRepository.
func (u *userRepository) CreateUserRepository(request *models.User) error {

	if err := u.db.Create(request).Error; err != nil {
		return err
	}
	return nil
}

// DeleteUserRepository implements UserRepository.
func (u *userRepository) DeleteUserRepository(id uint) error {

	query := models.User{ID: id}

	if err := u.db.Where("id = ?", id).Delete(&query).Error; err != nil {
		return err
	}

	return nil
}

// GetAllUserRepository implements UserRepository.
func (u *userRepository) GetAllUserRepository() ([]models.User, error) {

	var model []models.User
	query := u.db.Find(&model).Error
	if query != nil {
		return nil, query
	}
	return model, nil
}

// GetByIdUserRepository implements UserRepository.
func (u *userRepository) GetByIdUserRepository(id uint) (*models.User, error) {

	var model models.User
	if err := u.db.Where("id =?", id).First(&model).Error; err != nil {
		return nil, err
	}
	return &model, nil
}

// GetByUserNameRepository implements UserRepository.
func (u *userRepository) GetByPhoneRepository(phone string) (*models.User, error) {

	var model models.User
	query := u.db.First(&model, "phone =?", phone)

	if query.Error != nil {
		return nil, nil
	}
	return &model, nil
}

// GetByEmailRepository implements UserRepository.
func (u *userRepository) GetByEmailRepository(email string) (*models.User, error) {

	var model models.User
	query := u.db.First(&model, "email =?", email)

	if query.Error != nil {
		return nil, nil
	}
	return &model, nil
}

// UpdateUserRepository implements UserRepository.
func (u *userRepository) UpdateUserRepository(request *models.User) error {

	query := u.db.Model(&models.User{}).Where("id = ?", request.ID).Updates(request)

	if query.Error != nil {
		return query.Error
	}
	if query.RowsAffected == 0 {
		return errors.New("not id found")
	}
	return nil
}

func NewUserRepository(db *gorm.DB) UserRepository {
	// db.Migrator().DropTable(&models.User{})
	//db.AutoMigrate(&models.User{})
	return &userRepository{db: db}
}
