package resolvers

import (
	"ThaiLy/configs"
	"ThaiLy/models"
	"fmt"

	"github.com/graphql-go/graphql"
)

func GetCategoryResolver(p graphql.ResolveParams) (interface{}, error) {
	categoryID := p.Args["categoryID"]
	if categoryID != nil {
		fmt.Println(categoryID)
		var category []models.Category
		result := configs.GetDB().Where("id = ?", categoryID).Find(&category)
		fmt.Println(category)
		if result.Error != nil {
			return nil, result.Error
		}
		return category, nil
	}
	var categories []models.Category
	result := configs.GetDB().Find(&categories)
	if result.Error != nil {
		return nil, result.Error
	}
	return categories, nil
}

func GetProductByCategory(p graphql.ResolveParams) (interface{}, error) {
	productInput, ok := p.Args["ProductInput"].(map[string]interface{})
	offset, limit := 0, 10
	var featured *string
	if ok {
		if offsetVal, ok := productInput["offset"].(int); ok {
			offset = offsetVal
		}
		if limitVal, ok := productInput["limit"].(int); ok {
			limit = limitVal
		}
		if featuredVal, ok := productInput["featured"].(string); ok {
			featured = &featuredVal
		}
	}
	category, ok := p.Source.(models.Category)
	if !ok {
		return nil, fmt.Errorf("invalid category source")
	}
	categoryID := category.ID
	var products []models.Product
	result := configs.DB.Where("CategoryID = ?", categoryID).Limit(limit).Offset(offset)
	if featured != nil {
		result = result.Where("featured = ?", *featured)
	}
	if result.Error != nil {
		return nil, result.Error
	}
	result = result.Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}
	return products, nil
}
