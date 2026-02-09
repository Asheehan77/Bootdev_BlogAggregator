package main

import(
	"fmt"
	"os"
	"github.com/Asheehan77/Bootdev_BlogAggregator/internal/config"
	"github.com/Asheehan77/Bootdev_BlogAggregator/internal/database"
	"database/sql"
	_ "github.com/lib/pq"
)

type State struct {
	cfg		*config.Config
	db		*database.Queries
		
}

func main(){

	fcfg,err := config.Read()
	if err != nil {
		fmt.Printf("Error:%s\n",err)
		return
	}
	
	db, err := sql.Open("postgres", fcfg.DB_Url)
	if err != nil {
		fmt.Printf("Error:%s\n",err)
		return
	}
	quer := database.New(db)

	s := &State{
		cfg: &fcfg,
		db:  quer,
	}

	if len(os.Args) < 2 {
		fmt.Println("Error: Not enough arguments")
		os.Exit(1)
	}
	cmd := Command{
		Name: os.Args[1],
		Args: os.Args[2:],
	}

	list := Commands{
		commandList: map[string]func(*State,Command) error{},
	}

	list.Register("login",handlerLogin)
	list.Register("register",handlerRegister)
	list.Register("reset",handlerReset)
	list.Register("users",handlerUsers)
	err = list.Run(s,cmd)
	if err != nil {
		fmt.Println("Error:",err)
		os.Exit(1)
	}

	return
}