// Copyright 2016 David Lavieri.  All rights reserved.
// Use of this source code is governed by a MIT License
// License that can be found in the LICENSE file.

package goradix

import "sync"

// ----------------------- Inserts ------------------------ //

// Insert new string to the Radix Tree
func (r *Radix) Insert(s string, v ...interface{}) bool {
	return r.InsertBytes([]byte(s), v...)
}

// InsertBytes to the Radix Tree
func (r *Radix) InsertBytes(bs []byte, val ...interface{}) bool {
	var value interface{}

	if len(val) > 0 {
		value = val[0]
	}

	r.rLock()

	if len(r.Path) == 0 && len(r.nodes) == 0 {
		r.rUnlock()
		r.lock()
		r.Path = bs
		r.set(value)
		r.leaf = true
		r.key = true
		r.unlock()
		return true
	}

	var pathLen = len(r.Path)
	var bsLen = len(bs)

	match := 0
	i := 0
	var v byte

	for i, v = range r.Path {
		if i >= bsLen {
			// No more matches to check
			r.rUnlock()
			return false
		}

		if v == bs[i] && match == i {
			// continue as long it match the path
			match++
			continue
		}

		if v != bs[i] && match > 0 {
			// If the byte string does not match anymore but had
			// previous matches to the path then add the byte string
			// as children node
			r.rUnlock()
			r.lock()
			var prs = true

			if r.nodes == nil {
				// If there is no existing nodes then slice the path
				// until the last occurrence, add what is left of the path as
				// children and also add the byte string.
				r.addChildren(r.Path[i:], r.getNonBlocking(), nil, r.key)
				r.key = false
				r.set(nil)
				r.addChildren(bs[i:], value, nil, true)
				r.Path = r.Path[:i]
			} else {
				// Otherwise just add the new byte string as
				prs = r.pushChildren(bs, value, i, false)
			}

			r.unlock()

			return prs
		}
	}

	if match > 0 {
		// Check if it already exists
		if match == pathLen && pathLen == bsLen {
			r.rUnlock()
			r.lock()

			if value != nil {
				r.set(value)
			}
			r.key = true

			r.unlock()

			return true
		}

		r.rUnlock()
		r.lock()

		// If it matches all current node path and the byte string
		for _, c := range r.nodes {
			if c.InsertBytes(bs[i+1:], value) {
				r.unlock()

				return true
			}
		}

		// no match found on nodes
		r.addChildren(bs[i+1:], value, nil, true)

		r.unlock()

		return true
	}

	if r.master {
		// If there is NO match and the current node is the master Radix

		if r.Path != nil {
			r.rUnlock()
			r.lock()

			prs := r.pushChildren(bs, value, i, true)

			r.unlock()

			return prs
		}

		r.rUnlock()
		r.lock()

		for _, c := range r.nodes {
			if c.InsertBytes(bs, value) {
				r.unlock()
				return true
			}
		}

		// no match found on children nodes
		// add new byte string as node
		r.addChildren(bs, value, nil, true)

		r.unlock()
		return true
	}

	r.rUnlock()

	return false
}

// Add children node to the current Radix Tree node
func (r *Radix) addChildren(bs []byte, v interface{}, c []*Radix, k bool) {
	r.leaf = false
	var setLeaf bool

	if c == nil {
		setLeaf = true
	}

	r.nodes = append(r.nodes, &Radix{
		Path:  bs,
		nodes: c,
		value: v,
		leaf:  setLeaf,
		key:   k,
		mu:    &sync.RWMutex{},
		ts:    r.ts,
	})
}

// Push the current children nodes to a new node with the path
// of what is left from slicing of the current path
// and add the new byte string as children node
func (r *Radix) pushChildren(bs []byte, v interface{}, i int, master bool) bool {
	if (len(r.Path)) < i {
		return false
	}

	nodes := r.nodes
	r.nodes = nil
	r.addChildren(r.Path[i:], r.getNonBlocking(), nodes, r.key)
	r.key = false
	r.set(nil)

	if master {
		r.Path = nil
		r.addChildren(bs, v, nil, true)
	} else {
		r.Path = r.Path[:i]
		r.addChildren(bs[i:], v, nil, true)
	}

	return true
}
