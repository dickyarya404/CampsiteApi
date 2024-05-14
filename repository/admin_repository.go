package repository

// import (
// 	"campsite_reservation/middleware"
// 	"campsite_reservation/model"
// 	"errors"

// 	"golang.org/x/crypto/bcrypt"
// 	"gorm.io/gorm"
// )

// type AdminRepository interface {
// 	Login(email, password string) (model.User, string, error)
// 	Create(data model.Admin) error
// 	SelectAll() ([]model.Admin, error)
// }

// type adminRepository struct {
// 	db *gorm.DB
// }

// func NewAdminRepository(db *gorm.DB) *adminRepository {
// 	return &adminRepository{
// 		db: db,
// 	}
// }

// func (ur *adminRepository) Create(data model.Admin) error {
// 	return ur.db.Create(&data).Error
// }

// func (ur *adminRepository) Login(email, password string) (model.Admin, string, error) {
// 	var data model.Admin

// 	conf := ur.db.Where("email = ?", email).First(&data)
// 	if conf.Error != nil {
// 		return model.Admin{}, "", conf.Error

// 	}

// 	if conf.RowsAffected == 0 {
// 		return model.Admin{}, "", errors.New("Admin Not Found")
// 	}

// 	if err := bcrypt.CompareHashAndPassword([]byte(data.Password), []byte(password)); err != nil {
// 		return model.Admin{}, "", errors.New("Uncorrect Password")
// 	}

// 	token := ""
// 	var errorToken error
// 	token, errorToken = middleware.CreateToken(int(data.ID))
// 	if errorToken != nil {
// 		return model.Admin{}, "", errorToken

// 	}

// 	return data, token, nil
// }
