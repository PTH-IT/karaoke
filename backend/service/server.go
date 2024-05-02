package af

import (
	"fmt"
	"net/http"

	gormdb "karaoke/adapter/gormdb"
	config "karaoke/config"
	usecase "karaoke/usecase"
	"karaoke/utils"

	InforLog "karaoke/log/infor"

	"github.com/golang-jwt/jwt/v4"
	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Run() {
	InforLog.PrintLog(fmt.Sprintf("echo.New call"))
	e := echo.New()

	connectString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",

		config.Getconfig().PostgreDB.Host,
		config.Getconfig().PostgreDB.User,
		config.Getconfig().PostgreDB.Pass,
		config.Getconfig().PostgreDB.Db,
		config.Getconfig().PostgreDB.Port,
	)
	var err error
	gormDb, err := gorm.Open(postgres.Open(connectString), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}
	gormdb.Start(gormDb)
	userRepository := gormdb.NewUser()
	referrance := usecase.NewReferrance(userRepository)
	interactor := usecase.NewInteractor(gormDb, referrance)

	api := commonhandler{
		Interactor: &interactor,
	}
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"*"},
	}))

	e.POST("/login", AppV1PostLogin(api))
	e.POST("/register", AppV1RegisterUser(api))
	e.GET("/logout", AppV1GetLogout(api))
	e.GET("/swageer/*", echoSwagger.WrapHandler)
	e.Static("/*", "../web")
	fmt.Println("" + config.Getconfig().Port)
	// e.Logger.SetOutput(io.Discard)
	e.Logger.Fatal(e.Start(":" + config.Getconfig().Port))
}

func Checktoken(context echo.Context) error {
	authercations := context.Request().Header.Get("Authorization")
	user := utils.ParseToken(authercations)
	userID := user.Claims.(jwt.MapClaims)["userID"].(string)
	if !utils.GetToken(authercations, userID) {
		return context.String(http.StatusForbidden, "token awrong")
	}
	return nil
}
