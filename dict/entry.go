package dict

import (
   "nplsolver/transform"
   "errors"
)

// "enum" for what letter permutation is where
// when adding to this list, update the numPermutations constant, too!
const (
   word = iota            // the basic word. hello -> hello
   ordered                // the letters of the word, sorted. hello -> ehllo
   uniqueOrdered          // the unique letters of the word, sorted. hello -> ehlo
   consonantsOrdered      // just the consonants, sorted. hello -> hll
   pattern                // the letter pattern. e.g., hello -> 12334
)

const (
   // update this when adding to the above!
   numPermutations = 5
)


// An Entry represents a single dictionary entry, corresponding to a "row" of permutations of the letters in a word.
// Slots in the slice are identified by the word,ordered, etc. constants.
// An entry is immutable.
type Entry []string

// returns the basic word associated with this entry.
func (entry Entry) Word() string {
  return entry[word]
}

func (entry Entry) LettersOrdered() string {
  return entry[ordered]
}

// Return a new dictionary entry from the given base word
func NewEntryFromWord(newWord string) (Entry,error) {

   if len(newWord) == 0 {
      return nil, errors.New("Can not make a new entry from a nil or empty word")
   }
   
   // todo: check if the string is all whitespace
   
   entry := make([]string,numPermutations)
   entry[word] = newWord
   entry[ordered] = transform.SortAllCharacters(newWord)
   entry[uniqueOrdered] = transform.UniqueSortedCharacters(newWord)
   return entry,nil
}