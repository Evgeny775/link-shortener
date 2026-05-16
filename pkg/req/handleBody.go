package req

import (
	"fmt"
	"link-shortener/pkg/res"
	"net/http"
)

func HandleBody[payloadType any](w http.ResponseWriter, r *http.Request) (*payloadType, error) {

	body, err := Decode[payloadType](r.Body)
	if err != nil {
		fmt.Println("Decode err: ", err.Error())
		res.JSON(w, err.Error(), http.StatusBadRequest)
		return nil, err
	}

	err = Validate[payloadType](body)
	if err != nil {
		println("Validate err: ", err.Error())
		res.JSON(w, err.Error(), http.StatusBadRequest)
		return nil, err
	}
	return body, nil
}
