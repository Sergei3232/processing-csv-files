package postgres

import (
	"database/sql"
	sq "github.com/Masterminds/squirrel"
	_ "github.com/lib/pq"
)

type DataSCV struct {
	Id            int64
	IdFileStorage int64
	Height        int64
	Width         int64
}

type PostgresSQLDB struct {
	*sql.DB
	qb sq.StatementBuilderType
}

func NewDBConnect(sqlConnect string) (*PostgresSQLDB, error) {
	bd, err := sql.Open("postgres", sqlConnect) //postgres
	if err != nil {
		return &PostgresSQLDB{}, err
	}

	return &PostgresSQLDB{bd, sq.StatementBuilder.PlaceholderFormat(sq.Dollar)}, nil
}
