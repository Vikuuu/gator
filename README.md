# Gator 🐊

> This is RSS feed aggregator in Go!

## Feature 

- Add RSS feeds from across the internet to be collected.
- Store the collected posts in PostgreSQL databse.
- Follow and unfollow RSS feeds other users have added.

## Installing Go
For this CLI tool to work you will have to install GO(obviously..😜)

You will need to install Go(version 1.23+), you can follow this [[installation instruction](https://go.dev/doc/install) ]

- If you're getting a "command not found" error after following installation, it's most likely because the directory containing the `go` program isn't in yout `PATH`. You need to add the directory to your `PATH` by modifying your shell's configuration file. First, you need to know *where* the `go` command was installed. It might be in:

- `/usr/local/go/bin` (official installation)

- Somewhere else?

You can ensure it exists by attempting to run go using its full filepath. For example, if you think it's in ~/.local/opt/go/bin, you can run ~/.local/opt/go/bin/go version. If that works, then you just need to add ~/.local/opt/go/bin to your PATH and reload your shell:

```
# For Linux/WSL
echo 'export PATH=$PATH:$HOME/.local/opt/go/bin' >> ~/.bashrc
# next, reload your shell configuration
source ~/.bashrc
```

## Installing Gator

Use can install the gator by using `go install` command:
```bash
go install github.com/Vikuuu/gator
```

## Config

Gator is a multi-user CLI application. There's no server(other than the database), so it's only intended for local user.

- Create file `~/.gatorconfig.json` in your root directory, and add the following:

```json
{
    "db_url": "postgres://postgres:postgres@localhost:5432/gator",
}
```
For this you should have **PostgreSQL** installed.

> [!NOTE]
> You will have to create the database gator manually. Steps are shown [here](#installing-postgresql)

## Installing PostgreSQL

1. Linux/WSL(Debian):
```bash
sudo apt update
sudo apt install postgresql postgresql-contrib
```

2. Ensure it is installed by running the following command:
```bash
psql --version
```

3. Update postgres password(Linux only):
```bash
sudo passwd postgres
```
You can add the `postgres` as password so that you won't forget.

4. Start the Postgres server in the background
- Linux: `sudo service postgresql start`

5. Connect to the server, simply using `psql` client
- Linux: `sudo -u postgres psql`

You should see a new prompt that looks like this:
```bash
postgres=#
```

6. Now create the `gator` database in your postgresql
```bash
postgres=# CREATE DATABASE gator;
```

7. You can move to your database by using:
```bash
postgres=# \c gator
```

8. Check the table in the gator database.
```bash
postgres=# \dt
```

## Using the CLI

- You can register a user by following command:
```bash
gator register <name>
```

- You can login a user by following command:
```bash
gator login <username>
```

- You can add feeds by following command:
```bash
gator addfeed <feedName> <feedURL>
```

You can run the command `help` to get the information on all the available commands:
```bash
gator help
```
