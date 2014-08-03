// Copyright 2013, Amahi.  All rights reserved.
// Use of this source code is governed by the
// license that can be found in the LICENSE file.

// Functions for prefilling cache

package metadata

import (
	"os"
	"path/filepath"
	"time"
)

//prefill database cache using directory root, Hint, sleeptime after each read and a boolean denoting whether to prefill directory names
func (l *Library) Prefill(root string, Hint string, sleeptime time.Duration, includeDir bool) error {
	filepath.Walk(root, l.getWalkFunc(Hint, sleeptime, includeDir))
	return nil
}

//return function to be used on each file in the share for Tv/Movie
func (l *Library) getWalkFunc(Hint string, sleeptime time.Duration, includeDir bool) func(path string, info os.FileInfo, err error) error {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}
		if info.IsDir() != includeDir {
			return nil
		}
		//no errors returned
		l.GetMetadata(info.Name(), Hint)
		time.Sleep(sleeptime)
		return nil
	}
}
