package observable

import (
	"lld/notifyme/observer"
)

type QuantityObservable interface {
	Add(obs observer.NotifyObserver)
	Remove(obs observer.NotifyObserver)
	NotifyAll()
	SetQuantity(qty int)
}

type Observable struct {
	qty  int
	List []observer.NotifyObserver
}

func NewObservable() *Observable {
	obj := &Observable{qty: 0, List: make([]observer.NotifyObserver, 0)}
	return obj
}

func (obs *Observable) Add(no observer.NotifyObserver) {
	obs.List = append(obs.List, no)
}

func (oble *Observable) Remove(no observer.NotifyObserver) {
	for i, obs := range oble.List {
		if obs.GetId() == no.GetId() {
			oble.List = append(oble.List[:i],
				oble.List[i+1:]...)
		}
	}
}

func (obs *Observable) NotifyAll() {
	for _, observer := range obs.List {
		observer.Update(obs.qty)
	}
}

func (obs *Observable) SetQuantity(qty int) {
	obs.qty = qty
	obs.NotifyAll()
}
