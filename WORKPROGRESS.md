<h1>Stock Exchange Assets API</h1>
<h3>This is an API to consult user registeded assets.</h3>

Each user allowed to use should register on an onboarding process, allowing login on the system and consequently obtaining an authorization token.

<h3>Necessary Interfaces:</h3>
* user onboarding
* user login
* user management (of actives, add, remove, set order)
* list of actives of the logged user (based on identification token)

<h3>List of application pieces/containers (will be controled using swarm):</h3>
* api host
* database (postgres?)
(will the database have a previous payload? or will it start completely empty?)
