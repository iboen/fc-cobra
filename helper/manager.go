package helper

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

func AddKeyValuePair(key string, value string) {
	if findExistingKey(key) {
		fmt.Println("Key already exists")
		return
	}
	writeKeyValuePair(key, value)
	fmt.Println("Key added")
}

func UpdateKeyValuePair(key string, value string) {
	if !findExistingKey(key) {
		fmt.Println("Key does not exist")
		return
	}
	writeKeyValuePair(key, value)
	fmt.Println("Key updated")
}

func DeleteKeyValuePair(key string) {
	if !findExistingKey(key) {
		fmt.Println("Key does not exist")
		return
	}
	DeleteKeyHack(key)
	fmt.Println("Key deleted")
}

func DeleteKeyHack(key string) {
	settings := viper.AllSettings()
	delete(settings, key)

	var parsedSettings string
	for key, value := range settings {
		parsedSettings = fmt.Sprintf("%s\n%s: %s", parsedSettings, key, value)
	}

	d1 := []byte(parsedSettings)
	HandleError(os.WriteFile(viper.ConfigFileUsed(), d1, 0644))
}

func writeKeyValuePair(key string, value interface{}) {
	viper.Set(key, value)
	err := viper.WriteConfig()
	HandleError(err)
}

func findExistingKey(searchKey string) bool {
	return viper.IsSet(searchKey)
}

func HandleError(e error) {
	if e != nil {
		fmt.Println(e)
	}
}
