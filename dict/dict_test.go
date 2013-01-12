package dict_test

import (
  "nplsolver/dict"
  "testing"
)

func TestNewEntryFromWord(t *testing.T) {
  word := "hello"
  entry,_ := dict.NewEntryFromWord(word)
  if entry.Word() != word {
     t.Errorf("Entry word was not what was passed in: %v != %v",word,entry.Word())
  }
  
  if entry.LettersOrdered() != "ehllo" {
     t.Errorf("Entry's LettersOrdered method is incorrect. Should be ehllo, is %v",entry.LettersOrdered())
  }
}

