package balance

import (
	"errors"
	"fmt"
)

type BlanceMgr struct {
	allBalance map[string]Balance
}

var mgr = BlanceMgr{
	allBalance: make(map[string]Balance),
}

func (p *BlanceMgr) registerBalance(name string, b Balance) {
	p.allBalance[name] = b
}

func RegisterBalance(name string, b Balance) {
	mgr.registerBalance(name, b)
}

func DoBalance(name string, insts []*Instance) (inst *Instance, err error) {
	balance, ok := mgr.allBalance[name]
	if !ok {
		err = errors.New("not found")
		fmt.Println("not found :", name)
		return
	}
	inst, err = balance.DoBalance(insts)
	if err != nil {
		return
	}
	return
}
