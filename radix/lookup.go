// Copyright 2016 David Lavieri.  All rights reserved.
// Use of this source code is governed by a MIT License
// License that can be found in the LICENSE file.

package goradix

// ----------------------- Look Up ------------------------ //

// LookUp will return the node matching
func (r *Radix) LookUp(s string) (interface{}, error) {
	return r.LookUpBytes([]byte(s))
}

// LookUpBytes will return the node matching
func (r *Radix) LookUpBytes(bs []byte) (interface{}, error) {
	node, key, err := r.sLookUp(bs)

	if err != nil {
		return nil, err
	}

	if !key {
		return nil, ErrNoMatchFound
	}

	return node.get(), err
}

func (r *Radix) sLookUp(bs []byte) (*Radix, bool, error) {
	var traverseNode = r

	traverseNode.rLock()

	lbs, matches, _ := traverseNode.match(bs)

	// && ((!r.master && matches > 0) || r.master)
	if matches == len(traverseNode.Path) {
		if matches < len(bs) {
			for _, n := range traverseNode.nodes {
				if tn, nkey, err := n.sLookUp(lbs); tn != nil {
					traverseNode.rUnlock()

					return tn, nkey, err
				}
			}

			traverseNode.rUnlock()

			// Do not jump back to parent node
			return nil, false, ErrNoMatchFound
		}

		traverseNode.rUnlock()

		return traverseNode, traverseNode.key, nil
	}

	traverseNode.rUnlock()

	return nil, false, ErrNoMatchFound
}
