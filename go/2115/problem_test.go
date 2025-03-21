package problem

import (
	"reflect"
	"testing"
)

/*
2115. Find All Possible Recipes from Given Supplies
Medium
Topics
Companies
Hint

You have information about n different recipes. You are given a string array recipes and a 2D string array ingredients. The ith recipe has the name recipes[i], and you can create it if you have all the needed ingredients from ingredients[i]. A recipe can also be an ingredient for other recipes, i.e., ingredients[i] may contain a string that is in recipes.

You are also given a string array supplies containing all the ingredients that you initially have, and you have an infinite supply of all of them.

Return a list of all the recipes that you can create. You may return the answer in any order.

Note that two recipes may contain each other in their ingredients.
*/ 


func TestProblem(t *testing.T) {

  testCases := []struct {
    name string 
    recipes []string 
    ingredients [][]string 
    supplies []string
    want []string
  } {
    {
      "test1",
      []string{"bread"},
      [][]string{{"yeast", "flour"}},
      []string{"yeast", "flour", "corn"},
      []string{"bread"},
    },
    {
      "test2",
      []string{"bread", "sandwich"},
      [][]string{{"yeast", "flour"},{"bread","meat"}},
      []string{"yeast", "flour", "meat"},
      []string{"bread", "sandwich"},
    },
    {
      "test3",
      []string{"bread", "sandwich", "burger"},
      [][]string{{"yeast", "flour"},{"bread","meat"},{"sandwich", "meat", "bread"}},
      []string{"yeast", "flour", "meat"},
      []string{"bread", "sandwich", "burger"},
    },
  }

  for _, tc := range testCases {
    t.Run(tc.name, func(t *testing.T) {
      got := findAllRecipes(tc.recipes, tc.ingredients, tc.supplies)
      ok := reflect.DeepEqual(got, tc.want)
      if !ok {
        t.Errorf("findAllRecipes(%v, %v, %v) = %v; want = %v\n",
          tc.recipes, tc.ingredients, tc.supplies, got, tc.want)
      }
    })
  }
}
