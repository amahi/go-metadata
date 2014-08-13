// Copyright 2013, Amahi.  All rights reserved.
// Use of this source code is governed by the
// license that can be found in the LICENSE file.

// Functions for preprocess input

package metadata

import (
	"errors"
	"strings"
)

//Convert non-standard format strings to standard format
func (l *Library) preprocess(file_name string, hint string) (title string, mediatype string, err error) {
	//Send for hint Detection
	hint, err = l.detectTypeFromHint(file_name, hint)
	if err != nil {
		return file_name, hint, errors.New("Media Type Unknown")
	}
	if hint == "movie" {
		//strip off year name
		parts := strings.Split(file_name, "(")
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

		return strings.TrimSpace(result), hint, nil
	}

	if hint == "tv" {
		// strip off the full-stops
		result := ""
		parts := strings.Split(file_name, ".")
		if len(parts) > 1 {
			result := parts[0]
			for _, s := range parts[1:] {
				result += " " + s
			}
		}

		//strip off the underscores
		parts = strings.Split(result, "_")
		if len(parts) > 1 {
			result = parts[0]
			for _, s := range parts[1:] {
				result += " " + s
			}
		}
		result, _ = l.tvrage.UsableTVName(result)

		return strings.TrimSpace(result), hint, nil
	}

	return file_name, hint, errors.New("Media hint unknown")
}
