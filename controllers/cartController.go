package conterollers

import (
	initializers "e-com-be-go/initilizers"
	"e-com-be-go/models"

	"github.com/gin-gonic/gin"
)

func AddToCart(c *gin.Context) {
	var input struct {
		UserID   uint `json:"user_id"`
		Products []struct {
			ProductID uint `json:"product_id"`
			Quantity  int  `json:"quantity"`
		} `json:"products"`
	}
	var cart models.Cart
	result := initializers.DB.Where("user_id = ?", input.UserID).FirstOrCreate(&cart, models.Cart{UserID: input.UserID})
	if result.Error != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": result.Error.Error()})
		return
	}
	for _, p := range input.Products {
		var cartProduct models.CartProduct
		if err := initializers.DB.Where("cart_id = ? AND product_id = ?", cart.ID, p.ProductID).FirstOrCreate(&cartProduct, models.CartProduct{
			CartID:    cart.ID,
			ProductID: p.ProductID,
		}).Error; err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		// Update the quantity
		cartProduct.Quantity += p.Quantity
		initializers.DB.Save(&cartProduct)
	}
}
