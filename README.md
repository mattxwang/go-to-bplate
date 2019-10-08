# go-to-bplate (WIP)

go-based web crawler/API for parsing UCLA dining data! lots more work to do, including documentation - come back alter to see more :)

## dev setup & basic use

```
$ go get github.com/PuerkitoBio/goquery
$ go build
...
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
