package entity

import (
	"budgetapp/module/budget/enum"
	"budgetapp/module/db"
	"budgetapp/module/money"
	"time"
)

type Budget struct {
	db.BaseEntity
	ID            string          `json:"id"`
	UserId        string          `json:"user_id"`
	Amount        money.Money     `json:"amount"`
	CurrentAmount money.Money     `json:"current_amount"`
	Name          string          `json:"name"`
	Description   string          `json:"description"`
	Type          enum.BudgetType `json:"type"`
	StartAt       *time.Time      `json:"start_at"`
	EndAt         *time.Time      `json:"end_at"`
}
