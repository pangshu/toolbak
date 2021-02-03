package File

import (
	"errors"
	"os"
)

// AppendFile 插入文件内容.
func (toolFile *File)AppendFile(filePath string, data string) error {
	if filePath == "" {
		return errors.New("No path provided")
	}

	var file *os.File
	filePerm, err := toolFile.Mode(filePath)
	if err != nil {
		// create the file
		file, err = os.Create(filePath)
	} else {
		// open for append
		file, err = os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, filePerm)
	}
	if err != nil {
		// failed to create or open-for-append the file
		return err
	}

	defer func() {
		_ = file.Close()
	}()

	_, err = file.Write([]byte(data))

	return err
}
