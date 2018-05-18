package rand

import (
	"math/rand"
)

//ShuffleAlgorithm shuffle slice inside
type ShuffleAlgorithm interface {
	Shuffle(slice []interface{})
}

//fisherYatesShuffle implement using Durstenfeld's version of the algorithm
type fisherYatesShuffle struct {
	rand *rand.Rand
}

//NewFisherYatesShuffle create FisherYatesShuffle using Durstenfeld's version of the algorithm
func NewFisherYatesShuffle(seed int64) ShuffleAlgorithm {
	return &fisherYatesShuffle{
		rand: rand.New(rand.NewSource(seed)),
	}
}

func (fy *fisherYatesShuffle) Shuffle(slice []interface{}) {
	n := len(slice)
	for i := n - 1; i > 0; i-- {
		j := fy.rand.Intn(i + 1)
		slice[i], slice[j] = slice[j], slice[i]
	}
}
