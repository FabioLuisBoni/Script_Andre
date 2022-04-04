package filestreatment

import (
	"bufio"
	"io"
	"os"
	"strings"
)

func ExtractFileNames(dataPath string, index string) []string {
	dataFile, reader := OpenFile(dataPath, index)

	var files []string

	for i := 0; true; i++ {

		file, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}

			dataFile.Close()
			panic("Houve algum problema durante a leitura do arquivo em ExtractFileNames()")
		}

		file = strings.Trim(file, "\n")

		files = append(files, file)
	}

	return files
}

func OpenFile(dataPath string, fileName string) (*os.File, *bufio.Reader) {
	file, err := os.OpenFile(dataPath+fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic("Error opening file: " + err.Error())
	}

	reader := bufio.NewReader(file)

	return file, reader
}
