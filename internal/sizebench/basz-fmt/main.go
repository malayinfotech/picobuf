// Copyright (C) 2021 Storx Labs, Inc.
// See LICENSE for copying information.

package main

import (
	"fmt"

	"picobuf/internal/sizebench/basz"
)

func main() {
	fmt.Println([]byte{})
	basz.CallDynamicMethod(fmt.Println)
}
