# go-to-bplate (WIP)

go-based web crawler/API for parsing UCLA dining data! lots more work to do, including documentation - come back alter to see more :)

## dev setup & basic use

Dev Setup:

```
$ go get github.com/PuerkitoBio/goquery
$ go build
```

Server Usage:

```
$ ./go-to-bplate -s
$  curl "http://localhost:4242/date/2019-10-09?filters=Vegetarian&xfilters=Gluten&keywords=Cheese"
{"Date":"2019-10-09","Time":"2019-10-10T00:33:05.878415-07:00","Breakfast":{"Title":"Breakfast Menu for Today, October 9, 2019","Items":[{"Name":"Cheddar Cheese","RecipeLink":"http://menu.dining.ucla.edu/Recipes/061013/1","Location":"De Neve","DietaryInfo":["Vegetarian","Dairy"]},{"Name":"Feta Cheese","RecipeLink":"http://menu.dining.ucla.edu/Recipes/132026/1","Location":"De Neve","DietaryInfo":["Vegetarian","Dairy"]},{"Name":"Cheddar Cheese","RecipeLink":"http://menu.dining.ucla.edu/Recipes/061013/1","Location":"Bruin Plate","DietaryInfo":["Vegetarian","Dairy"]},{"Name":"Cottage Cheese","RecipeLink":"http://menu.dining.ucla.edu/Recipes/971486/1","Location":"Bruin Plate","DietaryInfo":["Vegetarian","Dairy"]}]},"Lunch":{"Title":"Lunch Menu for Today, October 9, 2019","Items":[{"Name":"Feta Cheese","RecipeLink":"http://menu.dining.ucla.edu/Recipes/975321/1","Location":"Covel","DietaryInfo":["Vegetarian","Dairy","Halal"]}]},"Dinner":{"Title":"Dinner Menu for Today, October 9, 2019","Items":[]}}

```

CLI Usage:

```
$ ./go-to-bplate -keywords chicken -xfilters Gluten
==========
Breakfast Menu for Today, October 7, 2019
-------
Chicken Apple Sausage at De Neve (http://menu.dining.ucla.edu/Recipes/111187/1)
==========
Lunch Menu for Today, October 7, 2019
-------
Pomegranate Walnut Chicken Stew at Covel (Tree Nuts) (http://menu.dining.ucla.edu/Recipes/111337/3)
Grilled Chicken Breast at Covel (http://menu.dining.ucla.edu/Recipes/977242/1)
Lemon Garlic Chicken Thighs & Drumsticks at De Neve (http://menu.dining.ucla.edu/Recipes/111005/1)
Grilled Chicken Breast at De Neve (http://menu.dining.ucla.edu/Recipes/977242/1)
Grilled Chicken Breast - Sun Dried Tomato Marinade at Bruin Plate (http://menu.dining.ucla.edu/Recipes/111136/2)
==========
Dinner Menu for Today, October 7, 2019
-------
Grilled Chicken Breast at Covel (http://menu.dining.ucla.edu/Recipes/977242/1)
Lemon Garlic Chicken Thighs & Drumsticks at De Neve (http://menu.dining.ucla.edu/Recipes/111005/1)
Grilled Chicken Breast at De Neve (http://menu.dining.ucla.edu/Recipes/977242/1)
Grilled Chicken Cilantro Jalapeno Marinade at Bruin Plate (http://menu.dining.ucla.edu/Recipes/111291/2)
```
