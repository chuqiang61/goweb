package main

import (
	"io"
	"net/http"
	"os/exec"
	"log"
)

func reLaunch()  {
	cmd :=exec.Command("sh","./deploy.sh")
	err :=cmd.Start()
	if err!=nil {
		log.Fatal(err)
	}
	err = cmd.Wait()
}

func firstpage(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, " <h1>hello,this is my deploy server</h1>")
	reLaunch()

}

func main() {
	http.HandleFunc("/", firstpage)
	http.ListenAndServe(":8000", nil)
}
