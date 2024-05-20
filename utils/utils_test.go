package utils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"testing"
)

func TestUtils(t *testing.T) {
	hypnogramDatasetDir := "../usecases/hypnogram/dataset/"
	dnaDatasetDir := "../usecases/dna/dataset/"

	//testDataNormalization(t, datasetDir)
	testReadHypnogramFile(t, hypnogramDatasetDir)
	testConvertDNASeq2Dinucleotide(t, dnaDatasetDir)
}

func testDataNormalization(t *testing.T, dir string) {
	t.Run("Normalize Dataset", func(t *testing.T) {
		NormalizeHypnogramDatasets(dir, 1)
	})
}

func testReadHypnogramFile(t *testing.T, dir string) {
	fileName := "/b000101.txt"
	t.Run("Read File", func(t *testing.T) {
		data := ReadHypnogramFile(dir + fileName)
		if data != nil {
			fmt.Println("=== Len:", len(data), ", Data: ", data)
		} else {
			panic("Something is wrong!")
		}
	})
}

func testConvertDNASeq2Dinucleotide(t *testing.T, dir string) {
	t.Run("Convert DNA Sequences to Dinucleotide", func(t *testing.T) {
		// read every dna file
		files, err := os.ReadDir(dir)
		if err != nil {
			fmt.Println("Error reading directory:", err)
			return
		}

		for _, fileInfo := range files {
			if filepath.Ext(fileInfo.Name()) == ".txt" {
				filePath := filepath.Join(dir, fileInfo.Name())
				data := ReadDNASeqFile(filePath)
				convData := ConvertDNASeq2Dinucleotide(data)
				intData := MapDinucleotideToInt(convData)
				log.Println("=== Len: ", len(intData), ", Data: ", intData)
			}
			break
		}
	})
}
