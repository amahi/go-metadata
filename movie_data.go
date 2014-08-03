// Copyright 2013, Amahi.  All rights reserved.
// Use of this source code is governed by the
// license that can be found in the LICENSE file.

// Functions for requesting Metadata

package metadata

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
)

var config = TmdbConfig{}

//The main call for getting movie data
func getMovieData(MediaName string) (string, error) {
	var met string
	results, err := searchMovie(MediaName)
	if err != nil {
		return met, err
	}
	if results.Total_results == 0 {
		return met, errors.New("No results found at TMDb")
	}
	if results.Results[0].Media_type == "person" {
		return met, errors.New("Metadata for persons not supported")
	} else if results.Results[0].Media_type == "tv" {
		return met, errors.New("Metadata for tv not supported inside a call for movie data")
	} else {

		movie_details, err := getMovieDetails(strconv.Itoa(results.Results[0].Id))
		if err != nil {
			return met, err
		}
		movie_details.Credits, err = getMovieCredits(strconv.Itoa(results.Results[0].Id))
		if err != nil {
			return met, err
		}
		movie_details.Config, err = getConfig()
		if err != nil {
			return met, err
		}
		movie_details.Id = results.Results[0].Id
		movie_details.Media_type = "movie"

		metadata, err := json.Marshal(movie_details)
		if err != nil {
			return met, err
		}
		met = string(metadata)
		return met, nil
	}
}

//search on TMDb for TV, persons and Movies with a given name
func searchTmdbMulti(MediaName string) (TmdbResponse, error) {
	res, err := http.Get("http://api.themoviedb.org/3/search/multi?api_key=" + tmdb_apikey + "&query=" + MediaName)
	var resp TmdbResponse
	if err != nil {
		return resp, err
	}
	if res.StatusCode != 200 {
		return resp, errors.New("Status Code 200 not recieved from TMDB")
	}
	body, err := ioutil.ReadAll(res.Body)
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return TmdbResponse{}, err
	}
	return resp, nil
}

//search on TMDb for Movies with a given name
func searchMovie(MediaName string) (TmdbResponse, error) {
	res, err := http.Get("http://api.themoviedb.org/3/search/movie?api_key=" + tmdb_apikey + "&query=" + MediaName)
	var resp TmdbResponse
	if err != nil {
		return resp, err
	}
	if res.StatusCode != 200 {
		return resp, errors.New("Status Code 200 not recieved from TMDB")
	}
	body, err := ioutil.ReadAll(res.Body)
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return TmdbResponse{}, err
	}
	return resp, nil
}

//search on TMDb for Tv Shows with a given name
func searchTmdbTv(MediaName string) (TmdbResponse, error) {
	res, err := http.Get("http://api.themoviedb.org/3/search/tv?api_key=" + tmdb_apikey + "&query=" + MediaName)
	var resp TmdbResponse
	if err != nil {
		return resp, err
	}
	if res.StatusCode != 200 {
		return resp, errors.New("Status Code 200 not recieved from TMDb")
	}
	body, err := ioutil.ReadAll(res.Body)
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return TmdbResponse{}, err
	}
	return resp, nil
}

//get configurations from TMDb
func getConfig() (TmdbConfig, error) {
	if config.Images.Base_url == "" {
		res, err := http.Get("http://api.themoviedb.org/3/configuration?api_key=" + tmdb_apikey)
		var conf TmdbConfig
		if err != nil {
			return conf, err
		}
		if res.StatusCode != 200 {
			return conf, errors.New("Status Code 200 not recieved from TMDb")
		}
		body, err := ioutil.ReadAll(res.Body)
		err = json.Unmarshal(body, &conf)
		if err != nil {
			return TmdbConfig{}, err
		}
		config = conf
		return conf, nil
	} else {
		return config, nil
	}
}

//get basic information for movie
func getMovieDetails(MediaId string) (MovieMetadata, error) {
	res, err := http.Get("http://api.themoviedb.org/3/movie/" + MediaId + "?api_key=" + tmdb_apikey)
	var met MovieMetadata
	if err != nil {
		return met, err
	}
	if res.StatusCode != 200 {
		return met, errors.New("Status Code 200 not recieved from TMDb")
	}
	body, err := ioutil.ReadAll(res.Body)
	err = json.Unmarshal(body, &met)
	if err != nil {
		return MovieMetadata{}, err
	}
	return met, nil
}

//get credits for movie
func getMovieCredits(MediaId string) (TmdbCredits, error) {
	res, err := http.Get("http://api.themoviedb.org/3/movie/" + MediaId + "/credits?api_key=" + tmdb_apikey)
	var cred TmdbCredits
	if err != nil {
		return cred, err
	}
	if res.StatusCode != 200 {
		return cred, errors.New("Status Code 200 not recieved from TMDb")
	}
	body, err := ioutil.ReadAll(res.Body)
	err = json.Unmarshal(body, &cred)
	if err != nil {
		return TmdbCredits{}, err
	}
	return cred, nil
}

//get basic information for Tv
func getTmdbTvDetails(MediaId string) (MovieMetadata, error) {
	res, err := http.Get("http://api.themoviedb.org/3/tv/" + MediaId + "?api_key=" + tmdb_apikey)
	var met MovieMetadata
	if err != nil {
		return met, err
	}
	if res.StatusCode != 200 {
		return met, errors.New("Status Code 200 not recieved from TMDb")
	}
	body, err := ioutil.ReadAll(res.Body)
	err = json.Unmarshal(body, &met)
	if err != nil {
		return MovieMetadata{}, err
	}
	return met, nil
}

//get credits for Tv
func getTmdbTvCredits(MediaId string) (TmdbCredits, error) {
	res, err := http.Get("http://api.themoviedb.org/3/tv/" + MediaId + "/credits?api_key=" + tmdb_apikey)
	var cred TmdbCredits
	if err != nil {
		return cred, err
	}
	if res.StatusCode != 200 {
		return cred, errors.New("Status Code 200 not recieved from TMDb")
	}
	body, err := ioutil.ReadAll(res.Body)
	err = json.Unmarshal(body, &cred)
	if err != nil {
		return TmdbCredits{}, err
	}
	return cred, nil
}

//filter out unwanted movie metadata before return to user
func filterMovieData(data string) (string, error) {
	var f filtered_output
	var det MovieMetadata
	err := json.Unmarshal([]byte(data), &det)
	if err != nil {
		return "", err
	}
	f.Title = det.Title
	f.Release_date = det.Release_date
	f.Release_date = f.Release_date[0:4]
	f.Artwork = det.Config.Images.Base_url + "original" + det.Poster_path

	metadata, err := json.Marshal(f)
	if err != nil {
		return "", err
	}
	return string(metadata), nil
}
