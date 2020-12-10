package cmd
import (
"bufio"
"bytes"
"fmt"
	"io"
	"log"
	"os"
"os/exec"
"strconv"
)

const log1 =  "/Users/jianghan/Hello_world/cmd1.log"
const log2 =  "/Users/jianghan/Hello_world/cmd2.log"

func ReadFile(){
	ch := make(chan int, 1)
	//write log
	go func() {
		writeLog(ch)
	}()
	cmdStr1 := "/Users/jianghan/Hello_world/test.log"
	cmd1 := exec.Command("tail","-f",cmdStr1)
	ch1 := make(chan []byte,4096)
	stdout1, err := cmd1.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	if err := cmd1.Start(); err != nil {
		log.Fatal(err)
	}

	go func() {
		reader := bufio.NewReader(stdout1)
		for {
			strline, err := readLine(reader)
			if err != nil && err != io.EOF {
				log.Fatal(err)
			}
			if len(strline) > 0 {
				fmt.Printf("Read message: %s\n", strline)
				ch1 <- []byte(strline)
			}

			if err == io.EOF {
				fmt.Println("EOF. No more messages")
				log.Println("[DEBUG] Exiting IO read")
				break
			}

			fmt.Println("Waiting for the next message...")
		}

		log.Println("[DEBUG] Exiting go func()")
	}()


	/*
	//cmd1 output2log
	go func() {
		for  {
			content,_ := buf1.ReadBytes('\n')
			content = append(content, 'b','y','c','m','d','1')
			fmt.Println(content)
			ch1 <- content
		}
	}()
	//cmd2 output2log
	go func() {
		for  {
			cmd2 := exec.Command(cmdStr2)
			cmd2.Stdout = &buf2
			content,_ := buf2.ReadBytes('\n')
			content = append(content,'b','y','c','m','d','2')
			fmt.Println(content)
			ch2 <- content
		}
	}()
	*/
	//write to log1
	go func() {
		file1, _ := os.OpenFile(log1,os.O_WRONLY,0666)
		for d := range ch1{
			_,err1 := fmt.Fprint(file1,d)
			if err1 != nil{
				fmt.Println(err1)
				file1.Close()
			}
		}
	}()
	/*
	//write to log2
	go func() {
		file2, _ := os.OpenFile(log2,os.O_WRONLY,0666)
		for d := range ch1{
			_,err := fmt.Fprint(file2,d)
			if err != nil{
				fmt.Println(err)
				file2.Close()
			}
		}
	}()
	*/
	<- ch
}



func writeLog(ch chan int) {
	path := "/Users/jianghan/Hello_world/test.log"
	logFile, err := os.OpenFile(path,os.O_WRONLY,0666)
	if err != nil {
		fmt.Println("文件打开失败", err)
	}
	defer logFile.Close()

	write := bufio.NewWriter(logFile)
	for i := 0;; i++ {
		str := strconv.Itoa(i) + "xxxxxxxxxxxxx \n"
		write.WriteString(str)
		write.Flush()
		if i==400{
			ch <- 1
			break
		}
	}
}


func readLine(reader *bufio.Reader) (strLine string, err error) {
	buffer := new(bytes.Buffer)
	for {
		var line []byte
		var isPrefix bool
		log.Println("[INFO] Start reading line...")
		line, isPrefix, err = reader.ReadLine()

		log.Printf("[DEBUG] Read Len: %d, isPrefix: %t, Error: %v\n", len(line), isPrefix, err)

		if err != nil && err != io.EOF {
			return "", err
		}

		buffer.Write(line)

		if !isPrefix {
			log.Println("[INFO] EOL found")
			break
		}
	}

	log.Println("[DEBUG] End of line")
	return buffer.String(), err
}


