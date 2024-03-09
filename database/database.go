package database

import(
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func DBSetup() (dbParams *sql.DB){
	err := godotenv.Load("config/.env")
	if err != nil {
		fmt.Println("Failed to load environmert file")
	} else {
		fmt.Println("Load environmert file success")
	}

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", 
		os.Getenv("PGHOST"), os.Getenv("PGPORT"), os.Getenv("PGUSER"), os.Getenv("PGPASSWORD"), os.Getenv("PGDATABASE"))

	DB, _ := sql.Open("postgres", psqlInfo)
	err = DB.Ping()
	if err != nil {
		fmt.Println("Database connection failed")
		panic(err)
		} else {
		fmt.Println("Database connection success")
	}
	
	return DB
}

// func DbMigrate(dbParams *sql.DB) {
// 	migrations := &migrate.PackrMigrationSource{
// 		Box: packr.New("migrations", "sql_migrations"),
// 	}

// 	n, errs := migrate.Exec(dbParams, "postgres", migrations, migrate.Up)
// 	if errs != nil {
// 		panic(errs)
// 	}

// 	DBConnection = dbParams

// 	fmt.Println("Applied", n, "Migrations!")
// }