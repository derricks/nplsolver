// The various methods, etc. for the actual server piece of the solver: socket listeners, message decoders, the like
package server

import (
  "fmt"
  "net"
  "time"
  "encoding/binary"
  "encoding/json"
  "nplsolver/properties"
  "nplsolver/solver"
  "nplsolver/dict"
)

const (
   connectionTimeoutProp = "server.connection.timeout.seconds"
)

// the query object we might expect from JSON on the socket
type query struct{
   puzzleType string // e.g., transposal or letterbank
   puzzleText string // the text to search for 
   dictionary string // which dictionary to use
   options map[string]interface{} // any solver specific additional options (e.g., which letter change in a letter change) or extra information needed by the query
}

type response struct {
   request query
   err     error
   results []string  
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
       return err
    }
    
    go handleConnection(conn)  
  }
  return nil
}

func handleConnection(conn net.Conn) {

    //close the connection, no matter what
    defer conn.Close()

   // make sure the connection won't hang with a timeout
   timeout, err := properties.GetInteger(connectionTimeoutProp)
   if err != nil {
      panic(fmt.Sprintf("Invalid connection timeout property: %v",properties.Get(connectionTimeoutProp)))
   }
   
   // set deadline to now plus timeout
   conn.SetDeadline(time.Now().Add(time.Duration(timeout) * time.Second))
   
   //parse the message: it assumes that the first 2 bytes gives you a size of the incoming string
   messageLength,err := getMessageLength(conn)
   if err != nil || messageLength == 0 {
      return
   }
   
   //  the string is assumed to be JSON
   buf := make([]byte,0,messageLength)
   bytesRead, err := conn.Read(buf)
   
   if err != nil || bytesRead == 0 {
      return
   }
   
   jsonData := new(query)
   err = json.Unmarshal(buf,&jsonData)
   if err != nil {
      return
   }
   
      
   solver,err := solver.GetSolver(solver.SolverType(jsonData.puzzleType))   
   if err != nil {
      err = writeResponse(conn, response{*jsonData,err,nil})
      if err != nil {
         return
      }
   }
   
   dict,err := dict.FindDictionaryByName(jsonData.dictionary)
   if err != nil {
      err = writeResponse(conn,response{*jsonData,err,nil})
      if err == nil {
         return
      }
   }
   
   // solve the puzzle!
   results := make([]string,0,100)
   resultsChan := make(chan string)
   doneChan := make(chan interface{})
   go solver.Solve(jsonData.puzzleText,dict,resultsChan,doneChan)
   
   done := false
   for !done {
       select {
           case curResult := <-resultsChan:
              results = append(results,curResult)
           case finishItem := <-doneChan:
               done = true
               switch finishItem.(type) {
                  case error:
                     writeResponse(conn,response{*jsonData,err,nil}) // no err trapping currently needed here; the connection's about to close anyway
                  default: // success at last!
                    writeResponse(conn,response{*jsonData,nil,results})   
               }
       }
   }
}

// takes a response object and writes it back down the socket with the same format (2 bytes for length and then json string)
func writeResponse(conn net.Conn, res response) error {
  bytes,err := json.Marshal(res)
  if err != nil {
     return err
  }
  err = binary.Write(conn,binary.BigEndian,int16(len(bytes)))
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
func getMessageLength(conn net.Conn) (int16,error) {
   var retval int16
   err := binary.Read(conn,binary.BigEndian,&retval)
   return retval,err
}