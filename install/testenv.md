# Testenvironment

## Database
- Start a docker container `docker run --name p2k12-db -e POSTGRES_PASSWORD=password -p 5432:5432 -d p2k12-db`

- psql `docker run -it --link p2k12-db:postgres --rm postgres sh -c 'exec psql -h "$POSTGRES_PORT_5432_TCP_ADDR" -p "$POSTGRES_PORT_5432_TCP_PORT" -U postgres'`
- create users

   create user p2k12 with password 'password';
   create user mortehu with password 'password';
   create user ccscanf with password 'password';

- restore database file `docker run -it --link p2k12-db:postgres -v (pwd)/p2k12.sql:/p2k12.sql --rm postgres sh -c 'exec psql -h "$POSTGRES_PORT_5432_TCP_ADDR" -p "$POSTGRES_PORT_5432_TCP_PORT" -U postgres -f p2k12.sql'`
