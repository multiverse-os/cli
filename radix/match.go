// Copyright 2016 David Lavieri.  All rights reserved.
// Use of this source code is governed by a MIT License
// License that can be found in the LICENSE file.

package goradix

// ----------------------- Match ------------------------ //

func (r Radix) match(bs []byte) ([]byte, int, []byte) {
	var i int
	var v byte
	var matches int

	for i < len(r.Path) {
		v = r.Path[i]
		if i >= len(bs) {
			break
		}

		if bs[i] == v && matches == i {
			matches++
		} else if bs[i] != v {
			break
		}

		i++
	}

	return bs[i:], matches, r.Path[i:]
}
