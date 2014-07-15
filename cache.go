// Copyright 2013, Amahi.  All rights reserved.
// Use of this source code is governed by the
// license that can be found in the LICENSE file.

// Functions for implementing cache

package metadata

import (
	"database/sql"
	"errors"
	_ "github.com/mattn/go-sqlite3"
	"os"
)

//FIXME - implement a cache here
func (l *Library) cache_lookup(MediaName string) (result string, mediatype string, err error) {
	db, err := sql.Open("sqlite3", l.dbpath)
	if err != nil {
		return "", "", err
	}
	defer db.Close()

	_, err = os.Open(l.dbpath)
	if err != nil {
		return "", "", err
	}

	rows, err := db.Query("SELECT data,type FROM metadata WHERE filename=?", MediaName)
	if err != nil {
		return "", "", err
	}

	defer rows.Close()

	for rows.Next() {
		var data string
		var mediatype string
		err = rows.Scan(&data, &mediatype)
		return data, mediatype, nil
	}

	return "", "", errors.New("No Data found in cache")
}

func (l *Library) add_to_cache(MediaName string, content string, mediatype string) error {
	db, err := sql.Open("sqlite3", l.dbpath)
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = os.Open(l.dbpath)
	if err != nil {
		sql := `
	        create table metadata (filename text not null primary key, data text not null, type text);
	        `
		_, err = db.Exec(sql)
		if err != nil {
			return err
		}
	}

	tx, err := db.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare("insert into metadata(filename, data, type) values(?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(MediaName, content, mediatype)
	tx.Commit()
	l.current_size++
	return nil
}
