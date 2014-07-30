// Copyright 2013, Amahi.  All rights reserved.
// Use of this source code is governed by the
// license that can be found in the LICENSE file.

// Functions for detecting hint

package metadata

import (
	"strings"
)

func (l *Library) detectTypeFromHint(MediaName string, Hint string) (string, error) {
	tags := strings.Split(Hint, " ")
	Hint = ""
	for _, tg := range tags {
		if strings.Contains(tg, "movie") {
			Hint = "movie"
		}
		if strings.Contains(tg, "tv") {
			if Hint == "" {
				Hint = "tv"
			}
		}
	}
	return Hint, nil
}
