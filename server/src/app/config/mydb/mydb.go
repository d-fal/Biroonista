package mydb

import(
	
	"fmt"
	"time"
    _ "github.com/lib/pq"
	"github.com/go-redis/redis"
	"github.com/go-pg/pg/orm"
	"github.com/go-pg/pg"
	"../../model"
	"log"
	"reflect"
)

const (
    dbhost = "localhost"
    dbport = "5432"
    dbuser = "postgres"
    dbpass = "0009887427"
    dbname = "bst"
)
var client *redis.Client
var pgdb *pg.DB


func CreateModels() {

	for _, model_ := range []interface{}{
			&model.Providers{},
			&model.Contactstype{},
			&model.Contacts{},
			&model.ProvidersContacts{},
			&model.Attachments{},
			&model.Flows{},
			&model.FlowProcedures{},
			&model.Exchange{},
			&model.User{},
			&model.UserContact{}, 
			&model.Membership{}, 
			&model.Events{},
			&model.Adventures{},
			&model.InvitationAux{},
			&model.ProvidersServices{},

			&model.Services{},
			&model.ProvidersServicesAttachments{}, 
			&model.EventsInvitations{},
			&model.Invitations{},
			&model.EventsUsers{},
			&model.UsersEventsAdventures{},
			&model.Voteoptions{},
			&model.Attachments{},
			&model.ProvidersAttachments{},
			&model.UsersEventsAdventuresVoteoptions{},
			&model.ProvidersAdventures{},
			} {
		
		err := pgdb.CreateTable(model_, &orm.CreateTableOptions{
			Temp: false,
			FKConstraints: true,
		})
		if err!=nil {
			log.Printf("Error with %s : %s", reflect.TypeOf(model_).Elem().Name(),err)
			
		}else{
			log.Printf("Model %s created!", reflect.TypeOf(model_).Elem().Name())
		}
	}


}
func TestSelect() {
	user2 := model.User{}
	_ , err := pgdb.Query(&user2 , `select * from users`)
	if err!=nil {
		panic("Cannot run query")
	}
	fmt.Println(user2)

}

func ExampleInsert() {
	user1 := model.User {
		FirstName : "Alireza",
		LastName: "baghban",
		UserName: "Alireza_Baghban",
	}
	err := pgdb.Insert(&user1)
	if err != nil {
		fmt.Println(err)
	}
	event1 := model.Events {
		Subject: "Party",
		UserID: user1.GetId(),
		// User: &user1,
		CreatedAt: time.Now(),
		
	}

	err = pgdb.Insert(&event1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Insert was successful")
}
func PgORMTest(){
	pgdb = pg.Connect(&pg.Options{
		User: "postgres",
		Password: "0009887427",
		Database: "bst",
	})
	

	

}

func CloseDb() {
	defer pgdb.Close()
}

func redisTest() {
    client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	// Output: PONG <nil>
}

func ExampleClient() {
	err := client.Set("key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := client.Get("key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := client.Get("key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
	// Output: key value
	// key2 does not exist
}