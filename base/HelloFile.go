package base

import (
	"fmt"
	"io/ioutil"
	"os"
)

var filesChannel = make(chan os.FileInfo, 100)
var stopChannel = make(chan bool)

func Encrypt() {
	var path string
	defer func() {
		if e := recover(); e != nil {
			fmt.Println("错误被捕获", e)
		}
	}()

	_, _ = fmt.Scanf("%s", &path)
	_, _ = fmt.Printf("扫描文件的路径为：%s\n", path)

	go addToChannel(path, filesChannel)
	readFromChannel(filesChannel)

}

var deep = 0
var fileCount = 0
var readCount = 0

func addToChannel(path string, filesChannel chan os.FileInfo) {
	deep++
	files, _ := ioutil.ReadDir(path)
	//fmt.Println("当前路径", path, "当前路径文件数", len(files))
	for _, file := range files {
		if file.IsDir() {
			addToChannel(path+"/"+file.Name(), filesChannel)
		}
		filesChannel <- file
		fileCount++
	}

	deep--
	if deep == 0 {
		fmt.Println("写入文件：", fileCount)
		stopChannel <- true
	}
}

func readFromChannel(filesChannel chan os.FileInfo) {
	for true {
		select {
		case file := <-filesChannel:
			encrypt(file)
			readCount++
		case <-stopChannel:
			for file := range filesChannel {
				encrypt(file)
				readCount++
			}
			goto end
		}
	}
end:
	fmt.Println("文件加密完毕", readCount)
	close(filesChannel)
	close(stopChannel)
}

func encrypt(file os.FileInfo) {
	logEncryptInfo(file.Name(), file.Name())
}

func logEncryptInfo(filename string, path string) {
	fmt.Printf("加密文文件名：%s\n", filename)
}
