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

// JsonResponse mapping json response from byte to selected T function this
// method also mapping status code, for example informational response (100 - 199),
// successful response (200 - 299) and redirection message (300 - 399) will not return error
// but for client error response (400 - 499) and server error response this method will
// return errors.
func JsonResponse[T any](response *http.Response) (T, error) {
	var sample T
	defer response.Body.Close()

	// read response body and transform it into []byte to unmarshall
	// with generic type
	rsp, err := io.ReadAll(response.Body)
	if err != nil {
		return sample, err
	}

	// unmarshalling json
	err = json.Unmarshal(rsp, &sample)
	if err != nil {
		return sample, err
	}

	return sample, nil
}
