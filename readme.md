# wallet

I use existing libs and tools :

 - Ozzo Validation, for input request validation
 - Godotenv, for env loader
 - jmoiron/sqlx for postgres driver
 - postgresql for DB

 # For setup after cloning/unzip the project:
> cd ex_service
> go mod tidy
> make changes in the .env file as the requierement there


# to do a unit test :
> i've made several unit testing but just in usecases layer
> go to the each usecase package you want to testing then run a command "go test"
> you can see the coverage testing in each usecase package by open the project with vscode, choose the testing file, right click then choose "Go:Toogle Test Coverage in Current Package"

# to run the project
after set the .env file with your database and redis credential, then stay still in root directory, then do "go run main.go" in terminal

