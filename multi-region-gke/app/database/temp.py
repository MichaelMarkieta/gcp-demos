import uuid
import lorem
from lorem.text import TextLorem
import json

def seed_tables(cats, prods, opts):
    categories = []

    for i in range(0, cats):
       products = []
       cat_id = str(uuid.uuid4())
       categories.append({
           "CategoryId": cat_id,
           "CategoryName": TextLorem(srange=(1,3)).sentence(),
           "CategoryDescription": lorem.sentence()
       })

       for j in range(0, prods):
           options = []
           prod_id = str(uuid.uuid4())
           products.append({
               "ProductId": prod_id,
               "ProductName": TextLorem(srange=(1,3)).sentence(),
               "ProductDescription": lorem.sentence()
           })

           for k in range(0, opts):
               opt_id = str(uuid.uuid4())
               options.append({
                   "OptionId": opt_id,
                   "OptionName": TextLorem(srange=(1,3)).sentence(),
                   "OptionDescription": lorem.sentence(),
                   "OptionValue": TextLorem(srange=(1,3)).sentence()
               })
           products[j]["Options"] = options
       categories[i]["Products"] = products

    with open('seed_3.txt', 'w') as outfile:
        json.dump(categories, outfile)

if __name__ == "__main__":
    cats = 100
    prods = 100
    opts = 100
    seed_tables(cats, prods, opts)
