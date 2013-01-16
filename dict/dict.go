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
   
   // iterate over each entry in the dictionary, applying the relevant function
   Iterate(handler func(entry Entry))
}

// refactored method for walking the dictionary entries and passing each entry to the function
func iterateOverDictionaryEntries(dictionary Dictionary, entryHandler func(entry Entry)) {
  for curEntry := dictionary.NextEntry(); curEntry != nil; curEntry = dictionary.NextEntry()   {
    entryHandler(curEntry)
  }
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



