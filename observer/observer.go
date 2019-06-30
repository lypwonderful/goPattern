package main

import "fmt"

type (
	MsgNotifer struct {
		observers map[Observer]Actioner
	}

	MsgObserver struct {
		Id int
	}

	Msg struct {
		Content string
	}

	Action struct {
		Actioner
	}
)

type (
	Observer interface {
		Event(msg Msg)
	}

	Notify interface {
		Register(o Observer)
		UnRegister(o Observer)
		Notify(msg string)
	}

	Actioner interface {
		Sleep()
		Movie()
	}
)

func (Action) Sleep() {
	fmt.Println("sleep")
}
func (Action) Movie() {
	fmt.Println("movie")
}

func (msgN *MsgNotifer) Register(o Observer) {
	msgN.observers[o] = Action{}
}

func (msgN *MsgNotifer) UnRegister(o Observer) {
	delete(msgN.observers, o)
}
func (msgN *MsgNotifer) Notify(msg Msg) {
	for msgO, acn := range msgN.observers {
		msgO.Event(msg)
		acn.Movie()
	}
}

func (o *MsgObserver) Event(msg Msg) {
	fmt.Printf("I'm Number:%d,it time to %s\n", o.Id, msg.Content)
}

func main() {
	msgN := MsgNotifer{
		observers: map[Observer]Actioner{},
	}
	ober1 := &MsgObserver{
		Id: 1,
	}
	ober2 := &MsgObserver{
		Id: 2,
	}

	msgN.Register(ober1)
	msgN.Register(ober2)

	msgN.Notify(Msg{Content: "play footbool."})

	msgN.UnRegister(ober1)
	msgN.Notify(Msg{Content: "sleep."})
}
