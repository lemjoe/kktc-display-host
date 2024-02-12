package internal

import (
	"github.com/lemjoe/kktc-display-host/internal/config"
)

type App struct {
}

func NewApp() *App {
	return &App{}
}

func (a *App) Run() error {
	_, err := config.InitConfig("./.env")
	if err != nil {
		return err
	}
	// db, err := repository.InitializeDB(confDB.DbType, confDB)
	// if err != nil {
	// 	return err
	// }
	// repos, err := db.NewRepository()
	// if err != nil {
	// 	return err
	// }
	// bundle := i18n.NewBundle(language.English)
	// services := service.NewService(repos)
	// err = services.MigrationService.Migrate()
	// if err != nil {
	// 	return err
	// }

	// handlers := handler.NewHandler(services, bundle)
	// err = handlers.Run(":" + confApp.Port)
	return err
}
