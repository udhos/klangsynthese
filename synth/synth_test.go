package synth

import (
	"testing"
	"time"

	"github.com/200sc/klangsynthese/wav"
	"github.com/stretchr/testify/assert"
)

func TestSinWav(t *testing.T) {
	testWave(t, Sin(A4, 2, 32))
}

func TestSquareWav(t *testing.T) {
	testWave(t, Square(A4, 2, 32))
}

func TestSawWav(t *testing.T) {
	testWave(t, Saw(A4, 2, 32))
}

func TestTriangleWav(t *testing.T) {
	testWave(t, Triangle(A4, 2, 32))
}

func TestPulse(t *testing.T) {
	testWave(t, Pulse(A4, 2, 32, 8))
}

func TestAdd(t *testing.T) {
	testWave(t,
		//	i.e harmonics
		Add(Sin(A4, 2, 16),
			Sin(A4*2, 2, 16),
			Sin(A4*3, 2, 16),
			Sin(A4*4, 2, 16),
			Sin(A4*5, 2, 16),
			Sin(A4*6, 2, 16),
		))
}

func testWave(t *testing.T, wave []byte) {
	a, err := wav.NewController().Wave(wave)
	assert.Nil(t, err)
	err = <-a.Play()
	assert.Nil(t, err)
	time.Sleep(2 * time.Second)
}
