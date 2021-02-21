# gopylists
A module for go that adds multiple type arrays.

## How to use
```go
package main

import (
  "fmt"
  "github.com/Wartori54/gopylists"
 )

func main() {
  //create a new blank list
  list := gopylists.NewPyList()
  
  //create a new list from arguments
  listfromarguments := gopylists.NewPyListFromArgs(1.0, "text", complex(1,2))
  
  //print the list
  fmt.Println(listfromarguments)
  
  //append to the list
  listfromarguments.Append(89)
  
  //or insert
  listfromarguments.Insert(-5, 2)
  
  //getting an element
  fmt.Println(listfromarguments.Get(1))
  
  //concatenate two lists
  list.Append("abc")
  listfromarguments.Concatenate(list)
  
  //get how many times appear an element
  listfromagruments.Count("text") // 1
  
  //and the "in" operator from python
  fmt.Println(gopylists.In([]int{1, 2, 3, 4}, 3) // true
}
```