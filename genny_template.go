package maps

import (
	"fmt"

	"github.com/cheekybits/genny/generic"
)

type SomeK generic.Type
type SomeV generic.Type

type _Prefix_SomeKNSomeVIf interface {
	Load(SomeK) (SomeV, error)
}

type _Prefix_SomeKNSomeVMutIf interface {
	Store(SomeK, SomeV)
	Delete(SomeK)
}

type _Prefix_SomeKNSomeVLooper interface {
	LoopItem(SomeK, SomeV) bool
}
type _Prefix_SomeKNSomeVLoopFunc func(SomeK, SomeV) bool

func (__ _Prefix_SomeKNSomeVLoopFunc) LoopItem(k SomeK, v SomeV) bool {
	return __(k, v)
}

type _Prefix_SomeKNSomeVIterIf interface {
	Range(_Prefix_SomeKNSomeVLooper)
}

type _Prefix_SomeKNSomeV map[SomeK]SomeV

func (__ _Prefix_SomeKNSomeV) Load(k SomeK) (SomeV, error) {
	v, ok := __[k]
	var err error
	if !ok {
		err = fmt.Errorf("not found for key %v", k)
	}
	return v, err
}

func (__ _Prefix_SomeKNSomeV) Store(k SomeK, v SomeV) {
	__[k] = v
}
func (__ _Prefix_SomeKNSomeV) Delete(k SomeK, v SomeV) {
	delete(__, k)
}
func (__ _Prefix_SomeKNSomeV) Range(f func(k SomeK, v SomeV) bool) {
	for k, v := range __ {
		if !f(k, v) {
			break
		}
	}
}

type __Prefix_SomeKNSomeV = _Prefix_SomeKNSomeV
type _Prefix_SomeKNSomeVSt struct {
	// do not want to export but want to use embedding method
	__Prefix_SomeKNSomeV
}

func New_Prefix_SomeKNSomeVSt(somes _Prefix_SomeKNSomeV) *_Prefix_SomeKNSomeVSt {
	return &_Prefix_SomeKNSomeVSt{
		__Prefix_SomeKNSomeV: somes,
	}
}
