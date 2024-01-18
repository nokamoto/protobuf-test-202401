package main

import (
	"fmt"

	"github.com/google/go-cmp/cmp"
	"github.com/nokamoto/protobuf-test-202401/gen/proto/v1"
	v2 "github.com/nokamoto/protobuf-test-202401/gen/proto/v2"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/testing/protocmp"
)

func main() {
	a := v1.A{
		A: "a",
		B: &v1.SuperSet{
			First: 100,
			Third: 300,
		},
		C: []*v1.SuperSet{
			{
				First: 100,
				Third: 300,
			},
			{
				First: 100 + 1,
				Third: 300 + 1,
			},
		},
		D: []*v1.Wrapper{
			{
				Set: &v1.SuperSet{
					First: 100 + 2,
					Third: 300 + 2,
				},
			},
		},
	}

	fmt.Println("diff v1.A -> v2.B")

	bytes, _ := proto.Marshal(&a)

	var b v2.B
	proto.Unmarshal(bytes, &b)
	fmt.Println(cmp.Diff(&a, &b, protocmp.Transform()))

	fmt.Println("diff v2.B -> v1.A")

	b.B.Second = 200
	b.C[0].Second = 200
	b.C[1].Second = 200 + 1
	b.D[0].Fourth = 400 + 2

	bytes, _ = proto.Marshal(&b)

	var a2 v1.A
	proto.Unmarshal(bytes, &a2)
	fmt.Println(cmp.Diff(&b, &a2, protocmp.Transform()))

	fmt.Println("a v1.A")
	fmt.Println(protojson.Format(&a))
	fmt.Println()

	fmt.Println("b v1.B")
	fmt.Println(protojson.Format(&b))
	fmt.Println()

	fmt.Println("a2 v1.A")
	fmt.Println(protojson.Format(&a2))
}
