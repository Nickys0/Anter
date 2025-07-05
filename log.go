package Anter

import (
	"errors"
	"fmt"
	"os"
)

func assert(condition bool, msg string) {
	if !condition {
		panic(msg);
	}
}

func Err(msg string) error{
	return errors.New(msg)
}

func ErrF(format string, n ...any) error{
	return errors.New(SErrF(format, n...))
}

func SErrF(format string, n ...any) string{
	return fmt.Sprintf(format, n...)
}

func UnimplFunc(f string){
	ErrF("Unimplemented func<%s>: Please wait", f)
	os.Exit(1)
}

func Unimpl(msg string){
    if msg == ""{
		ErrF("Unimplemented!")
	}else{
        ErrF("Unimplemented: %s", msg)
    }
}
