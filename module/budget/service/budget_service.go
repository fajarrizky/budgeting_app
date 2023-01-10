package service

import (
	"budgetapp/module/budget/dto"
	"budgetapp/module/budget/entity"
	"budgetapp/module/budget/repo"
	"budgetapp/module/logger"
	"budgetapp/module/money"
	"context"

	mon "github.com/Rhymond/go-money"
	"github.com/google/uuid"
)

type budgetService struct {
	log        logger.Logger
	budgetRepo repo.BudgetRepo
	uuid       uuid.UUID
}

type BudgetService interface {
	CreateBudget(context.Context, dto.CreateBudgetDto) (*entity.Budget, error)
	UpdateBudgetById(context.Context, string, dto.UpdateBudgetByIdDTO) (*entity.Budget, error)
	GetBudgetById(context.Context, string) (*entity.Budget, error)
	GetAllBudgetsByUserId(context.Context, string) ([]entity.Budget, error)
	DeleteBudgetById(context.Context, string) (*entity.Budget, error)
}

func NewBudgetService(
	log logger.Logger,
	budgetRepo repo.BudgetRepo,
) BudgetService {
	return &budgetService{
		log:        log,
		budgetRepo: budgetRepo,
		uuid:       uuid.New(),
	}
}

func (r *budgetService) CreateBudget(ctx context.Context,
	createBudget dto.CreateBudgetDto) (*entity.Budget, error) {
	budget := &entity.Budget{
		ID:            r.uuid.String(),
		UserId:        createBudget.UserId,
		Amount:        money.Money{Amount: *mon.NewFromFloat(createBudget.Amount.InexactFloat64(), createBudget.Currency)},
		CurrentAmount: money.Money{Amount: *mon.NewFromFloat(createBudget.Amount.InexactFloat64(), createBudget.Currency)},
		Name:          createBudget.Name,
		Description:   createBudget.Description,
		StartAt:       createBudget.StartAt,
		EndAt:         createBudget.EndAt,
	}

	result, err := r.budgetRepo.SaveBudgets(ctx, []entity.Budget{*budget})
	if err != nil {
		return nil, err
	}
	return &result[0], nil
}

func (r *budgetService) UpdateBudgetById(ctx context.Context, id string,
	updateBudget dto.UpdateBudgetByIdDTO) (*entity.Budget, error) {

	var amount, currAmount money.Money

	if updateBudget.Currency != nil {
		if updateBudget.Amount != nil {
			amount = money.Money{Amount: *mon.NewFromFloat(updateBudget.Amount.InexactFloat64(), *updateBudget.Currency)}
		}

		if updateBudget.CurrentAmount != nil {
			currAmount = money.Money{Amount: *mon.NewFromFloat(updateBudget.CurrentAmount.InexactFloat64(), *updateBudget.Currency)}
		}
	}

	budget := &entity.Budget{
		ID:            id,
		Amount:        amount,
		CurrentAmount: currAmount,
		Name:          *updateBudget.Name,
		Description:   *updateBudget.Description,
		StartAt:       updateBudget.StartAt,
		EndAt:         updateBudget.EndAt,
	}

	result, err := r.budgetRepo.UpdateBudgetById(ctx, budget)

	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *budgetService) GetBudgetById(ctx context.Context, id string) (*entity.Budget, error) {
	return r.budgetRepo.GetBudgetById(ctx, id)
}

func (r *budgetService) GetAllBudgetsByUserId(ctx context.Context, id string) ([]entity.Budget, error) {
	return r.budgetRepo.GetAllBudgetsByUserId(ctx, id)
}

func (r *budgetService) DeleteBudgetById(ctx context.Context, id string) (*entity.Budget, error) {
	return r.budgetRepo.DeleteBudgetById(ctx, id)
}
