// Copyright 2013, Amahi.  All f reserved.
// Use of this source code is governed by the
// license that can be found in the LICENSE file.

// Functions for getting Tv metadata

package metadata

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"io/ioutil"
	"net/http"
)

func getTvData(MediaName string) (string, error) {
	details, err := getSeriesDetails(MediaName)
	if err != nil {
		return "", err
	}
	tvmetadata, err := getTvMetadata(details)
	if err != nil {
		return "", err
	}
	data, err := json.Marshal(tvmetadata)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func getUsableTvName(MediaName string) (string, error) {
	res, err := http.Get("http://services.tvrage.com/myfeeds/search.php?key=" + TVRAGE_APIKEY + "&show=" + MediaName)
	if err != nil {
		return MediaName, err
	}
	body, err := ioutil.ReadAll(res.Body)
	var result tvrageResult
	err = xml.Unmarshal(body, &result)
	if err != nil {
		return MediaName, err
	}
	if result.ShowDetails == nil {
		return MediaName, errors.New("No result obtained from tvrage for filename string correction")
	} else {

		return result.ShowDetails[0].Name, nil
	}
	return MediaName, nil
}

func getSeriesDetails(MediaName string) (tvdbDetails, error) {
	var det tvdbDetails
	res, err := http.Get(gettvdbMirrorPath() + "api/GetSeries.php?seriesname=" + MediaName)
	if err != nil {
		return det, err
	}
	body, err := ioutil.ReadAll(res.Body)
	var results tvdbResult
	err = xml.Unmarshal(body, &results)
	if err != nil {
		return det, err
	}
	if results.SeriesDetails == nil {
		return det, errors.New("No result obtained from tvdb")
	}
	det = results.SeriesDetails[0]
	return det, nil
}

func getTvMetadata(Details tvdbDetails) (tvMetadata, error) {
	var met tvMetadata
	res, err := http.Get(gettvdbMirrorPath() + "api/" + TVDB_APIKEY + "/series/" + Details.SeriesId + "/all/" + Details.Language + ".xml")
	if err != nil {
		return met, err
	}
	body, err := ioutil.ReadAll(res.Body)
	err = xml.Unmarshal(body, &met)
	if err != nil {
		return met, err
	}
	met.Banner_Url = gettvdbMirrorPath() + "banners/"
	met.Media_type = "tv"
	return met, nil
}
func gettvdbMirrorPath() string {
	return "http://thetvdb.com/"
}
