# Numbers Classifier API (Golang)
This is a simple API that classifies numbers into even or odd. It is written in Golang and uses the [Gin]()

## Installation
1. Clone the repository
```bash
git clone https://github.com/HNG-12/number-classifier-golang.git
```
2. Change into the project directory
```bash
cd number-classifier-golang
```
3. Run the following command to install the dependencies
```bash
go mod download
```
4. Run the following command to start the server
```bash
go run main.go
```

## Usage
The API has only one endpoint which is a GET request to `/api/classify-number`. The request body should be a JSON object with a key `number` and a value of the number to be classified. The response will be a JSON object with a key `result` and a value of either `even` or `odd`.

Example request:
```bash
curl -X GET "http://localhost:8080/api/classify-number?number=666"
```

Response (Success Response):
```json
{
  "number":666,
  "is_prime":false,
  "is_perfect":false,
  "properties":["even"],
  "digit_sum":18,
  "fun_fact":"666 is the largest rep-digit triangular number."
}
```

Response (Error Response):
```json
{
  "number":"abc",
  "error":true
}
```

## Project Structure
The project is structured as follows:
```bash
number-classifier-golang
├── main.go
├── go.mod
├── go.sum
├── README.md
└── handlers
    └── number_classifier.go
├── utils
    └── number_classifier.go
├── models
    └── number_classifier.go
├── routes
    └── number_classifier.go
├── middlewares
    └── number_classifier.go
├── tests
    └── number_classifier_test.go
```

## Testing
To run the tests, run the following command:
```bash
go test ./...
```

## Contributing
1. Fork this repository
2. Clone the forked repository
3. Create your feature branch
4. Commit your changes
5. Push to the branch
6. Create a new Pull Request
7. Wait for your PR to be reviewed
8. Merge your PR
9. Congratulations 🎉