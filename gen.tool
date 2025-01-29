version: "0.1"
database:
  dsn : "root:123456@tcp(localhost:3306)/fan_club_api?charset=utf8mb4&parseTime=true&loc=Local"
  db  : "mysql"
  outPath :  "./internal/model"
  onlyModel : true
  modelPkgName : "model"
  fieldNullable : true
  # generate field with gorm index tag
  fieldWithIndexTag : true
  # generate field with gorm column type tag
  fieldWithTypeTag  : false