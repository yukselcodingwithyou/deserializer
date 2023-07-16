package deserializer

import (
	jsoniter "github.com/json-iterator/go"
	"gopkg.in/yaml.v3"
	"log"
)

// FileType is the supported file type in to deserialize
type FileType string

const (
	YAML FileType = "YAML"
	JSON FileType = "JSON"
)

// Deserializer deserializes file into an interface
type Deserializer interface {
	Deserialize(data []byte, value interface{}) error
}

// New creates a new deserializer
func New(fileType FileType) Deserializer {

	switch fileType {
	case YAML:
		return &yamlDeserializer{}
	case JSON:
		return &jsonDeserializer{}
	default:
		log.Fatalln("there is no such deserializer implementation")
	}
	return &jsonDeserializer{}
}

type jsonDeserializer struct{}

// Deserialize deserializes a json file into an interface
func (j jsonDeserializer) Deserialize(data []byte, value interface{}) error {
	err := jsoniter.Unmarshal(data, value)
	return err
}

type yamlDeserializer struct{}

// Deserialize deserializes a yml file into an interface
func (y yamlDeserializer) Deserialize(data []byte, value interface{}) error {
	err := yaml.Unmarshal(data, value)
	return err
}
