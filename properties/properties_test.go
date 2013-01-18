package properties

import (
   "testing"
)

const (
  testPropsFile = "./test.properties"
)

func TestPropertyLoading(t *testing.T) {
  loadProperties(t)  
  if Get("server.port") != "15023" {
     t.Errorf("Invalid property value for server.port: %v",Get("server.port"))
  }
}

func TestIntegerProperties(t *testing.T) {
   loadProperties(t)
   data,err := GetInteger("server.port")
   if err != nil {
      t.Errorf("Could not parse integer property %v",err)
   }
   if data != 15023 {
      t.Errorf("Incorrect value for server.port. Should be 15023, is %v",data)
   }
}

func loadProperties(t *testing.T) {
    err := LoadPropertiesFromFiles(testPropsFile)
    if err != nil {
       t.Errorf("Failed to load properties: %v",err)
    }
}