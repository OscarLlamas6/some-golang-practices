package main

import (
	"fmt"

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

	fmt.Println(vp.GetString("Oscar"))
	vp.Set("Vanessa", "Polanco")
	vp.WriteConfig()
}
