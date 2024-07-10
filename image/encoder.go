package image

import(
	"bytes"
	"encoding/base64"
	"log/slog"
	"image"
	"image/png"
)

func Encode(img image.Image) (string, error) {
	buffer := new(bytes.Buffer)

	if err := png.Encode(buffer, img); err != nil {
		slog.Error("jawi-ocr/lib", "package", "image", "function", "Encode")
		return "", err
	}

	encoded := base64.StdEncoding.EncodeToString(buffer.Bytes())

	return encoded, nil
}

func EncodeToBytes(img image.Image) ([]byte, error) {
	buffer := new(bytes.Buffer)

	if err := png.Encode(buffer, img); err != nil {
		slog.Error("jawi-ocr/lib", "package", "image", "function", "EncodeToBytes")
		return nil, err
	}

	return buffer.Bytes(), nil
}