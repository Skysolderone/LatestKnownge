package wisdom

import "golang.org/x/exp/rand"

type Wisdom interface {
	GetWisdom() string
}
type WisdomImpl struct{}

var quotes = []string{
	"Stay focused and keep shipping.",
	"Success is not final, failure is not fatal.",
	"Keep pushing your limits.",
}

func (w WisdomImpl) GetWisdom() string {
	return quotes[rand.Intn(len(quotes))]
}
