package util

import (
   "testing"
)

func TestIsStringEmpty(t *testing.T) {
  if !IsStringEmpty("   ") {
     t.Error("Empty string was not considered empty")
  }
  
  if IsStringEmpty("    stuff ") {
     t.Error("Non-empty string considered empty")
  }
}

func TestFileProcessing(t *testing.T) {
   var foundLine1, foundLine2 bool
   
   ProcessNonEmptyFileLines("./test_file.txt",'\n', func(line string) {
      if line == "line1" {
         foundLine1 = true
      } else if line == "line2" {
         foundLine2 = true
      } else {
         t.Errorf("Unexpected line in test file: %v",line)
      }
   })
   
   if !(foundLine1 && foundLine2) {
     t.Errorf("Did not receive line1 or line2 in test")
   }
}