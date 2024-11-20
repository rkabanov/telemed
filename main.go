package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	_ "github.com/lib/pq"
	"github.com/rkabanov/service/app"
	"github.com/rkabanov/service/store"
	"github.com/rkabanov/service/store/memory"
	"github.com/rkabanov/service/store/postgres"
	"github.com/rkabanov/service/web"
)

var buildDate string

type programArgs struct {
	store   *string
	pghost  *string
	pgport  *int
	pguser  *string
	pgpass  *string
	pgdb    *string
	webpath *string
	webport *int
}

func main() {
	log.Println("start, build:", buildDate)

	var args programArgs
	args.store = flag.String("store", "memory", "service store type: memory or postgres")
	args.pghost = flag.String("pghost", "localhost", "postgres host")
	args.pgport = flag.Int("pgport", 5432, "postgres port")
	args.pguser = flag.String("pguser", "root", "postgres user")
	args.pgpass = flag.String("pgpass", "secret", "postgres password")
	args.pgdb = flag.String("pgdb", "service", "postgres DB")
	args.webpath = flag.String("webpath", "", "web app URL path")
	args.webport = flag.Int("webport", 8180, "web app port")
	//	var source = "postgresql://root:secret@localhost:5433/simple_bank?sslmode=disable"
	flag.Parse()
	fmt.Println("using store1", *args.store)
	fmt.Println("pghost:", *args.pghost)

	var store app.Store
	switch *args.store {
	case "memory":
		store = getMemoryStore()
	case "postgres":
		store = getPostgresStore(args)
	default:
		log.Fatalf("store type '%v' not supported, exit", args.store)
	}
	store.Print()

	app := app.NewApp(store) // Business logic - depends on common store (or on all individual stores?).

	web := web.NewWebAPI(app) // Web API for the app - depends on business logic.

	var sb strings.Builder
	if *args.webpath != "" {
		sb.WriteString("/")
		sb.WriteString(*args.webpath)
	}
	path := sb.String()
	fmt.Println("Listen on path:", path, ", port:", *args.webport)

	http.HandleFunc(path+"/patient", web.HandlePatient) // GET and POST
	http.HandleFunc(path+"/patients", web.GetPatients)
	http.HandleFunc(path+"/doctor", web.HandleDoctor) // GET and POST
	http.HandleFunc(path+"/doctors", web.GetDoctors)
	log.Println(http.ListenAndServe(":"+strconv.Itoa(*args.webport), nil))

	log.Println("finish")
}

func getMemoryStore() *memory.Store {
	return memory.NewStore([]store.DoctorRecord{
		{ID: "9001", Name: "Dr. Paul", Email: "paul@yopmail.com", Role: "radiologist", Speciality: "dermatology"},
		{ID: "9002", Name: "Dr. Smith", Email: "smith@yopmail.com", Role: "admin", Speciality: ""},
		{ID: "9003", Name: "Dr. Tucker", Email: "tucker@yopmail.com", Role: "nurse", Speciality: ""},
	},
		[]store.PatientRecord{
			{ID: "1001", Name: "Evelynn Lang", Age: 20, External: false},
			{ID: "1002", Name: "Wells Maldonado", Age: 21, External: true},
			{ID: "1003", Name: "Elaina Davis", Age: 22, External: false},
			{ID: "1004", Name: "Lucas Wright", Age: 23, External: true},
			{ID: "1005", Name: "Lily Contreras", Age: 24, External: false},
		})
}

func getPostgresStore(args programArgs) *postgres.Store {
	var driver = "postgres"
	//	var source = "postgresql://root:secret@localhost:5433/simple_bank?sslmode=disable"
	source := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=disable",
		*args.pguser,
		*args.pgpass,
		*args.pghost,
		*args.pgport,
		*args.pgdb,
	)
	log.Println("getPostgresStore: source:", source)

	db, err := sql.Open(driver, source)
	if err != nil {
		log.Fatalf("failed to open DB connection: %v", err)
	}
	err = db.Ping()

	if err != nil {
		log.Fatalf("failed to ping DB: %v", err)
	}

	return postgres.NewStore(db)
}
