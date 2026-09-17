package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
)

func main() {
	fmt.Println("Initializing backup sequence...")
	ctx := context.Background()

	// 1. Point this to your renamed JSON key file
	sa := option.WithCredentialsFile("firebase_key.json")

	// 2. Connect to Firebase
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalf("Error initializing app: %v\n", err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalf("Error getting Firestore client: %v\n", err)
	}
	defer client.Close()

	// 3. List the files you want to back up
	filesToBackup := []string{
		"main.go",
		"server.go",
		"public/index.html",
		"classRuleset.go",
		"dieChecker.go",
		"doors.go",
		"entity.go",
		"map.go",
		"monsters.go",
		"path.go",
		"render.go",
		"theme.go",
		"traps.go",
		"treasure.go",
	}
	commitData := make(map[string]interface{})

	// 4. Read the text inside those files
	for _, filename := range filesToBackup {
		content, err := os.ReadFile(filename)
		if err != nil {
			fmt.Printf("Warning: Could not read %s. Skipping.\n", filename)
			continue
		}
		commitData[filename] = string(content)
	}

	// Add a timestamp so you know exactly when this backup was made
	commitData["timestamp"] = time.Now().Format(time.RFC1123)

	// 5. Push the whole package to a "commits" collection in Firestore
	_, _, err = client.Collection("commits").Add(ctx, commitData)
	if err != nil {
		log.Fatalf("Failed pushing backup to Firestore: %v\n", err)
	}

	fmt.Println("Success! Code snapshot safely vaulted in Firestore.")
}
