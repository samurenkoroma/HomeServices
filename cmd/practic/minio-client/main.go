package main

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"log"
)

func main() {
	endpoint := "minio-api.lab.raspi"
	accessKeyID := "minioadmin"
	secretAccessKey := "minioadmin"
	useSSL := false

	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:           credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure:          useSSL,
		TrailingHeaders: true,
	})
	if err != nil {
		log.Fatalln(err)
	}
	//buckets, err := minioClient.ListBuckets(context.Background())
	//if err != nil {
	//	return
	//}
	ctx := context.Background()

	uploadInfo, err := minioClient.FPutObject(ctx, "documents", "awdawd/my-objectname", "test.json", minio.PutObjectOptions{
		ContentType: "application/json",
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(uploadInfo)

	object, err := minioClient.GetObject(ctx, "documents", "my-objectname", minio.GetObjectOptions{})
	if err != nil {
		return
	}
	stat, err := object.Stat()
	if err != nil {
		return
	}

	buffer := make([]byte, stat.Size)
	object.Read(buffer)
	fmt.Printf("OBJECT: %+v \n", object)
	fmt.Printf("STAT: %+v \n", stat)
	fmt.Printf("%s", buffer)

}
