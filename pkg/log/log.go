package log

import (
	"fmt"
)

func Debug(l ...interface{}) {
	fmt.Println(l...)
}
