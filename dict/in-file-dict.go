package dict

import (
   "bufio"
)

// Creates a new dictionary based off of reading a file. Note that this assumes that said file is in our pre-defined format (see entry.go)
func NewDictionaryFromFile(filename string) (Dictionary,error) {
   return nil,nil
}

type fileDictionary struct {
   bufio.Reader
}



