package main

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
)

// ini parser
type any = interface{}

//mysql struct
type MysqlConfig struct {
	Address  string `ini:"address"`
	Port     int    `ini:"port"`
	Username string `ini:"username"`
	Password string `ini:"password"`
}

// redis struct
type RedisConfig struct {
	Host     string `ini:"host"`
	Prot     int    `ini:"port"`
	Password string `ini:"password"`
	Database int    `ini:"datase"`
}

// config
type Config struct {
	MysqlConfig `ini:"mysql"`
	RedisConfig `ini:"redis"`
}

func loadIni(fileName string, data any) (err error) {
	// 0. data must be a pointer type and  it must be a struct pointer
	t := reflect.TypeOf(data)
	if t.Kind() != reflect.Ptr || t.Elem().Kind() != reflect.Struct {
		err = fmt.Errorf("data para should be a strunct pointer")
		return
	}

	// 1. read files
	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		return
	}
	// 2. line by line
	lineSlice := strings.Split(string(b), "\r\n")
	fmt.Printf("%#v\n", lineSlice)

	var structName string
	for idx, line := range lineSlice {
		line = strings.TrimSpace(line)
		// 	2.1 if a comment, ignore
		if strings.HasPrefix(line, "#") || strings.HasPrefix(line, ";") {
			fmt.Printf("This is a comment: [%s]", line)
			continue
		} else if strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]") {
			// 	2.2 if [] it's a section
			fmt.Println("this is a section")
			sectionName := line[1 : len(line)-1]
			// using relfect to find sectionname in the data
			for i := 0; i < t.Elem().NumField(); i++ {
				field := t.Elem().Field(i)
				if sectionName == field.Tag.Get("ini") {
					// found the section name. recoded the name
					structName = field.Name
					fmt.Printf("found [%s] struct name %s\n", sectionName, structName)
				}

			}
		} else if index := strings.Index(line, "="); index > 1 && index < len(line)-2 {
			// 2.3 else if key=value and find use structName to get pair from data

			v := reflect.ValueOf(data)
			structObj := v.Elem().FieldByName(structName)
			if structObj.Kind() != reflect.Struct {
				err = fmt.Errorf("data %s is a strcut", structName)
				return
			}

			key := strings.TrimSpace(line[:index])
			value := strings.TrimSpace(line[index+1:])
			// loop through the struct. if tag == key, then set the value

			var fieldName string
			for i := 0; i < structObj.NumField(); i++ {
				field := structObj.Type().Field(i) // tag is in the type of reflect
				if field.Tag.Get("ini") == key {
					// find the key of the key/value pair
					fieldName = field.Name
					break
				}
			}

			// assign the value
			if len(fieldName) > 0 {
				fieldObj := structObj.FieldByName(fieldName)
				switch fieldObj.Type().Kind() {
				case reflect.String:
					fieldObj.SetString(value)
				case reflect.Int, reflect.Int16, reflect.Int32, reflect.Int64:
					var valueInt int64
					valueInt, err = strconv.ParseInt(value, 10, 64)
					if err != nil {
						err = fmt.Errorf("line:%d value type error", idx+1)
						return
					}
					fieldObj.SetInt(valueInt)
				case reflect.Float32, reflect.Float64:
					var valuefloat float64
					valuefloat, err = strconv.ParseFloat(value, 64)
					if err != nil {
						err = fmt.Errorf("line:%d value type error", idx+1)
						return
					}
					fieldObj.SetFloat(valuefloat)
				case reflect.Bool:
					var valueBool bool
					valueBool, err = strconv.ParseBool(value)
					if err != nil {
						err = fmt.Errorf("line:%d value type error", idx+1)
						return
					}
					fieldObj.SetBool(valueBool)
				}
			}

		} else {
			continue
		}
	}
	return
}

func main() {
	var cfg Config
	err := loadIni("./conf.ini", &cfg)
	if err != nil {
		fmt.Printf("loadini failed, err:%v", err)
		return
	}
	// fmt.Println(cfg.MysqlConfig.Address, cfg.MysqlConfig.Port, cfg.MysqlConfig.Username, cfg.MysqlConfig.Password)
	// fmt.Println(cfg.RedisConfig.Host, cfg.RedisConfig.Prot, cfg.RedisConfig.Password, cfg.RedisConfig.Database)
	fmt.Printf("%#v\n", cfg)

}
