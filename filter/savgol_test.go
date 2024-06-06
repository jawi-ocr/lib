package filter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSavgol_Success(t *testing.T) {
	// Given
	data := []uint16{0,0,0,0,0,0,3,4,4,7,10,10,10,13,14,19,16,19,18,26,24,27,20,21,26,41,41}
	url := "http://localhost:9998/filter/savitzky_golay"

	// When
	res, err := Savgol[float32](9,1,"interp",data, url)

	// Then
	assert.NotNil(t, res)
	assert.NoError(t, err)
	assert.Equal(t, len(data), len(res))
}