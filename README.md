# Kredit Plus Case Study

## Architecture

![Kredit Plus Architecture](Kredit%20Plus%20Architechture.drawio.png)

## Entity Relationship Diagram (ERD)

![Kredit Plus ERD](Kredit%20Plus%20ERD.png)

## Concurrent Transaction Handling

This system implements concurrent transaction safety in the following endpoint:
 ```sh
 🔗 POST {{base_url}}/transaction
 ```

When a customer submits a transaction request, the system checks their available credit limit for the selected tenor. To prevent race conditions when multiple requests hit the system at the same time, the logic uses:

- Row-level locking: using FOR UPDATE via GORM's clause
- Transaction block: using s.db.Begin() and tx.Commit()/tx.Rollback() for atomicity
- Limit validation: if the sum of OTR + Admin Fee exceeds the customer's limit, the request is aborted

📁 Code Location
 ```sh
 src/domain/transaction/service/create_transaction_svc.go
 ```

## Local Development

### How to run

1. Copy file `Makefile.example` to `Makefile`
  ```sh
  cp Makefile.example Makefile
  ```

2. Setup PostgreSQL then config database url on `Makefile` 
  ```
  MIGRATE_DB_URL=postgresql://username:password@127.0.0.1:5432/kredit-plus?sslmode=disable
  ```

3. Migrate database
  ```sh
  make migrate.up
  ```

4. Copy file `.env.example` to `.env`
  ```sh
  cp env.example .env
  ```

5. Setup the configuration to `.env`, but you're required to config the PostgreSQL & redis
  ```
  DB_HOST=127.0.0.1
  DB_PORT=5432
  DB_USERNAME=
  DB_PASSWORD=
  DB_SCHEMA=

  CACHE_HOST=127.0.0.1
  CACHE_PORT=6379
  CACHE_USERNAME=
  CACHE_PASSWORD=
  CACHE_DB=0
  ```
  
6. Install dependency go
  ```sh
  make deps
  ```

7. Run app for development
  ```sh
  make run
  ```

## Account
 ```sh
  email: admin@gmail.com
  password: password 
 ```

## API Documentation

[Download Postman API JSON File!](Kredit%20Plus.postman_collection.json)