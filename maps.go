package maps

//go:generate genny -in $GOFILE -out ${GOPACKAGE}_genny.go gen "SomeK=BUILTINS,interface{} SomeV=BUILTINS,interface{}"

import "github.com/cheekybits/genny/generic"

type SomeK generic.Type
type SomeV generic.Type

type OfSomeKNSomeVIf interface {
	Load(SomeK) (SomeV, bool)
}

type OfSomeKNSomeVMutIf interface {
	Store(SomeK, SomeV)
	Delete(SomeK)
}

type OfSomeKNSomeVIterIf interface {
	Range(func(SomeK, SomeV) bool)
}

type OfSomeKNSomeV map[SomeK]SomeV

func (__ OfSomeKNSomeV) Load(k SomeK) (SomeV, bool) {
	v, ok := __[k]
	return v, ok
}

func (__ OfSomeKNSomeV) Store(k SomeK, v SomeV) {
	__[k] = v
}
func (__ OfSomeKNSomeV) Delete(k SomeK, v SomeV) {
	delete(__, k)
}
func (__ OfSomeKNSomeV) Range(f func(k SomeK, v SomeV) bool) {
	for k, v := range __ {
		if !f(k, v) {
			break
		}
	}
}

type _OfSomeKNSomeV = OfSomeKNSomeV
type OfSomeKNSomeVSt struct {
	// do not want to export but want to use embedding method
	_OfSomeKNSomeV
}

func NewOfSomeKNSomeVSt(somes OfSomeKNSomeV) *OfSomeKNSomeVSt {
	return &OfSomeKNSomeVSt{
		_OfSomeKNSomeV: somes,
	}
}
