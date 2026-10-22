package main

import (
	"fmt"

	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/client-go/kubernetes/scheme"
)

func main() { fmt.Println(sets.New("a", "b").Len(), scheme.Scheme != nil) }
