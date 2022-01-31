package helpers

import "go.mongodb.org/mongo-driver/bson/primitive"

func UniqueObjectID(intSlice []primitive.ObjectID) []primitive.ObjectID {
	keys := make(map[primitive.ObjectID]bool)
	var list []primitive.ObjectID
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
