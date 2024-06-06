package UID

import pf "github.com/Bookshelf-Writer/PointerFactory"

/* Добавляем типы */

const (
	TypeN1 pf.TypeTag = iota + 1
	TypeN2
	TypeS
)

func init() {
	SYS.TypeMAP['b'] = TypeN1
	SYS.TypeMAP['k'] = TypeN2
	SYS.TypeMAP['r'] = TypeS
}
