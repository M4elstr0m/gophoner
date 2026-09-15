package modules

import "encoding/json"

func (m Module) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.String())
}
