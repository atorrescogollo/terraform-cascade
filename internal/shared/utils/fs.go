package utils

import "io/fs"

type DirEntry struct {
	name  string
	isDir bool
}

func (d DirEntry) Name() string {
	return d.name
}
func (d DirEntry) IsDir() bool {
	return d.isDir
}
func (d DirEntry) Type() fs.FileMode {
	if d.isDir {
		return fs.ModeDir
	}
	return fs.ModePerm
}
func (d DirEntry) Info() (fs.FileInfo, error) {
	return nil, nil
}

func NewDirEntry(name string, isDir bool) fs.DirEntry {
	return &DirEntry{
		name:  name,
		isDir: isDir,
	}
}
