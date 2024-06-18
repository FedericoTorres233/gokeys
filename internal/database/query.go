package db

import (
	"database/sql"

	"github.com/federicotorres233/gokeys/internal/database/query"
	"github.com/federicotorres233/gokeys/internal/types"
)

func AddRecord(db *sql.DB, record *types.Record) error {
	return query.AddRecord(db, record)
}

func GetRecordByWebsite(db *sql.DB, record *types.Record) error {
	return query.QueryByWebsite(db, record)
}
