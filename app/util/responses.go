package util

type response struct  {

	Status string    `json:"status"`
	Body   interface{}   `json:"body"`

}

func ResponseError(msg string)response  {
	var err response
	err.Status = "failed"
	err.Body =msg
	return err


}
func ResponseSuccess(body interface{})response {
	var succ response
	succ.Status = "success"
	succ.Body = body
	return succ


}

