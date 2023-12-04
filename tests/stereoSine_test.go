package tests

import (
	"math"
	"testing"
	"time"

	"github.com/oklookat/portaudio"
)

const sampleRate = 44100

func TestStereoSine(t *testing.T) {
	portaudio.Initialize()
	defer portaudio.Terminate()
	s := newStereoSine(256, 320, sampleRate, t)
	defer s.Close()
	chk(s.Start(), t)
	time.Sleep(2 * time.Second)
	chk(s.Stop(), t)
}

type stereoSine struct {
	*portaudio.Stream
	stepL, phaseL float64
	stepR, phaseR float64
}

func newStereoSine(freqL, freqR, sampleRate float64, t *testing.T) *stereoSine {
	s := &stereoSine{nil, freqL / sampleRate, 0, freqR / sampleRate, 0}
	var err error
	s.Stream, err = portaudio.OpenDefaultStream(0, 2, sampleRate, 0, s.processAudio)
	chk(err, t)
	return s
}

func (g *stereoSine) processAudio(out [][]float32) {
	for i := range out[0] {
		out[0][i] = float32(math.Sin(2 * math.Pi * g.phaseL))
		_, g.phaseL = math.Modf(g.phaseL + g.stepL)
		out[1][i] = float32(math.Sin(2 * math.Pi * g.phaseR))
		_, g.phaseR = math.Modf(g.phaseR + g.stepR)
	}
}
