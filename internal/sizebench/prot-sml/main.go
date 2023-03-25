// Copyright (C) 2021 Storx Labs, Inc.
// See LICENSE for copying information.

package main

import (
	"fmt"

	"picobuf/internal/sizebench/prot"
	"picobuf/internal/sizebench/prot/sml"
)

func main() {
	fmt.Println([]byte{})
	prot.Test(&sml.Nop{})
	prot.Test(&sml.Types{})
}
