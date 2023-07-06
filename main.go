package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"flag"
	"fmt"
	"hash"
	"hash/crc32"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const (
	SHA1          = "sha1"
	SHA256        = "sha256"
	SHA512        = "sha512"
	CRC32         = "crc32"
	MD5           = "md5"
	ErrorExitCode = 1
	MinArgsCount  = 3
	BufferSize    = 64 * 1024 * 1024 // 64 MB
)

var ModeHashConstructorMap = map[string]func() hash.Hash{
	SHA1:   sha1.New,
	SHA256: sha256.New,
	SHA512: sha512.New,
	MD5:    md5.New,
}

func getHashInstance(mode string) hash.Hash {
	if mode == CRC32 {
		return crc32.NewIEEE()
	}
	return ModeHashConstructorMap[mode]()
}

func calcHashSum(fileObj io.Reader, mode string) []byte {
	sumInstance := getHashInstance(mode)
	buf := make([]byte, BufferSize)
	for {
		bytesRead, err := fileObj.Read(buf)
		if err != nil {
			if err != io.EOF {
				panic(err)
			}
			break
		}
		sumInstance.Write(buf[:bytesRead])
	}
	return sumInstance.Sum(nil)
}

func timeIt(calcFunc func(fileObj io.Reader, mode string) []byte, fileObj io.Reader, mode string) []byte {
	start := time.Now()
	result := calcFunc(fileObj, mode)
	fmt.Printf("Calculated in %s.\n", time.Since(start))
	return result
}

func calculate(path string, mode string, verbose bool) {
	fileObj, err := os.Open(path)
	defer fileObj.Close()

	if err != nil {
		fmt.Printf("Error occurred when opening '%s': %s\n", path, err)
		return
	}

	fileInfo, err := fileObj.Stat()
	if err != nil {
		fmt.Printf("Error occurred when getting stats for file '%s': %s\n", path, err)
		return
	}
	if fileInfo.IsDir() {
		fmt.Println("Directory mode is not supported at the moment")
	} else {
		var result []byte
		if verbose {
			fmt.Printf("Calculating %s sum for file '%s'", mode, path)
			result = timeIt(calcHashSum, fileObj, mode)
		} else {
			result = calcHashSum(fileObj, mode)
		}
		fmt.Printf(hex.EncodeToString(result))
	}
}

func main() {
	supportedModes := []string{SHA1, SHA256, SHA512, CRC32, MD5}
	argsFull := os.Args
	if len(argsFull) < MinArgsCount {
		executableName := filepath.Base(argsFull[0])
		fmt.Println("Usage:", executableName, "MODE PATH [-v]")
		fmt.Println("Calculates hash sum for passed PATH with MODE.")
		fmt.Println("Available modes:", supportedModes)
		fmt.Println("Specifying -v option makes output verbose (statistic about calculation time, etc.)")
		os.Exit(ErrorExitCode)
	}
	mode := argsFull[1]
	path := argsFull[2]
	options := argsFull[3:]
	verbose := flag.Bool("v", false, "Verbose mode")
	err := flag.CommandLine.Parse(options)
	if err != nil {
		fmt.Println("Error occurred:", err)
		os.Exit(ErrorExitCode)
	}

	lowerMode := strings.ToLower(mode)
	_, found := ModeHashConstructorMap[lowerMode]
	if !found && lowerMode != CRC32 {
		fmt.Println(mode, "is not supported. Specify MODE one of:", supportedModes)
		os.Exit(ErrorExitCode)
	}
	calculate(path, lowerMode, *verbose)
}
