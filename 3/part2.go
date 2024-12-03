package main

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)


func parseMul(s string) int {
    strim := s[4:len(s)-1]
    ssplit := strings.Split(strim,",")
    l, _ := strconv.Atoi(ssplit[0])
    r, _ := strconv.Atoi(ssplit[1])
    return l*r
}

func main(){
    exp := "(mul\\((\\d+),(\\d+)\\))|(do(n't)?\\(\\))"
    re := regexp.MustCompile(exp)
    input, _ := io.ReadAll(os.Stdin)
    strs := re.FindAllString(string(input),-1)
    fmt.Println(strs)
    sum := 0
    enabled := true
    for _, s := range strs {
        if s[:2] == "do" {
            enabled = s[2] != 'n'
            continue
        }
        if s[:3] == "mul" && enabled {
            sum += parseMul(s)
            continue
        }
    }
    fmt.Println(sum)
}
