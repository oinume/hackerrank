// https://www.hackerrank.com/challenges/2d-array/problem
package datastructures

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func FindHourglass(r io.Reader) int {
	var values [][]int
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		var cols []int
		//fmt.Println(scanner.Text()) // Println will add back the final '\n'
		for _, s := range strings.Split(scanner.Text(), " ") {
			if v, err := strconv.ParseInt(s, 10, 32); err == nil {
				cols = append(cols, int(v))
			}
		}
		//fmt.Printf("cols = %+v\n", cols)
		values = append(values, cols)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	//fmt.Printf("values = %+v\n", values)

	largest := -1000000000
	for i, cols := range values {
		for j, col := range cols {
			// is center?
			if !(i >= 1 && i < len(values) -1 && j >= 1 && j < len(cols) - 1) {
				continue
			}
			s := sum(
				values[i-1][j-1], values[i-1][j], values[i-1][j+1],
				col,
				values[i+1][j-1], values[i+1][j], values[i+1][j+1],
			)
			if s > largest {
				largest = s
			}
			//fmt.Printf("%v ", col)
			//fmt.Printf("sum = %v\n", s)
		}
		//fmt.Println()
	}

	return largest
}

func sum(args ...int) int {
	s := 0
	for _, a := range args {
		s += a
	}
	return s
}