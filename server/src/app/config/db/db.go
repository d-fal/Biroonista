package db

import (
	"log"
	"reflect"
	"time"

	"../../lib"
	"../../model"
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	_ "github.com/paulmach/orb"
)

func DbModel() *pg.DB {

	return DbConnect()
}

func CreateModels() {
	pgdb := DbConnect()
	defer pgdb.Close()
	for _, model_ := range []interface{}{
		&model.Providers{},
		&model.Contactstype{},
		&model.Contacts{},
		&model.ProvidersContacts{},
		&model.Attachments{},
		&model.Flows{},
		&model.FlowProcedures{},
		&model.Exchange{},
		&model.Users{},
		&model.UsersContacts{},
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
			Temp:          false,
			FKConstraints: true,
		})
		if err != nil {
			// log.Printf("Error with %s : %s", reflect.TypeOf(model_).Elem().Name(),err)

		} else {
			log.Printf("Model %s created!", reflect.TypeOf(model_).Elem().Name())
		}
	}

}

func DbConnect() *pg.DB {
	pgdb := pg.Connect(&pg.Options{
		User:     lib.DB_USER,
		Password: lib.DB_PASS,
		Database: lib.DB_NAME,
	})

	pgdb.OnQueryProcessed(func(event *pg.QueryProcessedEvent) {
		query, err := event.FormattedQuery()
		if err != nil {
			panic(err)
		}

		log.Printf("%s %s", time.Since(event.StartTime), query)
	})

	return pgdb
}
