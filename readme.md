# Golang RSS Reader

Loosely inspired by Miniflux, another golang app.

## Goals

- RSS webapp
- Rest API
- SQLite

## Planned Features

- [ ] User accounts
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
