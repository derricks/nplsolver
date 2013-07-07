package solver

import (
   "nplsolver/dict"
   "nplsolver/transform"
)

type TransAdditionSolver struct{}
func (solver TransAdditionSolver) Solve(pattern string, dictionary dict.Dictionary, results chan<- string, done_channel chan<- interface{}) {

  // this works by finding transposals for the pattern plus each character
  extras := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
  findWords := make([]string, len(extras))
  
  for index, extra := range extras {
    // for each letter, look for transposals of that letter plus the original pattern
    // so make the transformed version of pattern + each letter
    findWords[index] = transform.SortAllCharacters(extra + pattern)
  }
 
  if matcher,err := newSameCharacterMatcher(); err != nil {
     done_channel <- err
  } else {
    // search for any item in the dictionary that has the same unordered set (using MatchTransformed)
    dictionary.Iterate(func(entry dict.Entry) {
       for _, word := range findWords {
         if matcher.MatchTransformed(word, entry) {
           results <- entry.Word()
         }
       }
    })
  }
  done_channel <- solver
}