<h3 align="center">exp.dev (backend)</h3>
<sup align="center">A backend for a <a target="_blank" href="github.com/tokiory/exp.dev">exp.dev</a> project</sup>


# Docker
To utilize infrastructure, use the following command:
```bash
docker compose up -d database migration
```

To run the application and the infrastructure, use:
```bash
docker compose up -d
```

To stop app and (or) infrastructure, use:
```bash
docker compose down
```

# Tools
- [`sqlc`](https://github.com/kyleconroy/sqlc): For generating db package code from SQL
- [`goose`](https://github.com/pressly/goose): For migrations

## SQL Generation
This application uses `sqlc` to generate Go code from SQL queries. To generate the code, run:

```bash
sqlc generate -f ./db/sqlc.yaml
```

This command will read the SQL files in the `db/migration` and `db/queries` directories and generate the corresponding Go code in the `db/report` directory.

## Migrations
All migrations are located in the `db/migration` directory. To apply migrations, use:

```bash
make migrate_up
```

# Architecture
1. We use immutable database records of reports to keep progress of every engineer;
2. This approach provide us a great way to analyze every developer path, 'cause we can simple `SELECT` all developers reports via SQL;
