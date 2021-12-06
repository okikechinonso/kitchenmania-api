package server

import "kitchenmaniaapi/infrastructure/persistence"

type Server struct{
	DB persistence.DbInterface
}

