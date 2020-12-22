package os

type Path string

type PermissionType int

const (
	UserAccess PermissionType = iota
	GroupAccess
	OtherAccess
	AllAccess
)

type Permission struct {
	Type       PermissionType
	Executable bool
}

// TODO: Output Permission chown and chmod's to apply settings defined in our
//       golang object representing permission/ownership. Also output and
//       marshal the `+x` type, the `7,5,4,0` type, and `rwxwx-wxw-" type
//
//  chmod [ugoa][+-=][rwxXst] fileORdirectoryName
//
// +	Add access
// -	Remove access
// =	Access explicitly assigned

type File struct {
	Path string

	Name       string
	Extension  string
	Executable bool

	Permissions Permissions
	Owner       User
	Group       Group

	// TODO: Should not load this by default, should just
	//       have metadata so everything is lighter weight
	Data []byte
}

type Directory struct {
	Path        Path
	Directories []*Directory
	Files       []*File
}
