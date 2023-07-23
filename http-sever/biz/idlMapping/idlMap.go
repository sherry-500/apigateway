package idlMapping

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
)

var IDLMap = make(map[string]string)

func AddIdl(svcName string, idlPath string) bool {
	

	return true
}

func DelIdl(svcName string)
//update both the IDLMap and idlPath.txt
func UpdateIdl(svcName string, idlPath string)

func GetIdl(svcName string) string

func InitMap() {
	idlFile, err := os.OpenFile("./idlPath.txt", os.O_RDWR|os.O_CREATE, 777)
	if err != nil {
		fmt.Println(err)
	}
	reader := bufio.NewReader(idlFile)
    //按行读取
    for {
       line, _, err := reader.ReadLine()
       if err == io.EOF {
           break
       }
       if err != nil {
           fmt.Println(err)
       }
       data := bytes.Split(line, []byte("="))
       IDLMap[string(data[0])] = string(data[1])
    }
}