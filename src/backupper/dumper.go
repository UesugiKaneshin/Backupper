package backupper

import (
	"bufio"
	"io"
	"log"
	"os"
	"os/exec"
)

type Dumper struct {
	host         string
	port         string
	user         string
	password     string
	databasename string
	filepath     string
}

func NewDumper() *Dumper {
	this := new(Dumper)

	return this
}

func (this *Dumper) SetDefaultValue(host string, port string, user string, password string, databasename string, filepath string) {
	this.host = host
	this.port = port
	this.user = user
	this.password = password
	this.databasename = databasename
	this.filepath = filepath
}

func (this *Dumper) Download() bool {
	result := true

	host := "-h" + this.host
	port := "-P" + this.port
	user := "-u" + this.user
	password := "-p" + this.password
	databasename := this.databasename
	filepath := this.filepath

	cmd := exec.Command("mysqldump", "--opt", "--quick", host, port, user, password, databasename)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		result = false
		log.Fatal(err)
	} else {
	}

	if err := cmd.Start(); err != nil {
		result = false
		log.Fatal(err)
	} else {
	}

	var writer *bufio.Writer
	if file, err := os.OpenFile(filepath, os.O_CREATE|os.O_RDWR, os.ModePerm); err != nil {
		result = false
		log.Fatal(err)
	} else {
		writer = bufio.NewWriter(file)

		defer file.Close()
	}

	reader := bufio.NewReader(stdout)
	for {
		bytes, _, err := reader.ReadLine()
		if err == io.EOF {
			cmd.Process.Kill()

			break
		} else if err != nil {
			result = false
			log.Fatal(err)
		} else {
			if _, err := writer.Write(bytes); err != nil {
				result = false
				log.Fatal(err)
			} else {
			}
		}
	}

	return result
}

func (this *Dumper) DeleteFile() bool {
	result := true

	err := os.Remove(this.filepath)

	if err != nil {
		result = false
		log.Fatal(err)
	} else {
	}

	return result
}
