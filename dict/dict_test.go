package dict

import (
  "testing"
)

func TestNewEntryFromWord(t *testing.T) {
  word := "hello"
  entry,_ := NewEntryFromWord(word)
  if entry.Word() != word {
     t.Errorf("Entry word was not what was passed in: %v != %v",word,entry.Word())
  }
  
  if entry.LettersOrdered() != "ehllo" {
     t.Errorf("Entry's LettersOrdered method is incorrect. Should be ehllo, is %v",entry.LettersOrdered())
  }
}

func TestInMemoryDictionary(t *testing.T) {
  hello := "hello"
  testing := "testing"
  
  dict,_ := NewDictionaryFromWords(hello,testing)
  
  shouldBeHello := dict.NextEntry().Word()
  if shouldBeHello != hello {
     t.Errorf("Words do not match: %v %v",shouldBeHello,hello)
  }
  
  shouldBeTesting := dict.NextEntry().Word()
  if shouldBeTesting != testing {
     t.Errorf("Words do not match: %v %v",shouldBeTesting,testing)
  }
  
  shouldBeNil := dict.NextEntry()
  if shouldBeNil != nil {
     t.Errorf("Dictionary should now return nil, returned %v",shouldBeNil)
  }
}

