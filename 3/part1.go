package main

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main(){
    exp := "mul\\((\\d+),(\\d+)\\)"
    re := regexp.MustCompile(exp)
    input, _ := io.ReadAll(os.Stdin)
    strs := re.FindAllString(string(input),-1)
    sum := 0
    for _, s := range strs {
        strim := s[4:len(s)-1]
        ssplit := strings.Split(strim,",")
        l, _ := strconv.Atoi(ssplit[0])
        r, _ := strconv.Atoi(ssplit[1])
        sum += l*r
    }
    fmt.Println(sum)
}
