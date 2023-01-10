package dto

import (
	"budgetapp/module/budget/enum"
	"time"

	"github.com/shopspring/decimal"
)

type CreateBudgetDto struct {
	UserId      string          `json:"user_id"`
	Amount      decimal.Decimal `json:"amount"`
	Currency    string          `json:"currency"`
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Type        enum.BudgetType `json:"type"`
	StartAt     *time.Time      `json:"start_at"`
	EndAt       *time.Time      `json:"end_at"`
}
