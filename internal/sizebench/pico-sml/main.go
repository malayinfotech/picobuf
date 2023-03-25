// Copyright (C) 2021 Storx Labs, Inc.
// See LICENSE for copying information.

package main

import (
	"fmt"

	"picobuf/internal/sizebench/pico"
	"picobuf/internal/sizebench/pico/sml"
)

func main() {
	fmt.Println([]byte{})
	pico.Test(&sml.Nop{})
	pico.Test(&sml.Types{})
}
