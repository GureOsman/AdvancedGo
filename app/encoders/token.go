package encoders
import(
	"io"
	"io/ioutil"
	"encoding/json"
	"log"
	//"github.com/gureosman/AdvancedGo/app/util"
	"github.com/gureosman/AdvancedGo/app/models"
)

func EncodeToken(body io.ReadCloser)(token models.Token)  {
	var data, _ = ioutil.ReadAll(body)
	if err :=json.Unmarshal(data ,&token); err !=nil{
		log.Println(err)
		return
	}
	return


}
