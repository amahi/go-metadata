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
func cache_lookup(MediaName string) (string, error) {
	db, err := sql.Open("sqlite3", "./metadata.db")
	if err != nil {
		return "", err
	}
	defer db.Close()

	_, err = os.Open("metadata.db")
	if err != nil {
		return "", err
	}

	rows, err := db.Query("SELECT data FROM metadata WHERE filename=?", MediaName)
	if err != nil {
		return "", err
	}

	defer rows.Close()

	for rows.Next() {
		var data string
		err = rows.Scan(&data)
		return data, nil
	}

	return "", errors.New("No Data found in cache")
}

func add_to_cache(MediaName string, content string) error {
	db, err := sql.Open("sqlite3", "./metadata.db")
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = os.Open("metadata.db")
	if err != nil {
		sql := `
	        create table metadata (filename text not null primary key, data string not null);
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
	stmt, err := tx.Prepare("insert into metadata(filename, data) values(?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(MediaName, content)
	tx.Commit()

	return nil
}
