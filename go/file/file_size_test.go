package file

import (
	"log"
	"os"
	"testing"
)

func TestFile(t *testing.T) {
	fileName := "/tmp/data.json"
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Error to read [file=%v]: %v", fileName, err.Error())
	}

	fi, err := f.Stat()
	size := fileSize(fi.Size())
	t.Logf("The [%s] is %d %s long\n", fileName, fi.Size(), size)
}
