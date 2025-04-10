package envconf

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

type parser struct {
	target interface{}
	set    map[string]string
}

func Load(filename string, target interface{}) error {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	p := parser{
		target: target,
		set:    make(map[string]string),
	}

	re := regexp.MustCompile("#.*")
	for _, line := range strings.Split(string(file), "\n") {
		line = re.ReplaceAllString(line, "")
		p.parseLine(line)
	}

	for _, line := range os.Environ() {
		line = re.ReplaceAllString(line, "")
		p.parseLine(line)
	}

	p.full("")

	return nil
}

func (me *parser) parseLine(line string) {
	line = strings.TrimSpace(line)
	p := strings.Index(line, "=")
	if p > 0 {
		me.set[strings.TrimSpace(line[0:p])] = strings.TrimSpace(line[p+1:])
	}
}

func (me *parser) full(prefix string) {
	v := reflect.ValueOf(me.target)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	t := v.Type()

	for i := 0; i < t.NumField(); i += 1 {
		f := t.Field(i)
		key := f.Tag.Get("cfg")
		if key == "" {
			key = f.Name
		}

		if prefix != "" {
			key = prefix + "_" + key
		}

		fv := v.Field(i)

		if fv.Kind() == reflect.Struct {
			subParser := &parser{
				target: fv.Addr().Interface(),
				set:    me.set,
			}
			subParser.full(key)
			continue
		}

		if s, ok := me.set[key]; ok {
			me.setValue(fv, key, s)
		} else {
			dft := f.Tag.Get("default")
			if dft != "" {
				me.setValue(fv, key, dft)
			}
		}
	}
}

func (me *parser) setValue(fv reflect.Value, key, s string) {
	switch fv.Type().Kind() {
	case reflect.String:
		fv.SetString(s)
	case reflect.Int, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int8:
		x, err := strconv.ParseInt(s, 10, 0)
		if err != nil {
			log.Printf("config error: %s required int, %s given\n", key, s)
		} else {
			fv.SetInt(x)
		}
	case reflect.Uint, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uint8:
		x, err := strconv.ParseUint(s, 10, 0)
		if err != nil {
			log.Printf("config error: %s required uint, %s given\n", key, s)
		} else {
			fv.SetUint(x)
		}
	case reflect.Bool:
		switch s {
		case "yes", "on":
			fv.SetBool(true)
		default:
			b, _ := strconv.ParseBool(s)
			fv.SetBool(b)
		}
	case reflect.Float32, reflect.Float64:
		x, err := strconv.ParseFloat(s, 0)
		if err != nil {
			log.Printf("config error: %s required float, %s given\n", key, s)
		} else {
			fv.SetFloat(x)
		}
	default:
		obj := reflect.New(fv.Type()).Interface()
		err := json.Unmarshal([]byte(s), obj)
		if err != nil {
			log.Printf("config error: %s required %s, JSON Error: %s\n", key, fv.Type().Kind(), err)
		} else {
			fv.Set(reflect.ValueOf(obj).Elem())
		}
	}
}
