package google

import "encoding/json"

// Enum implements an enumeration of string values.
type Enum struct {
	Basis     map[string]int
	BasisKeys []string
	data      []int
}

// UnmarshallJSON overrides to enable json encoding.
func (e *Enum) UnmarshallJSON(data []byte) (err error) {
	var keys []string
	if err = json.Unmarshal(data, &keys); err != nil {
		return err
	}
	for _, key := range keys {
		if e.data == nil {
			e.data = []int{}
		}
		e.data = append(e.data, e.Basis[key])
	}

	return nil
}

// MarshallJSON overrides to enable json decoding.
func (e *Enum) MarshallJSON() (data []byte, err error) {
	keys := []string{}
	for _, i := range e.data {
		keys = append(keys, e.BasisKeys[i])
	}

	if data, err = json.Marshal(keys); err != nil {
		return nil, err
	}

	return data, nil
}
