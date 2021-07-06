# CHEATSHEET

## Creating first index

```bash
PUT localhost:9200/<your-index>
```

e.g. `PUT localhost:9200/books`

```json
{
  "settings": {
    "number_of_shards": 1,
    "number_of_replicas": 1
  },
  "mappings": {
    "properties": {
      "name": {
        "type": "text"
      },
      "author": {
        "type": "integer"
      },
      "resume": {
        "type": "float"
      }
    }
  }
}
```

## Creating a document in your index

```bash
POST localhost:9200/<your-index>/_doc
```

e.g. `PUT localhost:9200/books/_doc`

```json
{
  "name": "Martine veut dissoudre l'Assemblée Nationale, feat Jean Castête",
  "author": "Jeanne Oskour"
}
```

## Query the data

```bash
POST localhost:9200/_search
```

```json
{
  "query": {
    "multi_match": {
      "query": "Jean",
      "fields": ["name", "author"]
    }
  }
}
```
