package models

type TransferRequest struct {
	FromUserID uint    `json:"from_user_id"`
	ToUserID   uint    `json:"to_user_id"`
	Amount     float64 `json:"amount"`
}
