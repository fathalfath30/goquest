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

package mapper

import (
	"encoding/json"
	"io"
	"net/http"
)

// JsonResponse will map any json response to expected struct
func JsonResponse[T any](response *http.Response) (T, error) {
	var obj T
	var err error
	var rsp []byte

	// close json body
	defer response.Body.Close()

	// read response body and transform it into []byte to unmarshall
	// with generic type
	rsp, err = io.ReadAll(response.Body)
	if err != nil {
		return obj, err
	}

	// unmarshalling json
	err = json.Unmarshal(rsp, &obj)
	if err != nil {
		return obj, err
	}

	return obj, StatusCode(response)
}
