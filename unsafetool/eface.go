package unsafetool

import (
	"unsafe"
)

type RType = uintptr

type Eface struct {
	rtype RType
	Data  unsafe.Pointer
}
