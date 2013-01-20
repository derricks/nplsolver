// The various methods, etc. for the actual server piece of the solver: socket listeners, message decoders, the like
package server

import (
  "fmt"
  "net"
  "time"
  "io"
  "encoding/binary"
  "encoding/json"
  "nplsolver/properties"
  "nplsolver/solver"
  "nplsolver/dict"
)

const (
   connectionTimeoutProp = "server.connection.timeout.seconds"
   maxMessageSize = 512
)

// the query object we might expect from JSON on the socket
type query struct{
   PuzzleType string // e.g., transposal or letterbank
   PuzzleText string // the text to search for 
   Dictionary string // which dictionary to use
   Options map[string]interface{} // any solver specific additional options (e.g., which letter change in a letter change) or extra information needed by the query
}

type response struct {
   Request query
   Err     error
   Results []string  
}

// this method will never return unless there's an error!
func StartServer() error {
  listener, err := net.Listen("tcp",fmt.Sprintf(":%v",properties.Get("server.port")))
  if err != nil {
     return err
  }
  
  for { // loop forever!
    conn, err := listener.Accept()
    if err != nil {
       continue
    }
    
    go handleConnection(conn)  
  }
  return nil
}

func handleConnection(conn net.Conn) {

   defer conn.Close()

   // make sure the connection won't hang with a timeout
   timeout, err := properties.GetInteger(connectionTimeoutProp)
   if err != nil {
      panic(fmt.Sprintf("Invalid connection timeout property: %v",properties.Get(connectionTimeoutProp)))
   }
   
   // set deadline to now plus timeout
   conn.SetDeadline(time.Now().Add(time.Duration(timeout) * time.Second))
   
   jsonData,err := readData(conn)
   if err != nil {
     logError("Did not unmarshal json data",err)
      return
   }
      
   solver,err := solver.GetSolver(solver.SolverType(jsonData.PuzzleType))   
   if err != nil {
      logError("Couldn't get solver",err)
      err = writeResponse(conn, response{jsonData,err,nil})
      if err != nil {
         return
      }
   }
   
   dict,err := dict.FindDictionaryByName(jsonData.Dictionary)
   if err != nil {
      err = writeResponse(conn,response{jsonData,err,nil})
      if err == nil {
         return
      }
   }
   
   // solve the puzzle!
   results := make([]string,0,100)
   resultsChan := make(chan string)
   doneChan := make(chan interface{})
   go solver.Solve(jsonData.PuzzleText,dict,resultsChan,doneChan)
   
   done := false
   for !done {
       select {
           case curResult := <-resultsChan:
              results = append(results,curResult)
           case finishItem := <-doneChan:
               done = true
               switch finishItem.(type) {
                  case error:
                     writeResponse(conn,response{jsonData,err,nil}) // no err trapping currently needed here; the connection's about to close anyway
                  default: // success at last!
                    err = writeResponse(conn,response{jsonData,nil,results})   
                    if err != nil {
                        logError("Couldn't write response to client",err)
                    }
               }
       }
   }
}

func logError(message string, err error) {
   fmt.Printf("Error: %v %v\n",message,err)
}

// pulls the json data off of a connection
func readData(conn net.Conn) (query,error) {
    
   buf := make([]byte,maxMessageSize)
   
   // read all the data off the connection
   bytesRead, err := conn.Read(buf)
   
   if (err != nil && err != io.EOF) || bytesRead == 0 {
      return *new(query),err
   }
   
   // unmarshal the relevant bytes into json structure
   jsonData := new(query)
   err = json.Unmarshal(buf[0:bytesRead],&jsonData)
   if err != nil {
      return *new(query),err
   }
   
   return *jsonData,nil
}

// takes a response object and writes it back down the socket with the same format (2 bytes for length and then json string)
func writeResponse(conn net.Conn, res response) error {
  fmt.Printf("Sending back %v\n",res)
  bytes,err := json.Marshal(res)
  if err != nil {
     return err
  }
  
  _,err = conn.Write(bytes)
  if err != nil {
     return err
  }
  
  return nil
}

// given a connection, reads the first two bytes and converts them to a 16-bit integer
func getMessageLength(conn net.Conn) (uint16,error) {   
   var retval uint16
   err := binary.Read(conn,binary.BigEndian,&retval)
   return retval,err
}