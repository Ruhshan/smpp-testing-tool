package main

import (
	"github.com/fiorix/go-smpp/smpp"
	"github.com/fiorix/go-smpp/smpp/pdu/pdufield"
	"github.com/fiorix/go-smpp/smpp/pdu/pdutext"
	"log"
)

func main() {
	tx := &smpp.Transmitter{
		Addr:   "localhost:8057",
		User:   "test",
		Passwd: "test",
	}
	// Create persistent connection, wait for the first status.
	conn := <-tx.Bind()
	if conn.Status() != smpp.Connected {
		log.Fatal(conn.Error())
	}

	sm, err := tx.Submit(&smpp.ShortMessage{
		Src:      "sender",
		Dst:      "recipient",
		Text:     pdutext.UCS2("আমার সোনার বাংলা আমি তোমাউয় ভালোবাস"),
		Register: pdufield.NoDeliveryReceipt,
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Message ID:", sm.RespID())

}
