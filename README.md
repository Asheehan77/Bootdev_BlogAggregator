Go Blog Aggregator Project from Boot.dev

Requirements:

1.  Go  
2.  Postgres v15 or later  
3.  Install Goose: go install github.com/pressly/goose/v3/cmd/goose@latest  
    From sql/schema directory run: goose postgres "postgres://postgres:postgres@localhost:5432/gator" up  
4.  Install SQLC: go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest  
    Run: sqlc generate  
    This will generate our sql models  

Setup:
Create a config file in the home directory: ~/.gatorconfig.json
```json
{
  "db_url": "protocol://username:password@host:port/database?sslmode=disable",
  "current_user_name": "username_goes_here"
}
```

Install:
Run: `go install` Bootdev_BlogAggregator
or
Run: `go build -o gator` and use gator

Usage:
1. Create an account  
> ./gator register *username*
2. If you have multiple accounts you can login  
> ./gator login *username*
3. List all users  
> ./gator users
4. Add a feed(you will automatically follow it  
> ./gator addfeed *name url*
5. List all feeds  
> ./gator feeds
6. Collect the feeds into your database  
> ./gator agg
7. Follow a feed  
> ./gator follow *url*
8. List your followed feeds  
> ./gator following
9. Unfollow a feed  
> ./gator unfollow *url*
10. Display the most recent posts from your followed feeds up to a chosen amount  
> ./gator browse *amount*
11. Clear your database  
> ./gator reset
