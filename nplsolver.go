// Entry point for the nplsolver program
package main

import (
   "fmt"
   _ "nplsolver/dict"
   _ "nplsolver/solver"
)

// main makes sure dictionaries and caches are set up and then sits on a socket waiting for commands
func main() {
  fmt.Println("Started NPL Solver Server")
}