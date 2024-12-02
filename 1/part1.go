package main

import (
	"fmt"
	"slices"
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


func main(){
    gauche, droite := getLists()
    slices.Sort(gauche)
    slices.Sort(droite)
    sumOfDiffs := 0
    for i := 0; i < len(gauche) && i < 1000; i++{
        a := gauche[i]
        b := droite[i]
        if a > b {
            sumOfDiffs += a-b
        } else {
            sumOfDiffs += b-a
        }
    }  
    fmt.Printf("%d\n", sumOfDiffs)
}
