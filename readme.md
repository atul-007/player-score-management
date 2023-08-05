Requirements

Every player will have an associated entry in the database, which will
contain four attributes: ID, Name, Country & Score. All the attributes are
mandatory. ID uniquely identifies each player. Name should have a cap of
15 characters. Country code will be a two letter code representing the
country (For e.g., IN, US). Score will be an integer representing player
score. The service should expose the following endpoints, which should
return JSON responses:


1. POST /players – Creates a new entry for a player

2. PUT /players/:id – Updates the player attributes. Only name and
score can be updated

3. DELETE /players/:id – Deletes the player entry

4. GET /players – Displays the list of all players in descending order

5. GET /players/rank/:val – Fetches the player ranked “val”

6. GET /players/random – Fetches a random player