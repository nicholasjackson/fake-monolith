handler {
  method = "GET"
  path = "/{id: [0-9]}"
  error_rate = ""

  payload {

    field "id" {
      type = "int"
      rand = "int"
    }
    
    field "name" {
      type = "string"
      rand = "hipsterword"
    }

  }

}
