package pkg

import "encoding/json"

func JSONEncoder(body []byte, dto interface{}) error {
	return json.Unmarshal(body, &dto)
}

func JSONDecoder(body interface{}) ([]byte, error) {
	return json.Marshal(body)
}
