#Zebrax BE Interview Test

A simple apps rest API with golang-gin

---------------------------------------

* Tech Stack Used.
* Explanation Application.
* Directory structure.
* Setup
	* Setup instructions
	* Application Deployment

---------------------------------------

## Tech Stack Used
	* Go Language
	* GIN Framework
	* mysql driver
	* crypto
	* godotenv

---------------------------------------

## Explanation Application
	This application is a Rest API for add, get, update and delete student data in the database. This application uses golang with gin framework on the server side and mysqli as database. Use Bearer Token to process the Authorization request and only receive json data for every request that is received.
	
	Aplikasi ini merupakan Rest API untuk melakukan penambahan, pengambilan, pengubahan dan penghapusan data student yang ada pada database. Aplikasi ini menggunakan golang dengan framework gin pada sisi server, mysql pada databasenya, Bearer Token untuk proses Authorization requestnya dan 3 layer arsitektur yaitu endpoint, model, dan 

---------------------------------------

## Directory structure.

-apps
--base
---endpoint
---model
---repository
--config
--services
-logger

---------------------------------------

## Setup
	## Setup instructions
		* Database
			* Create a database.
			* Import the zebrax.sql file in the Database directory into the database used.
			* Entry database driver, database user, database password, database host dan database name into the .env file
		* Application
			* run go mod tidy on terminal/command line to add missing modules
		* Running
			* go to apps directory
			* run go run main.go on erminal/command line to running the application
		* Testing With Postman
			* Import ZebraX Testing file on Postman directory to postman apps
			* Change host and port on URL field (default : host = 127.0.1.1 or localhost, port : 8080)
			* Select Bearer Token Type on Authorization navigation
			* Insert token $2a$08$1rzzWv6InEiwn8nDAnK0iuOdl6v0ySdYr8jir5mWYzglDmkXypr9W on Token fields
			* Go to body navigation and select raw body
			* Insert body data in json format
		* Testing With Golang Test
			* Go to apps directory
			* Run go test -v
			* Custom variable testArray value on test function
	## Application Deployment
		* Change variabel value GIN_MODE on .env file to be release

---------------------------------------
