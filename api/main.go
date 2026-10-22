package main

import (
	"fmt"

	"k8s.io/apimachinery/pkg/runtime/schema"
)

func main() { fmt.Println(schema.GroupVersion{Group: "apps", Version: "v1"}) }
