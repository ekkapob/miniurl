package v1

import (
	"miniurl/api/v1/handlers"
	"os"

	"github.com/go-pg/pg/v10"
	"github.com/gorilla/mux"
)

func NewRouter(r *mux.Router) {
	v1 := r.PathPrefix("/v1").Subrouter()

	opt, err := pg.ParseURL(os.Getenv("POSTGRES_URL"))
	if err != nil {
		panic(err)
	}
	db := pg.Connect(opt)

	ctx := handlers.Context{
		DB: db,
	}

	// var counter int
	// _, err = db.QueryOneContext(
	// 	context.Background(),
	// 	pg.Scan(&counter),
	// 	`SELECT nextval('url_counter')`,
	// )
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("sss")
	// fmt.Println(counter)

	v1.HandleFunc("/urls", ctx.CreateURL).Methods("POST")
	v1.HandleFunc("/urls/{shortURL}", ctx.GetURL).Methods("GET")
}
