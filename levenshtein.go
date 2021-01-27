package levenshtein

/*
#cgo CFLAGS: -I .

#include <stdint.h>
#include <stddef.h>
#include <stdlib.h>

#include "levenshtein.h"
*/
import "C"

import (
	"unsafe"
)

// LevenScore .
func LevenScore(a, b []byte) float32 {
	aLen := len(a)
	bLen := len(b)
	var aPtr, bPtr unsafe.Pointer
	if aLen > 0 {
		aPtr = unsafe.Pointer(&a[0])
	}
	if bLen > 0 {
		bPtr = unsafe.Pointer(&b[0])
	}

	return float32(C.leven_score(aPtr, C.size_t(aLen), bPtr, C.size_t(bLen)))
}
