package hashtable

import (
	"fmt"
	//"log"
	"strconv"
)

type Pair struct {
	key   int
	value string
}

func (p Pair) New(key int, value string) Pair {
	return Pair{key, value}
}

var table [][]Pair
var capacity int = 16
var size int = 0
var loadfactor float32 = 0.75

func Init(_capacity int) {

	capacity = _capacity
	for i := 0; i < capacity; i++ {
		table = append(table, []Pair{})
	}
}

func Clear() bool {
	if len(table) != 0 {
		for i := 0; i < len(table); i++ {
			table[i] = []Pair{}
		}
		size = 0
	}
	return true

}

func Put(pair Pair) string {

	//fmt.Printf("Put pair: %+v\n", pair)
	hashKey := hashcode(pair.key)
	//log.Printf("Capacity: %d, HashKey: %d\n", capacity, hashKey)
	if hashKey < capacity {
		for i := 0; i < len(table[hashKey]); i++ {
			if table[hashKey][i].key == pair.key {
				if table[hashKey][i].value != pair.value {
					tempValue := table[hashKey][i].value
					table[hashKey][i].value = pair.value
					return tempValue
				} else {
					return pair.value
				}
			}
		}

		table[hashKey] = append(table[hashKey], pair)
		size++
		return pair.value
	}
	return "null"
}

func Remove(key int) string {

	hashKey := hashcode(key)
	if hashKey < capacity {
		for i := 0; i < len(table[hashKey]); i++ {
			if table[hashKey][i].key == key {
				tempValue := table[hashKey][i].value
				table[hashKey] = append(table[hashKey][:i], table[hashKey][i+1:]...)
				size--
				return tempValue
			}
		}
	}

	return ""
}

func Get(key int) string {

	hashKey := hashcode(key)
	if hashKey < capacity {
		for i := 0; i < len(table[hashKey]); i++ {
			if table[hashKey][i].key == key {
				return table[hashKey][i].value
			}
		}
	}

	return "null"
}

func ContainsKey(key int) bool {

	hashKey := hashcode(key)
	if hashKey < capacity {
		for i := 0; i < len(table[hashKey]); i++ {
			if table[hashKey][i].key == key {
				return true
			}
		}
	}

	return false
}

func ContainsValue(value string) bool {

	for i := 0; i < capacity; i++ {
		for j := 0; j < len(table[i]); j++ {
			if table[i][j].value == value {
				return true
			}
		}
	}

	return false
}

func IsEmpty() bool {

	return size == 0
}

func Size() int {
	return size
}

func Print() {
	fmt.Printf("%+v\n", table)
}

func ToString() string {

	var sTable string = "["

	for i := range table {
		sTable += " ["
		for j := range table[i] {
			sTable += "{" + strconv.Itoa(table[i][j].key) + ", " + table[i][j].value + "} "
		}

		sTable += "]"
	}

	sTable += " ]"

	return sTable
}

func hashcode(key int) int {

	return key % capacity
}
