package main

import (
	"fmt"
)

type Portion int

const (
	Regular Portion = iota
	Small
	Large
)

type Udon struct {
	men      Portion
	aburaage bool
	ebiten   uint
}

func NewUdon(p Portion, aburaage bool, ebiten uint) *Udon {
	return &Udon{
		men:      p,
		aburaage: aburaage,
		ebiten:   ebiten,
	}
}

var tempuraUdon = NewUdon(Large, false, 2)

func main() {
	fmt.Printf("Udon details: %+v\n", tempuraUdon)
}
