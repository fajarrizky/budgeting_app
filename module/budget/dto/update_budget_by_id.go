package dto

import (
	"budgetapp/module/budget/enum"
	"time"

	"github.com/shopspring/decimal"
)

type UpdateBudgetByIdDTO struct {
	Amount        *decimal.Decimal `json:"amount"`
	CurrentAmount *decimal.Decimal `json:"current_amount"`
	Currency      *string          `json:"currency"`
	Name          *string          `json:"name"`
	Description   *string          `json:"description"`
	Type          *enum.BudgetType `json:"type"`
	StartAt       *time.Time       `json:"start_at"`
	EndAt         *time.Time       `json:"end_at"`
}
