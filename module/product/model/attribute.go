package productmodel

// Attribute represents a product attribute (e.g., Color, Size)
type Attribute struct {
	ID     int              `json:"id" gorm:"column:id;primaryKey"`
	Name   string           `json:"name" gorm:"column:name;type:varchar(255);not null"`
	Type   string           `json:"type" gorm:"column:type;type:varchar(50);default:'select'"`
	Values []AttributeValue `json:"values" gorm:"-"`
}

func (Attribute) TableName() string { return "attributes" }

// AttributeValue represents a value for an attribute
type AttributeValue struct {
	ID          int    `json:"id" gorm:"column:id;primaryKey"`
	AttributeID int    `json:"-" gorm:"column:attribute_id;not null"`
	Value       string `json:"value" gorm:"column:value;type:varchar(255);not null"`
}

func (AttributeValue) TableName() string { return "attribute_values" }

// CategoryAttribute links categories to attributes
type CategoryAttribute struct {
	CategoryID  int `gorm:"column:category_id;not null"`
	AttributeID int `gorm:"column:attribute_id;not null"`
}

func (CategoryAttribute) TableName() string { return "category_attributes" }

// CategoryAttributeWithValues for response (matches FE spec)
// {
//   "id": <attribute_id>,
//   "name": "<attribute_name>",
//   "values": [{"id": <attribute_value_id>, "value": "<value>"}]
// }
type CategoryAttributeWithValues struct {
	ID     int              `json:"id"`
	Name   string           `json:"name"`
	Values []AttributeValue `json:"values"`
}

// CategoryAttributeRow for database scanning
type CategoryAttributeRow struct {
	AttributeID    int    `gorm:"column:attribute_id"`
	AttributeName  string `gorm:"column:attribute_name"`
	ValueID        int    `gorm:"column:value_id"`
	AttributeValue string `gorm:"column:attribute_value"`
}
