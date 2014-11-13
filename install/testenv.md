# Testenvironment

## Database
- Start a docker container `docker run --name p2k12-db -p 5432:5432 -d postgres`
- restore database file `docker run -it --link p2k12-db:postgres -v (pwd)/p2k12.sql:/p2k12.sql --rm postgres sh -c 'exec psql -h "$POSTGRES_PORT_5432_TCP_ADDR" -p "$POSTGRES_PORT_5432_TCP_PORT" -U postgres -f p2k12.sql'`
