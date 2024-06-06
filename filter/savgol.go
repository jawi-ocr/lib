package filter

import (
	"encoding/json"
	"golang.org/x/exp/constraints"

	"github.com/mushoffa/go-library/http"
)

var (
	client = http.NewHttpClient()
)

type savgol struct {
	WindowLength uint `json:"window_length"`
	Polyorder uint `json:"polyorder"`
	Mode string `json:"mode"`
	Data []uint16 `json:"data"`
}

type savgolResponse[T constraints.Float] struct {
	Data []T `json:"data"`
}

func Savgol[T constraints.Float](window_length, polyorder uint, mode string, data []uint16, url string) ([]T, error) {

	request := savgol{
		WindowLength: window_length,
		Polyorder: polyorder,
		Mode: mode,
		Data: data,
	}

	response := savgolResponse[T]{}

	res, err := client.Post(url, &request, nil)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(res, &response); err != nil {
		return nil, err
	}

	return response.Data, nil
}