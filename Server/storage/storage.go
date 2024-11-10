package storage

import (
	"fmt"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"

	"github.com/Raffique/JL_Adventure_Tours/Server/config"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gin-gonic/gin"
)

var s3Client *s3.S3
var bucketName string

// InitStorage initializes storage based on environment
func InitStorage() {
    if config.IsProd {
        initS3()
    } else {
        createLocalUploadFolder()
    }
}

// Initialize S3 client for production
func initS3() {
    bucketName = os.Getenv("AWS_S3_BUCKET_NAME")
    s3Config := &aws.Config{
        Region:      aws.String(os.Getenv("AWS_REGION")),
        Credentials: credentials.NewStaticCredentials(os.Getenv("AWS_ACCESS_KEY_ID"), os.Getenv("AWS_SECRET_ACCESS_KEY"), ""),
    }
    sess := session.Must(session.NewSession(s3Config))
    s3Client = s3.New(sess)
}

// Create the local uploads folder if it doesn’t exist for development
func createLocalUploadFolder() {
    if _, err := os.Stat("uploads"); os.IsNotExist(err) {
        os.Mkdir("uploads", os.ModePerm)
    }
}

// UploadFile handles file uploads, choosing between S3 and local storage
func UploadFile(c *gin.Context, file *multipart.FileHeader) (string, error) {
    if config.IsProd {
        return uploadFileToS3(file)
    }
    return uploadFileToLocal(c, file)
}

// Uploads the file to AWS S3 and returns the file URL
func uploadFileToS3(file *multipart.FileHeader) (string, error) {
    src, err := file.Open()
    if err != nil {
        return "", err
    }
    defer src.Close()

    key := fmt.Sprintf("%d-%s", time.Now().Unix(), file.Filename)
    _, err = s3Client.PutObject(&s3.PutObjectInput{
        Bucket: aws.String(bucketName),
        Key:    aws.String(key),
        Body:   src,
        ContentType: aws.String(file.Header.Get("Content-Type")),
    })
    if err != nil {
        return "", err
    }

    fileURL := fmt.Sprintf("https://%s.s3.amazonaws.com/%s", bucketName, key)
    return fileURL, nil
}

// Uploads the file to local storage and returns the file URL
func uploadFileToLocal(c *gin.Context, file *multipart.FileHeader) (string, error) {
    filePath := filepath.Join("uploads", file.Filename)
    if err := c.SaveUploadedFile(file, filePath); err != nil {
        return "", err
    }
    fileURL := fmt.Sprintf("http://localhost:8080/uploads/%s", file.Filename)
    return fileURL, nil
}
