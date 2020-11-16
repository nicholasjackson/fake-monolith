# Fake Monolith

Simple Fake Monolith service, does not really do anything useful

## Handlers

### /products GET
Return an array of fake Product

```
curl localhost:9090/products
```

```json
 [ 
  {
    "ID": 397,
    "Name": "brunch",
    "ShortDescription": "Knausgaard umami vice farm-to-table bespoke blog tattooed skateboard put a bird on it seitan.",
    "LongDescription": "Line hide owe sugar determine construct afford lord rest treaty rugby hall meet pool emphasis check knock withdraw pollution session finance doubt talk opposition country end captain promise breath inform. Persuade distinction complete corporation committee react tell pull legislation study king hide distinction brown bring walk island wait design decline cross divide rise care return survey knock become minority welcome. Message variety form day conversation movement labour variety relate lee sum challenge line secure cause expression grant ball separate wan include assist let support sound emphasis bird climb explanation maintain. Martin john decision sleep ensure group northern worry pressure mind empire decade autumn say trip 're maintenance rail context imply survey establish fuel performance discipline creation cross impose enter play.",
    "PrimaryColor": "DarkOliveGreen"
  },
  {
    "ID": 621,
    "Name": "fixie",
    "ShortDescription": "Ethical sustainable meh selvage deep v bushwick waistcoat put a bird on it stumptown swag.",
    "LongDescription": "Explanation theory seem pain kill need test excuse advise eat motion shake waste attract spread piece specialist order sister afternoon intend debate receive close shut apply struggle sector generate pass. Soil contain research undertake expand choice conclude can wear distinguish recognize transfer replace explain lewis contrast transport pair start measure impact know influence attempt solve help star recognise object finish. Version address outcome source copy chairman stone research conference promote play appear note discuss hit night alter bet distinction account transport knock option activity risk mark formation seek strike observe. Wonder suit suggest cost launch remove detail deputy dance set flight travel procedure lead pick order school artist passage date office week notice study turn war income imagine collection base.",
    "PrimaryColor": "LightCoral"
  }
]
```

### /products/{id} GET
Return a single product referenced by it's ID

```
curl localhost:9090/products/397
```

```json
{
    "ID": 397,
    "Name": "brunch",
    "ShortDescription": "Knausgaard umami vice farm-to-table bespoke blog tattooed skateboard put a bird on it seitan.",
    "LongDescription": "Line hide owe sugar determine construct afford lord rest treaty rugby hall meet pool emphasis check knock withdraw pollution session finance doubt talk opposition country end captain promise breath inform. Persuade distinction complete corporation committee react tell pull legislation study king hide distinction brown bring walk island wait design decline cross divide rise care return survey knock become minority welcome. Message variety form day conversation movement labour variety relate lee sum challenge line secure cause expression grant ball separate wan include assist let support sound emphasis bird climb explanation maintain. Martin john decision sleep ensure group northern worry pressure mind empire decade autumn say trip 're maintenance rail context imply survey establish fuel performance discipline creation cross impose enter play.",
    "PrimaryColor": "DarkOliveGreen"
  }
```
