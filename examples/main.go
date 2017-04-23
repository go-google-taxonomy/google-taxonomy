package main

import (
	"fmt"

	"github.com/xreception/google-taxonomy-go/taxonomy"
)

func main() {
	tx, err := taxonomy.NewTaxonomy([]string{})
	if err != nil {
		panic(err)
	}
	infs, err := tx.GetRootsCategoryInfo(taxonomy.KeyLanguage)
	if err != nil {
		panic(err)
	}
	for _, inf := range infs {
		fmt.Printf("%d - %s\n", inf.ID, inf.String())
	}
	inf, err := tx.GetCategoryInfo(16, taxonomy.KeyLanguage)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%d - %s\n", inf.ID, inf.String())
}