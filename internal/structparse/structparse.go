// a simplified version of [caarlos0/env](https://github.com/caarlos0/env/blob/main/env.go)
// - no config allowed, default to ignore missing keys
// - only support parsing into a pointer to struct with string, int, float64, bool fields

package structparse

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

type Config struct {
	TagKey       string
	LookupFn     func(key string) (string, bool)
	AllowMissing bool
}

// will return error on missing values
func Parse(
	v any,
	config *Config,
) error {
	if config == nil {
		config = &Config{}
	}

	if reflect.ValueOf(v).Kind() != reflect.Pointer ||
		reflect.ValueOf(v).Elem().Kind() != reflect.Struct {
		return fmt.Errorf("Parse: expect v to be a pointer to struct")
	}

	refType := reflect.TypeOf(v).Elem()
	refVal := reflect.ValueOf(v).Elem()

	var errs []error
	for i := 0; i < refType.NumField(); i++ {
		refField := refType.Field(i)

		// parse tag
		tagVal, ok := refField.Tag.Lookup(config.TagKey)
		if !ok {
			continue
		}
		// parse using the tag's value
		valParsed, ok := config.LookupFn(tagVal)
		if !ok {
			if !config.AllowMissing {
				errs = append(errs, fmt.Errorf("Parse: value not found for tag %s", refField.Tag))
			}
			continue
		}
		// cast the parsed value
		switch refType.Field(i).Type.Kind() {
		case reflect.String:
			refVal.Field(i).SetString(valParsed)
		case reflect.Bool:
			refVal.Field(i).SetBool(valParsed == "1")
		case reflect.Int:
			if val, err := strconv.ParseInt(valParsed, 10, 64); err != nil {
				errs = append(errs, fmt.Errorf("Parse: error when parsing tag %s: %w", refField.Tag, err))
			} else {
				refVal.Field(i).SetInt(val)
			}
		case reflect.Float64:
			if val, err := strconv.ParseFloat(valParsed, 64); err != nil {
				errs = append(errs, fmt.Errorf("Parse: error when parsing tag %s: %w", refField.Tag, err))
			} else {
				refVal.Field(i).SetFloat(val)
			}
		default:
			errs = append(errs, fmt.Errorf("Parse: unsupported type %s for tag %s", refType.Field(i).Type.Kind(), refField.Tag))
		}
	}

	if len(errs) != 0 {
		return errors.Join(errs...)
	} else {
		return nil
	}
}
