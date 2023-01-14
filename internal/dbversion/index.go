package dbversion

import (
	"database/sql"
	"fmt"
	"path/filepath"

	"github.com/golang-migrate/migrate/v4"
	sqlite3 "github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattn/go-sqlite3"
	"github.com/sirupsen/logrus"
)

func StartVersion(dataPath string) error {
	if dataPath == "" {
		dataPath = "."
	}
	dbPath := fmt.Sprintf("%s%smetadata.db", dataPath, string(filepath.Separator))
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return err
	}
	defer db.Close()
	driver, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {
		return err
	}
	m, err := migrate.NewWithDatabaseInstance("file://migrations", "sqlite3", driver)
	if err != nil {
		return err
	}
	version, _, err := m.Version()
	if err != nil && err == migrate.ErrNilVersion {
		return m.Up()
	}
	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		return err
	}
	logrus.Infof("current database version is %d", version)
	return nil
}
