package service

import (
	"database/sql"
	"fmt"
	"github.com/mehulgohil/splitwise/models"
)

func (s *ServiceStruct) PostTransactionToDB(request models.TransactionModel) error {

	// add spender to people list
	peopleObjId, err := s.checkAndUpdatePeopleTable(request.SpentBy.Name, request.SpentBy.Mobile)
	if err != nil {
		return err
	}
	if peopleObjId != -1 {
		request.SpentBy.PersonID = peopleObjId
	}

	// add owed to person in db if not exist
	for i:=0; i<len(request.Split); i++ {
		peopleObjId, err := s.checkAndUpdatePeopleTable(request.Split[i].Name, request.Split[i].Mobile)
		if err != nil {
			return err
		}
		if peopleObjId != -1 {
			request.Split[i].PersonID = peopleObjId
		}
	}

	// insert place transaction
	result, err := s.DB.Exec("INSERT INTO transactions_per_place (total_amount, place, date) VALUES (?, ?, ?)", request.TotalAmount, request.Place, request.Date)
	if err != nil {
		return fmt.Errorf("addTransactionPerPlace: %v", err)
	}
	id, _ := result.LastInsertId()
	request.TransactionPerPlaceId = int(id)

	// add transactions
	for i:=0; i<len(request.Split); i++ {
		_, err = s.DB.Exec("INSERT INTO transactions (transaction_place_id, spent_by, owed_to, amount, status) VALUES (?, ?, ?, ?, ?)", request.TransactionPerPlaceId, request.SpentBy.PersonID, request.Split[i].PersonID, request.Split[i].ShareAmount, "pending")
		if err != nil {
			return fmt.Errorf("addTransactionPerPlace: %v", err)
		}
	}

	return nil
}

func (s *ServiceStruct) checkAndUpdatePeopleTable(name, mobileNumber string) (int, error) {
	var peopleObjectID int

	sqlErr := s.DB.QueryRow("SELECT id FROM people WHERE mobile_number = ?", mobileNumber).Scan(&peopleObjectID)
	if sqlErr != nil{
		if sqlErr == sql.ErrNoRows {

			result, err := s.DB.Exec("INSERT INTO people (name, mobile_number) VALUES (?, ?)", name, mobileNumber)
			if err != nil {
				return -1, fmt.Errorf("addPeople: %v", err)
			}
			id, _ := result.LastInsertId()
			return int(id), nil
		}
		return -1, fmt.Errorf("getPeople: %v", sqlErr.Error())
	}
	return peopleObjectID, nil
}
