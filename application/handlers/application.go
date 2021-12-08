package handlers

import "kitchenmaniaapi/infrastructure/persistence"

type app struct {
	DB persistence.DbInterface
}
