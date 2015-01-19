package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

func main() {
	type Say struct {
		Voice string `xml:"voice,attr"`
		Value string `xml:",chardata"`
	}
	type CallResponse struct {
		XMLName xml.Name `xml:"Response"`
		Play    string   `xml:"Play"`
		SayObj  Say      `xml:"Say"`
	}

	v := &CallResponse{Play: "http://demo.twilio.com/docs/classic.mp3"}
	//v.Comment = " Need more details. "
	//v.Address = Address{"Hanga Roa", "Easter Island"}
	v.SayObj = Say{Voice: "alice", Value: "Thanks for trying our documentation. Enjoy!"}

	enc := xml.NewEncoder(os.Stdout)
	enc.Indent("  ", "    ")
	if err := enc.Encode(v); err != nil {
		fmt.Printf("error: %v\n", err)
	}

}
