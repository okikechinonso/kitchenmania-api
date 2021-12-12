package handlers

import "kitchenmaniaapi/infrastructure/persistence"

type App struct {
	DB persistence.Database
}
