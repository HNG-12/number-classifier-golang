package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net/http"
	"strconv"
)

func IsPrime(n int) bool {
	if n <= 1 {
		return false
	}
	limit := int(math.Sqrt(float64(n)))
	for i := 2; i <= limit; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func IsPerfect(n int) bool {
	sum := 0
	for i := 1; i < n; i++ {
		if n%i == 0 {
			sum += i
		}
	}
	return sum == n
}

func IsArmstrong(n int) bool {
	sum := 0
	temp := n
	digits := len(strconv.Itoa(n))
	for temp != 0 {
		remainder := temp % 10
		sum += int(math.Pow(float64(remainder), float64(digits)))
		temp /= 10
	}
	return sum == n
}

func DigitSum(n int) int {
	if n < 0 {
		n = -n
	}
	sum := 0
	for n > 0 {
		sum += n % 10
		n /= 10
	}
	return sum
}

type NumberAPIResponse struct {
	Text string `json:"text"`
}

func GetFunFact(n int) (string, error) {
	url := fmt.Sprintf("http://numbersapi.com/%d/math?json", n)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("failed to close response body")
		}
	}(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to fetch fun fact: %d", resp.StatusCode)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var apiResponse NumberAPIResponse
	err = json.Unmarshal(body, &apiResponse)
	if err != nil {
		return "", err
	}
	return apiResponse.Text, nil
}
