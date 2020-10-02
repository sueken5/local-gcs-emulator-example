package main

import (
	"context"
	"io"
	"log"
	"os"

	"cloud.google.com/go/storage"
	"github.com/satori/uuid"
)

func main() {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatal(err)
	}

	bktName := "sample"
	objName := uuid.NewV4().String()

	bkt := client.Bucket(bktName)
	obj := bkt.Object(objName)

	w := obj.NewWriter(ctx)

	if _, err := w.Write([]byte("hello world")); err != nil {
		log.Fatal(err)
	}

	if err := w.Close(); err != nil {
		log.Fatal(err)
	}

	log.Println("success upload")

	r, err := obj.NewReader(ctx)
	if err != nil {
		log.Fatal(err)
	}

	if _, err := io.Copy(os.Stdout, r); err != nil {
		log.Fatal(err)
	}

	if err := client.Close(); err != nil {
		log.Fatal(err)
	}

	log.Println("success read")
	log.Println("public url: http://localhost:4443/" + bktName + "/" + objName)
}
