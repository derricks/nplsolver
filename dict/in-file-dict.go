package dict

import (
   "bufio"
   "os"
   "strings"
   "nplsolver/util"
)

const (
  tab = "\t"
  newline = '\n'
)

// Creates a new dictionary based off of reading a file. Note that this assumes that said file is in our pre-defined format (see entry.go)
func NewDictionaryFromFile(filename string) (Dictionary,error) {
   file,err := os.Open(filename)
   if err != nil {
       return nil,err
   }
   
   reader := bufio.NewReader(file)
   dictionary := fileDictionary{file,reader}
   return dictionary,nil
}

type fileDictionary struct {
   file *os.File
   *bufio.Reader
}

func (dict fileDictionary) NextEntry() Entry {
   line,err := dict.ReadString(newline) // use the embedded Reader's methods
   if err != nil {
      return nil;
   }
   trimmedLine := strings.Trim(line,util.Whitespace)
   entry := Entry(strings.Split(trimmedLine,tab))
   return entry
}

func (dict fileDictionary) Close() {
  dict.file.Close()
}

func (dict fileDictionary) Iterate(handler func(entry Entry)) {
  iterateOverDictionaryEntries(dict,handler)
}