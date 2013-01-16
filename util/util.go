// Useful utilities
package util

import (
   "io"
   "os"
   "bufio"
   "strings"
)

const (
   Whitespace = "\t\r\n "
)

// Process a byte-delimited text file by calling lineHandler for each non-empty (i.e., non-whitespace-containing) line
func ProcessNonEmptyFileLines(filePath string, delim byte, lineHandler func(line string)) error {
   file, err := os.Open(filePath)
   if (err != nil) {
      return err
   }
   defer file.Close()
   
   reader := bufio.NewReader(file)
   
   for line,err := reader.ReadString(delim); true; line,err = reader.ReadString(delim) {
       if IsStringEmpty(line) {
          continue
       }
       lineHandler(strings.Trim(line,Whitespace))
       if err == io.EOF {
          break
       }
   }
   return nil
}

func IsStringEmpty(text string) bool {
   return strings.Trim(text,Whitespace) == ""
}