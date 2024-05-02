package usecase

import (
	"fmt"
	"net/http"

	"karaoke/domain/model"
	errormessage "karaoke/log/error"
	"karaoke/utils"

	InforLog "karaoke/log/infor"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func NewInteractor(
	gormDb *gorm.DB,
	referrance Reference,

) Interactor {

	return Interactor{
		gormDb,
		referrance,
	}
}

type Interactor struct {
	gormDb     *gorm.DB
	referrance Reference
}

// LoginUser godoc
//	@Summary		LoginUser
//	@Description	login username
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			Content-Type	header		string		true	"application/json"	default(application/json)
//	@Param			Content-Length	header		string		true	"1000"				default(1000)
//	@Param			Host			header		string		true	"localhost"			default(localhost)
//	@Param			user			body		model.Login	true	"model.Login"
//	@Success		201				{object}	model.Token
//	@Failure		400				{object}	string
//	@Router			/login [post]
func (i *Interactor) LoginUser(context echo.Context) error {
	InforLog.PrintLog(fmt.Sprintf("LoginUser start"))

	var user model.Login
	err := context.Bind(&user)
	if err != nil {
		return context.String(http.StatusBadRequest, errormessage.PrintError("3", err).Error())
	}
	InforLog.PrintLog(fmt.Sprintf("reqest user: %v", user))

	result, err := i.referrance.GetUser(user.UserID, *utils.CryptPassword(user.Password))
	if err != nil {
		return context.String(http.StatusBadRequest, err.Error())
	}
	if result == nil {

		return context.JSON(http.StatusForbidden, model.MessageCheckUser{Type: "user", Message: "username or password is not correst"})
	}

	tokenString := utils.GenerateToken(result.UserID)
	token := &model.Token{
		Authorization: tokenString,
		Type:          "bearer",
	}
	InforLog.PrintLog(fmt.Sprintf("response token: %v", token))

	err = utils.SetToken(tokenString, user.UserID)
	if err != nil {
		return context.String(http.StatusBadRequest, errormessage.PrintError("5", err).Error())
	}
	return context.JSON(http.StatusOK, token)

}

// AddUser godoc
//	@Summary		AddUser
//	@Description	Add new user to database
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			Content-Type	header		string				true	"application/json"	default(application/json)
//	@Param			Content-Length	header		string				true	"1000"				default(1000)
//	@Param			Host			header		string				true	"localhost"			default(localhost)
//	@Param			token			body		model.RegisterUser	true	"model.RegisterUser"
//	@Success		200				{object}	string
//	@Failure		400				{object}	string
//	@Router			/adduser [post]
func (i *Interactor) RegisterUser(context echo.Context) error {
	InforLog.PrintLog(fmt.Sprintf("RegisterUser start"))

	var Adduser model.RegisterUser
	err := context.Bind(&Adduser)

	if err != nil || Adduser.Email == "" || Adduser.UserID == "" || Adduser.Password == "" {
		errData := map[string]interface{}{
			"message": "request body is invalid",
		}
		return context.JSON(http.StatusBadRequest, errData)
	}
	InforLog.PrintLog(fmt.Sprintf("reqest user: %v", Adduser))

	cryptPassword := utils.CryptPassword(Adduser.Password)
	result, err := i.referrance.CheckUserName(Adduser.UserID, Adduser.Email)
	if err != nil {
		return err
	}

	if result != nil {
		var messageError []model.MessageCheckUser
		for _, r := range result {
			if r.UserID == Adduser.UserID {
				messageError = append(messageError, model.MessageCheckUser{Type: "username", Message: "username is exist "})
			}
			if r.Email == Adduser.Email {
				messageError = append(messageError, model.MessageCheckUser{Type: "email", Message: "email is exist"})
			}

		}
		return context.JSON(http.StatusBadRequest, messageError)
	}
	err = i.referrance.AddUser(Adduser.UserID, *cryptPassword, Adduser.Email)
	if err != nil {
		return err
	}
	return context.NoContent(http.StatusOK)
}

// GetLogout godoc
//	@Summary		GetLogout
//	@Description	GetLogout
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			Content-Type	header		string	true	"application/json"	default(application/json)
//	@Param			Content-Length	header		string	true	"1000"				default(1000)
//	@Param			Host			header		string	true	"localhost"			default(localhost)
//	@Param			Authorization	header		string	true	"Authorization"
//	@Success		200				{object}	string
//	@Failure		400				{object}	error
//	@Router			/logout [get]
func (i *Interactor) GetLogout(context echo.Context) error {
	InforLog.PrintLog(fmt.Sprintf("GetLogout start"))

	authercations := context.Request().Header.Get("Authorization")
	InforLog.PrintLog(fmt.Sprintf("authercations = %s", authercations))

	user := utils.ParseToken(authercations)
	InforLog.PrintLog(fmt.Sprintf("user = %v", user))

	userID := user.Claims.(jwt.MapClaims)["userID"].(string)
	InforLog.PrintLog(fmt.Sprintf("userID = %s", userID))
	InforLog.PrintLog(fmt.Sprintf("utils.GetToken"))

	if !utils.GetToken(authercations, userID) {
		return context.String(http.StatusForbidden, "token awrong")
	}
	InforLog.PrintLog(fmt.Sprintf("utils.DeleteToken"))

	if !utils.DeleteToken(authercations, userID) {
		return context.String(http.StatusBadRequest, "Can not delete token")
	}
	InforLog.PrintLog(fmt.Sprintf("StatusOK"))

	return context.String(http.StatusOK, "susscess")
}
