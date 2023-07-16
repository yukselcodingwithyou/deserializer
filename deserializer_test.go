package deserializer

import (
	"io/ioutil"
	"log"
	"testing"
)

type Object struct {
	Name    string `yaml:"name" json:"name"`
	Surname string `yaml:"surname" json:"surname"`
}

func Test_jsonDeserializer_Deserialize(t *testing.T) {
	type args struct {
		data  []byte
		value interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "deserializes into object",
			args: args{
				data:  readFile("test.json"),
				value: &Object{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := jsonDeserializer{}
			if err := j.Deserialize(tt.args.data, tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("Deserialize() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_yamlDeserializer_Deserialize(t *testing.T) {
	type args struct {
		data  []byte
		value interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "deserializes into object",
			args: args{
				data:  readFile("test.yml"),
				value: &Object{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			y := yamlDeserializer{}
			if err := y.Deserialize(tt.args.data, tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("Deserialize() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func readFile(fileName string) []byte {
	content, err := ioutil.ReadFile(fileName)

	if err != nil {
		log.Fatal(err)
	}
	return content
}
