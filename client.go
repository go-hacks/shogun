package main
import (
  "fmt"
  "net"
  "bufio"
  "os"
  "strconv"
)

func main() {
  var conn [257]net.Conn
  var err error
  for i := 0; i < 257; i++ {
    addr := "127.0.0.1:" + strconv.Itoa(9000+i)
    conn[i], err = net.Dial("udp", addr)
    if err != nil {
     fmt.Printf("Some error %v", err)
     return
    }
    defer conn[i].Close()
  }
  fmt.Printf("Starting readers...")
  for i := 0; i < 256; i++ {
    go reader(conn[i], i)
  }
  fmt.Printf("Done.\n")
  fmt.Printf("Sending start signal...")
  for i := 0; i < 257; i++ {
    fmt.Fprintf(conn[i], "")
  }
  fmt.Printf("Done.\n")
  fmt.Printf("Receiving file...")
  outFile, _ := os.Create("testfile")
  outFile.Close()
  p :=  make([]byte, 1)
  _, _ = bufio.NewReader(conn[256]).Read(p)
  fmt.Printf("Complete!\n")
  return
}

func reader(conn net.Conn, val int) {
  for {
    p :=  make([]byte, 1)
    _, err := bufio.NewReader(conn).Read(p)
    if err == nil {
      writer(byte(val))
    }
  }
}

func writer(v byte) {
  val := make([]byte, 1)
  val[0] = v
  outFile, _ := os.OpenFile("testfile",os.O_APPEND|os.O_WRONLY,0666)
  _, _ = outFile.Write(val)
  outFile.Close()
}
