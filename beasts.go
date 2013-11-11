package main

import (
	"com.tengen/cm/util"
	"fmt"
	"github.com/robertkrimen/otto"
	//	"github.com/robertkrimen/otto/underscore"
	"time"
)

var _ fmt.Stringer = nil
var _ = util.Ignore()

type Javascript string

type Beast interface {
	DoAt(js Javascript, t time.Time) // Tell the beast to do something
	Run()
	Stop()
	Otto() otto.Otto
	BeastMaster() BeastMaster
}

type BeastMaster interface {
	BeastJustDid(js Javascript)
}

////////////////////////////////////////////////////////////////////////////////

type Shrew struct {
	ottoInterp  otto.Otto
	beastMaster BeastMaster
}

func (shrew *Shrew) Otto() otto.Otto {
	return shrew.ottoInterp
}

func (shrew *Shrew) BeastMaster() BeastMaster {
	return shrew.beastMaster
}

func (shrew *Shrew) DoAt(js Javascript, t time.Time) {
	fmt.Printf("ӜΘ7こ– I should schedule –  js=%v, t=%v\n", js, t)
}

func (shrew *Shrew) Run() {
	fmt.Printf("ӜΝ1や– I run!\n")
}

func (shrew *Shrew) Stop() {
	fmt.Printf("ӜΝ1や– I stop!!\n")
}

var _ Beast = &Shrew{}

////////////////////////////////////////////////////////////////////////////////

type Wombat struct {
}

func (wombat *Wombat) BeastJustDid(js Javascript) {
	fmt.Printf("Ӝε0け– I, Wombat, am ignoring –  js=%v\n", js)
}

////////////////////////////////////////////////////////////////////////////////

func main() {
	fmt.Printf("by Eurayle's sweet elk, I compile!\n")
}
