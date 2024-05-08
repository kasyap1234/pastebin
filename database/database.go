package database 
import (
	"context"
	"log"

	
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)
func GetMongoClient() *mongo.Client {
	return client
}
var client *mongo.Client 
func ConnectDB() *mongo.Client {
	
	clientOptions :=options.Client().ApplyURI("mongodb://localhost:80800")
	var err error 
	client ,err =mongo.Connect(context.Background(),clientOptions)
	if err != nil {
		log.Fatal(err); 

	}
	return client; 
}
func FindAll(collection *mongo.Collection) ([]interface{},error) {
	var results []interface {}; 
	cursor, err := collection.Find(context.Background(),bson.M{}); 
	if err !=nil {
		log.Fatal("Error while finding all documents ", err);
		return nil ,err; 
	}
	err =cursor.All(context.Background(),&results)
	if err !=nil {
		log.Fatal("Error while finding all documents ", err);
		return nil , err; 

	}
	return results, nil ; 
}
func FindOneByID(collection *mongo.Collection, filter interface {}) (interface {},error ){
	var result interface {}; 
	err :=collection.FindOne(context.Background(),filter). Decode(&result)
    if err !=nil {
		log.Fatal("Error while finding the document with the requried filter ", err ) 
		return nil , err ; 

	}
	return result, nil ; 


}
func InsertOne(collection *mongo.Collection, document interface {}) error {
 
 
	_,err := collection.InsertOne(context.Background(),document); 
	if err !=nil {
		log.Fatal("Error while inserting the document ", err )
		return err; 
	}
	return nil ; 
}

func UpdateOne(collection *mongo.Collection , filter interface {},update interface {})error{
	_,err := collection.UpdateOne(context.Background(),filter ,update);
	if err !=nil {
		log.Fatal(err); 

	}
	return err; 
	

}
func DeleteOne(collection *mongo.Collection , filter interface {})error {

	_, err := collection.DeleteOne(context.Background(),filter); 
	if err !=nil {
		log.Fatal(err); 
		return err ; 
	}
	return err; 

}





