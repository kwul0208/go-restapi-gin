package productController

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kwul0208/go-restapi-gin/models"
	"gorm.io/gorm"
)

func Index(c *gin.Context) {
	var products []models.Product
	models.DB.Find(&products)
	c.JSON(http.StatusOK, gin.H{"products": products})
}
func Show(c *gin.Context) {
	var product models.Product
	id := c.Param("id")

	if err := models.DB.First(&product, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Product not found"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"product": product})

}
func Create(c *gin.Context) {
	var product models.Product
	var inputMap map[string]interface{}

	// Bind JSON input ke map untuk memeriksa keberadaan field yang tidak diharapkan
	if err := c.ShouldBindJSON(&inputMap); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Invalid JSON format", "error": err.Error()})
		return
	}

	// Check for unexpected fields
	for key := range inputMap {
		if key != "name" && key != "description" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Unexpected field in input: " + key})
			return
		}
	}

	// Bind JSON input to the Product struct
	if err := c.ShouldBindJSON(&product); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Validation failed", "error": err.Error()})
		return
	}

	// Lakukan penyimpanan data jika tidak ada error
	models.DB.Create(&product)
	c.JSON(http.StatusCreated, gin.H{"message": "Product created", "product": product})
}

func Update(c *gin.Context) {
	var product models.Product
	id := c.Param("id")

	// Bind JSON input to the Product struct
	if err := c.ShouldBindJSON(&product); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Validation failed", "error": err.Error()})
		return
	}

	if models.DB.Model(&product).Where("id = ?", id).Updates(&product).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Failed to update data",
		})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "Success updated"})
}
func Delete(c *gin.Context) {
	var product models.Product

	var input struct {
		Id json.Number
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	id, _ := input.Id.Int64()
	if models.DB.Delete(&product, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Failed to update data",
		})
		return
	}

	c.AbortWithStatusJSON(http.StatusAccepted, gin.H{"message": "success delete data"})
}
