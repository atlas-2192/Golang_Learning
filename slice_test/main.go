package main

import (
	//"errors"
	"fmt"
	"strings"
)

type SmsRecipients struct {
	RecipientName   string `json:"recipient_name"`
	RecipientNumber string `json:"recipient_number"`
	Status          string
}

type SmsMsg struct {
	Contry     string          `json:"country"`
	Content    string          `json:"content"`
	Type       string          `json:"type"`
	Recipients []SmsRecipients `json:"recipients"`
}

func main() {

	//localSms := SmsMsg{ Contry:"jp", Content: "test content", Type:"record", Recipients: [{RecipientName: "test1", RecipientNumber:"1234"}, {RecipientName: "test2", RecipientNumber:"5678"}]}
	localSms := SmsMsg{Contry: "jp", Content: "test content", Type: "record"}
	localSms.Recipients = append(localSms.Recipients, SmsRecipients{RecipientName: "test2", RecipientNumber: "5678"})
	localSms.Recipients = append(localSms.Recipients, SmsRecipients{RecipientName: "test1", RecipientNumber: "1234"})
	fmt.Printf("struct = %+v \n", localSms)

	for index, _ := range localSms.Recipients {
		localSms.Recipients[index].Status = fmt.Sprintf("index %d, work", index)
	}
	fmt.Printf("struct = %+v \n", localSms)

	s := strings.Split("35.348,77.3483", ",")
	long, lat := s[0], s[1]
	fmt.Printf("long:%s_lat:%s \n", long, lat)
}
