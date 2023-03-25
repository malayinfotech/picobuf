// Copyright (C) 2021 Storx Labs, Inc.
// See LICENSE for copying information.

package main

import (
	"fmt"

	"picobuf/internal/sizebench/basz"
	"picobuf/internal/sizebench/pico"
	"picobuf/internal/sizebench/pico/one"
)

func main() {
	fmt.Println([]byte{})
	basz.CallDynamicMethod(fmt.Println)
	pico.Test(&one.Nop{})
	pico.Test(&one.Types{})
	basz.CallDynamicMethod(&one.Types{})
}
