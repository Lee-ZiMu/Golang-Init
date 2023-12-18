/**
 * @Description: 文件操作
 * @Author Lee
 * @Date 2023/12/13 13:26
 **/

package file

import (
	"io/fs"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

// IsExists 文件是否存在
func IsExists(path string) (os.FileInfo, bool) {
	f, err := os.Stat(path)
	return f, err == nil || os.IsExist(err)
}

// MkdirAll 创建文件夹
func MkdirAll(path string) error {
	if _, ok := IsExists(path); !ok {
		return os.MkdirAll(path, fs.ModePerm)
	}

	return nil
}

func SaveUploadFile(c *gin.Context, file *multipart.FileHeader, dst string) error {
	if err := MkdirAll(filepath.Dir(dst)); err != nil {
		return err
	}
	return c.SaveUploadedFile(file, dst)
}

func RemoveAll(path string) error {
	if _, ok := IsExists(path); !ok {
		return nil
	}

	return os.RemoveAll(path)
}

func Remove(path string) error {
	if _, ok := IsExists(path); !ok {
		return nil
	}

	return os.Remove(path)
}
