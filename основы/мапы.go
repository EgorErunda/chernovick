package main

//Хэш-таблица — структура данных,
//которая позволяет быстро и эффективно хранить и получать данные по ключу.

//Мапы в Go создаются с помощью функции make
var m map[keyType]valueType

//keyType и valueType — типы данных ключа и значения соответственно

//мапa со строками в качестве ключей и целыми числами в качестве значений
var m map[string]int
m = make(map[string]int)

//Также можно создать мапу с помощью литералa
m := map[keyType]valueType{
    key1: value1,
    key2: value2,
    // ...
}

