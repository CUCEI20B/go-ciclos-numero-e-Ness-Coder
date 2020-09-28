package main

import "fmt"

func main()  {
	var n float64
	var j float64
	var i float64
	var e float64
	var temp float64
	

	e = 2.71828
	fmt.Scan(&n)

	for j = 1; j <= n; j++ {
		temp = 1 // temp se usara como el factorial
		for i = 1; i <= j; i++ {
			temp = temp * j
		} 
		e = e + 1.0 /temp
	}
	fmt.Println(e)
}