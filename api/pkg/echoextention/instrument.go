package echoextention

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"strings"
)

//Root Level (After router)
//The following built-in middleware should be registered at this level:
//BodyLimit
//Logger
//Gzip
//Recover
//ServerHeader middleware adds a `Server` header to the response.
func RegisterGlobalMiddlewares(e *echo.Echo) {

	e.Use(middleware.BodyLimit("1M"))
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Skipper: Myskipper,
		Level:   -1,
	}))

}

func Myskipper(context echo.Context) bool {
	if strings.HasPrefix(context.Path(), "/status") ||
		strings.HasPrefix(context.Path(), "/swagger") ||
		strings.HasPrefix(context.Path(), "/metrics") {
		return true
	}

	return false
}

func addBetterErrorHandling(e *echo.Echo) {

	//Recover
	//https://echo.labstack.com/middleware/recover
	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize:         4 << 10, // 4 KB
		DisableStackAll:   false,
		DisablePrintStack: false,
	}))
}
