package base

import (
	"fmt"
	"io/ioutil"
	"os"
)

var dirs = make([]os.FileInfo, 10)
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

	files, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}

	copy(dirs, files)

	go addToChannel(path, filesChannel)
	readFromChannel(filesChannel)
	close(filesChannel)
	close(stopChannel)

}

var deep = 0

func addToChannel(path string, filesChannel chan os.FileInfo) {
	files, _ := ioutil.ReadDir(path)
	for _, file := range files {
		if file.IsDir() {
			deep++
			addToChannel(path+"/"+file.Name(), filesChannel)
		} else {
			filesChannel <- file
		}
	}
	deep--
	if deep == 0 {
		stopChannel <- true
	}
}

func readFromChannel(filesChannel chan os.FileInfo) {
	for true {
		select {
		case file := <-filesChannel:
			encrypt(file)
		case <-stopChannel:
			goto end
		}
	}
end:
	fmt.Println("文件加密完毕")

}

func encrypt(file os.FileInfo) {
	logEncryptInfo(file.Name(), file.Name())
}

func logEncryptInfo(filename string, path string) {
	fmt.Printf("加密文文件名：%s，路径：%s\n", filename, path)
}
