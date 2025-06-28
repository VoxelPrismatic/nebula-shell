package environ

import (
	"fmt"
	"os"
	"strings"
)

type Environment map[string]string

func Env() Environment {
	ret := Environment{}
	for _, line := range os.Environ() {
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 1 {
			ret[parts[0]] = ""
		} else {
			ret[parts[0]] = parts[1]
		}
	}
	return ret
}

func (env Environment) Get(name string) (string, error) {
	val, ok := env[name]
	if !ok || val == "" {
		return "", fmt.Errorf("environment variable $%s not set", name)
	}
	return val, nil
}

func (env Environment) GetList(name string) ([]string, error) {
	return env.GetArray(name, ":")
}

func (env Environment) GetArray(name, sep string) ([]string, error) {
	val, err := env.Get(name)
	if err != nil {
		return []string{val}, err
	}

	return strings.Split(val, sep), nil

}
