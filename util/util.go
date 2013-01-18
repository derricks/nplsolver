// Useful utilities
package util

import (
"fmt"
   "io"
   "os"
   "bufio"
   "strings"
)

const (
   Whitespace = "\t\r\n "
)

// Check to make sure the file exists
func DoesFileExist(path string) bool {
  _, err := os.Stat(path)
  fmt.Printf("does file exist: %v\n",err)
  if err != nil && os.IsNotExist(err) {
     return false
  } 
  return true
}

// Process a byte-delimited text file by calling lineHandler for each non-empty (i.e., non-whitespace-containing) line
func ProcessNonEmptyFileLines(filePath string, delim byte, lineHandler func(line string) error) error {
   file, err := os.Open(filePath)
   if (err != nil) {
      return err
   }
   defer file.Close()
   
   reader := bufio.NewReader(file)
   
   for line,err := reader.ReadString(delim); true; line,err = reader.ReadString(delim) {
       if IsStringEmpty(line) {
          if err == io.EOF {
             break
          } else {
            continue
          }
       }
       
       lineErr := lineHandler(strings.Trim(line,Whitespace))
       if lineErr != nil {
          return lineErr
       }
       
       if err == io.EOF {
          break
       }
   }
   return nil
}

func IsStringEmpty(text string) bool {
   return strings.Trim(text,Whitespace) == ""
}