# Blog Customer Backend

## Todo list
- [ ] User
   - [ ] Getting a profile by ID
   - [ ] Editing
   - [ ] Update settings
   - [x] Subscribe/unsubscribe
   - [ ] Getting a list of subsite categories
   - [ ] Request to change/confirm password
- [ ] Article
   - [ ] Create immediately
   - [ ] Create as a draft
   - [ ] Publication
   - [ ] List
   - [ ] Receiving by ID
   - [ ] Subscription
   - [ ] Unsubscribe
   - [ ] Editing
   - [ ] Comments
      - [ ] Receipt
      - [ ] Editing
      - [ ] Creation
      - [ ] Delete
- [x] Bookmark
   - [x] List
   - [x] Creation
   - [x] Delete
- [x] Auth
   - [x] Create authorization check middleware
   - [x] Token refresh
   - [x] Authorization
   - [x] Registration

### Features that were not added in an earlier version (those marked were implemented)
1. [ ] Moderator rights
2. [ ] Reactions to posts
2. [ ] Reactions to comments
3. [ ] Complaints
4. [ ] Hiding posts and users
5. [ ] Blur posts 18+
6. [ ] OAuth2 (yandex, telegram, and more)
7. [ ] Deleting an account
8. [ ] Admin
9. [ ] Notifications
10. [ ] Messages

# App
To launch the application do the following

1. Set up local/prod config files
2. Run the build command ```$ make build-app```
3. Run the application launch command ```$ make run-app-local``` or ```$ make run-app-prod```

# Migrations
To deploy migrations, you can use the cmd/migrator tool

It is based on the tool https://github.com/golang-migrate/migrate for this occasion you can use the commands from this library

For work, you need:
1. Set up the local/prod configuration file
    1. Create migrations if they don't exist yet
2. Build a tool for rolling up migrations with the command ```$ make build-migrator-up```
3. Run migrations using the command ```$ make run-migrate-up-local``` or ```$ make run-migrate-up-prod```

# Locale Development

### Docker Compose Local environments
Docker Compose Local file contains the following environment variables:

In this solution, local compose is used exclusively to create a database

* `POSTGRES_USER`
* `POSTGRES_PASSWORD`
* `PGADMIN_DEFAULT_EMAIL`
* `PGADMIN_DEFAULT_PASSWORD`

### Access to postgres:
* `localhost:5432`
* **Username:** postgres (as a default)
* **Password:** test_password (as a default)

### Access to PgAdmin:
* **URL:** `http://localhost:5050`
* **Username:** admin@service.example (as a default)
* **Password:** admin (as a default)

### Add a new server in PgAdmin:
* **Host name/address** `postgres`
* **Port** `5432`
* **Username** as `POSTGRES_USER`
* **Password** as `POSTGRES_PASSWORD`