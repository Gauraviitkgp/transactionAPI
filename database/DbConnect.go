package database

type (
	IDbConn interface {
		Connect()
		GetData()
		PutData()
		PushData()
		GetDataUsingPK(primaryKey string)
	}
)

type DbConn struct {
	connString string
	tempData
}

func (x DbConn) GetData

func main() {}

func callDb() {

}
