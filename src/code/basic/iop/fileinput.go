package iop

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

//FileIn 文件输入demo
func FileIn() {
	inputfile, inputerror := os.Open("E:/hello.txt")
	if inputerror != nil {
		fmt.Print(inputerror)
		return
	}
	//定义延迟调用函数用于关闭文件
	defer inputfile.Close()

	inputReader := bufio.NewReader(inputfile)

	for {
		inputString, readerError := inputReader.ReadString('\n')
		fmt.Print(inputString)
		if readerError == io.EOF {
			return
		}
	}
}

//FileInWithBuf 使用缓冲区读
func FileInWithBuf() {
	inputfile, inputerror := os.Open("E:/hello.txt")
	if inputerror != nil {
		fmt.Print(inputerror)
		return
	}
	//定义延迟调用函数用于关闭文件
	defer inputfile.Close()

	inputReader := bufio.NewReader(inputfile)

	for {
		buf := make([]byte, 1024)
		//Read 返回字节数
		inputString, _ := inputReader.Read(buf)
		if inputString == 0 {
			return
		}
		fmt.Print(string(buf))
	}
	//输出带有空行应该

}

//FileInWithCol 数据是按照列排列的
func FileInWithCol() {
	file, err := os.Open("E:/hello.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var col1, col2, col3 []string
	for {
		var v1, v2, v3 string
		_, err := fmt.Fscanln(file, &v1, &v2, &v3)
		// scans until newline
		if err != nil {
			break
		}
		col1 = append(col1, v1)
		col2 = append(col2, v2)
		col3 = append(col3, v3)
	}

	fmt.Println(col1)
	fmt.Println(col2)
	fmt.Println(col3)
}

//ReadAllCharsFromFileWriteToAnotherFile 将文件中的所有字符读到一个字符串中
//然后写到另一个文件中
func ReadAllCharsFromFileWriteToAnotherFile() {
	inputFile := "E:/hello.txt"
	outputFile := "E:/hellocopy.txt"

	buf, err := ioutil.ReadFile(inputFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "File Error:%s\n", err)
	}
	fmt.Printf("%s\n", string(buf))
	err = ioutil.WriteFile(outputFile, buf, 0644)
	if err != nil {
		panic(err.Error)
	}
}
