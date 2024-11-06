package res

import (
	"net/http" 
"encoding/json")

	func JsonRes(w http.ResponseWriter,data any,code int){
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(code)
		json.NewEncoder(w).Encode(data)
	}