package utils

import "io/fs"

// DirEntry represents a directory entry (fs.DirEntry)
type DirEntry struct {
	name  string
	isDir bool
}

// Name overrides fs.DirEntry.Name
func (d DirEntry) Name() string {
	return d.name
}

// IsDir overrides fs.DirEntry.IsDir
func (d DirEntry) IsDir() bool {
	return d.isDir
}

// Type overrides fs.DirEntry.Type
func (d DirEntry) Type() fs.FileMode {
	if d.isDir {
		return fs.ModeDir
	}
	return fs.ModePerm
}

// Info overrides fs.DirEntry.Info
func (d DirEntry) Info() (fs.FileInfo, error) {
	return nil, nil
}

// NewDirEntry DirEntry constructor
func NewDirEntry(name string, isDir bool) fs.DirEntry {
	return &DirEntry{
		name:  name,
		isDir: isDir,
	}
}
