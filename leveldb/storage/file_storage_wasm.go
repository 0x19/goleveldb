// Copyright (c) 2012, Suryandaru Triandana <syndtr@gmail.com>
// All rights reserved.
//
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

//go:build js || wasip1
// +build js wasip1

package storage

import (
	"os"
)

type unixFileLock struct {
	f *os.File
}

func (fl *unixFileLock) release() error {
	panic("wasm not supported")
}

func newFileLock(path string, readOnly bool) (fl fileLock, err error) {
	panic("wasm not supported")
}

func setFileLock(f *os.File, readOnly, lock bool) error {
	panic("wasm not supported")
}

func rename(oldpath, newpath string) error {
	panic("wasm not supported")
}

func isErrInvalid(err error) bool {
	panic("wasm not supported")
}

func syncDir(name string) error {
	panic("wasm not supported")
}
