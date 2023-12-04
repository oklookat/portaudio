package tests

import (
	"testing"
	"time"

	"github.com/oklookat/portaudio"
)

func TestEcho(t *testing.T) {
	portaudio.Initialize()
	defer portaudio.Terminate()
	e := newEcho(time.Second/3, t)
	defer e.Close()
	chk(e.Start(), t)
	time.Sleep(4 * time.Second)
	chk(e.Stop(), t)
}

type echo struct {
	*portaudio.Stream
	buffer []float32
	i      int
}

func newEcho(delay time.Duration, t *testing.T) *echo {
	h, err := portaudio.DefaultHostApi()
	chk(err, t)
	p := portaudio.LowLatencyParameters(h.DefaultInputDevice, h.DefaultOutputDevice)
	p.Input.Channels = 1
	p.Output.Channels = 1
	e := &echo{buffer: make([]float32, int(p.SampleRate*delay.Seconds()))}
	e.Stream, err = portaudio.OpenStream(p, e.processAudio)
	chk(err, t)
	return e
}

func (e *echo) processAudio(in, out []float32) {
	for i := range out {
		out[i] = .7 * e.buffer[e.i]
		e.buffer[e.i] = in[i]
		e.i = (e.i + 1) % len(e.buffer)
	}
}
