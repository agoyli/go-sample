package store

import "main/internal/store/pgx"

var store IStore

func Store() IStore {
	return store
}

func Connect() IStore {
	store = pgx.Connect()
	return store
}
