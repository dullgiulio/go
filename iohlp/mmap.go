/*
Copyright 2014 Tamás Gulácsi

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package iohlp

import (
	"io"
	"os"

	"gopkg.in/inconshreveable/log15.v2"
)

const MaxInt = int64(int(^uint(0) >> 1))

var Log = log15.New("lib", "iohlp")

func init() {
	Log.SetHandler(log15.DiscardHandler())
}

// MmapFile returns the mmap of the given path.
func MmapFile(fn string) ([]byte, io.Closer, error) {
	f, err := os.Open(fn)
	if err != nil {
		return nil, nil, err
	}
	p, closer, err := Mmap(f)
	f.Close()
	return p, closer, err
}
