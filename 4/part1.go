package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)



func findXMASh(y string, xi int) bool{
    if len(y) < xi + 4{
        return false
    }
    return y[xi: xi+4] == "XMAS"
}

func findXMASv(lines []string, xi, yi int) bool{
    if len(lines) < yi + 5{
        return false
    }
    return lines[yi][xi] == 'X' && lines[yi+1][xi] == 'M' && lines[yi+2][xi] == 'A' && lines[yi+3][xi] == 'S' 
}

func findXMASd(lines []string, xi, yi int) bool{
    if len(lines) < yi + 5 || len(lines[yi]) < xi + 4{
        return false
    }
    return lines[yi][xi] == 'X' && lines[yi+1][xi+1] == 'M' && lines[yi+2][xi+2] == 'A' && lines[yi+3][xi+3] == 'S' 
}

func findSAMXh(y string, xi int) bool{
    if len(y) < xi + 4{
        return false
    }
    return y[xi: xi+4] == "SAMX"
}

func findSAMXv(lines []string, xi, yi int) bool{
    if len(lines) < yi + 5{
        return false
    }
    return lines[yi][xi] == 'S' && lines[yi+1][xi] == 'A' && lines[yi+2][xi] == 'M' && lines[yi+3][xi] == 'X' 
}

func findSAMXd(lines []string, xi, yi int) bool{
    if len(lines) < yi + 5 || len(lines[yi]) < xi + 4{
        return false
    }
    return lines[yi][xi] == 'S' && lines[yi+1][xi+1] == 'A' && lines[yi+2][xi+2] == 'M' && lines[yi+3][xi+3] == 'X' 
}

func findXMASdu(lines []string, xi, yi int) bool{
    if yi-3 < 0 || len(lines[yi]) < xi + 4{
        return false
    }
    return lines[yi][xi] == 'X' && lines[yi-1][xi+1] == 'M' && lines[yi-2][xi+2] == 'A' && lines[yi-3][xi+3] == 'S' 
}

func findSAMXdu(lines []string, xi, yi int) bool{
    if yi-3 < 0 || len(lines[yi]) < xi + 4{
        return false
    }
    return lines[yi][xi] == 'S' && lines[yi-1][xi+1] == 'A' && lines[yi-2][xi+2] == 'M' && lines[yi-3][xi+3] == 'X' 
}

func main(){
    input, _ := io.ReadAll(os.Stdin)
    lines := strings.Split(string(input), "\n")
    count := 0
    for yi, y := range lines[:len(lines)-1] {
        for xi, x := range y {
            if x == 'X' {
                if findXMASh(y,xi) {
                    count ++
                }
                if findXMASv(lines, xi, yi){
                    count ++
                }
                if findXMASd(lines, xi, yi){
                    count ++
                }
                if findXMASdu(lines, xi, yi){
                    count ++
                }
            }
            if x == 'S' {
                if findSAMXh(y,xi){
                    count++
                }
                if findSAMXv(lines, xi, yi){
                    count++
                }
                if findSAMXd(lines, xi, yi){
                    count++
                }
                if findSAMXdu(lines, xi, yi){
                    count++
                } 
            }
        }
    }
    fmt.Println(count)
}
