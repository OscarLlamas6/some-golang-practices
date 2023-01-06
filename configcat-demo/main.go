package main

import (
	"fmt"

	configcat "github.com/configcat/go-sdk/v7"
	"github.com/spf13/viper"
)

func main() {

	vp := viper.New()

	vp.SetConfigName("config")
	vp.SetConfigType("json")
	vp.AddConfigPath(".")

	err := vp.ReadInConfig()
	if err != nil {
		fmt.Printf("Fatal error config file: %s \n", err)
	}

	CONFIG_CAT_ID := vp.GetString("ConfigCat_ID")
	USER_ID := vp.GetString("USER_ID")

	fmt.Println(USER_ID)

	user := &configcat.UserData{Identifier: USER_ID}

	client := configcat.NewClient(CONFIG_CAT_ID)

	isMyFirstFeatureEnabled := client.GetBoolValue("isMyFirstFeatureEnabled", false, user)

	if isMyFirstFeatureEnabled {
		fmt.Println("Feature activado desde ConfigCat :D")
	} else {
		fmt.Println("Feature desactivado desde ConfigCat :(")
	}
}
