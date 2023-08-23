/*
//
//  ______    _   _           _  __      _   _     ____   ___
// |  ____|  | | | |         | |/ _|    | | | |   |___ \ / _ \
// | |__ __ _| |_| |__   __ _| | |_ __ _| |_| |__   __) | | | |
// |  __/ _` | __| '_ \ / _` | |  _/ _` | __| '_ \ |__ <| | | |
// | | | (_| | |_| | | | (_| | | || (_| | |_| | | |___) | |_| |
// |_|  \__,_|\__|_| |_|\__,_|_|_| \__,_|\__|_| |_|____/ \___/
//
// Written by Fathalfath30.
// Email : fathalfath30@gmail.com
// Follow me on:
//  Github : https://github.com/fathalfath30
//  Gitlab : https://gitlab.com/Fathalfath30
//
*/

package goquest

import (
	"fmt"
	"net/http"
	"strings"
)

func (gq *GoQuest) Send(method, endpoint string, requestOption *RequestOption) (*Response, error) {
	req, err := http.NewRequest(method, gq.BaseUrl, strings.NewReader(""))
	if err != nil {
		return nil, err
	}

	// put default header
	header := http.Header{}
	if gq.header != nil {
		header = *gq.header
	}

	if requestOption != nil {
		// if request option has json we need to add "Content-Type: application/json" header if not exists
		if requestOption.Json != nil {
			// Loop over header names
			for name, values := range header {
				// Loop over all values for the name.
				for _, value := range values {
					fmt.Println(name, value)
				}
			}
		}
	}

	req.Header = header
	return nil, nil
}
