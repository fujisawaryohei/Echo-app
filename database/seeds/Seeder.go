package seeds

import (
	"flag"
	"log"
	"reflect"

	"gorm.io/gorm"
)

type Seed struct {
	dbConn *gorm.DB
}

func SeedHandler(db *gorm.DB) []string {
	flag.Parse()
	args := flag.Args()

	if len(args) >= 1 {
		switch args[0] {
		case "seed":
			Execute(db, args[1])
		}
	}
	return args
}

func Execute(db *gorm.DB, seedMethodNames ...string) {
	s := Seed{db}
	seedType := reflect.TypeOf(s)

	if len(seedMethodNames) == 0 {
		log.Println("Running all seeder...")

		for i := 0; i < seedType.NumMethod(); i++ {
			method := seedType.Method(i)
			seed(s, method.Name)
		}
	}

	for _, item := range seedMethodNames {
		seed(s, item)
	}
}

func seed(s Seed, seedMethodName string) {
	m := reflect.ValueOf(s).MethodByName(seedMethodName)

	if !m.IsValid() {
		log.Fatal("No method called ", seedMethodName)
	}

	log.Println("Seeding", seedMethodName, "...")
	m.Call(nil)

	log.Println("Seed", seedMethodName, "succedd")
}
