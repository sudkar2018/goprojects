package main

import ( "fmt"
  "flag"

)

func main(){
    s:=flag.String("s","default value","help text")
    flag.Parse()
    fmt.Printf("Hello:%s\n",*s)

}
