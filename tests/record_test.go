package tests

import (
	"encoding/binary"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"testing"

	"github.com/oklookat/portaudio"
)

func TestRecord(t *testing.T) {
	if len(os.Args) < 2 {
		fmt.Println("missing required argument:  output file name")
		return
	}
	fmt.Println("Recording.  Press Ctrl-C to stop.")

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)

	fileName := os.Args[1]
	if !strings.HasSuffix(fileName, ".aiff") {
		fileName += ".aiff"
	}
	f, err := os.Create(fileName)
	chk(err, t)

	// form chunk
	_, err = f.WriteString("FORM")
	chk(err, t)
	chk(binary.Write(f, binary.BigEndian, int32(0)), t) //total bytes
	_, err = f.WriteString("AIFF")
	chk(err, t)

	// common chunk
	_, err = f.WriteString("COMM")
	chk(err, t)
	chk(binary.Write(f, binary.BigEndian, int32(18)), t)               //size
	chk(binary.Write(f, binary.BigEndian, int16(1)), t)                //channels
	chk(binary.Write(f, binary.BigEndian, int32(0)), t)                //number of samples
	chk(binary.Write(f, binary.BigEndian, int16(32)), t)               //bits per sample
	_, err = f.Write([]byte{0x40, 0x0e, 0xac, 0x44, 0, 0, 0, 0, 0, 0}) //80-bit sample rate 44100
	chk(err, t)

	// sound chunk
	_, err = f.WriteString("SSND")
	chk(err, t)
	chk(binary.Write(f, binary.BigEndian, int32(0)), t) //size
	chk(binary.Write(f, binary.BigEndian, int32(0)), t) //offset
	chk(binary.Write(f, binary.BigEndian, int32(0)), t) //block
	nSamples := 0
	defer func() {
		// fill in missing sizes
		totalBytes := 4 + 8 + 18 + 8 + 8 + 4*nSamples
		_, err = f.Seek(4, 0)
		chk(err, t)
		chk(binary.Write(f, binary.BigEndian, int32(totalBytes)), t)
		_, err = f.Seek(22, 0)
		chk(err, t)
		chk(binary.Write(f, binary.BigEndian, int32(nSamples)), t)
		_, err = f.Seek(42, 0)
		chk(err, t)
		chk(binary.Write(f, binary.BigEndian, int32(4*nSamples+8)), t)
		chk(f.Close(), t)
	}()

	portaudio.Initialize()
	defer portaudio.Terminate()
	in := make([]int32, 64)
	stream, err := portaudio.OpenDefaultStream(1, 0, 44100, len(in), in)
	chk(err, t)
	defer stream.Close()

	chk(stream.Start(), t)
	for {
		chk(stream.Read(), t)
		chk(binary.Write(f, binary.BigEndian, in), t)
		nSamples += len(in)
		select {
		case <-sig:
			return
		default:
		}
	}
}
