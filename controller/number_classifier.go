package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"number-classifier/utils"
	"strconv"

	"number-classifier/models"
)

func ClassifyNumber(c *gin.Context) {
	numberString := c.Query("number")
	num, err := strconv.Atoi(numberString)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Number: numberString,
			Error:  true,
		})
		return
	}

	digitSum := utils.DigitSum(num)
	isArmstrong := utils.IsArmstrong(num)
	isPrime := utils.IsPrime(num)
	isPerfect := utils.IsPerfect(num)
	var properties []string
	if isArmstrong {
		properties = append(properties, "armstrong")
	}
	if num%2 != 0 {
		properties = append(properties, "odd")
	} else {
		properties = append(properties, "even")
	}
	funFact, err := utils.GetFunFact(num)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Could not fetch fun fact"})
		return
	}

	response := models.NumberResponse{
		Number:     num,
		IsPrime:    isPrime,
		IsPerfect:  isPerfect,
		Properties: properties,
		DigitSum:   digitSum,
		FunFact:    funFact,
	}

	c.JSON(http.StatusOK, response)
}
