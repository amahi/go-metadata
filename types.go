// Copyright 2013, Amahi.  All rights reserved.
// Use of this source code is governed by the
// license that can be found in the LICENSE file.

// Data structures and types for the metadata library

package metadata

import (
	"github.com/amahi/go-themoviedb"
	"github.com/amahi/go-tvrage"
)

type Library struct {
	max_size     int
	current_size int
	dbpath       string
	tmdb         *tmdb.TMDb
	tvrage       *tvrage.TVRage
}
