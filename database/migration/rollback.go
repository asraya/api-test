package migration

import (
	"fmt"
	"log"

	"api-test/database"
)

// RunRollback ..
func RunRollback() {

	db := database.PostsqlConn()

	if db.Error != nil {
		log.Fatalln(db.Error.Error())
	}

	if exist := db.HasTable("tag"); exist {
		fmt.Println("drop table tag")
		err := db.DropTable("tag")
		if err == nil {
			fmt.Println("success drop table tag")
		}
	}

}
