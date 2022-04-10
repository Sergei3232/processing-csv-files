package postgres

import (
	"database/sql"
	sq "github.com/Masterminds/squirrel"
	_ "github.com/lib/pq"
)

type PostgresSQL interface {
	GettingIdImageFileStorage(mapIdItems map[int64]DataSCV)
	GetImageHightWidth(mapIdItems map[int64]DataSCV) error
}

type DataSCV struct {
	Id            int64
	IdFileStorage int64
	Height        int64
	Width         int64
}

type SQLDB struct {
	*sql.DB
	qb sq.StatementBuilderType
}

func NewDBConnect(sqlConnect string) (*SQLDB, error) {
	bd, err := sql.Open("postgres", sqlConnect) //postgres
	if err != nil {
		return &SQLDB{}, err
	}

	return &SQLDB{bd, sq.StatementBuilder.PlaceholderFormat(sq.Dollar)}, nil
}
