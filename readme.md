### Install

```
go get -u github.com/andiksetyawan/config
```

### Example:
```
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/andiksetyawan/config"
)

type Config struct {
	DbName string `env:"DB_NAME"`
	DbHost string `env:"DB_HOST"`
}

func main() {
	os.Setenv("USER_SVC_DB_NAME", "user")
	os.Setenv("USER_SVC_DB_HOST", "localhost")

	cfg := Config{}
	if err := config.New(config.WithPrefix("USER_SVC_")).Load(&cfg); err != nil {
		log.Fatal(fmt.Errorf("failed to load env: %w", err))
	}

	//output: {DbName:user DbHost:localhost}
	fmt.Printf("%+v", cfg)
}
```
#### Load from .env file:
The default environment file name is .env, or you can customize the path of the environment file yourself by adding the option WithEnvPath("./folder/.env")
```
cfg := struct {
    DbName string `env:"DB_NAME"`
    DbHost string `env:"DB_HOST"`
}{}
if err := config.New(config.WithEnvPath("./examples/.env")).Load(&cfg); err != nil {
    log.Fatal(fmt.Errorf("failed to load env: %w", err))
}

//output: {DbName:user DbHost:example.com}
fmt.Printf("%+v", cfg)
```
