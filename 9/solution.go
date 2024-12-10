package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func solve(line string) []int {

	id := 0
	disk := []int{}
	for i, char := range line {
		if i%2 == 0 {
			for j := 0; j < int(char-'0'); j++ {
				disk = append(disk, id)
			}
			id++
		} else {
			for j := 0; j < int(char-'0'); j++ {
				disk = append(disk, -1)
			}
		}
	}

	forward := 0
	backward := len(disk) - 1
	for forward < backward {
		if disk[forward] == -1 && disk[backward] != -1 {
			disk[forward] = disk[backward]
			disk[backward] = -1
			forward++
			backward--
		} else if disk[backward] == -1 {
			backward--
		} else {
			forward++
		}
	}

	return disk
}

type file struct {
	id, size, idx  int
}

type free struct {
	size, idx int
}

type disk struct {
	Free []free
	Files []file
}

func (d *disk) move(file *file) {
	for i, free := range d.Free {
		if free.size >= file.size && free.idx <= file.idx {
			file.idx = free.idx
			if free.size == file.size {
				d.Free = append(d.Free[:i], d.Free[i+1:]...)
			} else {
				d.Free[i].idx += file.size
				d.Free[i].size -= file.size
			}
			break
		}
	}
}

func (d *disk) checksum() int {
	result := 0
	for _, file := range d.Files {
		for i := 0; i < file.size; i++ {
			result += (file.idx + i) * file.id
		}
	}
	return result
}


func solve2(line string) disk {
	id := 0
	idx := 0
	disk := disk{}

	for i, char := range line {
		size := int(char-'0')
		if i%2 == 0 {
			disk.Files = append(disk.Files, file{id, size, idx})
			id++
		} else {
			disk.Free = append(disk.Free, free{size, idx})
		}
		idx += size
	}

	for i := len(disk.Files) - 1; i >= 0; i-- {
		disk.move(&disk.Files[i])
	}


	return disk
}

func checksum(compact []int) int {
	result := 0
	for i, id := range compact {
		if id != -1 {
			result += i * id
		}
	}
	return result
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := strings.TrimSpace(scanner.Text())

	// compact := solve(line)
	// checksum := checksum(compact)
	disk := solve2(line)
	checksum := disk.checksum()

	fmt.Println(checksum)

}
