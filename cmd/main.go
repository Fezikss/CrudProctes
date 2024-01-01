package main

import (
	"log"
	"main/config"
	"main/controller"
	"main/storage/postgres"
)

func main() {
	cfg := config.Load()

	store, err := postgres.New(cfg)
	if err != nil {
		log.Fatalln("error while connecting to db err : ", err.Error())
		return
	}
	defer store.Db.Close()
	con := controller.New(store)

	//con.CreateCategory()
	//con.GetCategoryById()
	//con.GetCategoryList()
	//con.UpdateCategory()
	//con.DeleteCategory()

	//con.CreateProduct()
	//con.GetProductByID()
	//con.GetProductList()
	//con.UpdateProduct()
	//con.DeleteProduct()

	con.TopCategory()

}
