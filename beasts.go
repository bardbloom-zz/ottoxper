package main

import (
	"com.tengen/cm/util"
	"fmt"
	"github.com/robertkrimen/otto"
	//	"github.com/robertkrimen/otto/underscore"
	"math/rand"
	"time"
)

const NDeeds = 5

var _ fmt.Stringer = nil
var _ = util.Ignore()

var Rando = rand.New(rand.NewSource(123))

type Javascript string

type Beast interface {
	DoAt(js Javascript, t time.Time) // Tell the beast to do something
	Run()
	Stop()
	Otto() *otto.Otto
	BeastMaster() BeastMaster
}

type BeastMaster interface {
	BeastAboutToDo(beast Beast, js Javascript)
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

var Deedles []Deed = []Deed{
	Deed{"chirp", "chirp(\"%v\",%v)", chirp},
}

func chirp(shrew *Shrew) {
	fmt.Printf("Ӝχ0う– %v goes ‘%v’\n", shrew, "chirp")
}

////////////////////////////////////////////////////////////////////////////////

type Chore struct {
	StartTime time.Time
	Js        Javascript
	Done      bool
}

type Shrew struct {
	Name         string
	OttoInterp   *otto.Otto
	NDeedsToDo   int
	BeastMaster1 BeastMaster
	Deeds        []Deed
	Chores       []*Chore
}

func chirp_from_otto(ottow *otto.Otto) func(call otto.FunctionCall) otto.Value {
	return func(call otto.FunctionCall) otto.Value {
		shrewName, err := call.Argument(0).ToString()
		if err != nil {
			fmt.Printf("Ӝφ7マ– Error in ShrewName –  err=%v\n", err)
		}
		chirpNumb, err := call.Argument(1).ToInteger()
		if err != nil {
			fmt.Printf("ӜΞ5ゴ– Err in ChirpNumb–  err=%v\n", err)
		}

		fmt.Printf("ӜΙ7づ– chirp_from_otto –  shrewName=%v, chirpNumb=%v\n", shrewName, chirpNumb)
		return otto.UndefinedValue()
	}
}

func NewShrew(name string, bm BeastMaster) *Shrew {
	ottow := otto.New()
	ottow.Set("chirp", chirp_from_otto(ottow))
	return &Shrew{
		Name:         name,
		OttoInterp:   ottow,
		NDeedsToDo:   NDeeds,
		BeastMaster1: bm,
		Deeds:        Deedles,
		Chores:       []*Chore{},
	}
}

func (shrew *Shrew) Otto() *otto.Otto {
	return shrew.OttoInterp
}

func (shrew *Shrew) BeastMaster() BeastMaster {
	return shrew.BeastMaster1
}

func (shrew *Shrew) DoAt(js Javascript, t time.Time) {
	chore := &Chore{t, js, false}
	shrew.Chores = append(shrew.Chores, chore)
	fmt.Printf("Ӝς8ワ– DoAt –  len(shrew.Chores)=%v, shrew.Chores=%v\n", len(shrew.Chores), shrew.Chores)
}

func (this *Shrew) String() string {
	return fmt.Sprintf("Shrew(%v)", this.Name)
}

func (shrew *Shrew) Run() {
	util.DoThis("≈≈≈ Plz do the initial deeds via javascript!")
	for i := 0; i < shrew.NDeedsToDo; i++ {
		deed := randomDeed(shrew.Deeds)
		fmt.Printf("ӜΔ7に– %v doing %v\n", shrew, deed.Name)
		specificJsCall := Javascript(fmt.Sprintf(string(deed.Js), shrew, i))
		shrew.BeastMaster().BeastAboutToDo(shrew, specificJsCall)
		deed.Fun(shrew)
		time.Sleep(time.Second)
	}
	for {
		time.Sleep(time.Second / 10)
		shrew.DoAnUndoneChore()
	}
}

func (shrew *Shrew) DoAnUndoneChore() {
	util.DoThis("Ӝ I am a bit worried about atomicity on Chores")
	for _, ch := range shrew.Chores {
		if !ch.Done && time.Now().After(ch.StartTime) {
			fmt.Printf("Ӝυ8ゆ– I am about to javascriptly do this chore! –  ch=%v\n", ch)
			ch.Done = true
			shrew.Otto().Run(string(ch.Js))
			fmt.Printf("Ӝυ8ゆ– I am have done javascriptly this chore! –  ch=%v\n", ch)
		}
	}
}

func (shrew *Shrew) Stop() {
	fmt.Printf("ӜΝ1や– I stop!!\n")
}

var _ Beast = &Shrew{}

////////////////////////////////////////////////////////////////////////////////

type Event struct {
	T  time.Duration
	Js Javascript
}

type Wombat struct {
	Events    util.AtomicMap // really map[Beast][]Event
	StartTime time.Time
}

func NewWombat() *Wombat {
	return &Wombat{
		Events:    util.NewAtomicMap(),
		StartTime: time.Now(),
	}
}

func (wombat *Wombat) BeastAboutToDo(beast Beast, js Javascript) {
	fmt.Printf("Ӝη8ヂ– I, Wombat, am no longer ignoring–  beast=%v, js=%v\n", beast, js)
	updater := func(there bool, key, oldValue interface{}) (shouldUpdate bool, newValue interface{}) {
		ev := Event{time.Now().Sub(wombat.StartTime), js}
		if there {
			val := oldValue.([]Event)
			return true, append(val, ev)
		} else {
			return true, []Event{ev}
		}
	}

	wombat.Events.AtomicUpdate(beast, updater)
}

func (wombat *Wombat) PrintCurrentEventQueues() {
	keys := wombat.Events.AtomicKeys()
	for _, k := range keys {
		v, _ := wombat.Events.AtomicGet(k)
		fmt.Printf("ӜΗ9ボ– Status –  k=%v, v=%v\n", k, v)
	}
}

func (wombat *Wombat) Replay() {
	keys := wombat.Events.AtomicKeys()
	start := time.Now()
	for _, k := range keys {
		v, _ := wombat.Events.AtomicGet(k)
		evs := v.([]Event)
		beast := k.(Beast)
		for _, ev := range evs {
			beast.DoAt(ev.Js, start.Add(ev.T))
		}
	}
}

////////////////////////////////////////////////////////////////////////////////

func main() {
	fmt.Printf("by Eurayle's sweet elk, I compile!\n")
	wombat := NewWombat()
	shrew1 := NewShrew("ś1", wombat)
	shrew2 := NewShrew("ś2", wombat)
	go shrew1.Run()
	go shrew2.Run()
	fmt.Printf("Ӝη4に– Main: wait for shrews to finish\n")
	time.Sleep(time.Second * 20)
	fmt.Printf("ӜΗ5オ– Main: wombat replay\n")
	wombat.Replay()
	time.Sleep(time.Second * 20)
}
