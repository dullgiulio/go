/*
Copyright 2013 the Camlistore Authors.

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

package punchhole

import (
	"os"
	"syscall"
)

const (
	// FALLOC_FL_KEEP_SIZE: default is extend size
	falloc_fl_keep_size = 0x01

	// FALLOC_FL_PUNCH_HOLE: de-allocates range
	falloc_fl_punch_hole = 0x02
)

func init() {
	PunchHole = punchHoleLinux
}

// puncHoleLinux punches a hole into the given file starting at offset,
// measuring "size" bytes
func punchHoleLinux(file *os.File, offset int64, size int64) error {
	err := syscall.Fallocate(int(file.Fd()),
		falloc_fl_punch_hole|falloc_fl_keep_size,
		offset, size)
	if err == syscall.ENOSYS || err == syscall.EOPNOTSUPP {
		return errNoPunch
	}
	return err
}
