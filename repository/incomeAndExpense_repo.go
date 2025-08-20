package repository

import (
	"time"
	"vaqua/db"
	"vaqua/models"
)

type IncomeAndExpensesRepository interface {
	SumByType(accountID int64, kind string, from, to *time.Time) (float64, error)
}

type IncomeAndExpensesRepo struct {

}

//func NewIncomeAndExpenses(db *gorm.DB) IncomeAndExpensesRepository {
//	return &IncomeAndExpensesRepo{db: db}
//}

func (r *IncomeAndExpensesRepo) SumByType(accountID int64, kind string, from, to *time.Time) (float64, error) {
	var total float64 
	q := db.DB.Model(&models.IncomeAndExpenses{}).
    Select("COALESCE(SUM(amount), 0)").
    Where("account_id = ? AND type = ?", accountID, kind)


	if from != nil {
		q = q.Where("created_at >= ?", *from)
	}
	if to != nil {
		next := to.Add(24 * time.Hour)
		q = q.Where("created_at < ?", next)
	}

	err := q.Scan(&total).Error; 
	if err != nil {
		return 0, err
	}

	return total, nil
}