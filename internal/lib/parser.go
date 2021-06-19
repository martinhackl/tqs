package lib

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/go-playground/validator"
)

type Session struct {
	Version string            `json:"version" validate:"required"`
	Name    string            `json:"name" validate:"required"`
	Env     map[string]string `json:"env,omitempty"`
	Windows []Window          `json:"windows" validate:"required,min=1"`
}

type Window struct {
	Name string `json:"name" validate:"required"`
	Path string `json:"path" validate:"required"`
	Cmd  string `json:"cmd,omitempty" validate:"isdefault"`
}

func ParseJSONFile(filePath string) (*Session, error) {
	jsonFile, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var session Session
	json.Unmarshal(byteValue, &session)

	validate := validator.New()
	if err := validate.Struct(session); err != nil {
		return nil, err
	}

	return &session, nil
}

func Substitute(session *Session, variables map[string]string) {
	for k, v := range variables {
		for i := 0; i < len(session.Windows); i++ {
			session.Windows[i].Path = strings.Replace(session.Windows[i].Path, fmt.Sprintf("{%s}", k), v, -1)
		}
	}
}
