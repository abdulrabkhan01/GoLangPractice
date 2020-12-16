package main

import "fmt"

func main() {  
   var n1 int = 783
   var n2 int = 2
   var n3 int = 1

   var result int
   result = sum(n1, n2, n3);

   fmt.Printf( "Sum  : %d\n", result )
   x, y := multireturn()
   fmt.Println(x,y)
}


func sum(n1, n2, n3 int) int {
   var sum int

   sum = n1 + n2 + n3
   return sum 
}

func multireturn() (int, int) {
	return 1, 2
}