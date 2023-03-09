package postgres

import (
	"database/sql"
	"fmt"
	pb "grpc_todo/genproto/store"
	"log"

	"github.com/lib/pq"
)

const (
	PostgresHost     = "localhost"
	PostgresPort     = 5432
	PostgresUser     = "ismoiljon12"
	PostgresPassword = "12"
	PostgresDB       = "store_grpc"
)

func ConnnToDB() *sql.DB {
	conn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", PostgresHost, PostgresPort, PostgresUser, PostgresPassword, PostgresDB)

	db, err := sql.Open("postgres", conn)
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}

	return db

}

func CreateStore(store *pb.Store) (*pb.Store, error) {
	db := ConnnToDB()

	defer db.Close()
	newStore := pb.Store{}
	err := db.QueryRow("INSERT INTO stores(name,description,addresses,is_open) VALUES($1,$2,$3,$4) RETURNING id,name,description,addresses,is_open", store.Name, store.Description, pq.Array(store.Addresses), store.IsOpen).Scan(
		&newStore.Id,
		&newStore.Name,
		&newStore.Description,
		pq.Array(&newStore.Addresses),
		&newStore.IsOpen,
	)
	if err != nil {
		log.Fatalf("Failed to insert store: %v", err)
	}

	return &newStore, nil
}

func UpdateStore(store *pb.Store) error {
	db := ConnnToDB()
	defer db.Close()
	updateStore := pb.Store{
		Id:          2,
		Name:        "Store update",
		Description: "Update description",
		IsOpen:      false,
		Addresses: []string{
			"update_address-1",
			"update-address-2",
		},
	}
	_, err := db.Exec("UPDATE stores SET  name=$1,description=$2, addresses=$3,is_open=$4 WHERE id=$5", updateStore.Name, updateStore.Description, pq.Array(updateStore.Addresses), updateStore.IsOpen, updateStore.Id)
	if err != nil {
		log.Fatalf("Failed to update store: %v", err)
	}

	return nil
}

func DeleteStore(id int64) error {
	db := ConnnToDB()
	defer db.Close()
	_, err := db.Exec("DELETE FROM stores WHERE id=$1", id)
	if err != nil {
		log.Fatalf("Failed to delete store: %v", err)
	}

	return nil
}

func GetStore(id int64) (*pb.Store, error) {
	db := ConnnToDB()
	defer db.Close()

	store := pb.Store{}
	err := db.QueryRow("SELECT id,name, description,addresses,is_open FROM stores WHERE id=$1", id).Scan(&store.Id, &store.Name, &store.Description, pq.Array(&store.Addresses), &store.IsOpen)
	if err != nil {
		log.Fatalf("Failed to select stores: %v", err)
	}
	return &store, nil
}


