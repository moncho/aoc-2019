package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	width  = 25
	height = 6
)

type sif [][]byte

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var sif sif
	i := 0
	j := 0
	sif = append(sif, make([]byte, width*height))
	layer := sif[0]
	for scanner.Scan() {
		bb := scanner.Bytes()
		for _, b := range bb {
			if j == width*height {
				j = 0
				sif = append(sif, make([]byte, width*height))
				i++
				layer = sif[i]
			}
			layer[j] = b
			j++
		}
	}

	fmt.Printf("What is the number of 1 digits multiplied by the number of 2 digits? %d\n", sif.Checksum())

	fmt.Printf("What message is produced after decoding your image?\n\n%s\n", sif.Decode())
}

func (s sif) Checksum() int {
	minZeroCount, minLayer := 1<<32-1, 0
	var countPerLayer []map[byte]int
	for i, layer := range s {
		count := make(map[byte]int)
		countPerLayer = append(countPerLayer, count)
		for _, b := range layer {
			count[b]++
		}
		if count['0'] < minZeroCount {
			minZeroCount = count['0']
			minLayer = i
		}
	}
	return countPerLayer[minLayer]['1'] * countPerLayer[minLayer]['2']
}

func (s sif) Decode() string {
	decode := make([]byte, len(s[0]))
	for _, layer := range s {
		for i, pixel := range layer {
			if pixel == '2' || decode[i] != 0 {
				continue
			}
			switch pixel {
			case '0':
				decode[i] = ' '
			case '1':
				decode[i] = '*'
			}
		}
	}
	var result string

	for i := 0; i < width*height; i += width {
		result += string(decode[i : i+width])
		result += "\n"
	}

	return result
}
