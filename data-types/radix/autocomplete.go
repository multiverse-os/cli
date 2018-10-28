// Copyright 2016 David Lavieri.  All rights reserved.
// Use of this source code is governed by a MIT License
// License that can be found in the LICENSE file.

package goradix

// ----------------------- Autocomplete ------------------------ //

// AutoComplete will complete the word you are looking for
// and return them as string
func (r Radix) AutoComplete(s string, wholeWord bool) ([]string, error) {
	var stringWords []string
	byteWords, err := r.AutoCompleteBytes([]byte(s), wholeWord)

	if err == nil {
		for _, bs := range byteWords {
			stringWords = append(stringWords, string(bs))
		}
	}

	return stringWords, err
}

// AutoCompleteBytes will complete the word you are looking for
// and return them as slice of bytes
func (r Radix) AutoCompleteBytes(bs []byte, wholeWord bool) ([][]byte, error) {
	node, strip, err := r.acLookUp(bs)
	var ap []byte

	if err != nil {
		return nil, err
	}

	inWord := make(chan []byte)
	outWords := make(chan [][]byte)

	go buildWordsWorker(inWord, outWords)
	if wholeWord {
		ap = bs
	}
	buildWords(node, ap, strip, inWord, wholeWord)
	close(inWord)
	wordSlice := <-outWords
	close(outWords)

	return wordSlice, err
}

func (r *Radix) acLookUp(bs []byte) (*Radix, []byte, error) {
	var traverseNode = r

	lbs, matches, _ := traverseNode.match(bs)

	if matches > 0 || (matches == 0 && traverseNode.master) {
		if matches == len(traverseNode.Path) && matches < len(bs) {
			for _, n := range traverseNode.nodes {
				if tn, nlbs, err := n.acLookUp(lbs); tn != nil {
					return tn, nlbs, err
				}
			}
		}

		if len(lbs) == 0 {
			return traverseNode, bs, nil
		}
	}

	return nil, nil, ErrNoMatchFound
}
