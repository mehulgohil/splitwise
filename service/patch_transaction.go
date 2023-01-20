package service

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/mehulgohil/splitwise/models"
)

func (s *ServiceStruct) UpdatePaymentStatus(transactionId int, request models.PatchTransactionRequest) error {

	var peopleObjectID int

	// get people id by mobile number
	sqlErr := s.DB.QueryRow("SELECT id FROM people WHERE mobile_number = ?", request.Mobile).Scan(&peopleObjectID)
	if sqlErr != nil && sqlErr == sql.ErrNoRows{
		if sqlErr == sql.ErrNoRows {
			return errors.New("no transactions for given mobile number")
		}
		return fmt.Errorf("getPeopleByMobile: %v", sqlErr)
	}

	// update transaction status
	result, err := s.DB.Exec("UPDATE transactions set status=? where id=? and owed_to=?", "paid", transactionId, peopleObjectID)
	if err != nil {
		return fmt.Errorf("updateTransaction: %v", err)
	}
	noOfRowsAffected, _ := result.RowsAffected()
	if noOfRowsAffected < 1 {
		return errors.New("transaction not found")
	}

	return nil
}
