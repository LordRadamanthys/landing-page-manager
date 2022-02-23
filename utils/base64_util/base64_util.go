package base64_util

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"mime/multipart"
	"net/http"
)

func ConvertImageToBase64(file multipart.File) (string, error) {
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err

	}
	contentType := http.DetectContentType(data)
	switch contentType {
	case "image/png":
		fmt.Println("")
	case "image/jpeg":
		img, err := jpeg.Decode(bytes.NewReader(data))
		if err != nil {
			return "", errors.New("unable to decode jpeg")
		}

		var buf bytes.Buffer
		if err := png.Encode(&buf, img); err != nil {
			return "", errors.New("unable to decode jpeg")
		}
		data = buf.Bytes()
	default:
		return "", errors.New("unsuported jpeg")
	}
	return base64.StdEncoding.EncodeToString(data), nil
}
