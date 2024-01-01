package backend

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
)


type Repository struct {
	*sql.DB
}

func (r *Repository) Connect() error {
	db, err := sql.Open("postgres", os.Getenv("POSTGRE_URL"))
	if err != nil {
		return err
	}
	r.DB = db
	return nil
}
