package main

import (
	filestreatment "Script_Andre/FilesTreatment"
	"io"
	"strings"
)

const dataPath = "/home/fabio/Trinovati/go/src/Script_Andre/Data/"
const index = "index.txt"
const resultFileName = "d0_269_330"

func main() {
	resultFile, _ := filestreatment.OpenFile(dataPath, resultFileName)
	files := filestreatment.ExtractFileNames(dataPath, index)

	for fileIndex := range files {
		nameTrimed := strings.Trim(files[fileIndex], ".csv\n")
		nameSliced := strings.Split(nameTrimed, "_")

		file, reader := filestreatment.OpenFile(dataPath+"TableData/", files[fileIndex])
		for {
			var data []string

			dataLineString, err := reader.ReadString('\n')
			if err != nil {
				if err == io.EOF {
					break
				}

				file.Close()
				panic("Houve algum problema durante a leitura do arquivo em ExtractFileNames()")
			}

			dataLineString = strings.Trim(dataLineString, "\n")
			dataLine := strings.Split(dataLineString, ",")

			data = append(data, nameSliced[0], nameSliced[1])
			data = append(data, dataLine[3], dataLine[4], dataLine[5], dataLine[8], dataLine[11])

			for i := 0; i < 7; i++ {
				resultFile.WriteString(data[i])
				if i != 6 {
					resultFile.WriteString(";")

				} else {
					resultFile.WriteString("\n")
				}
			}
		}
	}
}
