package log

import (
	"io/ioutil"
	"log"
	"os"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

//TestInfoLog - testing info log
func TestInfoLog(t *testing.T) {
	Info("info log message")

	file, err := os.Open("/tmp/info.log")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	b, _ := ioutil.ReadAll(file)
	assert.Contains(t, bytesToStr(b), "info log message")
}

//TestErrorLog - testing error log
func TestErrorLog(t *testing.T) {
	Error("Error log message Sample")

	file, err := os.Open("/tmp/info.log")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	b, _ := ioutil.ReadAll(file)
	assert.Contains(t, bytesToStr(b), "Error log message Sample")

}

func bytesToStr(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
