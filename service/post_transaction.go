package service

import (
	"fmt"
	"github.com/mehulgohil/splitwise/models"
)

func (s *ServiceStruct) PostTransactionToDB() error {
	var allPeople []models.People

	rows, err := s.DB.Query("SELECT * FROM people WHERE id = 1")
	if err != nil {
		return err
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var alb models.People
		if err := rows.Scan(&alb.ID, &alb.Name, &alb.MobileNumber); err != nil {
			return err
		}
		allPeople = append(allPeople, alb)
	}
	if err := rows.Err(); err != nil {
		return err
	}
	fmt.Println(allPeople)
	return nil
}
