package config

import (
	"fmt"
	"reflect"
)

type Components int

const (
	SecretManager Components = iota
)

type SecretManagerConfig struct {
	Port     int
	Replicas int
}

type Config interface {
	Set()
	Get(string) (any, bool, error)
}

func New(c Components) Config {
	smc := &SecretManagerConfig{}

	switch c {
	case 0:
		return smc
	}
	return nil
}

func (smc *SecretManagerConfig) Set() {
	smc.Port = 5103
	smc.Replicas = 3
}

func (smc *SecretManagerConfig) Get(key string) (any, bool, error) {
	fmt.Println("START FINDING", key)
	v := reflect.ValueOf(smc)

	if v.Kind() != reflect.Pointer {
		return nil, false, fmt.Errorf("%s", "IS NOT POINTER TYPE")
	}

	v = v.Elem()

	sf, ok := v.Type().FieldByName(key)

	if !ok {
		return nil, false, nil
	}
	fmt.Println("ELEM:", v.FieldByName(sf.Name))
	return v.FieldByName(sf.Name), true, nil
}
