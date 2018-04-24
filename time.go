// Package cstrftime implements simple format date and time.
package cstrftime

// #include <time.h>
// #include <stdlib.h>
import "C"
import (
	"time"
	"unsafe"
)

// Format formats date according to the directives in the given format string.
// The directives begins with a percent (%) character. Any text not listed as a
// directive will be passed through to the output string.
func Format(s string, t time.Time) string {
	rawtime := C.long(t.Unix())
	var info *C.struct_tm
	C.ctime(&rawtime)
	info = C.localtime(&rawtime)
	buf := new(C.char)
	format := C.CString(s)
	defer C.free(unsafe.Pointer(format))
	maxsize := C.ulong(256)
	C.strftime(buf, maxsize, format, info)
	return C.GoString(buf)
}
