package service

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/mehulgohil/splitwise/models"
)

func (s *ServiceStruct) GetOweToTransactions(mobile string) (models.TransactionModel, error) {
	var finalResp models.TransactionModel

	// get people id by mobile number
	sqlErr := s.DB.QueryRow("SELECT * FROM people WHERE mobile_number = ?", mobile).Scan(&finalResp.SpentBy.PersonID, &finalResp.SpentBy.Name, &finalResp.SpentBy.Mobile)
	if sqlErr != nil && sqlErr == sql.ErrNoRows{
		if sqlErr == sql.ErrNoRows {
			return finalResp, errors.New("person not found with given mobile number")
		}
		return finalResp, fmt.Errorf("getPeopleByMobile: %v", sqlErr)
	}

	// get transactions of the person owe to
	rows, err := s.DB.Query("SELECT * FROM transactions WHERE spent_by = ?", finalResp.SpentBy.PersonID)
	if err != nil {
		return finalResp, fmt.Errorf("getTransactionsByMobile: %v", err)
	}

	var owedTo []models.PostTransactionRequestSplit

	for rows.Next() {
		eachSplit := models.PostTransactionRequestSplit{}
		if err := rows.Scan(&finalResp.TransactionId, &finalResp.TransactionPerPlaceId, &finalResp.SpentBy.PersonID, &eachSplit.PersonID, &eachSplit.ShareAmount, &eachSplit.PaymentStatus); err != nil {
			return finalResp, fmt.Errorf("couldn't scan row: %v", err)
		}

		sqlErr := s.DB.QueryRow("SELECT name, mobile_number FROM people WHERE id = ?", eachSplit.PersonID).Scan(&eachSplit.Name, &eachSplit.Mobile)
		if sqlErr != nil {
			return finalResp, fmt.Errorf("couldn't scan row: %v", err)
		}
		owedTo = append(owedTo, eachSplit)
	}
	finalResp.Split = owedTo

	// get no of people owed
	finalResp.NPeople = len(finalResp.Split)

	// get transaction per place details
	sqlErr = s.DB.QueryRow("SELECT * FROM transactions_per_place WHERE id = ?", finalResp.TransactionPerPlaceId).Scan(&finalResp.TransactionPerPlaceId, &finalResp.TotalAmount, &finalResp.Place, &finalResp.Date)
	if sqlErr != nil {
		return finalResp, fmt.Errorf("couldn't scan row: %v", err)
	}

	return finalResp, nil
}
