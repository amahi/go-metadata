// Copyright 2013, Amahi.  All rights reserved.
// Use of this source code is governed by the
// license that can be found in the LICENSE file.

// Data structures and types for the metadata library

package metadata

import (
	"encoding/xml"
)

const tmdb_apikey string = "dc10d9b00f8a4a777539655342cbb647"
const tvrage_apikey string = "L91ezivwoeU8DymX3Wtc"
const tvdb_apikey string = "89DA7ABD734DD427"

type Library struct {
	max_size     int
	current_size int
	dbpath       string
}

//response of search/multi
type TmdbResponse struct {
	Page          int
	Results       []TmdbResult
	Total_pages   int
	Total_results int
}

//results format from Tmdb
type TmdbResult struct {
	Adult          bool
	Name           string
	Backdrop_path  string
	Id             int
	Original_name  string
	Original_title string
	First_air_date string
	Release_date   string
	Poster_path    string
	Title          string
	Media_type     string
	Profile_path   string
}

//response of config
type TmdbConfig struct {
	Images ImageConfig
}

//Image configurtion
type ImageConfig struct {
	Base_url        string
	Secure_base_url string

	//possible sizes for images
	Backdrop_sizes []string
	Logo_sizes     []string
	Poster_sizes   []string
	Profile_sizes  []string
	Still_sizes    []string
}

//Movie metadata structure
type MovieMetadata struct {
	Id            int
	Media_type    string
	Backdrop_path string
	Poster_path   string
	Credits       TmdbCredits
	Config        TmdbConfig
	Imdb_id       string
	Overview      string
	Title         string
	Release_date  string
}

type TmdbCredits struct {
	Id   int
	Cast []TmdbCast
	Crew []TmdbCrew
}

type TmdbCast struct {
	Character    string
	Name         string
	Profile_path string
}

type TmdbCrew struct {
	Department   string
	Name         string
	Job          string
	Profile_path string
}

type tvrageResult struct {
	XMLName     xml.Name     `xml:"Results"`
	ShowDetails []tvrageShow `xml:"show"`
}

type tvrageShow struct {
	Id   int    `xml:"showid"`
	Name string `xml:"name"`
}

type tvdbResult struct {
	XMLName       xml.Name      `xml:"Data"`
	SeriesDetails []tvdbDetails `xml:"Series"`
}

type tvdbDetails struct {
	SeriesId string `xml:"seriesid"`
	Language string `xml:"language"`
	Name     string `xml:"SeriesName"`
}

type tvMetadata struct {
	Media_type string
	SeriesName string `xml:"Series>SeriesName"`
	Banner_Url string
	Actors     string `xml:"Series>Actors"`
	Overview   string `xml:"Series>Overview"`
	Banner     string `xml:"Series>banner"`
	FanArt     string `xml:"Series>fanart"`
	Poster     string `xml:"Series>poster"`
	Rating     string `xml:"Series>Rating"`
	FirstAired string `xml:"Series>FirstAired"`
}

type filtered_output struct {
	Media_type   string `json:"media_type"`
	Title        string `json:"title"`
	Artwork      string `json:"artwork"`
	Release_date string `json:"year"`
}
