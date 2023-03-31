package files

import (
	"bufio"
	"bytes"
	"image"
	"image/png"
	"io"
	"os"
)

func Load(path string) (raw []byte, err error) {
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		return raw, err
	}

	return io.ReadAll(file)
}

func LoadPng(path string) (img image.Image, err error) {
	raw, err := Load(path)
	if err != nil {
		return nil, err
	}

	buf := bytes.NewBuffer(raw)
	return png.Decode(buf)
}

func Save(path string, img image.Image) (err error) {
	var buf bytes.Buffer

	err = png.Encode(&buf, img)
	if err != nil {
		return err
	}

	file, err := os.Create(path)
	defer file.Close()
	if err != nil {
		return err
	}

	writer := bufio.NewWriter(file)
	writer.Write(buf.Bytes())
	return nil
}
