package dict

type inMemoryDictionary struct {
   index int
   entries []Entry
}


func (dict *inMemoryDictionary) NextEntry() Entry {
   if dict.index >= len(dict.entries)  {
      return nil
   }
   
   entry := dict.entries[dict.index]
   dict.index++
   return entry
}

func (dict *inMemoryDictionary) Close() { // no-op
}

// returns an in-memory dictionary with the specified words. primarily used for testing
func newDictionaryFromWords(words []string) (Dictionary,error) {
  // make a list of Entry objects derived from those words
  entries := make([]Entry,0,len(words)) // specify initial length (0) and capacity to ensure entries get written at 0

  for _,word := range words {
    entry,_ := NewEntryFromWord(word)
    entries = append(entries,entry)
  }
  
  // configure and return an inMemoryDictionary
  return &inMemoryDictionary{0,entries},nil    
}

