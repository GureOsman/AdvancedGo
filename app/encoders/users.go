package encoders

import(
	"io"
        "io/ioutil"
        "encoding/json"
        "log"
	"github.com/gureosman/AdvancedGo/app/models"
)

func EncodeSingleUsers(body io.ReadCloser)(user models.User)  {
	var data, _ = ioutil.ReadAll(body)
	if err :=json.Unmarshal(data ,&user); err !=nil{
		log.Println(err)
		return
	}
	return
}

