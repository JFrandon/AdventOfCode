package main

import (
	"fmt"
)


func getLists() ([]int, []int){
    gauche := make([]int,0,1000)
    droite := make([]int,0,1000)
    var a,b,n int
    var err error
    n, err = fmt.Scanf("%d   %d",&a, &b) 
    for i := 0 ; n == 2 && err == nil && i < 1000; i++ {
        gauche = append(gauche,a)
        droite = append(droite,b)
        n, err = fmt.Scanf("%d   %d",&a, &b) 
    }
    return gauche, droite
}

func countUnique(input []int) map[int]int{
    result := make(map[int]int)
    for _,i := range input{
        result[i] += 1
    }
    return result
}

func main(){
    gauche, droite := getLists()
    droiteMap := countUnique(droite)
    sumOfDiffs := 0
    for _,i := range gauche{
        sumOfDiffs += droiteMap[i] * i
    }  
    fmt.Printf("%d\n", sumOfDiffs)
}
