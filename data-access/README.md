# data access examples

This demonstrates a pattern for accessing data in a database.  This case specifically authenticating users.

Service objects are singletons that have business logic.  They use the DAO (data access object) singletons when they interact with the database.

