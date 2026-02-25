package example

import (
	"context"
	"fmt"
	"log"

	"github.com/PavelMkr/go-ai-framework/internal/client"
)

func main() {
	ai := client.New("http://localhost:11434/", "llama3")

	resp, err := ai.Generate(context.Background(), "Explain Go interfaces in one sentence")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("AI response:")
	fmt.Println(resp)
}
