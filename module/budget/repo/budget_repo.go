package repo

import (
	"budgetapp/module/budget/entity"
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type budgetRepo struct {
	db *gorm.DB
}

type BudgetRepo interface {
	SaveBudgets(context.Context, []entity.Budget) ([]entity.Budget, error)
	GetAllBudgetsByUserId(context.Context, string) ([]entity.Budget, error)
	GetBudgetById(context.Context, string) (*entity.Budget, error)
	UpdateBudgetById(context.Context, *entity.Budget) (*entity.Budget, error)
	DeleteBudgetById(context.Context, string) (*entity.Budget, error)
}

func NewBudgetRepo(db *gorm.DB) BudgetRepo {
	return &budgetRepo{
		db: db,
	}
}

func (r *budgetRepo) SaveBudgets(ctx context.Context,
	budgets []entity.Budget) ([]entity.Budget, error) {
	err := r.db.WithContext(ctx).Create(budgets).Error

	if err != nil {
		return nil, err
	}

	return budgets, nil
}

func (r *budgetRepo) GetBudgetById(ctx context.Context,
	id string) (*entity.Budget, error) {

	var budget *entity.Budget
	err := r.db.WithContext(ctx).
		Where("id = ?", id).First(budget).Error

	if err != nil {
		return nil, err
	}

	return budget, nil
}

func (r *budgetRepo) GetAllBudgetsByUserId(ctx context.Context,
	userId string) ([]entity.Budget, error) {

	var budgets []entity.Budget
	err := r.db.WithContext(ctx).
		Where("user_id = ?", userId).
		Find(&budgets).Error

	if err != nil {
		return nil, err
	}

	return budgets, nil
}

func (r *budgetRepo) UpdateBudgetById(ctx context.Context,
	budget *entity.Budget) (*entity.Budget, error) {

	updBudget := &entity.Budget{ID: budget.ID}
	err := r.db.WithContext(ctx).Model(updBudget).
		Updates(budget).Error

	if err != nil {
		return nil, err
	}

	return updBudget, nil
}

func (r *budgetRepo) DeleteBudgetById(ctx context.Context, id string) (*entity.Budget, error) {
	appl := &entity.Budget{
		ID: id,
	}
	err := r.db.WithContext(ctx).Clauses(clause.Returning{}).Delete(appl).Error

	if err != nil {
		return nil, err
	}

	return appl, nil
}
