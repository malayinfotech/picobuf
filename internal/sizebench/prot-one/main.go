// Copyright (C) 2021 Storx Labs, Inc.
// See LICENSE for copying information.

package main

import (
	"fmt"

	"picobuf/internal/sizebench/prot"
	"picobuf/internal/sizebench/prot/one"
)

func main() {
	fmt.Println([]byte{})
	prot.Test(&one.Nop{})
	prot.Test(&one.Types{})
}
