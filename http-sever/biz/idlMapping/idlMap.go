package idlMapping

import (
	"bufio"
	"fmt"
	"os"
    "strings"
    // "io"
    // "log"
    // "bytes"

	// "golang.org/x/tools/go/analysis/passes/nilfunc"
)

var IDLMap = make(map[string]string)

func AddIdl(svcName string, idlPath string) bool {
    _, ok := IDLMap[svcName]
    if ok {
        return false
    }
    IDLMap[string(svcName)] = string(idlPath)
    WriteBack()
	return true
}

func DelIdl(svcName string) bool {
    _, ok := IDLMap[svcName]
    if !ok  {
        return false
    }
    delete(IDLMap, svcName)
    WriteBack()
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
    WriteBack()
    return true
}

func GetIdl(svcName string) string {
    path, _ := IDLMap[svcName]
    return path
}

func InitMap() {
    idlFile, err := os.Open("./idlPath.txt")
    if err != nil{
        fmt.Println("Error opening file:", err)
        return
    }
    scanner := bufio.NewScanner(idlFile)
    for scanner.Scan() {
        line := scanner.Text()
        data := strings.Split(line, ",")
        IDLMap[data[0]] = data[1]
    }
    idlFile.Close()

	// idlFile, err := os.OpenFile("./idlPath.txt", os.O_RDWR|os.O_CREATE, 0777)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// reader := bufio.NewReader(idlFile)
    // //按行读取
    // for {
    //    line, _, err := reader.ReadLine()
    //    if err == io.EOF {
    //        break
    //    }
    //    if err != nil {
    //        fmt.Println(err)
    //    }
    //    data := bytes.Split(line, []byte(","))
    //    IDLMap[string(data[0])] = string(data[1])
    // }
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
        writer.WriteString(svcName + "," + idlPath + "\n")
    }
    writer.Flush()
}