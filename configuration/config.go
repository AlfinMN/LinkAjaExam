package configuration

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/spf13/viper"
)

func ReadEnv(key, defVal string) string {
	viper.SetConfigFile("./config.env")
	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Keyname : %v, not found, default key value : %v, has been loaded", key, defVal)
		return defVal
	}
	value, ok := viper.Get(key).(string)
	if !ok {
		log.Fatalf("invalid type assertion")
	}
	return value
}

func Connection() (*sql.DB, error, string, string) {

	dbUser := ReadEnv("DbUser", "root")
	dbPass := ReadEnv("DbPass", "Apinchocs98")
	dbHost := ReadEnv("DbHost", "localhost")
	dbPort := ReadEnv("DbPort", "3306")
	dbName := ReadEnv("DbName", "linkaja")
	serverPort := ReadEnv("ServerPort", "5555")
	serverHost := ReadEnv("ServerHost", "localhost")

	loadData := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	db, _ := sql.Open("mysql", loadData)

	err := db.Ping()
	if err != nil {
		log.Print(err)

	} else {
		fmt.Println("db connect")
	}

	return db, err, serverHost, serverPort
}
