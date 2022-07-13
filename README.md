2 new users to allow multiple collections on database
Generate SQL to create and populate the DB on the Docker image creation

# idealAPI
<h1>Ideal Code Challenge</h1>
<h1>Stock Exchange Assets API</h1>

Running the project:
go build .
docker-compose build
docker-compose up

Flow:
user login / get the token  id
manage assets (API)
list assets (API)

<h3>Project description:</h3>
This is an API to consult user registeded assets. A basic CRUD api for assets has been developed.
Each user allowed to use should register on an onboarding process, allowing login on the system and consequently obtaining an authorization token.
The database and API climb into individual containers using docker-compose. Access is via port 3000.

<h3>Current Status:</h3>
The project is incomplete with respect to authorization, implementation of some functions and also code maturity.
Pending points:
* API interface to select the list organized, if this task is to be done at server side
* API to obtain the quote
* Authorization/Authentication
* Unit testing
* Better code quality

<h3>Reasons for incomplete project:</h3>
This was the first project in which I had contact with GO. This took a heavy toll on time for minimal language knowledge. As a consequence, I had a direct impact on the proposed scope and on the quality of what was presented. Also, I'm in delivery week, with a very tight schedule. There was no free time for study and knowledge.

Improvement points:
* The authentication process was not addressed, being simulated in an extremely simplistic way with the userid being sent to the apis in auth, making a rough simulation of a keycloak. The ids were manually populated in the bank and provided in this document.
* Lack of middleware implementation for auth0.
* Lack of middleware for CORS and additional needs for the API to work on a reliable way.
* Focus on SOLID: the code should be much better prepared, with abstractions and focus on quality with SOLID. It was not possible to give this focus due to lack of knowledge of the language. Although I was delighted with my first contact with GO, the way the language works regarding structure surprised me with some differences with the traditional OOP approach used in other languages. I need a better understandment about the class composition with organization in structs and methods that resemble class extensions in other OOP languages. I also need to get to know the GO paradigm better because in the fundamentals of the language it aims to be simple and fast. Without proper knowledge of the model it will be easy to make a mess trying to achieve loosely coupled code.
* Unit tests. A reliable project MUST HAVE unit tests. Not only focused on coverage, but also a complete set of test cases.

<h3>Libraries you would like to know better and explore properly if I had time: </h3>
* Google Wire, I found the compile-time dependency injection approach offered quite interesting.
* GORM or another ORM library, this would make the separation of concerns better structured with repositories implementation. Together with IOC this would pprovide a better quality of code.
* Depending of what the selected ORM lib provides, maybe Goose would be a good idea to manage the DB schema.
* AuthBoss to help address authorization implementation.

