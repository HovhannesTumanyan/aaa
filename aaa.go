package main

import (
	"fmt"
	"math"
)

func main() {
	var num int
	fmt.Println("ծրագիրը ստուգում է մուտքագրված թվի պարզ լինելը")
	fmt.Println(" մուտքագրե՛ք թիվը")
	n, err := fmt.Scan(&num)
	if n < 1 {
		fmt.Println("մուտքային տվյալների սխալ", err)
		return
	}
	if !(num > 0) {
		fmt.Println("մուտքային տվյալների սխալ,պահանջվում է բնական թիվ")
		return
	}
    int b
	 b=num
	fmt.Println(num)

	a := math.Sqrt(b)

	k := 1
    l:= 2
	for ; k < a; k++ {
		
if a%l==0{
	fmt.Println("թիվը բաղադրյալ է")
	return
}

l++
	}

	return
}
