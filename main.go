package main

import (
	"fmt"

	"github.com/google/go-cmp/cmp"
	"github.com/nokamoto/protobuf-test-202401/gen/proto/v1"
	v2 "github.com/nokamoto/protobuf-test-202401/gen/proto/v2"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/testing/protocmp"
)

func main() {
	a := v1.A{
		A: "a",
	}

	bytes, _ := proto.Marshal(&a)

	var b v2.B
	proto.Unmarshal(bytes, &b)

	fmt.Println("diff", cmp.Diff(&a, &b, protocmp.Transform()))
}
