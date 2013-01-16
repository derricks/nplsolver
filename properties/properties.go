// Various methods for loading and accessing properties. Properties are stored by key and have a value
// TODO: Allow for hot reloading of properties using channels for thread-safe updates
package properties

import (
   "strings"
   "nplsolver/util"
)

const (
  whitespace = "\t\n "
)

type properties map[string]string

var props properties

// does the actual work of loading properties from the specified files. 
// Properties with the same name in different files are overwritten; in other words, properties in later files override properties in earlier files
// current properties format is key=value
func LoadPropertiesFromFiles(files ...string) error {
   props = make(properties)
   
   for _,curFile := range files {
      err := loadPropertiesFromFile(curFile)
      if err != nil {
        return err
      }
   }
   return nil
}

func loadPropertiesFromFile(propFile string) error {
  err := util.ProcessNonEmptyFileLines(propFile, '\n', func(line string) error {
     key, value := parsePropertyLine(line)  
     props[key] = value
     return nil
  })
  
  if err != nil {
     return err
  }
  return nil
}

func parsePropertyLine(line string) (key string, value string) {
   splitLine := strings.Split(line,"=")
   return strings.Trim(splitLine[0],whitespace),strings.Trim(splitLine[1],whitespace)
}

func Get(propertyName string) (propertyValue string) {
   return props[propertyName]
}