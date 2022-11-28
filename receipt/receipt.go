package receipt

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

var ReceiptDirectory string = filepath.Join("uploads")

type Receipt struct {
	ReceiptName string    `json:name`
	UploadDate  time.Time `json:uploadDate`
}

func GetReceipts() ([]Receipt, error) {
	reciepts := make([]Receipt, 0)
	files, err := os.ReadDir(ReceiptDirectory)
	if err != nil {
		return nil, err
	}
	for _, f := range files {
		var upload time.Time
		fileInfo, err := f.Info()
		if err != nil {
			fmt.Println("Error fetching file info")
		} else {
			upload = fileInfo.ModTime()
		}
		reciepts = append(reciepts, Receipt{ReceiptName: f.Name(), UploadDate: upload})
	}
	return reciepts, nil
}
