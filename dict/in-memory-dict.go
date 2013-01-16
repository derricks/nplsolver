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

func (dict *inMemoryDictionary) Close() { 
   dict.index = 0 // allow the dictionary to be re-used
}

func (dict *inMemoryDictionary) Iterate(handler func(entry Entry)) {
  iterateOverDictionaryEntries(dict,handler)
}

// returns an in-memory dictionary with the specified words. primarily used for testing
func NewDictionaryFromWords(words ...string) (Dictionary,error) {
  // make a list of Entry objects derived from those words
  entries := make([]Entry,0,len(words)) // specify initial length (0) and capacity to ensure entries get written at 0

  for _,word := range words {
    entry,_ := NewEntryFromWord(word)
    entries = append(entries,entry)
  }
  
  // configure and return an inMemoryDictionary
  return &inMemoryDictionary{0,entries},nil    
}

