package idgenerator

import (
	"math/big"
	"strconv"
	"time"
)

var (
	// Epoch is set to the twitter snowflake epoch of Nov 04 2010 01:42:54 UTC in milliseconds
	// You may customize this to set a different epoch for your application.
	Epoch int64 = 1715299200000

	// TopLevelMachineBits holds the number of bits to use for top level machine.
	TopLevelMachineBits uint8 = 17

	// LowLevelMachineBits holds the number of bits to use for low level machine.
	LowLevelMachineBits uint8 = 17

	// SequenceBits holds the number of bits to use for Step
	SequenceBits uint8 = 12

	// TimestampBits holds the number of bits to use for timestamp.
	TimestampBits uint8 = 42
)

type Generator interface {
	Generate() *big.Int
}

type generator struct {
	epoch time.Time
	time  int64
	step  int64

	topNode string
	lowNode string
}

func New(topNode, lowNode int64) (Generator, error) {
	g := &generator{
		topNode: formalize(strconv.FormatInt(topNode, 2), int(TopLevelMachineBits)),
		lowNode: formalize(strconv.FormatInt(lowNode, 2), int(LowLevelMachineBits)),
	}

	g.epoch = time.Unix(0, Epoch*int64(time.Millisecond))

	return g, nil
}

func (g *generator) Generate() *big.Int {
	i := big.NewInt(0)
	if _, ok := i.SetString(g.generateBits(), 2); !ok {
		panic("Invalid number!")
	}

	return i
}

func (g *generator) generateBits() string {
	return g.timestamp() + g.topNode + g.lowNode + g.getSequenceData()
}

func (g *generator) timestamp() string {
	now := time.Since(g.epoch).Milliseconds()
	if now == g.time {
		g.step += 1

		if g.step == 0 {
			for now <= g.time {
				now = time.Since(g.epoch).Milliseconds()
			}
		}
	} else {
		g.step = 0
	}

	g.time = now

	return formalize(strconv.FormatInt(now, 2), int(TimestampBits))
}

func (g *generator) getSequenceData() string {
	return formalize(strconv.FormatInt(g.step, 2), int(SequenceBits))
}

func formalize(str string, expectedBits int) string {
	diff := expectedBits - len(str)
	if diff == 0 {
		return str
	}
	for i := 0; i < diff; i++ {
		str = "0" + str
	}

	return str
}
