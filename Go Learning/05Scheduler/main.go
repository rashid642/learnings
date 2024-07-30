package main

import (
    "fmt"
    "log"
    "math/rand"
    "time"

    "github.com/robfig/cron/v3"
)

// Token represents a token with a creation time
type Token struct {
    ID        int
    CreatedAt time.Time
    Key       string
}

// generateNewKey generates a new API key (this is a placeholder)
func generateNewKey() string {
    return fmt.Sprintf("new-api-key-%d", rand.Intn(10000)) // Implement your key generation logic here
}

// checkAndUpdateTokens checks each token and updates if necessary
func checkAndUpdateTokens(tokens []Token) {
    sixMonthsAgo := time.Now().AddDate(0, -6, 0)
	fmt.Println("Checking....")
    for i, token := range tokens {
        if token.CreatedAt.Before(sixMonthsAgo) {
            newKey := generateNewKey()
            tokens[i].Key = newKey
            tokens[i].CreatedAt = time.Now()
            fmt.Printf("Token %d updated with new key: %s\n", token.ID, newKey)
        }
    }
}

func main() {
    // Create an in-memory array of tokens with random created_at dates
    tokens := []Token{
        {ID: 1, CreatedAt: time.Now().AddDate(0, -7, 0), Key: "old-key-1"},
        {ID: 2, CreatedAt: time.Now().AddDate(0, -5, 0), Key: "old-key-2"},
        {ID: 3, CreatedAt: time.Now().AddDate(0, -8, 0), Key: "old-key-3"},
    }

    // Create a new cron job
    c := cron.New()

    // Schedule the task to run every 30 seconds
    _, err := c.AddFunc("@every 5s", func() {
        checkAndUpdateTokens(tokens)
    })
    if err != nil {
        log.Fatalf("Failed to schedule task: %v", err)
    }

    // Start the cron scheduler
    c.Start()

    // Keep the program running
    select {}
}
