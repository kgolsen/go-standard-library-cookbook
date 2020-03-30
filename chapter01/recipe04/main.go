package main

import (
	"log"
	"os"
)

func main() {
	// Get an env var (may be unset)
	dbConn := os.Getenv("DB_CONN")
	log.Printf("DB_CONN: %s\n", dbConn)

	// Get an env var (and check to see if it's empty)
	varKey := "FU_CONN"
	connStr, ok := os.LookupEnv(varKey)
	if !ok {
		log.Printf("%v is not set\n", varKey)
	} else {
		log.Printf("%v: %v\n", varKey, connStr)
	}

	// Get/set with defaults
	newKey := "MU_CONN"

	// Set
	os.Setenv(newKey, "db://user:pass@foo.com?bar=baz")
	val := GetEnvDefault(newKey, "db://lame:dude@localhost?bar=qux")
	log.Printf("The set value is '%v'\n", val)

	os.Unsetenv(newKey)

	val = GetEnvDefault(newKey, "db://lame:dude@localhost?bar=qux")
	log.Printf("The default value is '%v'\n", val)
}

func GetEnvDefault(key, defVal string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		return defVal
	}
	return val
}
