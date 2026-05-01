package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"
)

func SaveImage(file multipart.File, fileHeader *multipart.FileHeader) (string, error) {
	data, err := io.ReadAll(file)
	if err != nil {
		return "", fmt.Errorf("ReadAll() failed,err:%w", err)
	}
	//生成年月目录
	now := time.Now()
	yearMonth := now.Format("2006/01")
	dirPath := filepath.Join("./uploads/images", yearMonth)
	newFilename := generateFileName(fileHeader.Filename)
	filePath := filepath.Join(dirPath, newFilename)

	dir := filepath.Dir(filePath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return "", fmt.Errorf("MkdirAll() failed,err:%w", err)
	}
	err = os.WriteFile(filePath, data, 0644)
	if err != nil {
		return "", fmt.Errorf("WriteFile() failed,err:%w", err)
	}
	return "/uploads/images/" + yearMonth + "/" + newFilename, nil
}

func generateFileName(BaseFilename string) string {
	//获取图片后缀
	ext := filepath.Ext(BaseFilename)

	//生成时间戳+MD5随机字符串
	timestamp := time.Now().UnixNano()
	hash := md5.Sum([]byte(fmt.Sprintf("%d%s", timestamp, BaseFilename)))
	randomStr := hex.EncodeToString(hash[:])[:16]

	//组合文件名
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
