# GATOR CLI

This is a CLI tool to fetch RSS feeds via url and store them in a database using user login (passwordless just for educational purposes).

## REQUIREMENTS
- Postgresql
- "gator" database created inside your database with access for postgres user
- Go

## How to install

Using **go install**:
```
go install github.com/rowsedgy/gator@latest
```

After that we need to create the file ~/.gatorconfig.json and put the Postgres database connection string inside along with a starting user, e.g:

```
{"db_url":"postgres://postgres:@localhost:5432/gator?sslmode=disable","current_user_name":"holgith"}
```

## Usage

- Create another user:

    `gator register <user>`

- Login (set that user as current user):
    
    `gator login <user>`

- Reset users (clears all registered users). This will delete everything in the database since all the data is dependent on the user that created it.

    `gator reset`

- List users

    `gator users`

- Add a feed for the current user. Also follows the feed after adding it.

    `gator addfeed <feed name> <feed url>`

- List all feeds

    `gator feeds`

- Follow a feed for current user. The feed must exist in the database.

    `gator follow <feed url>`

- Unfollow a feed for the current user

    `gator unfollow <feed url>`

- Lists followed feeds for current user

    `gator following`

- Browse posts for followed feeds for current user with an optional limit of posts (default is 2)

    `gator browse [limit]`

- Run aggregator every X duration. This will update the feeds, starting with the oldest ones first. **WARNING** Don't DOS the sites by using a very small duration.

    `gator agg 1m0s`

