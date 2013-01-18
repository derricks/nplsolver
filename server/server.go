// The various methods, etc. for the actual server piece of the solver: socket listeners, message decoders, the like
package server

import (
  "fmt"
  "net"
  "time"
  "nplsolver/properties"
)

const (
   connectionTimeoutProp = "server.connection.timeout.seconds"
)

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
   // make sure the connection won't hang with a timeout
   timeout, err := properties.GetInteger(connectionTimeoutProp)
   if err != nil {
      panic(fmt.Sprintf("Invalid connection timeout property: %v",properties.Get(connectionTimeoutProp)))
   }
   
   // set deadline to now plus timeout
   conn.SetDeadline(time.Now().Add(time.Duration(timeout) * time.Second))
   
   
   //parse the message: it assumes that the first 2 bytes gives you a size of the incoming string
   //  the string is assumed to be JSON
   
   //solve the query encoded in the json
   
   //return a json object that is the answers
      //- note that it should contain the original query as is (allowing the client to add extra fields to identify queries)
   
   //close the connection
   conn.Close()
  
}