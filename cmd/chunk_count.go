package cmd

import (
	"bufio"
	"log"
	"os"

	"split_cmd/utils"
)

func ExecuteByChunk(filename string, chunkCount int) {
	readFile, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer readFile.Close()

	fileSize, err := utils.GetFileSize(readFile)
	if err != nil {
		log.Fatal(err)
	}
	if fileSize < chunkCount {
		log.Fatalf("can't split into more than %d files", fileSize)
	}
	byteCount, rest := fileSize/chunkCount, fileSize%chunkCount

	reader := bufio.NewReader(readFile)
	filenameGenerator := utils.NewFilenameGenerator()

	for i := 1; i < chunkCount; i++ {
		chunks, _, err := utils.ReadChunksByByteCount(reader, byteCount)
		if err != nil {
			log.Fatal(err)
		}

		err = utils.CreateFileAndWrite("./tmp_dir/"+filenameGenerator.CurrentName, chunks)
		if err != nil {
			log.Fatal(err)
		}

		filenameGenerator.Increment()
	}

	chunks, _, err := utils.ReadChunksByByteCount(reader, byteCount+rest)
	if err != nil {
		log.Fatal(err)
	}

	err = utils.CreateFileAndWrite("./tmp_dir/"+filenameGenerator.CurrentName, chunks)
	if err != nil {
		log.Fatal(err)
	}
}
