package main

import (
	"UID"
	fWASM "github.com/Bookshelf-Writer/PointerFactory/compile/wasm"
)

//////////////////////////////////////////////////////////////////////////////

func main() {
	fWASM.InitWASM(&UID.SYS)
}
