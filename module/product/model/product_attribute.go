package productmodel

// ProductAttribute represents an attribute with its values for a product
type ProductAttribute struct {
	ID     int                     `json:"id"`
	Name   string                  `json:"name"`
	Values []ProductAttributeValue `json:"values"`
}

// ProductAttributeValue represents a value of an attribute
type ProductAttributeValue struct {
	ID    int    `json:"id"`
	Value string `json:"value"`
}

// AttributeValueRow is used for scanning from database
type AttributeValueRow struct {
	AttributeID    int    `gorm:"column:attribute_id"`
	AttributeName  string `gorm:"column:attribute_name"`
	ValueID        int    `gorm:"column:value_id"`
	AttributeValue string `gorm:"column:attribute_value"`
}
