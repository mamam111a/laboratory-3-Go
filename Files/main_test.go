package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)
func BenchmarkQPUSH(bench *testing.B) {
	queue := newQueue()

	for i := 0; i < 1000; i++ {
		queue.QPUSH("temp")
	}
}

func BenchmarkQPOP(bench *testing.B) {
	queue := newQueue()
	for i := 0; i < 1000; i++ {
		queue.QPUSH("temp")
	}
	bench.ResetTimer() 
	for i := 0; i < 1000; i++ {
		queue.QPOP()
	}
}
func TestQPUSH(test *testing.T) {
	queue := newQueue()
	queue.QREAD()
	queue.QPUSH("temp1")
	queue.QPUSH("temp2")
	queue.QPUSH("temp3")


	assert.Equal(test, "temp1", queue.head.cell)
	assert.Equal(test, "temp2", queue.head.next.cell)
	assert.Equal(test, "temp3", queue.tail.cell)
	
	queue.QREAD()
}

func TestQPOP(test *testing.T) {

	queue := newQueue()

	queue.QPUSH("temp1")
	queue.QPUSH("temp2")
	queue.QPUSH("temp3")

	assert.Equal(test, "temp1", queue.head.cell)

	queue.QPOP()

	assert.Equal(test, "temp2", queue.head.cell)
	assert.Equal(test, "temp3", queue.tail.cell)

	queue.QPOP()
	queue.QPOP()

	assert.Nil(test, queue.head)
	assert.Nil(test, queue.tail)
}

func TestQPOP_EMPTY(test *testing.T) {
	queue := newQueue()
	queue.QPOP()

	assert.Nil(test, queue.head)
	assert.Nil(test, queue.tail)
}

func TestBinarySerialization(test *testing.T) {
	queue := newQueue()
	queue.QPUSH("temp1")
	queue.QPUSH("temp2")
	queue.QPUSH("temp3")
	queue.BinarySerialization("/root/notExist.txt")
	err := queue.BinarySerialization("file.bin")
	assert.Nil(test, err)
	queue.BinaryDEserialization("notExist.bin")
	result, err := queue.BinaryDEserialization("file.bin")
	assert.Nil(test, err)
	assert.Equal(test, []string{"temp1", "temp2", "temp3"}, result)
}
func TestFileOperations(test *testing.T) {
	queue := newQueue()

	queue.QPUSH("temp1")
	queue.QPUSH("temp2")
	queue.WritingFromStructureToFile("/root/notExist.txt")
	queue.WritingFromStructureToFile("file.txt")

	newQueue := newQueue()
	newQueue.WritingFromFileToStructure("notExist.txt")
	newQueue.WritingFromFileToStructure("file.txt")
	
	assert.Equal(test, "temp1", newQueue.head.cell)
	assert.Equal(test, "temp2", newQueue.head.next.cell)
}







