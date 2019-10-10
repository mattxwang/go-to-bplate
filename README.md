# go-to-bplate (WIP)

`go-to-bplate` is a passion project I made to make the UCLA dining hall experience better. it's currently an API server that pulls data from the UCLA dining website and presents it in a more readable form, with caching, keyword search, and filtering by dietary needs (e.g. Vegan, Gluten free, etc.). it's very much a work in progress, so more's coming soon (including documentation)!

features right now:
* API server that pulls menu information by date; caches results in memory, performs keyword search and dietary filtering
* CLI that pulls menu information by date, performs keyword search and dietary filtering

roadmap:
- [x] parse UCLA dining website!
- [x] CLI
- [x] basic HTTP server
- [x] caching
- [ ] on-disk cache (JSON)
- [ ] deploy to Heroku!
- [ ] cache-flushing and auto-population on specific times
- [ ] cron job-like polling
- [ ] hooking up to a frontend (website, app, chrome extension)
- [ ] docs
- [ ] bot integration/webhooks
- [ ] real database?

and other things i need to do
* case-insensitive filtering and xfiltering
* polling take-out places
* make things more efficient
* explore options for nutrition facts/ingredients endpoints & data
* grab dining hall capacity data
* better search? (e.g. typo, plurals, word-analysis, ingredients)

## dev setup & basic use

setup environment

```
$ go get github.com/PuerkitoBio/goquery
$ go build
```

API server usage

```
$ ./go-to-bplate -s
$  curl "http://localhost:4242/date/2019-10-09?filters=Vegetarian&xfilters=Gluten&keywords=Cheese"
```

response format (JSON)

```json
{
    "Date": "2019-10-09",
    "Time": "2019-10-10T00:33:05.878415-07:00",
    "Breakfast": {
        "Title": "Breakfast Menu for Today, October 9, 2019",
        "Items": [{
            "Name": "Cheddar Cheese",
            "RecipeLink": "http://menu.dining.ucla.edu/Recipes/061013/1",
            "Location": "De Neve",
            "DietaryInfo": ["Vegetarian", "Dairy"]
        }, {
            "Name": "Feta Cheese",
            "RecipeLink": "http://menu.dining.ucla.edu/Recipes/132026/1",
            "Location": "De Neve",
            "DietaryInfo": ["Vegetarian", "Dairy"]
        }, {
            "Name": "Cheddar Cheese",
            "RecipeLink": "http://menu.dining.ucla.edu/Recipes/061013/1",
            "Location": "Bruin Plate",
            "DietaryInfo": ["Vegetarian", "Dairy"]
        }, {
            "Name": "Cottage Cheese",
            "RecipeLink": "http://menu.dining.ucla.edu/Recipes/971486/1",
            "Location": "Bruin Plate",
            "DietaryInfo": ["Vegetarian", "Dairy"]
        }]
    },
    "Lunch": {
        "Title": "Lunch Menu for Today, October 9, 2019",
        "Items": [{
            "Name": "Feta Cheese",
            "RecipeLink": "http://menu.dining.ucla.edu/Recipes/975321/1",
            "Location": "Covel",
            "DietaryInfo": ["Vegetarian", "Dairy", "Halal"]
        }]
    },
    "Dinner": {
        "Title": "Dinner Menu for Today, October 9, 2019",
        "Items": []
    }
}
```

CLI (not recommended)

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
