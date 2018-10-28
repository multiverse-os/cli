// Copyright 2016 David Lavieri.  All rights reserved.
// Use of this source code is governed by a MIT License
// License that can be found in the LICENSE file.

package goradix

// ----------------------- Remove ------------------------ //

// Remove an item, require string
func (r *Radix) Remove(s string) bool {
	return r.RemoveBytes([]byte(s))
}

// RemoveBytes an item, require slice of byte
func (r *Radix) RemoveBytes(bs []byte) bool {
	succeed, _ := r.cRemove(bs)
	return succeed
}

// return (succeed, delete child)

func (r *Radix) cRemove(bs []byte) (bool, bool) {
	r.lock()

	lbs, matches, _ := r.match(bs)

	if matches == len(r.Path) {
		if len(lbs) > 0 {
			for i, c := range r.nodes {
				if sd, dc := c.cRemove(lbs); sd {
					if dc {
						r.removeChild(i)
					}

					r.unlock()

					return sd, false
				}
			}

			r.unlock()

			return false, false
		}

		r.set(nil)
		r.key = false
		isLeft := r.leaf

		if !isLeft {
			r.mergeChild()
		}

		r.unlock()

		return true, isLeft
	}

	r.unlock()

	return false, false
}

func (r *Radix) removeChild(i int) {
	copy(r.nodes[i:], r.nodes[i+1:])
	r.nodes[len(r.nodes)-1] = nil
	r.nodes = r.nodes[:len(r.nodes)-1]

	r.mergeChild()
}

func (r *Radix) mergeChild() {
	if len(r.nodes) == 1 {
		c := r.nodes[0]
		c.mu.RLock()
		r.Path = append(r.Path, c.Path...)
		c.mu.RUnlock()
	}
}
