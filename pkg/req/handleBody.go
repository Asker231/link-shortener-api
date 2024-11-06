package req

import (
	"net/http"
	"github.com/Asker231/link-shortener-api.git/pkg/res"

)

func HandleBody[T any](w http.ResponseWriter, r *http.Request)(*T , error){
		body,err := Decode[T]( r.Body)
		if err != nil{
			res.JsonRes(w,err.Error(),402)
		}
		err = IsValid(body)
		if err != nil{
			res.JsonRes(w,err,402)
		}
		return &body,nil
	   
}