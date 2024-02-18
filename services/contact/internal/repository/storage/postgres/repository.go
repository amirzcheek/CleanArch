package postgres

import (
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Repository struct {
	db     *pgxpool.Pool
	genSql squirrel.StatementBuilderType
}

func New(db *pgxpool.Pool) *Repository {
	var rep = &Repository{
		db:     db,
		genSql: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}
	return rep
}
