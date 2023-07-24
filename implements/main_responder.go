package implements

import (
	"net/http"

	"github.com/starter-go/base/util"
	"github.com/starter-go/libgin"
)

// MainResponder 是统一的 Responder 主入口
type MainResponder struct {

	//starter:component
	_as func(libgin.Responder) //starter:as("#")

	Responders []libgin.Responder //starter:inject(".")

	cached []*libgin.ResponderRegistration
}

func (inst *MainResponder) _impl() {
	inst._as(inst)
}

// Registration ...
func (inst *MainResponder) Registration() *libgin.ResponderRegistration {
	return &libgin.ResponderRegistration{}
}

// Accept ...
func (inst *MainResponder) Accept(r *libgin.Response) bool {
	return true
}

// Send ...
func (inst *MainResponder) Send(res *libgin.Response) {

	rrlist := inst.getRegistrationList()

	for _, rr := range rrlist {
		if inst.isMatched(res, rr) {
			if rr.Responder.Accept(res) {
				rr.Responder.Send(res)
				return
			}
		}
	}

	c := res.Context
	c.String(http.StatusInternalServerError, "No Responder for the Response")
}

func (inst *MainResponder) isMatched(res *libgin.Response, rr *libgin.ResponderRegistration) bool {
	name1 := res.Responder
	name2 := rr.Name
	if name1 == "" {
		return true
	} else if name1 == name2 {
		return true
	}
	return false
}

func (inst *MainResponder) getRegistrationList() []*libgin.ResponderRegistration {
	list := inst.cached
	if list == nil {
		list = inst.loadRegistrationList()
		inst.cached = list
	}
	return list
}

func (inst *MainResponder) loadRegistrationList() []*libgin.ResponderRegistration {
	src := inst.Responders
	dst := make([]*libgin.ResponderRegistration, 0)
	for _, r1 := range src {
		if r1 == nil {
			continue
		}
		r2 := r1.Registration()
		if inst.isReady(r2) {
			dst = append(dst, r2)
		}
	}
	return inst.sort(dst)
}

func (inst *MainResponder) isReady(rr *libgin.ResponderRegistration) bool {
	if rr == nil {
		return false
	}
	if !rr.Enabled {
		return false
	}
	if rr.Responder == nil {
		return false
	}
	return true
}

func (inst *MainResponder) sort(list []*libgin.ResponderRegistration) []*libgin.ResponderRegistration {
	s := &util.Sorter{}
	s.OnLen = func() int { return len(list) }
	s.OnSwap = func(i1, i2 int) { list[i1], list[i2] = list[i2], list[i1] }
	s.OnLess = func(i1, i2 int) bool {
		n1 := list[i1].Priority
		n2 := list[i2].Priority
		return n1 < n2
	}
	s.Sort()
	return list
}
