package main

import (
	"fmt"
	"generics/array"
	"math/rand"
	"time"
)

// import "fmt"

// type Array[T any] struct {
// 	values []T
// }

// func NewArray[T any](values []T) *Array[T] {
// 	return &Array[T]{
// 		values,
// 	}
// }

// func (a Array[T]) ForEach(callback func(T)) {
// 	for _, v := range a.values {
// 		callback(v)
// 	}
// }

// func (a *Array[T]) Map(callback func(T) T) *Array[T] {
// 	for i, v := range a.values {
// 		value := callback(v)
// 		a.values[i] = value
// 	}
// 	return a
// }

// func (a *Array[T]) Filter(callback func(T) bool) *Array[T] {
// 	var newArray []T
// 	for _, v := range a.values {
// 		if callback(v) {
// 			newArray = append(newArray, v)
// 		}
// 	}
// 	a.values = newArray
// 	return a
// }

// func (a *Array[T]) Concat(arrays ...[]T) *Array[T] {
// 	for _, v := range arrays {
// 		a.Push(v...)
// 	}
// 	return a
// }

// func (a *Array[T]) Push(values ...T) *Array[T] {
// 	for _, v := range values {
// 		a.values = append(a.values, v)
// 	}
// 	return a
// }

type Gender string

const (
	Male   Gender = "Male"
	Female        = "Female"
)

type User struct {
	Name   string
	Age    int
	Gender Gender
}

// func main() {
// 	array := NewArray([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
// 	users := NewArray([]User{
// 		{
// 			Name:   "Kevin",
// 			Age:    21,
// 			Gender: Male,
// 		},
// 		{
// 			Name:   "Juan",
// 			Age:    22,
// 			Gender: Male,
// 		},
// 		{
// 			Name:   "Juana",
// 			Age:    12,
// 			Gender: Female,
// 		},
// 		{
// 			Name:   "Anggie",
// 			Age:    22,
// 			Gender: Female,
// 		},
// 	})
// 	// strings := NewArray([]string{"hello", "world"})
// 	mayoresDeEdad := users.
// 		Filter(func(u User) bool { return u.Age > 18 }).
// 		Filter(func(u User) bool { return u.Gender == Female })
// 	fmt.Println("mayoresDeEdad", mayoresDeEdad)

// 	fmt.Println(array.values)
// 	array.Map(func(i int) int {
// 		return i + 1
// 	}).Map(func(i int) int {
// 		return i + 2
// 	}).Filter(func(i int) bool {
// 		return i%2 == 0
// 	}).Filter(func(i int) bool {
// 		return i%3 == 0
// 	})
// 	fmt.Println(array.values)

// 	n := NewArray([]int{1, 2})
// 	fmt.Println(n.Concat([]int{3, 4}, []int{5, 6}))
// 	fmt.Println(n.Push(7, 8))
// }

func main() {
	var users []User
	// users := []User{
	// 	{
	// 		Name:   "Kevin",
	// 		Age:    21,
	// 		Gender: Male,
	// 	},
	// 	{
	// 		Name:   "Juan",
	// 		Age:    22,
	// 		Gender: Male,
	// 	},
	// 	{
	// 		Name:   "Juana",
	// 		Age:    12,
	// 		Gender: Female,
	// 	},
	// 	{
	// 		Name:   "Anggie",
	// 		Age:    22,
	// 		Gender: Female,
	// 	},
	// }
	for i := 0; i < 600000; i++ {
		random := rand.Int()
		name := fmt.Sprintf("Nombre random %d", random)
		users = append(users, User{
			Name:   name,
			Age:    random,
			Gender: Male,
		})
	}
	now := time.Now()

	// arrayNoConcurrent := array.NewArrayNoConcurrent(users)
	// arrayNoConcurrent.Map(func(u User) User {
	// 	// time.Sleep(1 * time.Second)
	// 	u.Name = "Kevin"
	// 	return u
	// })
	// fmt.Println("Time for no concurrent array", float64(time.Since(now)))

	concurrentArray := array.NewArrayConcurrent(users)
	concurrentArray.Map(func(u User) User {
		// time.Sleep(1 * time.Second)
		u.Name = "Kevin concurrente"
		return u
	})
	fmt.Println("Time for concurrent array", float64(time.Since(now)))
}
