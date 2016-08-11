package encoders

import (
	"io"
	"github.com/gureosman/AdvancedGo/app/models"
	"io/ioutil"
	"encoding/json"
	"log"
)

func EncodeTicket(body io.ReadCloser)(tickets models.Tickets)  {
	var ticket, _ = ioutil.ReadAll(body)
	if err :=json.Unmarshal(ticket,&tickets); err !=nil{
		log.Println(err)
		return
	}
	return


}
