package route

import (
	"net/http"
	"lurcury/http/model"
)

func Router(){
	//http.HandleFunc("/keyStore", model.KeyStoreProduct)
	http.HandleFunc("/testparams", model.TestParams)
	http.HandleFunc("/testbodys", model.TestBodys)
}
