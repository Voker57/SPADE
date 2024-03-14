package utils

import (
	"bufio"
	"fmt"
	"math/big"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
)

// HandleError check the error if err then panic
func HandleError(e error) {
	if e != nil {
		panic(e)
	}
}

// NormalizeDatasets this function removes zeros from hypnogram values
// by adding 1 to all elements
// Do not run this twice :)
func NormalizeDatasets() {
	dir := "../usecases/dataset"

	// Read all files in the directory
	files, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}
	print("number of files: ", len(files))

	// Loop through each file in the directory
	for _, fileInfo := range files {
		// Check if the file is a .txt file
		if filepath.Ext(fileInfo.Name()) == ".txt" {
			// Open the file
			filePath := filepath.Join(dir, fileInfo.Name())
			file, err := os.Open(filePath)
			if err != nil {
				fmt.Printf("Error opening file %s: %v\n", filePath, err)
				continue
			}
			defer file.Close()

			scanner := bufio.NewScanner(file)
			modifiedIntegers := make([]int, 0)

			for scanner.Scan() {
				number, err := strconv.Atoi(scanner.Text())
				if err != nil {
					fmt.Printf("Error parsing integer from file %s: %v\n", filePath, err)
					continue
				}
				modifiedIntegers = append(modifiedIntegers, number-1)
			}

			if err := scanner.Err(); err != nil {
				fmt.Printf("Error scanning file %s: %v\n", filePath, err)
				continue
			}

			outputFile, err := os.Create(filePath)
			if err != nil {
				fmt.Printf("Error creating file %s: %v\n", filePath, err)
				continue
			}
			defer outputFile.Close()

			writer := bufio.NewWriter(outputFile)
			for _, number := range modifiedIntegers {
				_, err := fmt.Fprintln(writer, number)
				if err != nil {
					fmt.Printf("Error writing to file %s: %v\n", filePath, err)
					continue
				}
			}

			err = writer.Flush()
			if err != nil {
				fmt.Printf("Error flushing buffer for file %s: %v\n", filePath, err)
				continue
			}

			fmt.Println("Successfully modified and saved file:", filePath)
		}
	}

}

func CountValues(data []*big.Int, v int) {

}

// GenDummyData generate dummy a vector of integers with "ptVecSize" elements,
// for each user, the values are going to be between 1 and "maxValue"
func GenDummyData(numUsers, ptVecSize int, maxValue int64) [][]int {
	// generate random data for test
	dummyData := make([][]int, numUsers)
	for j := 0; j < numUsers; j++ {
		dummyData[j] = make([]int, ptVecSize)
		for i := 0; i < ptVecSize; i++ {
			dummyData[j][i] = int(rand.Int63n(maxValue) + 1)
		}
	}

	return dummyData
}
