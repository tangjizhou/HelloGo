package base

import (
	"fmt"
)

func Encrypt() {
	var path string
	defer func() {
		if e := recover(); e != nil {
			fmt.Println("错误被捕获", e)
		}
	}()

	_, err := fmt.Scanf("请输入要加密的位置: %s", &path)
	fmt.Println("当前扫描的路径为", path)
	if err != nil {
		panic("文件路径下不存在任何文件")
	}

}
