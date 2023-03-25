// Copyright (C) 2021 Storx Labs, Inc.
// See LICENSE for copying information.

package main

import (
	"fmt"

	"picobuf/internal/sizebench/pico"
	"picobuf/internal/sizebench/pico/one"
)

func main() {
	fmt.Println([]byte{})
	pico.Test(&one.Nop{})
	pico.Test(&one.Types{})
}
