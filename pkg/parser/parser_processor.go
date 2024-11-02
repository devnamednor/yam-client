package parser

import (
	"os"

	"github.com/goccy/go-yaml"
)

func parse(filePath string)(YamlParserConfig,error){
    f, err :=os.Open(filePath)

    if err!=nil {
        return nil,err
    }

    defer f.Close()

    decoder :=yaml.NewDecoder(f)

    var parserConfig YamlParserConfig
    err = decoder.Decode(&parserConfig)

    if err!=nil {
        eturn nil,err
    }
}