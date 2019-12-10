package main

import (
	"bufio"
	"fmt"
	"os"
)

type sif [][]byte

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	width := 25
	height := 6
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

	fmt.Printf(" what is the number of 1 digits multiplied by the number of 2 digits? %d\n", sif.Checksum())
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
