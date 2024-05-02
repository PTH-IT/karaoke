package af

import (
	"karaoke/usecase"

	"github.com/labstack/echo/v4"
)

type commonhandler struct {
	Interactor *usecase.Interactor
}

func AppV1PostLogin(api commonhandler) echo.HandlerFunc {

	return func(context echo.Context) error {
		return api.Interactor.LoginUser(context)
	}

}
func AppV1RegisterUser(api commonhandler) echo.HandlerFunc {

	return func(context echo.Context) error {
		return api.Interactor.RegisterUser(context)
	}
}

func AppV1GetLogout(api commonhandler) echo.HandlerFunc {

	return func(context echo.Context) error {
		return api.Interactor.GetLogout(context)
	}
}
