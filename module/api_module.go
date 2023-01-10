package module

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"budgetapp/config"
	"budgetapp/module/budget"
	"budgetapp/module/db"
	"budgetapp/module/logger"
	"budgetapp/router"

	"gorm.io/gorm"

	chimddlwr "github.com/go-chi/chi/v5/middleware"
)

type ApiModule struct {
	router       router.Router
	dB           *gorm.DB
	configModule *config.ConfigModule
	budgetModule *budget.BudgetModule
}

func NewApiModule(ctx context.Context) *ApiModule {

	configModule := config.NewConfigModule()

	configService := configModule.GetConfigService()

	loggerFactory := logger.NewFactory(configService.GetEnv())

	db.RunMigrations(configService.GetDbConfig(), "migrations")

	dB := db.InitDb(configService.GetDbConfig(), configService.GetEnv())

	r := router.NewChiRouter()
	r.AddMiddleWare(chimddlwr.Logger)
	r.AddMiddleWare(chimddlwr.Recoverer)
	r.AddMiddleWare(chimddlwr.RealIP)

	//modules
	budgetMod := budget.NewBudgetModule(ctx, r, dB, loggerFactory, configService)

	return &ApiModule{
		router:       r,
		dB:           dB,
		configModule: configModule,
		budgetModule: budgetMod,
	}
}

func (m *ApiModule) GetMux() http.Handler {
	return m.router.GetMux()
}

func (m *ApiModule) GetConfigService() config.ConfigService {
	return m.configModule.GetConfigService()
}

func (m *ApiModule) CleanUp() {

	// closing db connection
	if m.dB != nil {
		sqlDb, err := m.dB.DB()

		if err != nil {
			log.Printf("failed to close db connection | err: %s", err.Error())
		} else {
			err = sqlDb.Close()

			if err != nil {
				log.Printf("error while closing db connection | err: %s", err.Error())
			} else {
				fmt.Println("successfully closed db connection")
			}
		}
	}
}
