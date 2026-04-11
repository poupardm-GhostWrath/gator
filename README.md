# Gator

### Table of Contents
- [About](#about)
- [Getting Started](#getting-started)
    - [Requirements](#requirements)
    - [Install](#install)
        - [Postgresql](#postgresql)
        - [Gator](#gator-1)

## About

Guided project part of Boot.dev. Goal to build an RSS feed aggregator in Go.

## Getting Started

#### Requirements

- Postgresql 18.3+ or Higher
- Go 1.26.0+ or Higher

#### Install

##### Postgresql

1. Install Postgresql

Linux / WSL (debian)
```
sudo apt update && sudo apt install postgresql postgresql-contrib
```

Linux (Arch)
```
sudo pacman -S postgresql
```

macOS with brew
```
brew install postgresql@15
```

2. (Linux / WSL only) Set postgres password

```
sudo passwd postgres
```

Make passwd `postgres`.

3. Start Postgresql

Mac: `brew services start postgresql@15`

Linux (debian): `sudo service postgresql start`

Linux (Arch): `systemctl enable -- now postgresql.service && systemctl start postgresql.service`

4. Check postgresql working

Mac: `psql postgres`

Linux: `sudo -u postgres psql`

You should see a new prompt:
```
postgres=#
```

5. Create gator database

```
CREATE DATABASE gator;
```

6. Connect to the new database

```
\c gator
```

You should see a new prompt:
```
gator=#
```

7. Set user password (Linux/WSL only)

```
ALTER USER postgres PASSWORD 'postgres';
```

Type `exit` to leave shell.

##### Gator

1. Clone repository:

```
git clone https://github.com/poupardm-GhostWrath/gator
```

2. Change into the project directory:

```
cd gator
```

3. Install gator

```
go install
```

4. Create Gator Config

```
echo '{"db_url":"postgres://postgres:postgres@localhost:5432/gator?sslmode=disable"}' > ~/.gatorconfig.json
```

5. Install Goose

```
go install https://github.com/pressly/goose/v3/cmd/goose@latest
```

Run `goose -version` to make sure it's installed correctly.

6. Create DB

`cd sql/schema`

Run the next command 5 times to build DB:
```
goose postgres "postgres://postgres:postgres@localhost:5432/gator" up
```

7. Run Program

```
usage: gator <command> [args...]
help: gator help
```