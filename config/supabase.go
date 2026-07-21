package config

import (
	"log"
	"os"

	supabase "github.com/supabase-community/supabase-go"
)

var Supabase *supabase.Client

func InitSupabase() {
	url := os.Getenv("SUPABASE_URL")
	key := os.Getenv("SUPABASE_SERVICE_ROLE_KEY")

	if url == "" || key == "" {
		log.Fatal("SUPABASE_URL or SUPABASE_SERVICE_ROLE_KEY is not set")
	}

	client, err := supabase.NewClient(url, key, nil)
	if err != nil {
		log.Fatalf("failed to initialize Supabase client: %v", err)
	}

	Supabase = client

	log.Println("Supabase client initialized")
}
