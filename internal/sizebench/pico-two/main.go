// Copyright (C) 2021 Storx Labs, Inc.
// See LICENSE for copying information.

package main

import (
	"fmt"

	"picobuf/internal/sizebench/pico"
	"picobuf/internal/sizebench/pico/two"
)

func main() {
	fmt.Println([]byte{})
	pico.Test(&two.Nop{})
	pico.Test(&two.Types{})
	pico.Test(&two.Types2{})
}
