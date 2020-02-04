package main

import (
	"bufio"
	"log"
	"net"
	"os"
)

func main() {
	cn, err := net.Dial("tcp", ":8000")
	if err != nil {
		log.Println(err)
		return
	}

	reader := bufio.NewReader(os.Stdin)

	input := make([]byte, 1024)
	var inputStr string
	for {
		// 标准输入读取
		n, _ := reader.Read(input)
		if n > 0 {
			k, _ := cn.Write(input[:n])
			if k > 0 {
				inputStr = string(input[:k])
				log.Println(inputStr)
			}
		}

	}

}
