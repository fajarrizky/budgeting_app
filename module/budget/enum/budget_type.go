package enum

import (
	"database/sql/driver"
	"errors"
)

type BudgetType struct {
	slug string
}

var (
	UNKNOWN        = BudgetType{""}
	INVESTMENT     = BudgetType{"INVESTMENT"}
	GROCERIES      = BudgetType{"GROCERIES"}
	TRANSPORTATION = BudgetType{"TRANSPORTATION"}
	UTILITIES      = BudgetType{"UTILITIES"}
	RENT           = BudgetType{"RENT"}
	ENTERTAINMENT  = BudgetType{"ENTERTAINMENT"}
	OTHER          = BudgetType{"OTHER"}
)

func (a *BudgetType) String() string {
	return a.slug
}

func BudgetTypeFromString(s string) (BudgetType, error) {
	switch s {
	case INVESTMENT.slug:
		return INVESTMENT, nil
	case GROCERIES.slug:
		return GROCERIES, nil
	case TRANSPORTATION.slug:
		return TRANSPORTATION, nil
	case UTILITIES.slug:
		return UTILITIES, nil
	case RENT.slug:
		return RENT, nil
	case ENTERTAINMENT.slug:
		return ENTERTAINMENT, nil
	case OTHER.slug:
		return OTHER, nil
	}

	return UNKNOWN, errors.New("unknown budget type")
}

func (a *BudgetType) Value() (driver.Value, error) {
	if a == nil {
		return nil, nil
	}
	return a.slug, nil
}

func (a *BudgetType) Scan(val any) error {
	valString, ok := val.(string)
	if !ok {
		return errors.New("could not scan BudgetType type | assertion failed")
	}
	addrType, err := BudgetTypeFromString(valString)

	if err != nil {
		return err
	}

	*a = addrType

	return nil
}

// func (a *BudgetType) MarshalJSON() ([]byte, error) {
// 	return json.Marshal(a.slug)
// }

// func (a *BudgetType) UnmarshalJSON(data []byte) error {
// 	var slug string

// 	if err := json.Unmarshal(data, &slug); err != nil {
// 		return err
// 	}

// 	addrType, err := BudgetTypeFromString(slug)

// 	if err != nil {
// 		return nil
// 	}

// 	*a = addrType

// 	return nil
// }
