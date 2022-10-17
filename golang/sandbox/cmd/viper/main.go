package main

import (
	"bytes"
	"embed"
	_ "embed"
	"fmt"
	"io/fs"
	"log"

	"github.com/spf13/viper"
)

//go:embed conf.yaml
var conf1 []byte

//go:embed conf.yaml
var conf2 embed.FS

func example1(conf []byte) {
	viper.SetConfigType("yaml") // or viper.SetConfigType("YAML")

	viper.ReadConfig(bytes.NewBuffer(conf))

	fmt.Println(viper.Get("hoge"))
	fmt.Println(viper.GetStringMap("clothing"))
}

func example2(f fs.ReadFileFS) {
	viper.SetConfigType("yaml") // or viper.SetConfigType("YAML")

	file, err := f.Open("conf.yaml")
	if err != nil {
		log.Fatal(err)
	}
	if err := viper.ReadConfig(file); err != nil {
		log.Fatal(err)
	}

	fmt.Println(viper.Get("hoge"))
	fmt.Println(viper.GetStringMap("clothing"))
}
func main() {
	example2(conf2)
}
