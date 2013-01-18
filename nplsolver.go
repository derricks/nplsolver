// Entry point for the nplsolver program
package main

import (
   "fmt"
   "strings"
   "os"
   "errors"
   "path/filepath"
   "nplsolver/dict"
   "nplsolver/properties"
   "nplsolver/util"
   "nplsolver/solver"
)

const (
   FILE_PERMS = 0777
)

// main makes sure dictionaries and caches are set up and then sits on a socket waiting for commands
func main() {
  fmt.Println("Started NPL Solver Server")
  
  // load all the properties
  err := properties.LoadPropertiesFromFiles("./solver.properties")
  assertError("Could not load properties",err)
  
  fmt.Println("Properties loaded")
  assertError("Could not load properties",err)
    
  // ensure the caches directory exists
  cachesDir,err := getCachesDirectory()
  assertError("Error getting directory",err)
  
  err = os.MkdirAll(cachesDir,FILE_PERMS)
  assertError("Could not make caches directory",err)
  
  err = checkDictionaries()
  assertError("Could not make dictionaries",err)
  
  
 }

// Will panic with the given message if err is not nil
func assertError(message string,err error) {
  if err != nil {
     fmt.Printf("Error: %v, %v\n",message,err)
  }
}

func getCachesDirectory() (string,error) {
    return filepath.Abs(properties.Get("server.cache.dir"))
}

func checkDictionaries() error {
    
    cachesDir,err := getCachesDirectory()
    if err != nil {
      return err
    }
    
    dictionaryAliases := getDictionaryNames()
    if len(dictionaryAliases) == 0 {
       return errors.New("No dictionaries found in properties")
    }
    
    fmt.Printf("dictionaries: %v\n",dictionaryAliases)
    
    // for the given dictionaries, make sure each dictionary has a pre-made version available in the caches directory
    for _,alias := range dictionaryAliases {
      fmt.Printf("Checking status of dictionary %v\n",alias)
    
      cachedFileName := strings.Join([]string{cachesDir,alias},string(os.PathSeparator))
 
      if !util.DoesFileExist(cachedFileName) {
         fmt.Printf("Pre-made dictionary file does not exist. Creating at %v\n",cachedFileName)
         dictSourceFileKey := fmt.Sprintf("dictionaries.%v.path",alias)
         err = dict.MakeDictionaryFromTextFile(properties.Get(dictSourceFileKey),cachedFileName)
         if err != nil {
            return err
         }
         fmt.Printf("Created dictionary cache at %v\n",cachedFileName)
      } 
    }
    return nil
}

// under the current version of Go, strings.Split will return the empty string if the given string does not contain any instance of the
// separator
func getDictionaryNames() []string {
    comma := ","
    dictionaryAliasesProp := properties.Get("dictionaries.aliases")
    if dictionaryAliasesProp == "" {
       return make([]string,0,0)
    }
    
   return strings.Split(dictionaryAliasesProp,comma)
}