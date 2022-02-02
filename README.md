# Cards API
REST API to create and interact with decks of playing cards.

## Tech Stack
 - Go 1.16

## Usage
 1. Either run directly using ```go run .```
	or build and run using ```go build ```
 2. Run all unit tests using `go test ./...`

## REST API
`All the params need to be provided in body as JSON`
### 1. Create new Deck
 `POST /deck` Creates a new deck according to the provided params.
#### Request Params  
| param | type | default | description|
| --- | --- | --- | --- |
| shuffle | boolean, optional | false | If true, the deck will be created in shuffled order |

To create custom deck user need to provide cards in Query.
ex. http://localhost:8000/deck?cards=AS,KD,AC,2C,KH

#### Response
| param | type | description|
| --- | --- | --- |
| deck_id | string | UUID of the created deck|
| shuffled | boolean | Indicates whether the deck was shuffled during creation |
| remaining | integer | The number of cards remaining in the deck |


### 2. Get Deck
 `GET /deck/{deck_uuid}` Returns the deck corresponding to the provided deck UUID
 
 #### Response
| param | type | description|
| --- | --- | --- |
| deck_id | string | UUID of the returned deck |
| shuffled | boolean | Indicates whether the deck was shuffled during creation |
| remaining | integer | remaining cards in the deck |
| cards | array of card objects `{suit string, value string, code string}` | Deck of cards |
 
 
 ### 3. Draw Cards
  `PUT /deck/{deck_uuid}` Returns _n_ cards from the top of the deck corresponding to the provided deck UUID
  
  #### Request Params
  | param | type | default | description|
  | --- | --- | --- | --- |
  |numberofcards| integer | N/A | The number of cards to draw. Must be greater than `0`.|
  
   #### Response
| param | type | description|
| --- | --- | --- |
| cards | array of card objects `{suit string, value string, code string}` | The drawn cards. |

## Testing API
 1. For testing api user has to provide the newely generated `UUID` wherever its required in the request url.
 2. Start the server by building the package or execute command `go run .` for testing
 3. Command for running test is `go test ./...`