// Copyright (C) 2021 Storx Labs, Inc.
// See LICENSE for copying information.

package main

import (
	"fmt"

	"picobuf"
)

func main() {
	var z uint32
	enc := picobuf.NewEncoder()
	enc.Uint32(1, &z)
	fmt.Println(enc.Buffer())
}
