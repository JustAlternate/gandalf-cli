package main

import (
	"fmt"
	"encoding/json"
	"os"
	"math/rand"
	"github.com/spf13/cobra"
)

type Quote struct {
	Quote string
	Book  string
}

func parseQuotesJson() []Quote {
	content, err := os.ReadFile("quotes.json")
	if err != nil {
		fmt.Printf("Error reading file quotes.json: %v", err)
		os.Exit(1)
	}
	var payload []Quote
	err = json.Unmarshal(content, &payload)
	if err != nil {
		fmt.Printf("Error parsing quotes.json: %v", err)
		os.Exit(1)
	}
	return payload
}

func talk(cmd *cobra.Command, args []string){
	quotes := parseQuotesJson()
	fmt.Printf("%s\n", quotes[rand.Intn(len(quotes))].Quote)
}

func list(cmd *cobra.Command, args []string){
	quotes := parseQuotesJson()
	for i := 0; i<len(quotes) ; i++ {
		fmt.Printf("%s\n", quotes[i])
	}
}

func main(){
	var rootCmd = &cobra.Command{
	 Use:   "gandalf",
	 Short: "Say a random gandalf quote.",
	 Long: `This command will print a randomized popular gandalf quote`,
	 Run: talk,
	}

	var listCmd = &cobra.Command{
		Use: "list",
		Short: "List all gandalf quotes.",
		Long: "List all gandalf quotes.",
		Run: list,
	}

	rootCmd.AddCommand(listCmd)

	err := rootCmd.Execute()
	if err != nil {
		fmt.Printf("Error executing cli: %v", err)
		os.Exit(1)
	}
}
