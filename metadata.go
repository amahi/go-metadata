// Copyright 2013, Amahi.  All rights reserved.
// Use of this source code is governed by the
// license that can be found in the LICENSE file.

// Functions for requesting Metadata

package metadata

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"os"
)

func Init(sz int, dbpath string) (*Library, error) {
	db, err := sql.Open("sqlite3", dbpath)
	if err != nil {
		return &Library{}, err
	}
	defer db.Close()

	_, err = os.Open(dbpath)
	if err != nil {
		sql := `
	        create table metadata (filename text not null primary key, data text not null, type text);
	        `
		_, err = db.Exec(sql)
		if err != nil {
			return &Library{}, err
		}
	}

	return &Library{max_size: sz, dbpath: dbpath, current_size: 0}, nil

}

//get metadata for Tv shows and movies from a given medianame and hint as to whether the media is "tv" or "movie"
func (l *Library) GetMetadata(MediaName string, Hint string) (json string, err error) {
	var met string
	met, typ, err := l.cache_lookup(MediaName)
	if err == nil {
		if typ == "tv" {
			res, err := filterTvData(met)
			return res, err
		} else {
			res, err := filterMovieData(met)
			return res, err
		}

	}
	processed_string, mediatype, err := l.preprocess(MediaName, Hint)
	if err != nil {
		return met, err
	}
	if mediatype == "tv" {
		met, err := getTvData(processed_string)
		if err != nil {
			return "", err
		}
		err = l.add_to_cache(MediaName, met, "tv")
		met, err = filterTvData(met)
		if err != nil {
			return "", err
		}
		return met, nil
	} else if mediatype == "movie" {
		met, err := getMovieData(processed_string)
		if err != nil {
			return "", err
		}
		err = l.add_to_cache(MediaName, met, "movie")
		met, err = filterMovieData(met)
		if err != nil {
			return "", err
		}
		return met, nil
	}
	return met, err
}
