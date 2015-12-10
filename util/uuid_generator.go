package util

import (
	"fmt"
	"github.com/docker/docker/pkg/random"
	"github.com/Sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var (
	names            	[]string
	adverbs            	[]string
	adjectives          []string
)

func init() {
	//read yaml file to load the words
	
	bytes, err := ioutil.ReadFile("./util/names.yml")
	if err != nil {
		logrus.Fatalf("Failed to read the names.yml: %v", err)
	}
	yaml.Unmarshal(bytes, &names)
	
	bytes, err = ioutil.ReadFile("./util/adverbs.yml")
	if err != nil {
		logrus.Fatalf("Failed to read the adverbs.yml: %v", err)
	}
	yaml.Unmarshal(bytes, &adverbs)
	
	bytes, err = ioutil.ReadFile("./util/adjectives.yml")
	if err != nil {
		logrus.Fatalf("Failed to read the adjectives.yml: %v", err)
	}
	yaml.Unmarshal(bytes, &adjectives)
}

// GenerateUUID generates a random name from the list of advebs, adjectives and names in this package
// formatted as "{{adverb}}-{{adjective}}-{{name}}". For example 'abnormally-abandoned-einstein'. 
func GenerateUUID() string {
	rnd := random.Rand
	uuid := fmt.Sprintf("%s-%s-%s", adverbs[rnd.Intn(len(adverbs))], adjectives[rnd.Intn(len(adjectives))], names[rnd.Intn(len(names))])
	return uuid
}