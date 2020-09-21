package utils

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/sirupsen/logrus"
)

// EndpointExec execute command with HTTP on Keepalived host
func EndpointExec(u string) (*bytes.Buffer, error) {
	response, err := http.Get(u)
	if err != nil {
		logrus.WithField("url", u).WithError(err).Error("Error sending request to endpoint")
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		logrus.WithField("statuscode", response.StatusCode).Error("Request was not successfull")
		return nil, errors.New("Request was not successfull")
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		logrus.WithError(err).Error("Error parsing response")
		return nil, err
	}

	return bytes.NewBuffer(body), nil
}
