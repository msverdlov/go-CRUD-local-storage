# Simple API implementation in Go

## To Get Started
#### Clone repo
1. Make sure you have [Go](https://go.dev/) installed on your computer.
2. Run the command to clone this github repository and change the directory to the project's folder:
```bash
git clone https://github.com/vsevdrob/api-go-gin-viper.git && cd api-go-gin-viper/src
```
#### Before we run the programme, let's download a couple of Go dependencies first.
3. This will download one of the coolest HTTP web framework [gin](https://github.com/gin-gonic/gin) written in Go (Golang).
```bash
go get github.com/gin-gonic/gin
```
4. [Viper](https://github.com/spf13/viper) helps us to to operate with the predefined `config.yaml` file in order to export some required values from it.
```bash
go get github.com/spf13/viper
```
# Usage 
## Start server
Run the command that starts the server on host `127.0.0.1` and port `8080`.
```bash
go run main.go
```
After that open another terminal window and insure the path of current working directory is `*/api-go-gin-viper/src/`
___
## Add data
Adds a data to the storage. Assuming that *address* and *amount* keys/values in `addData.json`
```bash
curl localhost:8080/addData --header "Content-Type: application json" -d @examples/addData.json --request "POST"
```
## Add dataset
Adds a dataset. The same keys/values from above are also predefined in `addDataset.json`.
```bash
curl localhost:8080/addDataset --header "Content-Type: application json" -d @examples/addDataset.json --request "POST"
```
## Fetch data
Get a data by assigning his/her **`id`** in the query.
```bash
curl localhost:8080/fetchData?id=1 --request "GET"
```
## Fetch dataset
Get a addDataset from the storage.
```bash
curl localhost:8080/fetchDataset --request "GET"
```
## Update date
Update data amount by increasing it. Assign the **`id`** in the query to update a specific data. 
```bash
curl localhost:8080/updateData?id=1 --request "PATCH"
```
## Delete data
Delete a data from storage. Assign data **`id`** in the query to delete the preferred one. 
```bash
curl localhost:8080/deleteData?id=1 --request "DELETE"
```
# Licence
**MIT**
