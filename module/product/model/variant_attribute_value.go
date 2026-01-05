package productmodel

type VariantAttributeValue struct {
	VariantID        int `gorm:"column:variant_id;not null"`
	AttributeValueID int `gorm:"column:attribute_value_id;not null"`
}

func (VariantAttributeValue) TableName() string {
	return "variant_attribute_values"
}
