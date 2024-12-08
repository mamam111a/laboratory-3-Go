package main //принадлежит пакету мэйн

import (
	"bufio"  // для ввод/вывод
	"fmt"    // для форматированного вывода
	"os"     // для операцимй системы, такими как открытие файлов
	"strings"
	"encoding/binary"
)

type SingleNode struct {
	cell string
	next *SingleNode
}

type Queue struct {
	head *SingleNode
	tail *SingleNode
}

func newQueue() *Queue {
	return &Queue{}
}

func (queue *Queue) QPUSH(cell string) {
	node := &SingleNode{cell: cell, next: nil}

	if queue.head == nil {
		queue.head = node
		queue.tail = node
	} else {
		queue.tail.next = node
		queue.tail = node
	}
}

func (queue *Queue) QPOP() {
    if queue.head == nil {
        return
    }
    queue.head = queue.head.next
    if queue.head == nil {
        queue.tail = nil
    }
}

func (queue *Queue) QREAD() {
	current := queue.head
	if current == nil {
		fmt.Println("QUEUE IS EMPTY!!!")
	} else {
		for current != nil {
			fmt.Print(current.cell, " ")
			current = current.next
		}
		fmt.Println()
	}
}

func (queue *Queue) WritingFromFileToStructure(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("ERROR OPENING FILE:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		elements := strings.Fields(line)
		for _, element := range elements {
			queue.QPUSH(element)
		}
	}
}

func (queue *Queue) WritingFromStructureToFile(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("ERROR CREATING FILE", err)
		return
	}
	defer file.Close()

	current := queue.head
	for current != nil {
		_, err := fmt.Fprint(file, current.cell+" ")
		if err != nil {
			fmt.Println("ERROR WRITING TO FILE", err)
			return
		}
		current = current.next
	}
}

func (queue *Queue)BinarySerialization(filename string) error {

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	current := queue.head
	for current != nil {//запись в файл. littleEndian записывает младший байт раньше
		len := int32(len(current.cell)) //находим длину и записываем в 32бит = 4байта
		if err := binary.Write(file, binary.LittleEndian, len); err != nil { 
			return err
		}


		_, err := file.Write([]byte(current.cell)) //записывает байтовый срез
		if err != nil {
			return err
		}


		current = current.next
	}

	return nil
}


func (queue *Queue)BinaryDEserialization(filename string) ([]string, error) { //аргументы и возвращаемый результат

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var result []string

	for {
		var len int32
		err := binary.Read(file, binary.LittleEndian, &len)
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			return nil, err
		}

		buffer := make([]byte, len) //мэйк это выделение памяти
		_, err = file.Read(buffer)
		if err != nil {
			return nil, err
		}

		text := string(buffer)
		result = append(result, text)
	}

	return result, nil
}
/*
func main() {

	scanner := bufio.NewScanner(os.Stdin)
	queue := newQueue()
	var filename string
	for {
		
		fmt.Println()
		fmt.Print("Enter command ==>> ")
		
		scanner.Scan() 
		command := scanner.Text()
		parts := strings.Fields(command)
		
		if parts[0] == "EXIT" {
			return
		}
		
		if len(parts) > 1 {
			filename = parts[1] + ".txt"
		}

		if _, err := os.Stat(filename); os.IsNotExist(err) {
			fmt.Println("FILE DOESNT EXIST!!!")
			return
		}


		
		fmt.Print("Choose serialization:")
		fmt.Println()
		fmt.Print("1 - binary")
		fmt.Println()
		fmt.Print("2 - text")
		fmt.Println()
		fmt.Print("==>> ")
		scanner.Scan() 
		serialization := scanner.Text()
		
		if serialization == "2" {
		queue.WritingFromFileToStructure(filename)
		} else {
			queue.BinaryDEserialization(filename)
		}
		
		switch parts[0] {

		case "QPUSH":
			if len(parts) == 3 {
				queue.QPUSH(parts[2])
				if serialization == "2" {
					queue.WritingFromStructureToFile(filename)
				} else {
					queue.BinarySerialization(filename)
				}
				
			} else {
				fmt.Println("ERROR INPUT!!")
			}

		case "QREAD":
			if len(parts) == 2 {
				queue.QREAD()
			} else {
				fmt.Println("ERROR INPUT!!")
			}

		case "QPOP":
			if len(parts) == 2 {
				queue.QPOP()
			} else {
				fmt.Println("ERROR INPUT!!")
			}
			
			if serialization == "2" {
					queue.WritingFromStructureToFile(filename)
				} else {
					queue.BinarySerialization(filename)
				}
			

		default:
			fmt.Println("Unknown command:")
		}
	}
	
}
*/
