package main

import(
	"errors"
	"fmt"
	"time"
	"context"
	"html"
	"github.com/Asheehan77/Bootdev_BlogAggregator/internal/database"
	"github.com/Asheehan77/Bootdev_BlogAggregator/internal/rss"
	"github.com/google/uuid"
)

const(
	default_feed = "https://www.wagslane.dev/index.xml"
)


type Command struct {
	Name	string
	Args	[]string
}

type Commands struct {
	commandList		map[string]func(*State,Command) error
}

func middlewareLoggedIn(handler func(s *State, cmd Command, user database.User) error) func(*State, Command) error {
    return func(s *State, cmd Command) error {
        user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
        if err != nil {
            return err
        }

        return handler(s, cmd, user)
    }
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

func handlerAgg(s *State,cmd Command) error{

	rssf,err := rss.FetchFeed(context.Background(),default_feed)
	if err != nil {
		return err
	}

	rssf.Channel.Title = html.UnescapeString(rssf.Channel.Title)
	rssf.Channel.Description = html.UnescapeString(rssf.Channel.Description)

	for i := range rssf.Channel.Item{
		rssf.Channel.Item[i].Title = html.UnescapeString(rssf.Channel.Item[i].Title)
		rssf.Channel.Item[i].Description = html.UnescapeString(rssf.Channel.Item[i].Description)
	}

	fmt.Println(rssf)
	return nil
}

func handlerAddFeed(s *State,cmd Command, user database.User) error{
	if len(cmd.Args) < 2 {
		return errors.New("Missing arguments for addfeed command. \nUsage: addfeed <name> <url>")
	}

	feed_id := uuid.New()
	feed := database.AddFeedParams{
		ID: feed_id,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name: cmd.Args[0],
		Url: cmd.Args[1],
		UserID: user.ID,
	}

	f, err := s.db.AddFeed(context.Background(),feed)
	if err != nil {
		return err
	}

	feed_f := database.CreateFeedFollowParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID: user.ID,
		FeedID: feed_id,
	}
	_,err = s.db.CreateFeedFollow(context.Background(),feed_f)
	if err != nil {
		return err
	}

	fmt.Println("Name: ",f.Name,"\nUrl: ",f.Url)
	return nil
}

func handlerGetFeeds(s *State,cmd Command) error{
	feeds,err := s.db.GetFeeds(context.Background())
	if err != nil {
		return err
	}
	fmt.Printf("Saved Feeds:\n")
	for _,f := range feeds {
		fmt.Printf(" - Name: %s\n   Url: %s\n   Saved By: %s\n",f.Name,f.Url,f.Name_2)
	}
	return nil
}

func handlerFollow(s *State,cmd Command, user database.User) error{
	if len(cmd.Args) < 1 {
		return errors.New("Missing arguments for follow command. \nUsage: follow <url>")
	}
	feed,err := s.db.FeedByUrl(context.Background(),cmd.Args[0])
	if err != nil {
		return err
	}

	follow := database.CreateFeedFollowParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID: user.ID,
		FeedID: feed.ID,
	}

	created_follow,err := s.db.CreateFeedFollow(context.Background(),follow)
	if err != nil {
		return err
	}

	fmt.Printf("%s Followed Feed: %s",created_follow.UserName,created_follow.FeedName)
	return nil
}

func handlerFollowing(s *State,cmd Command, user database.User) error{
	feeds, err := s.db.FeedFollowersForUser(context.Background(),user.Name)
	if err != nil {
		return nil
	}
	fmt.Printf("Current User %s Follows:\n",s.cfg.CurrentUserName)
	for _,feed := range feeds {
		fmt.Printf("- %s\n",feed.FeedName)
	}
	return nil
}

func (c *Commands) Run(s *State,cmd Command) error{
	if com,exists := c.commandList[cmd.Name]; exists{
		err := com(s,cmd)
		if err != nil{
			return err
		}
	}else{
		return errors.New("Unknown command")
	}
	return nil
}

func (c *Commands) Register(name string,f func(*State, Command) error){
	c.commandList[name] = f
}