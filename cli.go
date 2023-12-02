package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func main() {
	year := flag.Int("year", 2023, "Year of the Advent of Code")
	day := flag.Int("day", 1, "Day of the Advent of Code")
	flag.Parse()

	sessionCookie := os.Getenv("TOKEN")
	if sessionCookie == "" {
		log.Fatal("Session cookie not set")
	}

	// Create directory path
	dirPath := fmt.Sprintf("%d/%d", *year, *day)
	err := os.MkdirAll(dirPath, 0755)
	if err != nil {
		log.Fatal("Error creating directory:", err)
	}

	// Function to make a request and save the content
	makeRequestAndSave := func(url, filename string) {
		client := &http.Client{}
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			log.Fatal(err)
		}

		req.Header.Set("Cookie", "session="+sessionCookie)
		resp, err := client.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal("Error reading response body:", err)
		}

		err = ioutil.WriteFile(filepath.Join(dirPath, filename), body, 0644)
		if err != nil {
			log.Fatal("Error writing to file:", err)
		}

		log.Println("Saved content to:", filepath.Join(dirPath, filename))
	}

	// // Downloading and parsing description
	// descURL := fmt.Sprintf("https://adventofcode.com/%d/day/%d", *year, *day)
	// makeRequestAndSave(descURL, "description.html")

	// // Optional: Parse and save a clean version of the description
	// // Note: Parsing HTML can be complex and might require adjustments based on the structure of the webpage
	// parseAndSaveDescription(filepath.Join(dirPath, "description.html"), filepath.Join(dirPath, "description.txt"))

	// Downloading input
	inputURL := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", *year, *day)
	makeRequestAndSave(inputURL, "input.txt")

	// Basic template for main.go
	mainGoContent := []byte("package main\n\nfunc main() {\n\t// Your code here\n}\n")

	// Write the basic main.go template to file
	part1GoFilename := filepath.Join(dirPath, "part1.go")
	err = ioutil.WriteFile(part1GoFilename, mainGoContent, 0644)
	if err != nil {
		log.Fatal("Error writing main.go to file:", err)
	}
	log.Println("Saved content to:", part1GoFilename)

	part2GoFilename := filepath.Join(dirPath, "part2.go")
	err = ioutil.WriteFile(part2GoFilename, mainGoContent, 0644)
	if err != nil {
		log.Fatal("Error writing main.go to file:", err)
	}
	log.Println("Saved content to:", part2GoFilename)

}

// parseAndSaveDescription parses the HTML file to extract the challenge description from the <main> tag
func parseAndSaveDescription(htmlFile, txtFile string) {
	data, err := ioutil.ReadFile(htmlFile)
	if err != nil {
		log.Fatal("Error reading HTML file:", err)
	}

	// Adjusted regex to target content within <main> tags
	// Note: This regex is basic and might not work perfectly for all kinds of HTML content.
	re := regexp.MustCompile(`(?s)<article>(.*?)</article>`)
	matches := re.FindSubmatch(data)
	if len(matches) > 1 {
		content := strings.TrimSpace(string(matches[1]))

		// Further processing to clean up HTML content
		// This can be adjusted or enhanced based on specific needs
		content = cleanHTML(content)

		err := ioutil.WriteFile(txtFile, []byte(content), 0644)
		if err != nil {
			log.Fatal("Error writing description to file:", err)
		}
		log.Println("Parsed and saved description to:", txtFile)
	} else {
		log.Println("No description found in HTML file")
	}
}

// cleanHTML attempts to remove HTML tags and convert some HTML entities to plain text
func cleanHTML(html string) string {
	// Remove script and style elements
	html = regexp.MustCompile(`(?s)<(script|style).*?</\1>`).ReplaceAllString(html, "")

	// Remove all HTML tags, leaving only the content
	html = regexp.MustCompile(`(?s)<.*?>`).ReplaceAllString(html, "")

	// Replace HTML entities with their plain text equivalents
	html = strings.ReplaceAll(html, "&lt;", "<")
	html = strings.ReplaceAll(html, "&gt;", ">")
	html = strings.ReplaceAll(html, "&amp;", "&")
	html = strings.ReplaceAll(html, "&quot;", "\"")
	html = strings.ReplaceAll(html, "&#39;", "'")

	// Additional replacements for specific formatting can be added here

	return strings.TrimSpace(html)
}
