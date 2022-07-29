package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"gocloud.dev/blob"
	_ "gocloud.dev/blob/gcsblob"
)

func main() {
	ctx := context.Background()
	bucket, err := blob.OpenBucket(ctx, "gs://nvoss-jan-problem")
	if err != nil {
		log.Fatal(err)
	}
	defer bucket.Close()
	log.Print(bucket)
	r, err := bucket.NewReader(ctx, "dog.png", nil)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()
	log.Println("Content-Type:", r.ContentType())

	iter := bucket.List(nil)
	for {
		obj, err := iter.Next(ctx)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(obj.Key)
	}
}
