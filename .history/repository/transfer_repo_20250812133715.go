package repository

//import "vaqua/models"

type TransferRepository interface {
	//CreateTransfer(transfer *models.Transfer) error
	//GetTransferByID(id uint) (models.Transfer, error)
}

type TransferRepo struct{}

func (r *TransferRepo) CreateTransfer(transaction *models.Transaction) error{
	err :=db.DB.Create(transaction).Errorif err != nil{
		return error
	}
	return nil
}
func (r *TransferRepo) GetTransferByID(id uint)(transaction *models.Transaction)error{
	err :=db.DB.Where("id = ?", id).Find(&transaction).Error
	if errors.Is(err, gorm.ErrRecordNotFound){
		return nil, nil 
	}
	if err != nil{
		return &models.transfer{}, err
	}
	return &transfer, nil
}
