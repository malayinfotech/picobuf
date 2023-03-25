// Copyright (C) 2021 Storx Labs, Inc.
// See LICENSE for copying information.

package main

import (
	"fmt"

	"picobuf/internal/sizebench/prot"
	"picobuf/internal/sizebench/prot/two"
)

func main() {
	fmt.Println([]byte{})
	prot.Test(&two.Nop{})
	prot.Test(&two.Types{})
	prot.Test(&two.Types2{})
}
