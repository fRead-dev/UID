package UID

import pf "github.com/Bookshelf-Writer/PointerFactory"

//////////////////////////////////////////////////////////////////////////////

const (
	TypeNone = pf.TypeNone

	TypeBook   pf.TypeTag = 30 // Книга
	TypeAuthor pf.TypeTag = 32 // Автор

	TypeUser pf.TypeTag = 50 // Пользователь
)

func init() {
	SYS.TypeMAP['0'] = TypeNone

	SYS.TypeMAP['b'] = TypeBook
	SYS.TypeMAP['a'] = TypeAuthor

	SYS.TypeMAP['u'] = TypeUser
}
