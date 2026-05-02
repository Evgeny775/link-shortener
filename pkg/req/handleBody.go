package req

import (
	"link-shortener/pkg/res"
	"net/http"
)

func HandleBody[payloadType any](w http.ResponseWriter, r *http.Request) (*payloadType, error) {

	body, err := Decode[payloadType](r.Body)
	if err != nil {
		res.JSON(w, err.Error(), 400)
		return nil, err
	}

	err = Validate[payloadType](body)
	if err != nil {
		res.JSON(w, err.Error(), 400)
		return nil, err
	}
	return body, nil
}
