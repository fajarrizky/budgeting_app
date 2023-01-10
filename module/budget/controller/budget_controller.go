package controller

import (
	"budgetapp/exception"
	"budgetapp/module/budget/dto"
	"budgetapp/module/budget/service"
	"budgetapp/request"
	"budgetapp/response"
	"errors"
	"net/http"

	"github.com/google/uuid"
)

type BudgetController struct {
	budgetService service.BudgetService
}

func NewBudgetController(
	budgetService service.BudgetService,
) *BudgetController {

	return &BudgetController{
		budgetService: budgetService,
	}
}

func (c *BudgetController) CreateBudget(
	rw http.ResponseWriter, r *http.Request) {
	var dto dto.CreateBudgetDto

	err := request.DecodeJSONBody(r, &dto)

	if err != nil {
		response.RespondWithException(rw, exception.BadRequestException(err, err.Error()))
		return
	}

	budget, err := c.budgetService.CreateBudget(r.Context(), dto)

	if err != nil {
		response.RespondWithException(rw, err)
		return
	}

	response.RespondWithJSON(rw, http.StatusOK, response.Response{
		Data: budget,
	})

}

func (c *BudgetController) UpdateBudgetById(
	rw http.ResponseWriter, r *http.Request) {
	var dto dto.UpdateBudgetByIdDTO
	err := request.DecodeJSONBody(r, &dto)

	if err != nil {
		response.RespondWithException(rw, exception.BadRequestException(err))
		return
	}

	id := request.GetUrlParam(r, "id")
	//validate id
	_, parseErr := uuid.Parse(id)
	if parseErr != nil {
		response.RespondWithException(rw, exception.BadRequestException(errors.New("invalid id"), "invalid uuid"))
		return
	}

	appl, err := c.budgetService.UpdateBudgetById(r.Context(), id, dto)

	if err != nil {
		response.RespondWithException(rw, err)
		return
	}

	response.RespondWithJSON(rw, http.StatusOK, response.Response{
		Data: appl,
	})
}

func (c *BudgetController) GetBudgetsByUserId(
	rw http.ResponseWriter, r *http.Request) {

	userId := r.URL.Query().Get("user_id")

	budgets, err := c.budgetService.GetAllBudgetsByUserId(
		r.Context(), userId)

	if err != nil {
		response.RespondWithException(rw, err)
		return
	}

	response.RespondWithJSON(rw, http.StatusOK, response.Response{
		Data: budgets,
	})
}

func (c *BudgetController) GetBudgetById(
	rw http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	budget, err := c.budgetService.GetBudgetById(
		r.Context(), id)

	if err != nil {
		response.RespondWithException(rw, err)
		return
	}

	response.RespondWithJSON(rw, http.StatusOK, response.Response{
		Data: budget,
	})
}

func (c *BudgetController) DeleteBudgetById(
	rw http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	budget, err := c.budgetService.DeleteBudgetById(
		r.Context(), id)

	if err != nil {
		response.RespondWithException(rw, err)
		return
	}

	response.RespondWithJSON(rw, http.StatusOK, response.Response{
		Data: budget,
	})
}
