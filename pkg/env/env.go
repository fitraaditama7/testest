package env

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

func Load(prefix string, out interface{}, filename ...string) error {
	err := godotenv.Load(filename...)
	if err != nil {
		log.Println(err)
		return err
	}

	err = envconfig.Process(prefix, out)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
