package turso

// See https://docs.turso.tech/sdk/go/quickstart#local-embedded-replicas

// var Db *sql.DB

// func Init() *sql.DB {
// 	dbName := "local.db"
// 	TURSO_DB_URL := os.Getenv("TURSO_DB_URL")
// 	TURSO_DB_TOKEN := os.Getenv("TURSO_DB_TOKEN")

// 	dir, err := os.MkdirTemp("", "libsql-*")
// 	if err != nil {
// 		fmt.Println("Error creating temporary directory:", err)
// 		os.Exit(1)
// 	}
// 	// defer os.RemoveAll(dir)

// 	dbPath := filepath.Join(dir, dbName)

// 	syncInterval := time.Minute

// 	connector, err := libsql.NewEmbeddedReplicaConnector(dbPath, TURSO_DB_URL,
// 		libsql.WithAuthToken(TURSO_DB_TOKEN),
// 		libsql.WithSyncInterval(syncInterval),
// 	)

// 	if err := connector.Sync(); err != nil {
// 		fmt.Println("Error syncing database:", err)
// 	}

// 	if err != nil {
// 		fmt.Println("Error creating connector:", err)
// 		os.Exit(1)
// 	}
// 	// defer connector.Close()

// 	Db = sql.OpenDB(connector)
// 	// defer db.Close()

// 	return Db
// }

// type User struct {
// 	ID   int
// 	Name string
// }

// func GetUsers() {
// 	rows, err := Db.Query("SELECT * FROM users")
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "failed to execute query: %v\n", err)
// 		os.Exit(1)
// 	}
// 	defer rows.Close()

// 	var users []User

// 	for rows.Next() {
// 		var user User

// 		if err := rows.Scan(&user.ID, &user.Name); err != nil {
// 			fmt.Println("Error scanning row:", err)
// 			return
// 		}

// 		users = append(users, user)
// 		fmt.Println(user.ID, user.Name)
// 	}

// 	if err := rows.Err(); err != nil {
// 		fmt.Println("Error during rows iteration:", err)
// 	}
// }
