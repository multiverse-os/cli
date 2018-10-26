// Copyright 2016 David Lavieri.  All rights reserved.
// Use of this source code is governed by a MIT License
// License that can be found in the LICENSE file.

package main

import (
	"fmt"

	"github.com/falmar/goradix"
)

func main() {
	radix := goradix.New(false)
	radix.Insert("romanus", 1)
	radix.Insert("romane", 100)
	radix.Insert("romulus", 1000)

	// Return remaining text
	words, err := radix.AutoComplete("ro", false)

	if err != nil { // No Match Found
		return
	}

	// AutoComplete 'rom'; Words: [manus mane mulus]
	fmt.Printf("AutoComplete: '%s'; Words: %v\n", "ro", words)

	// Return whole words
	words, _ = radix.AutoComplete("ro", true)

	// AutoComplete 'rom'; Words: [romanus romane romulus]
	fmt.Printf("AutoComplete: '%s'; Words: %v\n", "ro", words)
}
