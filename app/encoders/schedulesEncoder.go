package encoders

import (
	"io"
	"github.com/gureosman/AdvancedGo/app/models"
	"io/ioutil"
	"encoding/json"
	"log"
)

func EncodeSchedules(body io.ReadCloser)(schedules models.Schedule)  {
	var schedule, _ = ioutil.ReadAll(body)
	if err :=json.Unmarshal(schedule,&schedules); err !=nil{
		log.Println(err)
		return
	}
	return


}
