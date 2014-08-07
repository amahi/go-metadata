// Copyright 2013, Amahi.  All rights reserved.
// Use of this source code is governed by the
// license that can be found in the LICENSE file.

// Golang library for requesting and caching Movies and TV metadata
package metadata

import (
	"database/sql"
	"errors"
	_ "github.com/mattn/go-sqlite3"
	"github.com/amahi/go-themoviedb"
	"os"
)

// Initiate the Libray with a valid database path and size.
// Size must not change on subsequent calls
func Init(sz int, dbpath, tmdb_api string) (*Library, error) {
	db, err := sql.Open("sqlite3", dbpath)
	if err != nil {
		return &Library{}, err
	}
	defer db.Close()

	_, err = os.Open(dbpath)
	if err != nil {
		sql := `
	        create table metadata (filename text not null primary key, data text not null, type text, timestamp integer not null);
	        `
		_, err = db.Exec(sql)
		if err != nil {
			return &Library{}, err
		}
	}
	rows, err := db.Query("SELECT COUNT(filename) FROM metadata;")
	if err != nil {
		return &Library{}, err
	}

	defer rows.Close()

	cs := 0
	for rows.Next() {
		err = rows.Scan(&cs)
	}

	tmdb := tmdb.Init(tmdb_api)

	return &Library{max_size: sz, dbpath: dbpath, current_size: cs, tmdb: tmdb}, nil

}

// Get metadata for TV shows and movies from a given medianame (a
// filename, typically) and hint as to whether the media is "tv" or
// "movie". The hint can also be a space separated list of tags that
// contain keywords like "tv" and "movie". First match for tv or movies
// (for now) wins.
func (l *Library) GetMetadata(media_name string, Hint string) (json string, err error) {
	var met string
	met, typ, err := l.cache_lookup(media_name)
	if err == nil {
		if typ == "tv" {
			res, err := filterTvData(met)
			return res, err
		} else {
			res, err := l.tmdb.ToJSON(met)
			return res, err
		}

	}
	processed_string, mediatype, err := l.preprocess(media_name, Hint)
	if err != nil {
		return "{}", err
	}
	if mediatype == "tv" {
		met, err := getTvData(processed_string)
		if err != nil {
			return "{}", err
		}
		err = l.add_to_cache(media_name, met, "tv")
		met, err = filterTvData(met)
		if err != nil {
			return "{}", err
		}
		return met, nil
	} else if mediatype == "movie" {
		met, err := l.tmdb.MovieData(processed_string)
		if err != nil {
			return "{}", err
		}
		err = l.add_to_cache(media_name, met, "movie")
		met, err = l.tmdb.ToJSON(met)
		if err != nil {
			return "{}", err
		}
		return met, nil
	}
	return "{}", errors.New("No Results. MediaType should be tv or movie. Metadata library got: " + mediatype)
}
