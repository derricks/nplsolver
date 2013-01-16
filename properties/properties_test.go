package properties

import (
   "testing"
)

func TestPropertyLoading(t *testing.T) {
  err := LoadPropertiesFromFiles("./test.properties")
  if err != nil {
     t.Errorf("Failed to load properties: %v",err)
  }
  
  if Get("server.port") != "15023" {
     t.Errorf("Invalid property value for server.port: %v",Get("server.port"))
  }
}