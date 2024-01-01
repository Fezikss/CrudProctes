package controller

import (
	"fmt"
	"github.com/google/uuid"
	"main/models"
)

func (c Controller) CreateCategory() {
	name := ""
	fmt.Print(" Enter Category name  : ", name)
	fmt.Scan(&name)

	ctg := models.Category{
		Name: name,
	}

	id, err := c.Store.CategoryStorege.Insert(ctg)
	if err != nil {
		fmt.Println("error while creating category inside controller err : ", err.Error())
	}

	fmt.Println("Id : ", id)

}

func (c Controller) GetCategoryById() {
	idStr := ""
	fmt.Print("Id : ")
	fmt.Scan(&idStr)

	id, err := uuid.Parse(idStr)
	if err != nil {
		fmt.Println("id type is not uuid !")
		return
	}

	category, err := c.Store.CategoryStorege.GetById(id)
	if err != nil {
		fmt.Println("error while geting by id category inside controller err : ", err.Error())
		return
	}

	fmt.Println(category)

}

func (c Controller) GetCategoryList() {
	categories, err := c.Store.CategoryStorege.GetList()
	if err != nil {
		fmt.Println("error while getting category list inside controller err : ", err.Error())
		return
	}

	fmt.Println(categories)

}

func (c Controller) UpdateCategory() {
	var (
		idStr string
		name  string
	)
	fmt.Print("Enter category id : ")
	fmt.Scan(&idStr)
	fmt.Print("Enter category name : ")
	fmt.Scan(&name)
	id, err := uuid.Parse(idStr)
	if err != nil {
		fmt.Println("It is not uuid type !")
		return
	}
	category := models.Category{
		Id:   id,
		Name: name,
	}
	if err := c.Store.CategoryStorege.Update(category); err != nil {
		fmt.Println("error while updating data inside controller err : ", err.Error())
		return
	}

	fmt.Println("Data successfully updated !")
}

func (c Controller) DeleteCategory() {
	idStr := ""
	fmt.Print("Enter category id : ", idStr)
	fmt.Scan(&idStr)

	id, err := uuid.Parse(idStr)
	if err != nil {
		fmt.Println("It is not uuid Type !")
		return
	}

	if err := c.Store.CategoryStorege.Delete(id); err != nil {
		fmt.Println("error while deleting data inside controller err : ", err.Error())
		return
	}

	fmt.Println("Data successfully deleted !")
}
