package main

import (
   "bufio"
   "fmt"
   "io"
   "os"
   "unicode"
)

func main() {
    file, err := os.Open("input.txt")
    if err != nil {
        return
    }
    defer file.Close()

    var count int = 0
    var num int = 50
    var clicks int = 0

    reader := bufio.NewReader(file)
    for {
    	line, err := reader.ReadString('\n')
	if err == io.EOF {
	   break
	}
	if err != nil {
	   return
	}
	var direction int = 0
	var cur int = 0
	for _, ch := range line {
	    if ch == 'L' {
	       direction = -1
	    } else if ch == 'R' {
	      direction = 1
	    } else if unicode.IsDigit(ch) {
	      if direction == 0 {
	      	 fmt.Printf("WE GOT A BIG PROBLEM\n")
		 return
	      }
	      cur = cur * 10
	      cur = cur + (int(ch) - '0')
	    } else {
	      if cur == 0 {
	      	 break
	      }
	      cur = cur * direction

	      var prevcnt = mod(num, 100)
	      var dist int = 0
	      fmt.Printf("The dial is at %d and we are turning it ", prevcnt)
	      if (direction < 0) {
	      	 fmt.Printf("Left")
		 dist = prevcnt
	      } else {
	      	   fmt.Printf("Right")
		   dist = 100 - prevcnt
	      }

	      if (dist == 0) {
	      	 dist = 100
	      }
	      fmt.Printf(" by %d\n", cur)

	      fmt.Printf("The minimum distance to a click is %d\n", dist*direction);

	      if cur*direction >= dist {
	      	 fmt.Printf("Clicking\n")
		 var metacnt int = 0
		 metacnt = cur - dist*direction
		 fmt.Printf("  Leftover is %d\n", metacnt)
		 fmt.Printf("  that means that %d clicks remains\n", (int(float32(metacnt*direction)/100.0)))
		 clicks = clicks + 1 + int(float32(metacnt*direction)/100.0)
		 fmt.Printf("  the number of clicks is %d\n", clicks)
	      }

	      num = num + cur
	      fmt.Printf("The internal count (num) is %d\n", num)
	      var wheelloc int = mod(num, 100)
	      fmt.Printf("wheel position left on %d\n", wheelloc)
	      fmt.Printf("dist is %v\n", dist)
	      fmt.Printf("prevcnt is %d\n", prevcnt)

	      if (num % 100) == 0 {
	      	 count += 1
	      }
	      fmt.Printf("The count is %d\n", count)
	      fmt.Printf("--------------------\n")
     	      direction = 0
	    }

	    //fmt.Printf("%s is %t\n",string(ch), unicode.IsDigit(ch))
	}
    }
    fmt.Printf("The number is %d\n", num)
    fmt.Printf("The count is %d\n", count)
    fmt.Printf("the number of clicks is %d\n", clicks)
    fmt.Printf("the number of clicks and counts is %d\n", clicks+count)
}

func mod(a, b int) int {
    m := a % b
    if a < 0 && b < 0 {
        m -= b
    }
    if a < 0 && b > 0 {
        m += b
    }
    return m
}