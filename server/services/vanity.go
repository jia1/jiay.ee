package services

import (
	"encoding/json"
	"fmt"
	"log"
	"vanity.go/database"
)

type Vanity struct {
	TargetURL string `json:"target_url,omitempty"`
	VanityURL string `json:"vanity_url,omitempty"`
}

type UserVanity struct {
	VanityURLs []Vanity `json:"vanity_urls,omitempty"`
}

func GetAllVanityURLs(userID string) UserVanity {
	_, redisJSONHandler := database.RedisClients()

	data := fmt.Sprint(redisJSONHandler.JSONGet(userID, "."))
	log.Print(data)

	// https://stackoverflow.com/questions/49006594/interface-to-byte-conversion-in-golang#comment124966799_49006594
	bytes := []byte(data)

	userVanity := UserVanity{}
	err := json.Unmarshal(bytes, &userVanity)
	if err != nil {
		log.Print("Failed to JSON Unmarshal")
		return UserVanity{VanityURLs: make([]Vanity, 0)}
	}

	return userVanity
}
