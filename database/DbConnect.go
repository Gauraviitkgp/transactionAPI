package database

import (
	"PBL_Proj/models"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type (
	IDbConn interface {
		Connect()
		GetData()
		DeleteData()
		PushData()
		GetDataUsingPK(primaryKey string)
	}
)

type DbConn struct {
	connString string
	dataObj    *mongo.Collection
	client     *mongo.Client
}

var DbClient DbConn = DbConn{}

func (x *DbConn) GetData() {

	cursor, err := x.dataObj.Find(
		context.Background(),
		bson.D{},
	)
	if err != nil {
		println(err)
	}
	models.Transactions = nil

	for cursor.Next(context.TODO()) {
		elem := &models.Transaction{}
		k := &bson.D{}
		cursor.Decode(k)
		if err := cursor.Decode(elem); err != nil {
			log.Fatal(err)
		}
		elem.TransactionId = uint32(k.Map()["_id"].(int64))
		models.Transactions = append(models.Transactions, *elem)
	}
}

func (x *DbConn) PushData(transaction models.Transaction) error {
	_, err := x.dataObj.InsertOne(
		context.TODO(),
		bson.D{{"_id", transaction.TransactionId},
			{"timeStamp", transaction.TimeStamp},
			{"isCredit", transaction.IsCredit},
			{"accountId", transaction.AccountId},
			{"transactionType", transaction.TransactionType},
			{"purchaseType", transaction.PurchaseType},
			{"amount", transaction.Amount},
		})
	return err
}

func (x *DbConn) DeleteData(tranId uint64) error {
	_, err := x.dataObj.DeleteOne(context.TODO(), bson.D{{"_id", tranId}})

	return err
}

func (x *DbConn) UpdateData(tranId uint64, transaction models.Transaction) error {
	_, err := x.dataObj.DeleteOne(context.TODO(), bson.D{{"_id", tranId}})
	if err != nil {
		return err
	}
	err = x.PushData(transaction)
	return err
}

func (x *DbConn) Connect() {
	ctx := context.TODO()
	pathToCertificate := "/Users/suryawanshi.gaurav/Downloads/X509-cert-4058268142071853748.pem"
	x.connString = "mongodb+srv://mfcustomer.kqrmo.mongodb.net/MFMoney?authSource=%24external&authMechanism=MONGODB-X509&retryWrites=true&w=majority&tlsCertificateKeyFile=" + pathToCertificate
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().
		ApplyURI(x.connString).
		SetServerAPIOptions(serverAPIOptions)
	var err error
	x.client, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	x.dataObj = x.client.Database("MFMoney").Collection("MFTransactions")
	_, err = x.dataObj.CountDocuments(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	err = x.client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
}

func main() {}

func callDb() {

}
