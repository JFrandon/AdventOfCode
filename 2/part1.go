package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
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
    fmt.Println(report)
    derivee := make([]int, len(report)-1)
    for i := 0 ; i < len(derivee) ; i++{
        derivee[i] = report[i+1] - report[i]
    } 
    fmt.Println(derivee)
    decreasing := derivee[0] < 0
    for dy := range derivee{
        if dy < 0 != decreasing {
            fmt.Println("Non monotone")
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
    safe := 0
    for report, err := getNextReport(reader); err == nil ; report, err = getNextReport(reader){
        safe := isSafe(report)
        for i := 0 ; i < len(report) || !safe; i++{
            slices.Delete(report, i, i+1)
            safe = isSafe(report)
        } 
        if safe{
            fmt.Println("Safe")
            safe++
        }
    }
    fmt.Println(safe)
}
