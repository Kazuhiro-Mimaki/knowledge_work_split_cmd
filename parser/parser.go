package parser

import (
	"errors"
	"flag"
)

type SplitType int

const (
	Line SplitType = iota + 1
	Chunk
	Byte
)

type CmdArgs struct {
	SplitType    SplitType
	LineCount    int
	ChunkCount   int
	ByteCount    int
	SuffixLength int
	IsNumeric    bool
	Filepath     string
	Prefix       string
}

func ParseCmdArgs() (*CmdArgs, error) {
	lineCount := flag.Int("l", 0, "line_count")
	chunkCount := flag.Int("n", 0, "chunk_count")
	byteCount := flag.Int("b", 0, "byte_count")

	// suffixLengthはデフォルトを2とする
	suffixLength := flag.Int("a", 2, "suffix_length")
	isNumeric := flag.Bool("d", false, "is_numeric")

	flag.Parse()

	// lineCount, chunkCount, byteCount のいずれか1つのみ指定できる
	splitType, err := getSplitType(*lineCount, *chunkCount, *byteCount)
	if err != nil {
		return nil, err
	}

	filepath, prefix, err := parseFlagArgs(flag.Args())
	if err != nil {
		return nil, err
	}

	return &CmdArgs{
		SplitType:    splitType,
		LineCount:    *lineCount,
		ChunkCount:   *chunkCount,
		ByteCount:    *byteCount,
		SuffixLength: *suffixLength,
		IsNumeric:    *isNumeric,
		Filepath:     filepath,
		Prefix:       prefix,
	}, nil
}

func getSplitType(lineCount, chunkCount, byteCount int) (SplitType, error) {
	switch true {
	case lineCount > 0 && chunkCount == 0 && byteCount == 0:
		return Line, nil
	case lineCount == 0 && chunkCount > 0 && byteCount == 0:
		return Chunk, nil
	case lineCount == 0 && chunkCount == 0 && byteCount > 0:
		return Byte, nil
	case lineCount == 0 && chunkCount == 0 && byteCount == 0:
		return 0, errors.New("split type must be specified")
	default:
		return 0, errors.New("split type must be one")
	}
}

func parseFlagArgs(flagArgs []string) (string, string, error) {
	var (
		filepath string
		prefix   string
	)

	if len(flagArgs) == 1 {
		filepath, prefix = flagArgs[0], ""
	} else if len(flagArgs) == 2 {
		filepath, prefix = flagArgs[0], flagArgs[1]
	} else {
		return "", "", errors.New("invalid args")
	}

	return filepath, prefix, nil
}
