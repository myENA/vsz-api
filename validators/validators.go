package validators

import (
	"github.com/satori/go.uuid"
	"net"
	"reflect"
	"strconv"
	"strings"
	"errors"
	"fmt"
)

func contains(haystack interface{}, needle interface{}) bool {
	v := reflect.ValueOf(haystack)

	if reflect.Slice == v.Kind() {
		l := v.Len()
		for i := 0; i < l; i++ {
			if v.Index(i).Interface() == needle {
				return true
			}
		}
	}

	return false
}

func StrNotEmpty(v string) error {
	if strings.TrimSpace(v) == "" {
		return errors.New("input string cannot be empty")
	}
	return nil
}

// InRange will return an error if the expected list of values does not contain the actual value provided
func InRange(allowed interface{}, actual interface{}) error {
	if contains(allowed, actual) {
		return nil
	}
	return fmt.Errorf("\"%#v\" is not in provided range", actual)
}

// StrIsInt will return an error if the input is not parseable as an int
func StrIsInt(v string) error {
	if err := StrNotEmpty(v); err != nil {
		return err
	} else if _, err = strconv.Atoi(v); err != nil {
		return fmt.Errorf("\"%s\" is not parseable as an int: %s", v, err)
	}
	return nil
}

// StrIsPositiveInt will return an error if the input is not both an integer and greater than or equal to 0
func StrIsPositiveInt(v string) error {
	if err := StrNotEmpty(v); err != nil {
		return err
	} else if i, err := strconv.Atoi(v); err != nil {
		return fmt.Errorf("\"%s\" is not parseable as an int: %s", v, err)
	} else if i < 0 {
		return fmt.Errorf("\"%d\" is not >= 0", i)
	} else {
		return nil
	}
}

// StrIsIntInSet will return an error if the input is not an integer or if it is not in the provided list of expected ints
func StrIsIntInSet(allowable []int, v string) error {
	if err := StrNotEmpty(v); err != nil {
		return err
	} else if i, err := strconv.Atoi(v); err != nil {
		return fmt.Errorf("\"%s\" is not parseable as an int: %s", v, err)
	} else if !contains(allowable, i) {
		return fmt.Errorf("\"%d\" is not in provided range", i)
	} else {
		return nil
	}
}

func StrIsUUIDv4(v string) error {
	if err := StrNotEmpty(v); err != nil {
		return err
	} else if _, err = uuid.FromString(v); err != nil{
		return fmt.Errorf("\"%s\" is not parseable as an UUIDv4: %s", v, err)
	} else {
		return nil
	}
}

func StrIsMAC(v string) error {
	if err := StrNotEmpty(v); err != nil {
		return err
	} else if _, err = net.ParseMAC(v); err != nil {
		return fmt.Errorf("\"%s\" is not parseable as a MAC: %s", v, err)
	} else {
		return nil
	}
}

func StrIsBool(v string) error {
	if err := StrNotEmpty(v); err != nil {
		return err
	} else if _, err = strconv.ParseBool(v); err != nil {
		return fmt.Errorf("\"%s\" is not parseable as a bool: %s", v, err)
	} else {
		return nil
	}
}

func StrIsIP(v string) error {
	if err := StrNotEmpty(v); err != nil {
		return err
	} else if ip := net.ParseIP(v); ip == nil {
		return fmt.Errorf("\"%s\" is not parseable as an IP Address: %s", v, err)
	} else {
		return nil
	}
}
