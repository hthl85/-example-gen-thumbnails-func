package thumbnails

import (
	"cloud.google.com/go/storage"
	"context"
	"fmt"
	"github.com/disintegration/imaging"
	"image/png"
	"log"
)

// Global API clients used across function invocations.
var (
	storageClient *storage.Client
)

func init() {
	// Declare a separate err variable to avoid shadowing the client variables.
	var err error

	storageClient, err = storage.NewClient(context.Background())
	if err != nil {
		log.Fatalf("storage.NewClient: %v", err)
	}
}

// GCSEvent is the payload of a GCS event.
type GCSEvent struct {
	Name           string `json:"name"`
	Bucket         string `json:"bucket"`
	Metageneration string `json:"metageneration"`
}

// GenThumbnails consumes a GCS event.
func GenThumbnails(ctx context.Context, e GCSEvent) error {
	name := e.Name
	inputBucket := e.Bucket
	outputBucket := "thumbnails-storage"

	inputBlob := storageClient.Bucket(inputBucket).Object(name)
	r, err := inputBlob.NewReader(ctx)
	if err != nil {
		return fmt.Errorf("NewReader: %v", err)
	}

	outputBlob := storageClient.Bucket(outputBucket).Object(name)
	w := outputBlob.NewWriter(ctx)
	defer w.Close()

	img, err := png.Decode(r)
	if err != nil {
		return fmt.Errorf("decode: %v", err)
	}

	thumb := imaging.Resize(img, 100, 0, imaging.CatmullRom)

	err = png.Encode(w, thumb)
	if err != nil {
		return fmt.Errorf("encode: %v", err)
	}

	return nil
}
