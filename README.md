# Gator: RSS Feed Aggregator CLI

A RSS feed aggregator CLI tool implemented in Go with PostgreSQL backend. Gator allows you to manage RSS feeds, follow your favourite sources, and browse aggregated posts from the command line.

![gator](/gator.png)
*AI generated img*
## Prerequisites

Before installing and using Gator, ensure you have the following installed:

- **Go** (version 1.19 or later) - [Download Go](https://golang.org/dl/)
- **PostgreSQL** (version 12 or later) - [Download PostgreSQL](https://www.postgresql.org/download/)

## Installation

Install Gator using Go's built-in package manager:

```bash
go install github.com/Cheemx/gator@latest
```

## Configuration

Create a configuration file named `.gatorconfig.json` in your home directory with the following structure:

```json
{
  "db_url": "postgres://<postgres_user_name>:<postgre_user_passowrd>@localhost:5432/gator?sslmode=disable",
  "current_user_name": ""
}
```

**Note:** The `current_user_name` field will be automatically filled when you run the register command.

### Database Setup

Ensure your PostgreSQL database is running and accessible with the connection string provided in your config file. Gator will create the necessary tables automatically.

## Development

If you're developing or contributing to Gator, you can run commands directly using:

```bash
go run . <command_name> [args]
```

## Commands Reference

### User Management

#### Register/Login
```bash
gator register <username>
gator login <username>
```
Registers a new user or logs in an existing user. Sets the `current_user_name` in the configuration file.

#### Users
```bash
gator users
```
Lists all registered usernames from the users table.

### Database Management

#### Reset
```bash
gator reset
```
Resets all four database tables: users, feeds, posts, and feed_follows. **Use with caution** - this will delete all data.

### Feed Management

#### Add Feed
```bash
gator addfeed <name> <url>
```
Adds a new RSS feed to the feeds table and automatically follows it for the current user.

**Example:**
```bash
gator addfeed "TechCrunch" "https://techcrunch.com/feed/"
```

#### List Feeds
```bash
gator feeds
```
Displays all available feeds with their name, URL, and creator information.

#### Follow Feed
```bash
gator follow <url>
```
Follows an existing feed by adding it to the feed_follows table for the current user.

**Example:**
```bash
gator follow "https://techcrunch.com/feed/"
```

#### Following
```bash
gator following
```
Lists the names of all feeds that the current user is following.

#### Unfollow Feed
```bash
gator unfollow <url>
```
Unfollows a feed for the currently logged-in user.

**Example:**
```bash
gator unfollow "https://techcrunch.com/feed/"
```

### Content Aggregation

#### Aggregate Feeds
```bash
gator agg <duration>
```
Starts periodic aggregation of feeds at the specified interval. Accepts duration formats like:
- `1s` (1 second)
- `1m` (1 minute) 
- `1h` (1 hour)
- `2h45m` (2 hours 45 minutes)

**Example:**
```bash
gator agg 30m
```

#### Browse Posts
```bash
gator browse [limit]
```
Displays posts from your followed feeds stored in the database. Posts are sorted in reverse chronological order (most recent first). Optional limit parameter controls the number of posts displayed.

**Examples:**
```bash
gator browse        # Browse all posts
gator browse 10     # Browse latest 10 posts
```

## Workflow Example

Here's a typical workflow for using Gator:

1. **Register as a new user:**
   ```bash
   gator register johndoe
   ```

2. **Add some RSS feeds:**
   ```bash
   gator addfeed "Hacker News" "https://news.ycombinator.com/rss"
   gator addfeed "Reddit Programming" "https://www.reddit.com/r/programming/.rss"
   ```

3. **Start aggregating feeds:**
   ```bash
   gator agg 1h
   ```

4. **Browse your aggregated posts:**
   ```bash
   gator browse 20
   ```

## Database Schema

Gator uses four main tables:

- **users**: Stores registered user information
- **feeds**: Contains RSS feed URLs, names, and metadata  
- **posts**: Stores individual posts/articles from feeds
- **feed_follows**: Manages user-feed relationships (which users follow which feeds)



## Support

If you encounter any issues or have questions, please [open an issue](https://github.com/yourusername/gator/issues) on GitHub.