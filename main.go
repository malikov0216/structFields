package main

import (
	"fmt"
	"reflect"
	"testing"
)

/* Assign Функция присавивания одинаковых полей
Задача реализовать функцию которая сравнивает две структуры,
находит одинаковые по названию и типу поля и из источника sour
присваевает эти поля в получателя dist.
Также возвращает количество совпадающих нахваний полей с типами.

Если будет возможность то постараться пройтись по всем уровням вложенности
*/
type A struct {
	A string
	B uint
	C string
}

type B struct {
	AA string
	B  int
	C  string
}

func structNameTypes (StructValue interface{}) ([]interface{}, []string, []reflect.Value) {

		objectStruct := reflect.ValueOf(StructValue).Elem()
		fieldTypes := make([]interface{}, objectStruct.NumField())
		fieldNames := make([]string, objectStruct.NumField())
		fieldValues := make([]reflect.Value, objectStruct.NumField())

		for i:= 0; i < objectStruct.NumField(); i++ {
			fieldTypes[i] = objectStruct.Field(i).Type()
			fieldNames[i] = objectStruct.Type().Field(i).Name
			fieldValues[i] = objectStruct.Field(i)
		}
		return fieldTypes, fieldNames, fieldValues
}

func Assign(sour interface{}, dist interface{}) uint {

	sourTypes, sourNames, sourValues := structNameTypes(sour)
	distTypes, distNames, _ := structNameTypes(dist)
	var result uint = 0
	ps := reflect.ValueOf(dist)
	// struct
	s := ps.Elem()
	fmt.Println(s.Kind())

	for i := range sourTypes {
		for j := range distTypes {
			if sourTypes[i] == distTypes[j] && sourNames[i] == distNames[j] {
				result++
				if s.Kind() == reflect.Struct {
					f := s.FieldByName(distNames[j])
					if f.IsValid() {
						if f.CanSet() {
							if f.Kind() == reflect.String {
								f.SetString(sourValues[j].String())
							} else if f.Kind() == reflect.Int {
								f.SetInt(sourValues[j].Int())
							} else if f.Kind() == reflect.Bool {
								f.SetBool(sourValues[j].Bool())
							} else {
								fmt.Println("Invalid Data Type")
							}
						}
					}
				}
			}
		}
	}

	return result
}

// TestAssign тестирование функции Assign
func TestAssign(t *testing.T) {
	type A struct {
		A string
		B uint
		C string
	}
	type B struct {
		AA string
		B  int
		C  string
	}
	var (
		a = A{
			A: "Тест A",
			B: 55,
			C: "Test C",
		}
		b = B{
			AA: "OKOK",
			B:  10,
			C:  "FAFA",
		}
	)
	result := Assign(a, b)
	switch true {
	case b.B != 10:
		t.Errorf("b.B = %d; необходимо 10", b.B)
	case b.C != "Test C":
		t.Errorf("b.C = %v; необходимо 'Test C'", b.C)
	case result != 1:
		t.Errorf("Assign(a,b) = %d; необходимо 1", result)
	}
}

func main () {
	//var t testing.T
	//TestAssign(&t)
	var a = A{"a", 2, "cdcsdcs"}
	var b = B{"w", 5, "l"}
	Assign(&a, &b)
	fmt.Println("Result value",b.C)
}