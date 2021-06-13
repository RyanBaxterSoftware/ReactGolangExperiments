package models
import "go.mongodb.org/mongo-driver/bson/primitive"
type FoodList struct {
 
  ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
  ItemName   string             `json:"itemname,omitempty"`
  Status bool               `json:"status,omitempty"`
}