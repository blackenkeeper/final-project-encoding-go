package encoding

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/Yandex-Practicum/final-project-encoding-go/models"
	"gopkg.in/yaml.v3"
)

// JSONData тип для перекодирования из JSON в YAML
type JSONData struct {
	DockerCompose *models.DockerCompose
	FileInput     string
	FileOutput    string
}

// YAMLData тип для перекодирования из YAML в JSON
type YAMLData struct {
	DockerCompose *models.DockerCompose
	FileInput     string
	FileOutput    string
}

// MyEncoder интерфейс для структур YAMLData и JSONData
type MyEncoder interface {
	Encoding() error
}

// Encoding перекодирует файл из JSON в YAML
func (j *JSONData) Encoding() error {
	jsonFileContent, err := os.ReadFile(j.FileInput)
	if err != nil {
		fmt.Printf("Ошибка чтения файла: %s", err.Error())
	}

	err = json.Unmarshal(jsonFileContent, &j.DockerCompose)
	if err != nil {
		fmt.Printf("Ошибка десериализации: %s", err.Error())
	}

	fileInputDataConvertInYaml, err := yaml.Marshal(&j.DockerCompose)
	if err != nil {
		fmt.Printf("Ошибка при сериализации в yaml: %s", err.Error())
	}

	err = os.WriteFile(j.FileOutput, fileInputDataConvertInYaml, 0644)
	if err != nil {
		fmt.Printf("Ошибка записи сериализованных данных в файл: %s", err.Error())
	}

	return nil
}

// Encoding перекодирует файл из YAML в JSON
func (y *YAMLData) Encoding() error {
	yamlFileContent, err := os.ReadFile(y.FileInput)
	if err != nil {
		fmt.Printf("Ошибка чтения файла: %s", err.Error())
	}

	err = yaml.Unmarshal(yamlFileContent, &y.DockerCompose)
	if err != nil {
		fmt.Printf("Ошибка десериализации: %s", err.Error())
	}

	fileInputDataConvertInJson, err := json.Marshal(&y.DockerCompose)
	if err != nil {
		fmt.Printf("Ошибка при сериализации в yaml: %s", err.Error())
	}

	err = os.WriteFile(y.FileOutput, fileInputDataConvertInJson, 0644)
	if err != nil {
		fmt.Printf("Ошибка записи сериализованных данных в файл: %s", err.Error())
	}

	return nil
}
