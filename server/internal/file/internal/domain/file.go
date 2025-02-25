package domain

import "github.com/chenmingyong0423/go-mongox/v2"

type File struct {
	mongox.Model `bson:",inline"`
	Type         string
	Url          string
	Dst          string
}
