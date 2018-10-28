// Copyright 2016 David Lavieri.  All rights reserved.
// Use of this source code is governed by a MIT License
// License that can be found in the LICENSE file.

package goradix

import "bytes"

func buildWords(rt *Radix, bs, strip []byte, words chan<- []byte, ww bool) {
	var npath []byte
	npath = append(bs, rt.Path...)

	if len(rt.nodes) > 0 {
		if rt.key {
			addWord(rt, npath, strip, ww, words)
		}

		for _, n := range rt.nodes {
			buildWords(n, npath, strip, words, ww)
		}
	} else if rt.key {
		addWord(rt, npath, strip, ww, words)
	}
}

func addWord(rt *Radix, npath, strip []byte, ww bool, words chan<- []byte) {
	if w := bytes.Replace(npath, strip, nil, 1); len(w) > 0 {
		if !ww {
			words <- w
		} else if _, matches, _ := rt.match(strip); matches != len(rt.Path) && ww {
			words <- w
		}
	}
}

func buildWordsWorker(inWords <-chan []byte, outWords chan<- [][]byte) {
	var wordSlice [][]byte

	for v := range inWords {
		wordSlice = append(wordSlice, v)
	}

	outWords <- wordSlice
}
