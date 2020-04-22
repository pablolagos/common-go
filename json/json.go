package json

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
)

/*
 * Read a json file and unpack it to an object
 * Usage: LoadJsonFromFile(filename,&myObject)
 */
func LoadJsonFomFile(filename string, object interface{}) (err error) {
	/* Leer rawdata */
	rawJsonData, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	/* Validaciones */
	if !json.Valid(rawJsonData) {
		return errors.New("No valid json data")
	}

	/* Obtener datos */
	err = json.Unmarshal(rawJsonData, &object)
	return err
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
