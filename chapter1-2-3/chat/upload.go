package main

import (
	"net/http"
	"io"
	"io/ioutil"
	"path"
)


func uploaderHandler(w http.ResponseWriter, req *http.Request) {
	userID := req.FormValue("userid")
	file, header, err := req.FormFile("avatarFile")
	if err != nil {
		io.WriteString(w, err.Error())
		return
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		io.WriteString(w, err.Error())
		return
	}
	filename := path.Join("avatars", userID + path.Ext(header.Filename))
	err = ioutil.WriteFile(filename, data, 0777)
	if err != nil {
		io.WriteString(w, err.Error())
		return
	}
	io.WriteString(w, "Successful")
}