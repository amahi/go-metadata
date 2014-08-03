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
	"time"
)

//performs cache lookup and return error if data not found in cache
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

//adds data to cache
func (l *Library) add_to_cache(MediaName string, content string, mediatype string) error {
	if l.current_size > l.max_size {
		return errors.New("size exceeded")
	}
	db, err := sql.Open("sqlite3", l.dbpath)
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = os.Open(l.dbpath)
	if err != nil {
		sql := `
	        create table metadata (filename text not null primary key, data text not null, type text, timestamp integer not null);
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
	stmt, err := tx.Prepare("insert into metadata(filename, data, type, timestamp) values(?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(MediaName, content, mediatype, int(time.Now().Unix()))
	err = tx.Commit()
	if err != nil {
		return err
	}
	l.current_size++

	if l.current_size > l.max_size {
		err = l.removeDB_LeastRecentlyUsed(db)
		if err != nil {
			return err
		}
	}
	return nil
}

// LRU policy implemented here
func (l *Library) removeDB_LeastRecentlyUsed(db *sql.DB) error {

	rows, err := db.Query("SELECT MIN(timestamp) FROM metadata;")
	if err != nil {
		return err
	}

	defer rows.Close()

	var ts int
	for rows.Next() {
		err = rows.Scan(&ts)
		if err != nil {
			return err
		}
	}

	_, err = db.Exec("DELETE from metadata where timestamp=?", ts)
	if err != nil {
		return err
	}

	return nil
}
