package unsafetool

import (
	"unsafe"
)

type Itab struct {
	ignore uintptr
	RType  RType
}

type Iface struct {
	Itab *Itab
	Data unsafe.Pointer
}
