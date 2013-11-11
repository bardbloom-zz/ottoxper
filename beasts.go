package main

import (
	"com.tengen/cm/util"
	"fmt"
	"time"
)

var _ fmt.Stringer = nil
var _ = util.Ignore()

type Javascript string

type Beast interface {
	DoAt(js Javascript, t time.Time)
	Perform(actionAsJs string, beastmaster BeastMaster, action func())
	Run()
	Stop()
}

type BeastMaster interface {
}

////////////////////////////////////////////////////////////////////////////////

func main() {
	fmt.Printf("by Eurayle's sweet elk, I compile!\n")
}
