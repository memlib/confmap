package confmap

import (
	"reflect"

	"github.com/spf13/viper"
)

type HasDefaultEnvs interface {
	DefaultEnvs() map[string]any
}

func MapEnvsTyped[T any](prefix string) (*T, error) {
	configObj := new(T)
	if err := MapEnvs(prefix, configObj); err != nil {
		return nil, err
	}
	return configObj, nil
}

// MapEnvs struct fields must contain tag mapstructure
func MapEnvs(prefix string, objPointer any) error {
	v := viper.New()
	v.SetEnvPrefix(prefix)
	v.AutomaticEnv()
	rt := reflect.Indirect(reflect.ValueOf(objPointer)).Type()

	defaults := make(map[string]any)
	if def, ok := objPointer.(HasDefaultEnvs); ok {
		defaults = def.DefaultEnvs()
	}

	for i := 0; i < rt.NumField(); i++ {
		tagVal := rt.Field(i).Tag.Get("mapstructure")
		if tagVal == "" {
			continue
		}

		err := v.BindEnv(tagVal)
		if err != nil {
			return err
		}

		if defaultVal, hasDefault := defaults[tagVal]; hasDefault {
			v.SetDefault(tagVal, defaultVal)
		}
	}
	return v.Unmarshal(objPointer)
}
