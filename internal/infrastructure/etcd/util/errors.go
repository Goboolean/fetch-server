package etcdutil

import "errors"



var ErrFieldNotFount = errors.New("field is not found")

var ErrGivenNotAPointer = errors.New("given is not a pointer")