package x14nfile

import "time"

type File struct {
	Name         string
	Path         string
	Size         int64
	CreateTime   time.Time
	ModifyTime   time.Time
	AccessedTime time.Time
	IsDirectory  bool
}
