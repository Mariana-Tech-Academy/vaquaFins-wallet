package service

import (
	"time"
	"vaqua/repository"
)

//it’s a DTO (Data Transfer Object) that exists only in Go memory.
type Summary struct {
	AccountID    int64    `json:"account_id"`
	From         *string  `json:"from,omitempty"`
	To           *string  `json:"to,omitempty"`
	IncomeTotal  float64  `json:"income_total"`
	ExpensesTotal float64 `json:"expenses_total"`
	Net          float64  `json:"net"`
}

/*type IncomeAndExpensesRepository interface {
	SumByType(accountID int64, kind string, from, to *time.Time) (float64, error)
}*/

type IncomeAndExpensesService struct {
	Repo repository.IncomeAndExpensesRepository
}



func (s *IncomeAndExpensesService) GetSummary(accountID int64, from, to *time.Time, fromStr, toStr *string) (*Summary, error) {
	income, err := s.Repo.SumByType(accountID, "income", from, to)
	if err != nil {
		return nil, err
	}
	expenses, err := s.Repo.SumByType(accountID, "expense", from, to)
	if err != nil {
		return nil, err
	}
	sum := &Summary{
		AccountID:     accountID,
		From:          fromStr,
		To:            toStr,
		IncomeTotal:   income,
		ExpensesTotal: expenses,
		Net:           income - expenses,
	}
	return sum, nil
}





