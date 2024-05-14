package usercase

// import (
// 	"campsite_reservation/dto"
// 	"campsite_reservation/model"
// 	"campsite_reservation/repository"
// 	"errors"

// 	// "errors"

// 	"golang.org/x/crypto/bcrypt"
// )

// type AdminUseCase struct {
// 	Login(email, password string) (model.Admin, string, error)
// 	Create(payloads dto.CreateAdminrRequest) error
// 	SelectAll() ([]model.Admin, error)
// }

// type adminUsecase struct {
// 	adminRepository repository.AdminRepository
// }

// func AdminUsecase(ur repository.UserRepository) *userUsecase {
// 	return &userUsecase{userRepository: ur}
// }

// func (us *userUsecase) Create(payload dto.CreateUserRequest) error {
// 	if payload.Email == "" || payload.Password == "" {
// 		return errors.New("error. email dan password harus diisi")
// 	}
// 	securePassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
// 	if err != nil {
// 		return err
// 	}

// 	user := model.User{
// 		Email:    payload.Email,
// 		Password: string(securePassword),
// 	}

// 	if err := us.userRepository.Create(user); err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (us *userUsecase) SelectAll() ([]model.User, error) {
// 	users, err := us.userRepository.SelectAll()
// 	if err != nil {
// 		return nil, err
// 	}
// 	return users, nil
// }

// func (us *userUsecase) Login(email, password string) (model.User, string, error) {

// 	user := model.User{
// 		Email:    email,
// 		Password: password,
// 	}

// 	data, token, err := us.userRepository.Login(user.Email, user.Password)
// 	if err != nil {
// 		return model.User{}, "", err
// 	}
// 	return data, token, nil
// }
