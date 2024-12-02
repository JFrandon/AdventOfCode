package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)


func getNextReport(reader *bufio.Reader) ([]int, error){
    line, err := reader.ReadString('\n')
    if err != nil {
        return nil, err
    }
    line2 := strings.Split(line[:len(line)-1], " ")    
    report := make([]int,0,1000)
    for _,s := range line2 {
        i, _ := strconv.Atoi(s)
        report = append(report, i)
    }
    return report, nil
}


func isSafe(report []int) bool{
    derivee := make([]int, len(report)-1)
    for i := 0 ; i < len(derivee) ; i++{
        derivee[i] = report[i+1] - report[i]
    } 
    fmt.Println(derivee)
    decreasing := derivee[0] < 0
    for _, dy := range derivee{
        if dy < 0 != decreasing {
            fmt.Printf("Non monotone %v: %d < 0\n", decreasing, dy)
            return false
        }
        if dy < 0 {
            dy = -dy
        }
        if dy < 1 || dy > 3 {
            fmt.Println("too steep")
            return false
        }
    }
    return true
}

func main(){
    reader := bufio.NewReader(os.Stdin)
    safeN := 0
    for report, err := getNextReport(reader); err == nil ; report, err = getNextReport(reader){
        safe := isSafe(report)
        for i := 0 ; i < len(report) && !safe; i++{
            fmt.Println(report)
            newr := make([]int, 0, len(report)-1)
            newr = append(newr,report[0:i]...)
            newr = append(newr,report[i+1:len(report)]...)
            fmt.Println(newr)
            safe = isSafe(newr)
        } 
        if safe{
            fmt.Println("Safe")
            safeN++
        }
    }
    fmt.Println(safeN)
}
