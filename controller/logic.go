package controller

import (
	"fmt"
)

func (c Controller) TopCategory() {

	productList, err := c.Store.ProductStorege.GetList()
	if err != nil {
		fmt.Println("error while getting product list in top logic inside controller err : ", err.Error())
		return
	}
	categoryList, err := c.Store.CategoryStorege.GetList()
	if err != nil {
		fmt.Println("error while getting category list in top logic inside controller err : ", err.Error())
		return
	}
	//maxName := ""
	maxSum := 0
	caunterMap := make(map[string]int)

	for _, product := range productList {
		for _, category := range categoryList {
			if product.CategoryID == category.Id {
				caunterMap[category.Name] += product.Price
			}

		}
	}
	name := ""
	for i, v := range caunterMap {
		if v > maxSum {
			maxSum = v
			name = i
		}
	}
	fmt.Println(name, maxSum)

}
