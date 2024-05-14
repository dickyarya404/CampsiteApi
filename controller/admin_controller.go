package controller

// import (
// 	"campsite_reservation/dto"
// 	"campsite_reservation/helper"
// 	"campsite_reservation/usercase"
// 	"net/http"

// 	"github.com/labstack/echo/v4"
// )

// type adminController struct {
// 	adminCase usercase.AdminUsecase // usercase.AdminUsecase
// }

// func NewAdminController(adminUsecase usercase.AdminUsecase) *adminController {
// 	return &adminController{adminCase: adminUsecase}

// }

// func (uc *adminController) Create(e echo.Context) error {
// 	var payload dto.CreateAdminRequest

// 	if err := e.Bind(&payload); err != nil {
// 		return err
// 	}

// 	err := uc.adminCase.Create(payload)
// 	if err != nil {
// 		return e.JSON(http.StatusBadRequest, helper.FailedResponse(err.Error()))
// 	}

// 	return e.JSON(http.StatusOK, helper.SuccessResponse("Succes Create Admin"))
// }

// func (uc *adminController) Login(e echo.Context) error {
// 	var payload dto.CreateAdminRequest

// 	if err := e.Bind(&payload); err != nil {
// 		return e.JSON(http.StatusBadRequest, helper.FailedResponse("Invalid Request"))
// 	}

// 	user, token, err := uc.adminCase.Login(payload.Email, payload.Password)
// 	if err != nil {
// 		return e.JSON(http.StatusUnauthorized, helper.FailedResponse("Unauthorized"))
// 	}

// 	response := dto.ResponseLogin{

// 		Email: user.Email,
// 		Token: token,
// 	}

// 	return e.JSON(http.StatusOK, helper.SuccessWithDataResponse("Succes Login", response))
// }
