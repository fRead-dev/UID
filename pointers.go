package UID

import pf "github.com/Bookshelf-Writer/PointerFactory"

/* Добавляем типы */

const (
	TypeN1 pf.TypeTag = iota + 1
	TypeN2
	TypeS
)

func init() {
	SYS.TypeMAP['1'] = TypeN1
	SYS.TypeMAP['2'] = TypeN2
	SYS.TypeMAP['s'] = TypeS
}
