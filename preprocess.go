// Copyright 2013, Amahi.  All rights reserved.
// Use of this source code is governed by the
// license that can be found in the LICENSE file.

// Functions for preprocess input

package metadata

import (
	"errors"
	"strconv"
	"strings"
)

//Convert non-standard format strings to standard format
func (l *Library) preprocess(MediaName string, Hint string) (title string, mediatype string, err error) {
	//Send for Hint Detection
	Hint, err = l.detectTypeFromHint(MediaName, Hint)
	if err != nil {
		return MediaName, Hint, errors.New("Media Type Unknown")
	}
	if Hint == "movie" {
		//strip off year name
		parts := strings.Split(MediaName, "(")
		result := parts[0]
		for _, s := range parts[1:] {
			yearparts := strings.Split(s, ")")
			l := len(yearparts)
			if l > 1 {
				result += " " + yearparts[l-1]
			} else {
				result += " " + s
			}
		}

		//strip off the extension and full-stops
		parts = strings.Split(result, ".")
		if len(parts) > 1 {
			result = parts[0]
			for _, s := range parts[1 : len(parts)-1] {
				result += " " + s
			}
			if len(parts[len(parts)-1]) > 4 {
				result += parts[len(parts)-1]
			}
		}

		//replace spaces by %20
		parts = strings.Split(result, " ")
		result = parts[0]
		a, err := strconv.Atoi(parts[0])
		if err == nil {
			//remove year
			if a > 999 && a < 3000 {
				result = ""
			}

		}
		if len(parts) > 1 {
			for _, s := range parts[1:] {
				a, err = strconv.Atoi(s)
				if err == nil {
					//remove year
					if a <= 999 || a >= 3000 {
						result += "%20" + s
					}
				} else {
					result += "%20" + s
				}
			}
		}

		return result, Hint, nil
	} else if Hint == "tv" {
		//replace spaces by %20
		parts := strings.Split(MediaName, " ")
		result := parts[0]
		if len(parts) > 1 {
			for _, s := range parts[1:] {
				result += "%20" + s
			}
		}

		//strip off the full-stops
		parts = strings.Split(result, ".")
		if len(parts) > 1 {
			result = parts[0]
			for _, s := range parts[1:] {
				result += "%20" + s
			}
		}

		//strip off the underscores
		parts = strings.Split(result, "_")
		if len(parts) > 1 {
			result = parts[0]
			for _, s := range parts[1:] {
				result += "%20" + s
			}
		}
		result, _ = getUsableTvName(result)

		//replace spaces by %20
		parts = strings.Split(result, " ")
		result = parts[0]
		if len(parts) > 1 {
			for _, s := range parts[1:] {
				result += "%20" + s
			}
		}

		return result, Hint, nil
	}

	return MediaName, Hint, errors.New("Media Type Unknown")
}
