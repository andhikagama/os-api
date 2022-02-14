package migrations

import (
	"database/sql"
	"fmt"

	"github.com/andhikagama/os-api/shared"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type (
	Migration struct {
		resource shared.Resource
	}
)

func New(resource shared.Resource) *Migration {
	return &Migration{
		resource: resource,
	}
}

func (m *Migration) Up() error {
	db, _ := sql.Open("mysql",
		fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?multiStatements=true", m.resource.Config.MySQLUser,
			m.resource.Config.MySQLPassword, m.resource.Config.MySQLHost, m.resource.Config.MySQLPort, m.resource.Config.MySQLDatabase))
	driver, _ := mysql.WithInstance(db, &mysql.Config{})
	mg, _ := migrate.NewWithDatabaseInstance(
		"file://./migrations/sql",
		"mysql",
		driver,
	)

	return mg.Up()
}
