package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func readNumberLists(filename string) ([]uint32, []uint32, error) {
	var (
		list1 = make([]uint32, 0, 1000)
		list2 = make([]uint32, 0, 1000)
	)

	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Error closing file: %v\n", err)
		}
	}(file)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		num1, num2, _ := getIntegers(line)

		list1 = append(list1, num1)
		list2 = append(list2, num2)
	}

	return list1, list2, nil
}

func getIntegers(line string) (uint32, uint32, error) {
	fields := strings.Fields(line)
	if len(fields) != 2 {
		return 0, 0, errors.New("incorrect line format")
	}
	num1, _ := strconv.ParseInt(fields[0], 10, 32)
	num2, _ := strconv.ParseInt(fields[1], 10, 32)
	return uint32(num1), uint32(num2), nil
}

func day01() {
	var (
		list1 []uint32
		list2 []uint32
	)
	list1, list2, _ = readNumberLists("inputs/day01/01.txt")
	sort.Sort(Uint32Slice(list1))
	sort.Sort(Uint32Slice(list2))

	var sum uint32
	for i := 0; i < len(list1); i++ {
		sum += abs(list1[i], list2[i])
	}
	println(sum)

	list2Cntr := make(map[uint32]uint32)
	for _, v := range list2 {
		list2Cntr[v]++
	}

	var sum2 uint32
	for _, v := range list1 {
		sum2 += v * list2Cntr[v]
	}

	println(sum2)
}

func day02() {
	
}

func main() {
	// day01()
	day02()
}

func abs(a, b uint32) uint32 {
	if a > b {
		return a - b
	}
	return b - a
}

type Uint32Slice []uint32

func (x Uint32Slice) Len() int           { return len(x) }
func (x Uint32Slice) Less(i, j int) bool { return x[i] < x[j] }
func (x Uint32Slice) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }
