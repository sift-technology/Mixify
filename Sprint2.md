# Work we have completed

## Frontend

Help

## Backend

For sprint 2 we have been able to host the frontend and backend on the same port. Furthermore, we have created a post method to retrieve user responses 
from the frontend and store those responses in our backend database. Then we created a method which uses the survey responses and weights
them to coorespond with music metrics for the spotify API and stores them in the database.

# Unit Tests

## Frontend

* Help

## Backend

* *TestValidateJSONResponse()* - Test to see if we can popuated the user defined Response type with JSON numbers
* *TestCreateUniqueUUID()* - Test to see if user ID generates unique ID values for each user
* *TestWeightFunct()* - Test to see if the music metrics corresponding to all possible user responses are within an appropriate range

# Backend API Documentation

**Server.go**

* *NewServer()* - Creates new gorrila mux router and database and stores in server, calls the *routes()*, and returns new server.
* *routes()* - Adds routes and the http method associated with each route. Last "route" host frontend on same port as backend.
* *CreateResponse()* - Called on "/results" route. Decodes JSON argument from survey response, calls *Weights()*, and stores 
UUID, response, and music metrics in the database.
* *Weights()* - Assigns specific weight values cooresponsing to each response to assign the user's music metrics with a specific value.
* *ListResponses()* - Called on "/response" route. Encodes all data from the database to JSON and outputs that data. 
* *removeResponse()* - Called on "/response/{id}" route. Removes response cooresponding to the UUID passed as a route argument (id). 

**angular_live.go**

File responsible for hosting front end assets on backend port (8080). *routes()* uses the **AngularHandler** variable to make the connection.

**main.go**

Creates new server by calling *NewServer()* and host that server on port 8080 using http's method *ListenAndServe()*.
