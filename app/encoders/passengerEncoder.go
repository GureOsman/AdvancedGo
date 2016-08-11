package encoders
import (
	"io"
	"github.com/gureosman/AdvancedGo/app/models"
	"io/ioutil"
	"encoding/json"
	"log"
)
func EncodePassengers(body io.ReadCloser)(passengers models.Passengers)  {
	var Passenger, _ = ioutil.ReadAll(body)
	if err :=json.Unmarshal(Passenger,&passengers); err !=nil{
		log.Println("passengers: ",passengers )
		log.Println(err)
		return
	}
	return
}
