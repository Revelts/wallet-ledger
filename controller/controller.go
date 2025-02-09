package controller

import (
	"encoding/json"
	"go-testing/constants"
	"go-testing/database"
	"go-testing/dto"
	"go-testing/helper"
	"net/http"
	"time"

	_ "github.com/lib/pq"
)

func DepositHandler(w http.ResponseWriter, r *http.Request) {
	var req dto.TransactionRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helper.HttpResponseError(w, r, err.Error(), http.StatusBadRequest)
		return
	}

	_, err := database.DbConn.Exec(constants.QueryDeposit, req.UserID, req.Amount, time.Now())
	if err != nil {
		helper.HttpResponseError(w, r, err.Error(), http.StatusInternalServerError)
		return
	}

	helper.HttpResponseSuccess(w, r, http.StatusOK)
}

func WithdrawHandler(w http.ResponseWriter, r *http.Request) {
	var req dto.TransactionRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helper.HttpResponseError(w, r, err.Error(), http.StatusBadRequest)
		return
	}

	_, err := database.DbConn.Exec(constants.QueryWithdraw, req.UserID, req.Amount, time.Now())
	if err != nil {
		helper.HttpResponseError(w, r, constants.ErrInsufficientFunds, http.StatusBadRequest)
		return
	}

	helper.HttpResponseSuccess(w, r, http.StatusOK)
}

func TransferHandler(w http.ResponseWriter, r *http.Request) {
	var req dto.TransferRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helper.HttpResponseError(w, r, err.Error(), http.StatusBadRequest)
		return
	}

	tx, err := database.DbConn.Begin()
	if err != nil {
		helper.HttpResponseError(w, r, constants.ErrTransactionStart, http.StatusInternalServerError)
		return
	}

	_, err = tx.Exec(constants.QueryTransferOut, req.SenderID, req.Amount, time.Now())
	if err != nil {
		tx.Rollback()
		helper.HttpResponseError(w, r, constants.ErrInsufficientTransfer, http.StatusBadRequest)
		return
	}

	_, err = tx.Exec(constants.QueryTransferIn, req.ReceiverID, req.Amount, time.Now())
	if err != nil {
		tx.Rollback()
		helper.HttpResponseError(w, r, constants.ErrProcessingTransfer, http.StatusInternalServerError)
		return
	}

	tx.Commit()
	helper.HttpResponseSuccess(w, r, http.StatusOK)
}

func WalletBalanceHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get(constants.UserId)

	var balance string
	err := database.DbConn.QueryRow(constants.QueryWalletBalance, userID).Scan(&balance)
	if err != nil {
		helper.HttpResponseError(w, r, constants.ErrFetchingBalance, http.StatusInternalServerError)
		return
	}
	helper.HttpResponseSuccess(w, r, dto.WalletBalanceResponse{Balance: balance})
}

func WalletHistoryHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get(constants.UserId)

	rows, err := database.DbConn.Query(constants.QueryWalletHistory, userID)
	if err != nil {
		helper.HttpResponseError(w, r, constants.ErrFetchingHistory, http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var transactions []dto.WalletHistoryItem
	for rows.Next() {
		var txn dto.WalletHistoryItem
		if err = rows.Scan(&txn.Amount, &txn.Balance, &txn.Type, &txn.CreatedAt); err != nil {
			helper.HttpResponseError(w, r, constants.ErrScanningRow, http.StatusInternalServerError)
			return
		}
		transactions = append(transactions, txn)
	}

	helper.HttpResponseSuccess(w, r, transactions)
}
