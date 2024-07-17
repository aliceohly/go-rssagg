# RSS Feed Aggregator

This is a Go project where we build an RSS feed aggregator in Go! It's a web server that allows clients to:

- Add RSS feeds to be collected
- Follow and unfollow RSS feeds that other users have added
- Fetch all of the latest posts from the RSS feeds they follow

RSS feeds are a way for websites to publish updates to their content. You can use this project to keep up with your favorite blogs, news sites, podcasts, and more!

## Getting Started

### Local Development

`go build & ./go-rssagg`

### Database Migration

Migrate the DB to the most recent version available

`cd sql`

`goose postgres postgres://xxx:@localhost:5432/rssagg up`

Roll back the version by 1
`goose postgres postgres://xxx:@localhost:5432/rssagg down`

### Parse SQL to Go code

`cd rssagg`
`sqlc generate`
