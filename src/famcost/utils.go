package main

import "net/http"
import "crypto/md5"
import "encoding/hex"

func checkInternalServerError(err error, w http.ResponseWriter) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func isAuthenticated(w http.ResponseWriter, r *http.Request) {
	if !authenticated {
		http.Redirect(w, r, "/login", 301)
	}
}

func md5V(str string) string  {
    h := md5.New()
    h.Write([]byte(str))
    return hex.EncodeToString(h.Sum(nil))
}