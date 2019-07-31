package model

import "github.com/gogf/gf/g/os/gtime"

type Model struct {
	ID uint
	UpdateAt gtime.Time
	CreateAt gtime.Time
}

