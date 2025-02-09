package dto

type TransactionRequest struct {
	UserID string `json:"user_id"`
	Amount string `json:"amount"`
}

type TransferRequest struct {
	SenderID   string `json:"sender_id"`
	ReceiverID string `json:"receiver_id"`
	Amount     string `json:"amount"`
}
