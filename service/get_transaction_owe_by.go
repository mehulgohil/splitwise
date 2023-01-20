package service

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/mehulgohil/splitwise/models"
)

func (s *ServiceStruct) GetOweByTransactions(mobile string) ([]models.TransactionModel, error) {
	var finalResp []models.TransactionModel
	var oweByPeopleId int
	// get people id by mobile number
	sqlErr := s.DB.QueryRow("SELECT id FROM people WHERE mobile_number = ?", mobile).Scan(&oweByPeopleId)
	if sqlErr != nil && sqlErr == sql.ErrNoRows{
		if sqlErr == sql.ErrNoRows {
			return finalResp, errors.New("person not found with given mobile number")
		}
		return finalResp, fmt.Errorf("getPeopleByMobile: %v", sqlErr)
	}

	// get transactions of the person owe to
	rows, err := s.DB.Query("SELECT transaction_place_id, spent_by, amount, status FROM transactions WHERE owed_to = ?", oweByPeopleId)
	if err != nil {
		return finalResp, fmt.Errorf("getTransactionsByMobile: %v", err)
	}
	for rows.Next() {
		eachTransaction := models.TransactionModel{}
		if err := rows.Scan(&eachTransaction.TransactionPerPlaceId, &eachTransaction.SpentBy.PersonID, &eachTransaction.MyShare, &eachTransaction.PaymentStatus); err != nil {
			return finalResp, fmt.Errorf("couldn't scan row: %v", err)
		}

		// get transaction per place details
		sqlErr = s.DB.QueryRow("SELECT * FROM transactions_per_place WHERE id = ?", eachTransaction.TransactionPerPlaceId).Scan(&eachTransaction.TransactionPerPlaceId, &eachTransaction.TotalAmount, &eachTransaction.Place, &eachTransaction.Date)
		if sqlErr != nil {
			return finalResp, fmt.Errorf("couldn't scan row: %v", err)
		}

		// get transaction for who owed details
		sqlErr := s.DB.QueryRow("SELECT name, mobile_number FROM people WHERE id = ?", eachTransaction.SpentBy.PersonID).Scan(&eachTransaction.SpentBy.Name, &eachTransaction.SpentBy.Mobile)
		if sqlErr != nil {
			return finalResp, fmt.Errorf("couldn't get spent by details: %v", err)
		}

		finalResp = append(finalResp, eachTransaction)
	}

	return finalResp, nil
}
