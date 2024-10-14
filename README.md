# RoastAdvisor

## Setup

Start postgres

```
brew services start postgresql@14
```

Compile the binary and run it

```
make run
```

Enter a psql shell

```
psql -U postgres -h localhost -d postgres
```

Stop postgres running after use

```
brew services stop postgresql@14
```
