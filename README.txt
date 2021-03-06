== REQUIREMENT ==
Docker

== RUN SERVER ==
1. Open terminal in project folder
2. Run "make container" for build (include running server and database) in the docker container 
3. Run "make stop" for stopping and deleting server container and database container
4. Run "make test" for unit testing

== ENDPOINT API ==
(GET)   http://localhost:9898/account/{id} 
(POST)  http://localhost:9898/account/{from_account_number}/transfer
{
 "to_account_number" : "555001",
 "amount" : 1000
}

==Revision ==
Makefile & Unit testing

