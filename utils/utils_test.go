package utils

import (
	"fmt"
	"testing"
)

func TestUtils(t *testing.T) {
	datasetDir := "../usecases/dataset"
	//testDataNormalization(t, datasetDir)
	testReadFile(t, datasetDir)
}

func testDataNormalization(t *testing.T, dir string) {
	t.Run("Normalize Dataset", func(t *testing.T) {
		NormalizeDatasets(dir, 1)
	})
}

func testReadFile(t *testing.T, dir string) {
	fileName := "/b000101.txt"
	t.Run("Read File", func(t *testing.T) {
		data := ReadFile(dir + fileName)
		if data != nil {
			fmt.Println("=== Len:", len(data), ", Data: ", data)
		} else {
			panic("Something is wrong!")
		}
	})
}
