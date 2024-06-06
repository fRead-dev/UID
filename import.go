package UID

import pf "github.com/Bookshelf-Writer/PointerFactory"

type WrapperObj struct {
	pf.WrapperObj
}

var SYS = pf.CreateWrapper()

type CreatorObj struct {
	pf.CreatorObj
}

type PointerObj struct {
	pf.PointerObj
}
