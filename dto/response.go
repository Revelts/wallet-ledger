package dto

type TransactionResponse struct {
	Status string `json:"status"`
}

type WalletBalanceResponse struct {
	Balance string `json:"balance"`
}

type WalletHistoryItem struct {
	Amount    string `json:"amount"`
	Balance   string `json:"balance"`
	Type      string `json:"type"`
	CreatedAt string `json:"created_at"`
}
