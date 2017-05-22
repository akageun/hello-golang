package main

import "fmt"

/**
 * RESULT

mp1 :  map[key2:8 key1:10]
map :  map[key1:10 key2:5]
key1 :  10
map :  map[key2:5]
key1 is not exist
map :  5 true
new map: map[hello:1 world:2]
snsCode :  map[FB:FaceBook GGP:Google Plus TWT:TWITTER]
key :  GGP  value :  Google Plus
key :  TWT  value :  TWITTER
key :  FB  value :  FaceBook

 */
func main() {
	var mp1 map[string]int //Map is a reference type. This is 'nil'

	mp1 = make(map[string]int)

	mp1["key1"] = 10
	mp1["key2"] = 8
	fmt.Println("mp1 : ", mp1)

	mp := make(map[string]int)

	mp["key1"] = 10
	mp["key2"] = 5

	fmt.Println("map : ", mp)

	key1 := mp["key1"]
	fmt.Println("key1 : ", key1);

	delete(mp, "key1") //remove key1
	fmt.Println("map : ", mp)

	_, exists := mp["key1"]
	if exists == false {
		fmt.Println("key1 is not exist")
	}

	value, exists := mp["key2"]
	fmt.Println("map : ", value, exists)

	n := map[string]int{"hello": 1, "world": 2}
	fmt.Println("new map:", n)

	//literal
	snsCode := map[string]string{
		"GGP": "Google Plus",
		"TWT": "TWITTER",
		"FB":   "FaceBook",
	}
	fmt.Println("snsCode : ", snsCode)

	for key, value := range snsCode {
		fmt.Println("key : ", key, " value : ", value)
	}
}
