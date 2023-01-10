package route

import (
	"budgetapp/module/budget/controller"
	"budgetapp/router"
	"net/http"
)

func BudgetV1Routes(c *controller.BudgetController) router.RouteGroup {
	return router.RouteGroup{
		BasePath: "/v1/budgets",
		Routes: []router.Route{
			{
				Path:    "/",
				Method:  http.MethodPost,
				Handler: c.CreateBudget,
			},
			{
				Path:    "/{id}",
				Method:  http.MethodPut,
				Handler: c.UpdateBudgetById,
			},
			{
				Path:    "/{id}",
				Method:  http.MethodGet,
				Handler: c.GetBudgetById,
			},
			{
				Path:    "/{user_id}",
				Method:  http.MethodGet,
				Handler: c.GetBudgetsByUserId,
			},
			{
				Path:    "/{id}",
				Method:  http.MethodDelete,
				Handler: c.DeleteBudgetById,
			},
		},
	}
}