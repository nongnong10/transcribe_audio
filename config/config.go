package config

import (
	"encoding/json"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
	"os"
	"reflect"
	"strings"
)

type Config struct {
	Files  Files  `mapstructure:"files"`
	APIKey string `mapstructure:"api_key"`
}

type Files struct {
	AudioFiles []string `mapstructure:"audio_files"`
	VideoFiles []string `mapstructure:"video_files"`
	TextFiles  []string `mapstructure:"text_files"`
}

const (
	LocalProjectPath = "/home/huy/pipe/transcribe_and_detect_speech"
	ConstConfig      = "config"
	Yml              = "yml"
	RootPath         = "."
	Docker           = "docker"
)

func Load() Config {
	os.Chdir(LocalProjectPath)

	// Load config from config.yml
	// With common function in carbon/common
	vip := viper.New()
	vip.SetConfigName(ConstConfig)
	vip.SetConfigType(Yml)
	vip.AddConfigPath(RootPath) // ROOT

	vip.SetEnvPrefix(Docker)
	vip.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	vip.AutomaticEnv()

	err := vip.ReadInConfig()
	if err != nil {
		panic(err)
	}

	// Workaround https://github.com/spf13/viper/issues/188#issuecomment-399518663
	// to allow read from environment variables when Unmarshal
	for _, key := range vip.AllKeys() {
		var (
			js     interface{}
			val    = vip.Get(key)
			valStr = fmt.Sprintf("%v", val)
		)

		err := json.Unmarshal([]byte(valStr), &js)

		if err != nil {
			vip.Set(key, val)
		} else {
			vip.Set(key, js)
		}
	}

	fmt.Printf("===== Config file used: %+v \n", vip.ConfigFileUsed())

	cfg := Config{}
	err = vip.Unmarshal(&cfg, func(dc *mapstructure.DecoderConfig) {
		dc.DecodeHook = mapstructure.ComposeDecodeHookFunc(
			StringToStructHookFunc(),
			StringToSliceWithBracketHookFunc(),
			dc.DecodeHook)
	})
	if err != nil {
		panic(err)
	}

	return cfg
}

func StringToStructHookFunc() mapstructure.DecodeHookFunc {
	return func(
		f reflect.Type,
		t reflect.Type,
		data interface{},
	) (interface{}, error) {
		if f.Kind() != reflect.String ||
			(t.Kind() != reflect.Struct && !(t.Kind() == reflect.Pointer && t.Elem().Kind() == reflect.Struct)) {
			return data, nil
		}
		raw := data.(string)
		var val reflect.Value
		// Struct or the pointer to a struct
		if t.Kind() == reflect.Struct {
			val = reflect.New(t)
		} else {
			val = reflect.New(t.Elem())
		}

		if raw == "" {
			return val, nil
		}
		err := json.Unmarshal([]byte(raw), val.Interface())
		if err != nil {
			return data, nil
		}
		return val.Interface(), nil
	}
}

// DecodeHookFunc for converting string to struct
// Support load env json array to array of struct
func StringToSliceWithBracketHookFunc() mapstructure.DecodeHookFunc {
	return func(
		f reflect.Kind,
		t reflect.Kind,
		data interface{}) (interface{}, error) {
		if f != reflect.String || t != reflect.Slice {
			return data, nil
		}

		raw := data.(string)
		if raw == "" {
			return []string{}, nil
		}
		var slice []json.RawMessage
		err := json.Unmarshal([]byte(raw), &slice)
		if err != nil {
			return data, nil
		}

		var strSlice []string
		for _, v := range slice {
			strSlice = append(strSlice, string(v))
		}
		return strSlice, nil
	}
}
