package main

import (
	"fmt"
	"math"
	"os"

	"github.com/jfreymuth/pulse"
)

func main() {
	c, err := pulse.NewClient()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer c.Close()

	s, err := c.DefaultSink()
	if err != nil {
		fmt.Println(err)
		return
	}

	stream, err := c.NewPlayback(synth, pulse.PlaybackLowLatency(s))
	if err != nil {
		fmt.Println(err)
		return
	}

	stream.Start()

	fmt.Print("Press enter to stop...")
	os.Stdin.Read([]byte{0})

	stream.Close()
}

var t, phase float32

func synth(out []float32) {
	for i := range out {
		x := float32(math.Sin(2 * math.Pi * float64(phase)))
		out[i] = x * 0.1
		f := [...]float32{440, 550, 660, 880}[int(2*t)&3]
		phase += f / 44100
		if phase >= 1 {
			phase--
		}
		t += 1. / 44100
	}
}
