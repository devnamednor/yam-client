package parser

type YamlRequest struct{
	Name string 	`yaml:"name"`
	Method string	`yaml:'method'`
	Url string		`yaml:"url"`                         
	Headers map[string]string `yaml:"headers,omitempty"`  
	ContentType string        `yaml:"contentType"`   
	RequestBody string		  `yaml:"body,omitempty"`  
}

type YamlParserConfig struct {
	Requests []YamlRequest  `yaml:"requests"`
}