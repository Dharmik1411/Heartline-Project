# Heartline-Project

# Simple Auth API (Go + PostgreSQL)

Beginner-friendly REST API project built with Golang and PostgreSQL.
Supports user registration, login, profile fetching, and profile updating with JWT-based authentication.

## Endpoints
- `POST /register`
- `POST /login`
- `GET /profile` (protected)
- `PATCH /profile` (protected)

### Database Setup
Run the SQL script in `database.sql` to create the `users` table:
```bash
psql -U postgres -d simple_auth -f database.sql
```
### How to Run
1. Create PostgreSQL DB and table (`users`).
2. Update DB connection in `config/config.go`.
3. Run:
```bash
go mod tidy
go run main.go
```


### Working days: Mon - Fri (35-40 hrs per week). 
### Time Zone: EST Eastern Time	UTC -5:00 / -4:00 or MST Mountain Time	UTC -7:00 / -6:00




