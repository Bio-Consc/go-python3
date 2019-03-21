package python3

/*
#include "Python.h"
#include "macro.h"
#include "type.h"
*/

import "C"

import (
	"unsafe"
)

// PyString_FromString (const char *v)
// Return value: New reference.
// Return a new string object with a copy of the string v as value on success, and NULL on failure.
// The parameter v must not be NULL; it will not be checked.
func PyString_FromString(v string) *PyObject {
	cstr := C.CString(v)
	defer c.free(unsafe.Pointer(cstr))
	return togo(C.PyString_FromString(cstr))
}
