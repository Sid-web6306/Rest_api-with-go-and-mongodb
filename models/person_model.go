package Person
import
(
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)
type Person struct {
	ID          primitive.ObjectID 	`json:"_id,omitempty" bson:"_id,omitempty"`
	Name 		string 				`json:"Name" bson:"name"`
	Address 	string				`json:"Address bson:"Address"`
	CreatedAt 	time.Time			`json:"CreatedAt" bson:"CreatedAt"`
	UpdatedAt 	time.Time 			`json:"UpdatedAt" bson:"UpdatedAt"`
}