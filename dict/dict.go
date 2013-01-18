// Various utilities for working with dictionaries, including iterating through them
package dict

import (
   "os"
   "bufio"
   "strings"
   "fmt"
   "nplsolver/util"
   "nplsolver/properties"
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

  // open the new file (trunc if it already exists)
  destFile, destErr := os.Create(dest)
  if destErr != nil {
     return destErr
  }
  defer destFile.Close()
  destWriter := bufio.NewWriter(destFile)
  
  writeErr := util.ProcessNonEmptyFileLines(source, '\n', func(line string) error {
      entry, entryErr := NewEntryFromWord(line)
      if entryErr != nil {
         return entryErr
      }
  
      destLine := fmt.Sprintf("%v\n",strings.Join(entry,"\t"))
      destWriter.WriteString(destLine)
      return nil
  })
  
  if writeErr != nil {
     return writeErr
  }

  destWriter.Flush()

  return nil  
}

func getDictionaryNamePropForName(name string) string {
  return fmt.Sprintf("dictionaries.%v.name",name)
}

func getDictionaryPathPropForName(name string) string {
  return fmt.Sprintf("dictionaries.%v.path",name)
}
// Given a Dictionary name, returns the prop for the dictionary name and the dictionary cache file
func GetDictionaryPropsForName(name string) (string,string) {
  return getDictionaryNamePropForName(name),getDictionaryPathPropForName(name)
}

// figures out the cached file for the given dictionary name
func FindDictionaryByName(dictName string) (dict Dictionary,err error) {
   dictName,dictPath := GetDictionaryPropsForName(dictName)
   return NewDictionaryFromFile(properties.Get(dictPath))
}

