package handlers

import (
	"math/rand"
	"net/http"
	"time"
)

func HardOp(w http.ResponseWriter, r *http.Request) {
	timer1 := time.Duration(10+rand.Int()%11) * time.Second
	time.Sleep(timer1)
	if rand.Int()%2 == 0 {
		http.Error(w, "Internal Server Error", 500+rand.Int()%5)
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Operation completed"))
	}
}
