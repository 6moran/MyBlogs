package minio

import (
	"bytes"
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"mime/multipart"
	"path/filepath"
	"time"

	"MyBlogs/pkg/config"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var minioClient *minio.Client
var bucketName string

// InitMinIO 初始化 MinIO 客户端
func InitMinIO(cfg *config.MinioConfig) error {
	var err error
	minioClient, err = minio.New(cfg.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.AccessKey, cfg.SecretKey, ""),
		Secure: cfg.UseSSL,
	})
	if err != nil {
		return fmt.Errorf("创建 MinIO 客户端失败: %w", err)
	}

	bucketName = cfg.BucketName

	// 检查 bucket 是否存在，不存在则创建
	exists, err := minioClient.BucketExists(context.Background(), bucketName)
	if err != nil {
		return fmt.Errorf("检查 bucket 失败: %w", err)
	}
	if !exists {
		err = minioClient.MakeBucket(context.Background(), bucketName, minio.MakeBucketOptions{})
		if err != nil {
			return fmt.Errorf("创建 bucket 失败: %w", err)
		}
	}

	return nil
}

// SaveImage 上传图片到 MinIO
func SaveImage(file multipart.File, fileHeader *multipart.FileHeader, articleID int) (string, error) {
	data, err := io.ReadAll(file)
	if err != nil {
		return "", fmt.Errorf("读取文件失败: %w", err)
	}

	newFilename := generateFileName(fileHeader.Filename)
	objectName := fmt.Sprintf("articles/%d/%s", articleID, newFilename)

	contentType := fileHeader.Header.Get("Content-Type")
	if contentType == "" {
		contentType = "application/octet-stream"
	}

	_, err = minioClient.PutObject(
		context.Background(),
		bucketName,
		objectName,
		bytes.NewReader(data),
		int64(len(data)),
		minio.PutObjectOptions{ContentType: contentType},
	)
	if err != nil {
		return "", fmt.Errorf("上传到 MinIO 失败: %w", err)
	}

	return fmt.Sprintf("/%s/%s", bucketName, objectName), nil
}

func generateFileName(BaseFilename string) string {
	ext := filepath.Ext(BaseFilename)
	timestamp := time.Now().UnixNano()
	hash := md5.Sum([]byte(fmt.Sprintf("%d%s", timestamp, BaseFilename)))
	randomStr := hex.EncodeToString(hash[:])[:16]
	return fmt.Sprintf("%d%s%s", timestamp, randomStr, ext)
}

func IsImageFile(filename string) bool {
	ext := filepath.Ext(filename)
	switch ext {
	case ".jpg", ".jpeg", ".png", ".webp", ".gif":
		return true
	}
	return false
}
