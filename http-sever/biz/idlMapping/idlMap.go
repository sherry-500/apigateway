package idlMapping

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"

	// "golang.org/x/tools/go/analysis/passes/nilfunc"
)

var IDLMap = make(map[string]string)

func AddIdl(svcName string, idlPath string) bool {
    IDLMap[string(svcName)] = string(idlPath)
	return true
}

func DelIdl(svcName string) bool {
    _, ok := IDLMap[svcName]
    if !ok  {
        return false
    }
    delete(IDLMap, svcName)
    return true
}
//update both the IDLMap and idlPath.txt
func UpdateIdl(svcName string, idlPath string) bool {
    _, ok := IDLMap[svcName]
    if !ok  {
        return false
    }
    delete(IDLMap, svcName)
    IDLMap[string(svcName)] = string(idlPath)

    return true
}

func GetIdl(svcName string) string {
    path, _ := IDLMap[svcName]
    return path
}

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

func WriteBack() {
    idlFile, err := os.OpenFile("./idlPath.txt", os.O_RDWR|os.O_CREATE, 777)
	if err != nil {
		fmt.Println(err)
	}
    idlFile.Truncate(0)
    idlFile.Seek(0, 0)
    writer := bufio.NewWriter(idlFile)
    for svcName, idlPath := range IDLMap {
        writer.WriteString(svcName + "=" + idlPath)
    }
    writer.Flush()
}