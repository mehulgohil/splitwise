package service

import "database/sql"

type ServiceStruct struct {
	DB *sql.DB
}