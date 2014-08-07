// Copyright 2013, Amahi.  All rights reserved.
// Use of this source code is governed by the
// license that can be found in the LICENSE file.

// Data structures and types for the metadata library

package metadata

import (
	"encoding/xml"
	"github.com/amahi/go-themoviedb"
)

const tvrage_apikey string = "L91ezivwoeU8DymX3Wtc"
const tvdb_apikey string = "89DA7ABD734DD427"

type Library struct {
	max_size     int
	current_size int
	dbpath       string
	tmdb         *tmdb.TMDb
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
	Title        string `json:"title"`
	Artwork      string `json:"artwork"`
	Release_date string `json:"year"`
}
