# Sprint 3: Work we have completed

## Frontend

For sprint 2, we worked on developing and adding the actual questions of the survey instead of the placement ones that we had in previous Sprints. This included finding images, developing questions, changing the layout of the frotend, and changing logic for the array that was passed into the backend. Additionally, we implemented more of the /results page, which created space for the output of the survey (spotify song recommendations). Though this has not been 100% integrated with the backend, we have the foundation for it there.
## Backend

For sprint 2, we refactored our code to include the new questions. We had to update our logic for a lot of our code, most notably, in the Weights() function that calculates attributes based on user responses. We also implemented the connection to the Spotify API, which involved registering our application, authenticating it, and using built-in methods to get recommendations. For the last part, we had to add logic to translate the numbers from our Weights() function to values the Spotify API could understand. We implemented a Recommend() function that actually used our translated inputs to ouput song recommendations using the Spotify API algorithm. We are currently working on translating the Spotify Recommendations (which we can print in JSON format) to URLS that can easily be embedded in the frontend.
# Unit Tests

## Frontend

#### Unit
* *title of app should be "Mixify"* - verifies app title
* *title in toolbar should be "Mixify"* - checks that header contains title
* *navigate to "results" takes you to /results* - checks that setting path to '/results' routes app to results component
* *title in h2 tag is "Questions :"* - verifies title of questions component
* *click "Submit" calls onSubmit function* - checks that onSubmit function is called when user clicks submit button
* *getList function makes post request* - checks that calling getList function in service makes a post request

#### Cypress
* *clicking "submit" navigates to a new url* - checks that when user clicks the submit button, the app navigates to '/results' url
* *clicking all of the answer choices in the survey* - checks that when a user clicks answer choices and then submits, the app navigates to '/results' url
* *clicking a song title* - verifies that when a user clicks on a Spotify URL, it takes them to the Spotify site

## Backend

* *TestValidateJSONResponse()* - Test to see if we can populate the user defined Response type with JSON numbers, updated for updated Response data structure
* *TestCreateUniqueUUID()* - Test to see if user ID generates unique ID values for each user
* *TestWeightFunct()* - Test all 163,840 possible responses and sees if Spotify will recommend tracks for all possible answer combinations

# Backend API Documentation

**spotify.go**
* *Recommend()* - Takes in a Response data structure of user responses and a Spotify client, generates Spotify recommendations based on user Responses and returns the Tracks in a native Spotify API format

**authenticate.go**
* *Authenticate()* - Authenticates our application to use the Spotify API using private keys (for this reason, this file is not committed to GitHub)
**server.go**

* *NewServer()* - Creates new gorrila mux router and database and stores in server, calls the *routes()*, and returns new server.
* *routes()* - Adds routes and the http method associated with each route. Last "route" host frontend on same port as backend.
* *CreateResponse()* - Called on "/results" route. Decodes JSON argument from survey response, calls *Weights()*, and stores 
UUID, response, and music metrics in the database. Also calls Recommend(), then stores recommendations in a database as well.
* *Weights()* - Assigns specific weight values cooresponsing to each response to assign the user's music metrics with a specific value.
* *ListResponses()* - Called on "/response" route. Encodes all data from the database to JSON and outputs that data. 
* *removeResponse()* - Called on "/response/{id}" route. Removes response cooresponding to the UUID passed as a route argument (id). 

**angular_live.go**

File responsible for hosting front end assets on backend port (8080). *routes()* uses the **AngularHandler** variable to make the connection.

**main.go**

Creates new server by calling *NewServer()* and host that server on port 8080 using http's method *ListenAndServe()*.
