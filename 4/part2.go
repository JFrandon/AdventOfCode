package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)



func findX_MAS(lines []string, xi,yi int) bool{
    if yi < 1 || yi >= len(lines)-1 || xi < 1 || xi + 1 >= len(lines[yi+1]) {
        return false
    }
    fmt.Println(lines[yi-1][xi-1],
        lines[yi+1][xi+1],
        lines[yi+1][xi-1],
        lines[yi-1][xi+1])
    if ! (lines[yi-1][xi-1] == 'M' && lines[yi+1][xi+1] == 'S' ) && !(lines[yi-1][xi-1] == 'S' && lines[yi+1][xi+1] == 'M' ){
        return false
    }
    if ! (lines[yi+1][xi-1] == 'M' && lines[yi-1][xi+1] == 'S' ) && !(lines[yi+1][xi-1] == 'S' && lines[yi-1][xi+1] == 'M' ){
        return false
    }
    return true
}

func main(){
    input, _ := io.ReadAll(os.Stdin)
    lines := strings.Split(string(input), "\n")
    count := 0
    fmt.Println(len(lines), len(lines[0]))
    for yi, y := range lines[:len(lines)-1] {
        for xi, x := range y {
            if x != 'A'{
                continue
            }
            if findX_MAS(lines, xi,yi) {
                count++
            }
        }
    }
    fmt.Println(count)
}
