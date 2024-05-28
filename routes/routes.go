package routes

import (
	"crud/database"
	"crud/models"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateAnimal route for creating new animal

// @Summary Create a new animal
// @Description Create a new animal with the given information
// @Accept  json
// @Produce  json
// @Param   animal  body    models.Pet  true  "Animal Info"
// @Success 200 {string} string "Animal inserted successfully"
// @Failure 400 {string} string "Some error occur. Please retry."
// @Router /create-pet [post]
func CreateAnimal(c *gin.Context) {
	var data models.Pet
	err := c.BindJSON(&data)

	if err != nil {
		log.Println("Error in Reading the body ", err)
		return
	}

	err = database.InsertIntoPet(data)

	if err != nil {
		log.Println("Error in inserting to Pets table ", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Some error occur.Please retry.",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Animal inserted successfully",
	})
}

// FindByID route for getting the pet for particular id

// @Summary Get pet by ID
// @Description Get the pet by ID of the pet
// @Param   id  path    int  true  "Pet ID"
// @Success 200 {object} models.Pet
// @Failure 400 {string} string "Some error occurred. Please retry."
// @Router /pet/id/{id} [get]
func FindByID(c *gin.Context) {
	id := c.Param("id")

	newID, _ := strconv.Atoi(id)

	data, err := database.GetPetByID(newID)
	if err != nil {
		log.Println("Errorin getting the pet for particular id ", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Some error occured.Please retry.",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

// FindByStatus route to get pet by status

// @Summary Get pet by status
// @Description Get the pet by status of the pet
// @Param   status  path    string  true  "Pet Status"
// @Success 200 {array} models.Pet
// @Failure 400 {string} string "Some error occurred. Please retry."
// @Router /pet/status/{status} [get]
func FindByStatus(c *gin.Context) {

	status := c.Param("status")
	var statusID int

	if status == "available" {
		statusID = 1
	} else if status == "pending" {
		statusID = 2
	} else {
		statusID = 3
	}

	data, err := database.GetPetByStatus(statusID)
	if err != nil {
		log.Println("Error in getting the Pet by status ", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Some Error  occured.Please retry",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})

}

// DeleteByID route will delete the Pet with given ID

// @Summary Delete pet by ID
// @Description Delete the pet by ID
// @Param   id  path    int  true  "Pet ID"
// @Success 200 {string} string "Pet deleted Successfully."
// @Failure 400 {string} string "Some error occurred. Please retry."
// @Router /pet/{id} [delete]
func DeleteByID(c *gin.Context) {
	id := c.Param("id")
	newID, _ := strconv.Atoi(id)

	count, err := database.GetCountOfID(newID)
	if err != nil {
		log.Println("Error in getting the total count ", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "SOme error occured.Please retry.",
		})
		return
	}

	if count == 0 {
		c.JSON(http.StatusOK, gin.H{
			"message": "There is no Pet of particular ID",
		})
		return
	}

	err = database.DeletePetByStatus(newID)
	if err != nil {
		log.Println("Error in deleting the certain Pet ", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Some error occured.Please retry",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Pet deleted Successfully.",
	})

}

// FullUpdate route will update the full data of the particular Pet

// @Summary Update full data of pet by ID
// @Description Update the full data of the pet by ID
// @Accept  json
// @Produce  json
// @Param   id    path    int        true  "Pet ID"
// @Param   pet   body    models.Pet true  "Pet Data"
// @Success 200 {string} string "Pet updated successfully"
// @Failure 400 {string} string "Some error occurred. Please retry."
// @Router /pet/fullupdate/{id} [put]
func FullUpdate(c *gin.Context) {
	var data models.Pet

	id := c.Param("id")
	newID, _ := strconv.Atoi(id)

	err := c.BindJSON(&data)
	if err != nil {
		log.Println("Error in reading the body ", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Some error occured.Please retry",
		})
		return
	}

	err = database.FullUpdateByID(newID, data)
	if err != nil {
		log.Println("Unable to update the record of particular id ", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Some error occured.Please retry",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Pet updated successfully",
	})
}

// PartialUpdate update the partial record of particular id
// @Summary Partially update pet by ID
// @Description Partially update the pet by ID
// @Accept  json
// @Produce  json
// @Param   id    path    int        true  "Pet ID"
// @Param   pet   body    models.Pet true  "Pet Data"
// @Success 200 {string} string "Partial updation successful"
// @Failure 400 {string} string "Some error occurred. Please retry."
// @Router /pet/partialupdate/{id} [patch]
func PartialUpdate(c *gin.Context) {
	id := c.Param("id")

	newID, _ := strconv.Atoi(id)

	mp := models.Pet{}

	c.BindJSON(&mp)

	fmt.Println(mp)

	data, err := database.GetPetByID(newID)
	if err != nil {
		log.Println("Error in getting the data of particular id ", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Some error occur.Please retry",
		})
		return
	}

	if len(mp.PhotoUrls) != 0 {
		data.PhotoUrls = append(data.PhotoUrls, mp.PhotoUrls...)
	}
	if len(mp.Tags) != 0 {
		data.Tags = append(data.Tags, mp.Tags...)
	}
	if len(mp.Name) != 0 {
		data.Name = mp.Name
	}

	if mp.Status != 0 {
		data.Status = mp.Status
	}

	err = database.FullUpdateByID(newID, data)
	if err != nil {
		log.Println("Error in updating the certain field ", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Some error occured.Please retry",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Partial updation successfull",
	})

}
