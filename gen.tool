version: "0.1"
database:
  dsn : "root:123456@tcp(localhost:3306)/auth?charset=utf8mb4&parseTime=true&loc=Local"
  db  : "mysql"
  outPath : "./internal/model"
  onlyModel : true
  modelPkgName : "model"
  fieldNullable : true
  fieldWithIndexTag : true
  fieldWithTypeTag : false
  fieldSignable : true