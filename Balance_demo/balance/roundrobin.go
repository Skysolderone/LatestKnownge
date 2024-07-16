package balance

import (
	"errors"
)

type RoundrobinBalance struct {
	curIndex int
}

func init() {
	RegisterBalance("round", &RoundrobinBalance{})
}

func (r *RoundrobinBalance) DoBalance(insts []*Instance, key ...string) (inst *Instance, err error) {
	if len(insts) == 0 {
		err = errors.New("no instance")
		return
	}
	lens := len(insts)
	if r.curIndex >= lens {
		r.curIndex = 0
	}

	inst = insts[r.curIndex]
	r.curIndex++
	return
}
