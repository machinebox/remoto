package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"

	"github.com/machinebox/remoto/examples/files/client/files"
	"github.com/pkg/errors"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

func run() error {
	// see https://medium.com/@matryer/make-ctrl-c-cancel-the-context-context-bd006a8ad6ff
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	defer func() {
		signal.Stop(c)
		cancel()
	}()
	go func() {
		select {
		case <-c:
			cancel()
		case <-ctx.Done():
		}
	}()

	f, err := os.Open(os.Args[1])
	if err != nil {
		return err
	}
	defer f.Close()
	images := files.NewImagesClient("http://localhost:8080", http.DefaultClient)
	request := &files.FlipRequest{}
	ctx = request.SetImage(ctx, filepath.Base(os.Args[1]), f)
	resp, err := images.Flip(ctx, request)
	if err != nil {
		return errors.Wrap(err, "images.Flip")
	}
	flipped, err := resp.FlippedImage.Open(ctx)
	if err != nil {
		return errors.Wrap(err, "open file from response")
	}
	defer flipped.Close()
	outfile := filepath.Join(os.Args[1], "flipped")
	out, err := os.Create(outfile)
	if err != nil {
		return err
	}
	defer out.Close()
	if _, err := io.Copy(out, flipped); err != nil {
		return errors.Wrap(err, "writing file")
	}
	log.Println("flipped image saved to", outfile)
	return nil
}
