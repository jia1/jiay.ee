package services

import (
	"encoding/json"
	"fmt"
	"log"
	"vanity.go/database"
)

type Vanity struct {
	TargetURL  string `json:"target_url,omitempty"`
	VanityPath string `json:"vanity_path,omitempty"`
}

type UserVanity struct {
	VanityPaths []string `json:"vanity_paths,omitempty"`
}

func GetUserVanity(userID string) UserVanity {
	_, redisJSONHandler := database.RedisClients()

	data := fmt.Sprint(redisJSONHandler.JSONGet(userID, "."))
	log.Print(data)

	// https://stackoverflow.com/questions/49006594/interface-to-byte-conversion-in-golang#comment124966799_49006594
	bytes := []byte(data)
	log.Print(bytes)

	userVanity := UserVanity{}
	err := json.Unmarshal(bytes, &userVanity) // FIXME: Figure this out
	if err != nil {
		log.Print("Failed to JSON Unmarshal")
		return UserVanity{VanityPaths: make([]string, 0)}
	}

	return userVanity
}

func CreateVanity(userID string, vanity Vanity) bool {
	_, redisJSONHandler := database.RedisClients()

	existingUserVanity := GetUserVanity(userID)
	vanityPaths := existingUserVanity.VanityPaths
	updatedVanityPaths := append(vanityPaths, vanity.VanityPath)

	res, err := redisJSONHandler.JSONSet(userID, ".", updatedVanityPaths)
	if err != nil {
		log.Fatalf("Failed to JSONSet")
		return false
	}
	if res.(string) != "OK" {
		fmt.Println("Failed to Set: ")
		return false
	}
	fmt.Printf("Success: %s\n", res)
	return true
}
