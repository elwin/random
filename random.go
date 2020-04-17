package random

/*
#include <stdlib.h>
*/
import "C"

func Random() int {
	return int(C.random())
}

