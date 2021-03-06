package maps_test

import (
	"fmt"
	"testing"

	"github.com/mark-ahn/maps"
)

func TestMap(t *testing.T) {
	m := maps.OfStringNInterface{}

	m["test"] = 2

	st := maps.NewOfStringNInterfaceSt(m)
	st.Range(maps.OfStringNInterfaceLoopFunc(func(k string, v interface{}) bool {
		fmt.Println(k, v)
		return true
	}))
}
