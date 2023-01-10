package budget

import (
	"budgetapp/config"
	"budgetapp/module/budget/controller"
	"budgetapp/module/budget/repo"
	"budgetapp/module/budget/route"
	"budgetapp/module/budget/service"
	"budgetapp/module/logger"
	"budgetapp/router"
	"context"

	"gorm.io/gorm"
)

type BudgetModule struct {
	budgetService service.BudgetService
}

func NewBudgetModule(
	ctx context.Context,
	r router.Router,
	db *gorm.DB,
	loggerFactory logger.Factory,
	configService config.ConfigService,
) *BudgetModule {

	//repos
	budgetRepo := repo.NewBudgetRepo(db)

	//services
	budgetService := service.NewBudgetService(loggerFactory.NewLogger(), budgetRepo)

	//controllers
	budgetController := controller.NewBudgetController(budgetService)

	//routes
	r.RegisterRouteGroup(route.BudgetV1Routes(budgetController))

	return &BudgetModule{
		budgetService: budgetService,
	}
}
