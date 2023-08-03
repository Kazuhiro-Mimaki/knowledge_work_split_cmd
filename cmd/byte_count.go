package cmd

import (
	"bufio"
	"log"
	"os"

	"split_cmd/utils"
)

func ExecuteByteCount(filename string, byteCount int) {
	readFile, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer readFile.Close()

	reader := bufio.NewReader(readFile)

	filenameGenerator := utils.NewFilenameGenerator()

	for {
		chunks, cursor, err := utils.ReadChunksByByteCount(reader, byteCount)
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			log.Fatal(err)
		}

		err = utils.CreateFileAndWrite("./tmp_dir/"+filenameGenerator.CurrentName, chunks)

		writeFile, err := os.Create("./tmp_dir/" + filenameGenerator.CurrentName)
		if err != nil {
			log.Fatal(err)
		}

		filenameGenerator.Increment()

		writeFile.Close()

		// Check for EOF
		if err == nil && cursor < byteCount {
			break
		}
	}
}
