package p

type Mutex struct {
	l chan struct{}
}

func NewMutex() *Mutex {
	m := &Mutex{
		l: make(chan struct{}, 1),
	}
	m.l <- struct{}{}
	return m
}

func (m Mutex) Lock() {
	<-m.l
}

func (m Mutex) UnLock() {
	m.l <- struct{}{}
}
