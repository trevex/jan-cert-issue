package main

import (
	"context"
	"log"

	"gocloud.dev/blob"
	_ "gocloud.dev/blob/gcsblob"
)

func main() {
	ctx := context.Background()
	bucket, err := blob.OpenBucket(ctx, "gs://my-bucket")
	if err != nil {
		log.Fatal(err)
	}
	defer bucket.Close()
	log.Print(bucket)
	r, err := bucket.NewReader(ctx, "foo.txt", nil)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()
	log.Println("Content-Type:", r.ContentType())
}
