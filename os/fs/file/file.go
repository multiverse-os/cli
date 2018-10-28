package file

import (
	"time"
)

// TODO: Implement a radix-trie (patricia) prefix-sorted storage of file
// collections. This can help us do auto-completion, and a variety of other
// useful features.

type File struct {
	name         string
	path         Path
	permissions  int
	owner        string
	group        string
	data         []byte
	createdAt    time.Time
	lastModified time.Time
}

// Name returns the name of the file as presented to Open.
func (self *File) Name() string { return self.name }

func (self *File) IsExecutable() bool { return false }
func (self *File) IsHidden() bool     { return false }

func (self *File) Size() {}
func (self *File) Type() {} // MIME Type, based off magic numbers, extension and more

func (self *File) Create()      {}
func (self *File) Open()        {}
func (self *File) Rename()      {}
func (self *File) Move()        {}
func (self *File) Remove()      {}
func (self *File) ChangeOwner() {}
func (self *File) ChangeGroup() {}

func (self *File) Head(lines int) {}
func (self *File) Tail(lines int) {}

func (self *File) Checksum() {}

func (self *File) Compress()   {}
func (self *File) Decompress() {}
