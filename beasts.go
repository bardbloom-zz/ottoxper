package main

import (
	"com.tengen/cm/util"
	"fmt"
	"github.com/robertkrimen/otto"
	//	"github.com/robertkrimen/otto/underscore"
	"math/rand"
	"time"
)

var _ fmt.Stringer = nil
var _ = util.Ignore()

var Rando = rand.New(rand.NewSource(123))

type Javascript string

type Beast interface {
	DoAt(js Javascript, t time.Time) // Tell the beast to do something
	Run()
	Stop()
	Otto() otto.Otto
	BeastMaster() BeastMaster
}

type BeastMaster interface {
	BeastAboutToDo(js Javascript)
}

////////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////////

type Deed struct {
	Name string
	Js   Javascript
	Fun  func(*Shrew)
}

func randomDeed(deeds []Deed) Deed {
	k := Rando.Int31n(int32(len(deeds)))
	return deeds[k]
}

////////////////////////////////////////////////////////////////////////////////

type Shrew struct {
	Name         string
	OttoInterp   otto.Otto
	NDeedsToDo   int
	BeastMaster1 BeastMaster
	Deeds        []Deed
}

func (shrew *Shrew) Otto() otto.Otto {
	return shrew.OttoInterp
}

func (shrew *Shrew) BeastMaster() BeastMaster {
	return shrew.BeastMaster1
}

func (shrew *Shrew) DoAt(js Javascript, t time.Time) {
	fmt.Printf("ӜΘ7こ– I should schedule –  js=%v, t=%v\n", js, t)
}

func (this *Shrew) String() string {
	return fmt.Sprintf("Shrew(%v)", this.Name)
}

func (shrew *Shrew) Run() {
	for i := 0; i < shrew.NDeedsToDo; i++ {
		deed := randomDeed(shrew.Deeds)
		fmt.Printf("ӜΔ7に– %v doing %v\n", shrew, deed.Name)
		shrew.BeastMaster().BeastAboutToDo(deed.Js)
		deed.Fun(shrew)
	}
}

func (shrew *Shrew) Stop() {
	fmt.Printf("ӜΝ1や– I stop!!\n")
}

var _ Beast = &Shrew{}

////////////////////////////////////////////////////////////////////////////////

type Wombat struct {
}

func (wombat *Wombat) BeastAboutToDo(js Javascript) {
	fmt.Printf("Ӝε0け– I, Wombat, am ignoring –  js=%v\n", js)
}

////////////////////////////////////////////////////////////////////////////////

func main() {
	fmt.Printf("by Eurayle's sweet elk, I compile!\n")
}
