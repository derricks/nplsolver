// Various utilities for working with dictionaries, including iterating through them
package dict

import (
//  "io"
  "bufio"
)

type Dictionary interface {
   // returns the next entry in the dictionary or nil if the dictionary is at its end
   NextEntry() Entry
   
   // tells the dictionary to shut down any resources it was using
   Close()
}

// Given a text file of words, one per line, make a new file that has optimized lookups for different kinds of word patterns
func MakeDictionaryFromTextFile(source string, dest string) error {
  return nil  
}

// Creates a new dictionary based off of reading a file
func NewDictionaryFromFile(filename string) (Dictionary,error) {
   return nil,nil
}

type fileDictionary struct {
   bufio.Reader
}



