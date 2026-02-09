package main

import(
	"errors"
	"fmt"
	"time"
	"context"
	"github.com/Asheehan77/Bootdev_BlogAggregator/internal/database"
	"github.com/google/uuid"
)


type Command struct {
	Name	string
	Args	[]string
}

type Commands struct {
	commandList		map[string]func(*State,Command) error
}

func handlerLogin(s *State, cmd Command) error {
	if len(cmd.Args) == 0 {
		return errors.New("No arguments for login command")
	}
	u, err := s.db.GetUser(context.Background(),cmd.Args[0])
	if err != nil {
		return errors.New("That username doesn't exist")
	}
	s.cfg.SetUser(cmd.Args[0])
	fmt.Printf("User has been set to: %s\n",u.Name)
	return nil
}

func handlerRegister(s *State, cmd Command) error {
	if len(cmd.Args) == 0 {
		return errors.New("No arguments for register command")
	}
	user := database.CreateUserParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name: cmd.Args[0],
	}
	u, err := s.db.GetUser(context.Background(),cmd.Args[0])
	if err == nil {
		return errors.New("That username already exists")
	}
	
	u, err = s.db.CreateUser(context.Background(),user)
	if err != nil {
		return err
	}
	s.cfg.CurrentUserName = cmd.Args[0]
	fmt.Println("User Created: ",u.Name)
	s.cfg.SetUser(u.Name)
	fmt.Printf("User has been set to: %s\n",u.Name)
	return nil
}

func handlerReset(s *State,cmd Command) error{
	err := s.db.Reset(context.Background())
	if err != nil {
		return err
	}
	fmt.Println("Database Successfully Reset")
	return nil
}

func handlerUsers(s *State,cmd Command) error{
	ulist, err := s.db.GetUsers(context.Background())
	if err != nil {
		return err
	}

	for _,u := range ulist {
		curr := ""
		if u.Name == s.cfg.CurrentUserName{
			curr = " (current)"
		}
		fmt.Printf("* %s%s\n",u.Name,curr)
	}

	return nil
}

func (c *Commands) Run(s *State,cmd Command) error{
	if com,exists := c.commandList[cmd.Name]; exists{
		err := com(s,cmd)
		if err != nil{
			return err
		}
	}
	return nil
}

func (c *Commands) Register(name string,f func(*State, Command) error){
	c.commandList[name] = f
}