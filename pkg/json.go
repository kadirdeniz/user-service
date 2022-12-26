package pkg

import "encoding/json"

func JSONEncoder(body []byte, dto interface{}) error {
	encodeError := json.Unmarshal(body, &dto)
	if encodeError != nil {
		return encodeError
	}
	return nil
}

func JSONDecoder(body interface{}) ([]byte, error) {
	marshaledJSON, decodeError := json.Marshal(body)
	if decodeError != nil {
		return nil, decodeError
	}

	return marshaledJSON, nil
}
