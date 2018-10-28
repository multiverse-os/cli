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

	value, err := radix.LookUp("romane")

	if err != nil { // No Match Found
		return
	}

	// Output: Found node, Value: 100
	fmt.Println("Found node, Value: ", value)
}
