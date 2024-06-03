package db

import (
	"database/sql"

	"github.com/federicotorres233/gokeys/internal/db/query"
	"github.com/federicotorres233/gokeys/internal/types"
)

func AddRecord(db *sql.DB, record types.Record) error {
	return query.AddRecord(db, record)
}

func GetRecordByWebsite(db *sql.DB, website string) (types.Record, error) {
	return query.QueryByWebsite(db, website)
}
