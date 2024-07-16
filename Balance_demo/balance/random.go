package balance

import (
	"errors"
	"math/rand"
)

type RanomBalance struct{}

func init() {
	RegisterBalance("random", &RanomBalance{})
}

func (r *RanomBalance) DoBalance(insts []*Instance, key ...string) (inst *Instance, err error) {
	if len(insts) == 0 {
		err = errors.New("no instance")
		return
	}
	lens := len(insts)
	index := rand.Intn(lens)
	inst = insts[index]
	return
}
