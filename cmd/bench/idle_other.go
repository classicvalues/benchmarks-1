// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !linux
// +build !linux

package main

// waitForIdle is only implemented for Linux.
func waitForIdle() error {
	return nil
}
