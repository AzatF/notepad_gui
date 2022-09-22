package database

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"notepad/config"
	"notepad/pkg/logging"
	"path"
	"strings"
)

type DB struct {
	db     *sql.DB
	logger *logging.Logger
}

var TitleList []string

type TableStr struct {
	ID    int
	Title string
	Text  string
}

func NewDataBase(cfg *config.Config, logger *logging.Logger) (DB, error) {

	sqlite, err := sql.Open("sqlite3", path.Join(cfg.DataPath, "notes.db"))
	_, _ = sqlite.Exec(fmt.Sprintf("CREATE TABLE IF NOT EXISTS note (id INTEGER NOT NULL PRIMARY KEY, title VARCHAR (30), text TEXT DEFAULT ('undefined'))"))
	if err != nil {
		logger.Error(err)
	}

	return DB{
		db:     sqlite,
		logger: logger,
	}, nil
}

func (d *DB) CreateNewTable(newTitle string) error {

	d.logger.Infof("CREATING NEW NOTE %s", newTitle)

	if newTitle == "" {
		return nil
	}

	s, _ := d.GetTitles()
	for _, v := range s {
		if v == newTitle {
			return nil
		}
	}

	_, err := d.db.Exec(fmt.Sprintf("INSERT INTO note (title) VALUES ('%s')", newTitle))
	if err != nil {
		d.logger.Error(err)
		return err
	}

	return nil
}

func (d *DB) GetNote(title string) (text string, err error) {

	query, err := d.db.Query(fmt.Sprintf("SELECT text FROM note WHERE title = ('%s')", strings.TrimSpace(title)))
	if err != nil {
		d.logger.Error(err)
		return "", err
	}

	for query.Next() {
		err = query.Scan(&text)
		if text == "undefined" {
			text = ""
		}
		if err != nil {
			d.logger.Error(err)
		}
	}

	return

}

func (d *DB) GetTitles() (titles []string, err error) {

	var ss TableStr
	query, err := d.db.Query("SELECT title FROM note")
	if err != nil {
		return nil, err
	}

	for query.Next() {
		err = query.Scan(&ss.Title)
		titles = append(titles, ss.Title)
		if err != nil {
			d.logger.Error(err)
		}
	}

	TitleList = titles

	return

}

func (d *DB) WriteNote(title, text string) error {

	_, err := d.db.Exec(fmt.Sprintf("UPDATE note SET (text) = ('%s') WHERE title = ('%s')", text, title))
	if err != nil {
		d.logger.Error(err)
		return err
	}

	return nil
}
