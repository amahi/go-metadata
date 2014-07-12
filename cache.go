// Copyright 2013, Amahi.  All rights reserved.
// Use of this source code is governed by the
// license that can be found in the LICENSE file.

// Functions for implementing cache

package metadata

import (
	"errors"
)

//FIXME - implement a cache here
func cache_lookup(MediaName string) (string, error) {
	return "", errors.New("Cache not implemented")
}

func add_to_cache(MediaName string, content string) error {
	return errors.New("Cache not implemented")
}
