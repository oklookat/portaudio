package tests

import (
	"math/rand"
	"testing"
	"time"

	"github.com/oklookat/portaudio"
)

func TestNoise(t *testing.T) {
	portaudio.Initialize()
	defer portaudio.Terminate()
	h, err := portaudio.DefaultHostApi()
	chk(err, t)
	stream, err := portaudio.OpenStream(portaudio.HighLatencyParameters(nil, h.DefaultOutputDevice), func(out []int32) {
		for i := range out {
			out[i] = int32(rand.Uint32())
		}
	})
	chk(err, t)
	defer stream.Close()
	chk(stream.Start(), t)
	time.Sleep(time.Second)
	chk(stream.Stop(), t)
}
