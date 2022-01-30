package handlers

import "kitchenmaniaapi/infrastructure/persistence/pgsql"

type App struct {
	DB pgsql.DbInterface
}
