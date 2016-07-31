package raytracer

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"strings"
)

type ObjectMap map[string]Object

type LightMap map[string]Light

type Reader struct {
	Shapes ObjectMap `json:"shapes"`
	Camera Camera    `json:"camera"`
	Lights LightMap  `json:"lights"`
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
			break
		case strings.Contains(k, "plan"):
			s := &Plane{}
			err := json.Unmarshal(v, &s)
			if err != nil {
				return err
			}
			result[k] = s
			break
		default:
			return errors.New("Unrecognized shape")
		}
	}
	*sm = result
	return nil
}

func (sm *LightMap) UnmarshalJSON(data []byte) error {
	lights := make(map[string]json.RawMessage)
	err := json.Unmarshal(data, &lights)
	if err != nil {
		return err
	}
	result := make(LightMap)
	for k, v := range lights {
		switch true {
		case strings.Contains(k, "dist"):
			s := &DistantLight{}
			err := json.Unmarshal(v, &s)
			if err != nil {
				return err
			}
			result[k] = s
		case strings.Contains(k, "point"):
			s := &PointLight{}
			err := json.Unmarshal(v, &s)
			if err != nil {
				return err
			}
			result[k] = s
		default:
			return errors.New("Unrecognized light")
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

func ReadScene(scene string) (*Camera, ObjectMap, LightMap) {

	n1, err := ioutil.ReadFile(scene)
	check(err)

	r := Reader{}
	err2 := json.Unmarshal(n1, &r)
	if err2 != nil {
		fmt.Println(err2)
		return nil, nil, nil
	}
	return &r.Camera, r.Shapes, r.Lights
}
