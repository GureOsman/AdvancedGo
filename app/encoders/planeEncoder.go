package encoders

import (
	"io"
	"github.com/gureosman/AdvancedGo/app/models"
	"io/ioutil"
	"encoding/json"
	"log"
)

func EncodePlane(body io.ReadCloser)(planes models.Planes)  {
	var Planes, _ = ioutil.ReadAll(body)
	if err :=json.Unmarshal(Planes,&planes); err !=nil{
		log.Println(err)
		return
	};
	return


}
