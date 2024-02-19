package postgres

import (
	"database/sql"
	"github.com/Masterminds/squirrel"
)

type Repository struct {
	db     *sql.DB
	genSql squirrel.StatementBuilderType
}

func New(db *sql.DB) *Repository {
	var rep = &Repository{
		db:     db,
		genSql: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}
	return rep
}
