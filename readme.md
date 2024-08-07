# Golang RSS Reader

Loosely inspired by Miniflux, another golang app.

> [Just an idea for now] Maybe this app should be more than just an RSS reader. Maybe it should be a tool in aiding consuming all content. (Maybe I should split it up into separate modules/packages too but that sounds hard so I'll do that much later.) So while I get my RSS feeds, theres also sparse items that I find and want to consume. I'd like this app to have a "later" functionality, to save articles/posts/videos/whatever. It would also be nice if I could build out a viewing(reading, listening, and/or watching) queue. All hopefully to curb my incessent tab opening and never finishing content I start.

## Goals

- RSS webapp
- Rest API
- SQLite

## Planned Features

- [ ] User accounts
- [ ] Bring in DaisyUI
- [ ] Infinitely nested folders for organizing Feeds
  - Allowing for grouping multiple feeds owned by same person/entity (e.g. youtube & blog)
- [ ] Viewing a folder will view all contained feeds
- [ ] RSS Categories to be tracked
- [ ] Read Later label to not loose entries of interest
- [ ] Different view modes
- [ ] Support/integration for embedding youtube(/similar) videos for easy viewing and direct links.
- [ ] Searchable viewing history
- [ ] Mark as unread
- [ ] (would be cool to have a) [Wish](https://github.com/charmbracelet/wish) App

## Development

Set `TURSO_DB_URL` & `TURSO_DB_TOKEN` environment variables.

Run `make start`

## Reources

- [Templ Syntax & Usage](https://templ.guide/syntax-and-usage)
- [ ] [Turso Railway](https://docs.turso.tech/features/embedded-replicas/with-railway)
