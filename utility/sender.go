package utility

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

type HxAgentData struct {
	hostname string `json:"hostname"`
}

const(
	hxAgentUrl = "http://xxxx"
	contentType = "application/json;charset=UTF-8"
)

// in job.Handler()
//func (j *job)conventToHxAgentData() HxAgentData{
// add a filed named HxAgentData for job!
func JobToHxAgentData(v struct{}) HxAgentData{
	return HxAgentData{}
}

func SendData(data HxAgentData) error {
	bytesData, err := json.Marshal(data)
	if err != nil {
		return errors.New("E] interrupted data sent, err: " + err.Error())
	}
	formData := url.Values{
		//change name!!!
		"mode":{"hxAgent"},
		"sendData":{string(bytesData)},
	}

	resp, err := http.PostForm(hxAgentUrl, formData)
	if err != nil{
		return errors.New("E] fail to get response, err: " + err.Error())
	}
	defer 	resp.Body.Close()
	//重用TCP
	io.Copy(ioutil.Discard, resp.Body)
	return nil
}
