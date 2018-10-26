// Copyright 2016 David Lavieri.  All rights reserved.
// Use of this source code is governed by a MIT License
// License that can be found in the LICENSE file.

package main

import "github.com/falmar/goradix"

func main() {
	radix := goradix.New(false)

	radix.Insert("romanus")
	radix.Insert("romane")
	radix.Insert("romulus")
}
