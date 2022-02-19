package fileUpload

import (
	"errors"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

/*
	UpLoad 文件上传和存储
		参数说明:
			r: http request
			parentDir: 存放文件的子目录
*/
func UpLoad(r *http.Request, parentDir string) (err error, fileName string, filePath string, dir string) {
	err = nil
	for {
		if r == nil {
			err = errors.New("nil request")
			break
		}
		errP := r.ParseMultipartForm(32 << 20)
		if errP != nil {
			log.Printf("parse file form failed, err: %v", errP)
			err = errors.New("parse file failed")
			break
		}

		file, handler, errU := r.FormFile("file")
		if errU != nil {
			log.Printf("file mdl get form data failed, err: %v", errU)
			err = errors.New("get form data failed")
			break
		}
		defer file.Close()

		log.Printf("file mdl header: %v", handler.Header)
		// 文件路径(推荐使用配置文件进行配置)
		curPath, _ := filepath.Abs(filepath.Dir(os.Args[0]))
		dir = curPath
		fileName = handler.Filename
		filePath = filepath.Join(dir, parentDir, fileName)
		f, errU := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666) // 如果文件存在会覆盖写
		if errU != nil {
			log.Printf("file open or create file failed, err: %v", errU)
			err = errors.New("write file failed")
			break
		}
		defer f.Close()

		// 使用io.Copy的方式优化内存使用
		_, errU = io.Copy(f, file)
		if errU != nil {
			log.Printf("store file failed, err: %v", errU)
			err = errors.New("store file data failed")
			break
		}
		err = nil

		break
	}
	return err, fileName, filePath, dir
}
