rest api skeleton in go

**Backend:** Go  
**Persistence:** MariaDB No ORM

### Project Structure

there is a migration folder which contains the db script for user and role tables

all the code is in the src folder, which contains a main.go file that acts as entry point to the program,
and the following packages:

- database: database initialization and query methods
- entity: structs used both as dto and db entities
- handler: methods that are called by the multiplex and produce a response to the client
- reply: function that parses to json the response and adds headers to it
- server : contains both the implementation of the server and the multiplex which manages endpoints