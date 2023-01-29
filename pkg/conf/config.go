package conf

import (
	"encoding/json"
	"fmt"
	"goFrame/pkg/encoding"
	"os"
	"path"
	"strings"
)

var loaders = map[string]func([]byte, any) error{
	".json": LoadFromJsonBytes,
	".yaml": LoadFromYamlBytes,
	".yml":  LoadFromYamlBytes,
}

func LoadConfig(file string, v any) error {
	content, err := os.ReadFile(file)
	if err != nil {
		return err
	}

	loader, ok := loaders[strings.ToLower(path.Ext(file))]
	if !ok {
		return fmt.Errorf("unrecognized file type: %s", file)
	}

	return loader(content, v)
}

func LoadFromJsonBytes(content []byte, v any) error {
	return json.Unmarshal(content, v)
}

func LoadFromYamlBytes(content []byte, v any) error {
	b, err := encoding.YamlToJson(content)
	if err != nil {
		return err
	}

	return LoadFromJsonBytes(b, v)
}
