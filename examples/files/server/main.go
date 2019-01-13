package main

import (
	"bytes"
	"context"
	"fmt"
	"image/jpeg"
	"os"

	"github.com/disintegration/imaging"
	"github.com/machinebox/remoto/examples/files/server/files"
	"github.com/machinebox/remoto/remototypes"
	"github.com/pkg/errors"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

func run() error {
	addr := "0.0.0.0:8080"
	fmt.Println("listening on", addr)
	err := files.Run(addr, service{})
	if err != nil {
		return err
	}
	return nil
}

type service struct{}

func (service) Flip(ctx context.Context, req *files.FlipRequest) (*remototypes.FileResponse, error) {
	f, err := req.Image.Open(ctx)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	img, err := imaging.Decode(f)
	if err != nil {
		return nil, err
	}
	rotated := imaging.Rotate180(img)
	var buf bytes.Buffer
	if err := jpeg.Encode(&buf, rotated, &jpeg.Options{Quality: 100}); err != nil {
		return nil, errors.Wrap(err, "encode jpeg")
	}
	resp := &remototypes.FileResponse{
		Filename:      "flipped.jpg",
		Data:          &buf,
		ContentLength: buf.Len(),
		ContentType:   "image/jpeg",
	}
	return resp, nil
}
