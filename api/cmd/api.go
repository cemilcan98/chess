package cmd

import (
	"github.com/cemilcan98/chess/internal/chess/controller"
	"github.com/cemilcan98/chess/pkg/echoextention"
	"github.com/cemilcan98/chess/pkg/log"
	mongohelper "github.com/cemilcan98/chess/pkg/mongoextentions"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/cobra"
	"time"
)

// apiCmd represents the api command
var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "A brief description of your command",
}

func init() {
	rootCmd.AddCommand(apiCmd)

	apiCmd.RunE = func(cmd *cobra.Command, args []string) error {

		instance := echo.New()
		echoextention.RegisterGlobalMiddlewares(instance)

		mongoDb, err := mongohelper.NewDatabase(dbconn, dbName)
		if err != nil {
			log.Logger.Fatalf("Failed to connect database. Error :%s", err.Error())
		}
		controller.RegisterHandlers(instance, mongoDb)

		instance.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: []string{"*"},
			AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		}))

		log.Logger.Infof("Service is starting...")
		go func() {
			if err := instance.Start(":5001"); err != nil {
				log.Logger.Fatalf("Failed to shutting down the server. Error:%v", err)
			}
		}()
		echoextention.Shutdown(instance, 3*time.Second)

		return nil
	}
}
