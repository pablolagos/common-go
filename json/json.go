package json

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
)

func LoadJsonFomFile(filename string) (result interface{}, err error) {
	/* Leer rawdata */
	rawJsonData, err := ioutil.ReadFile(filename)
	if err != nil {
		return result, err
	}

	/* Validaciones */
	if !json.Valid(rawJsonData) {
		return result, errors.New("No valid json data")
	}

	/* Obtener datos */
	err = json.Unmarshal(rawJsonData, &result)
	return result, err
}

func SaveJson(object interface{}, filename string, perm os.FileMode) error {

	/* Convertir a texto */
	rawdata, err := json.Marshal(object)
	if err != nil {
		return err
	}

	/* Write to file */
	err = ioutil.WriteFile(filename, rawdata, perm)
	if err != nil {
		return nil
	}
	return nil
}
