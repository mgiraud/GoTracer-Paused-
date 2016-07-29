package raytracer

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"strings"
)

type ObjectMap map[string]Object

type Reader struct {
	Shapes ObjectMap `json:"shapes"`
	Camera Camera    `json:"camera"`
}

func (sm *ObjectMap) UnmarshalJSON(data []byte) error {
	shapes := make(map[string]json.RawMessage)
	err := json.Unmarshal(data, &shapes)
	if err != nil {
		return err
	}
	result := make(ObjectMap)
	for k, v := range shapes {
		switch true {
		case strings.Contains(k, "sphere"):
			s := &Sphere{}
			err := json.Unmarshal(v, &s)
			if err != nil {
				return err
			}
			result[k] = s
		default:
			return errors.New("Unrecognized shape")
		}
	}
	*sm = result
	return nil
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadScene(scene string) (*Camera, ObjectMap) {

	n1, err := ioutil.ReadFile(scene)
	check(err)

	r := Reader{}
	err2 := json.Unmarshal(n1, &r)
	if err2 != nil {
		fmt.Println(err)
		return nil, nil
	}
	return &r.Camera, r.Shapes
}
