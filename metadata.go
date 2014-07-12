// Copyright 2013, Amahi.  All rights reserved.
// Use of this source code is governed by the
// license that can be found in the LICENSE file.

// Functions for requesting Metadata

package metadata

import ()

//get metadata for Tv shows and movies from a given medianame and hint as to whether the media is "tv" or "movie"
func GetMetadata(MediaName string, Hint string) (json string, err error) {
	var met string
	met, err = cache_lookup(MediaName)
	if err == nil {
		return met, err
	}
	processed_string, mediatype, err := preprocess(MediaName, Hint)
	if err != nil {
		return met, err
	}
	if mediatype == "tv" {
		met, err := getTvData(processed_string)
		if err != nil {
			return "", err
		}
		add_to_cache(MediaName, met)
		return met, err
	} else if mediatype == "movie" {
		met, err := getMovieData(processed_string)
		if err != nil {
			return "", err
		}
		add_to_cache(MediaName, met)
		return met, err
	}
	return met, err
}
