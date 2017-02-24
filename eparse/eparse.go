package eparse

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

func readWords(path string) ([]string, error) {
	// open input file
	fi, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer fi.Close()

	var words []string
	scanner := bufio.NewScanner(fi)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}
	return words, scanner.Err()
}

func writeWords(words []string, path string) error {
	// open output file
	fo, err := os.Create(path)
	if err != nil {
		return err
	}
	defer fo.Close()

	w := bufio.NewWriter(fo)
	fmt.Fprintf(w, "\"E-mail Address\"\r\n")
	for _, word := range words {
		fmt.Fprintf(w, "\"%s\"\r\n", word)
	}
	return w.Flush()
}

func removeDuplicates(xs *[]string) (dupes int) {
	found := make(map[string]bool)
	dupes = 0
	j := 0
	for i, x := range *xs {
		if !found[x] {
			found[x] = true
			(*xs)[j] = (*xs)[i]
			j++
		} else {
			dupes++
		}
	}
	*xs = (*xs)[:j]
	return
}

// filterEmails looks for Valid Email Addresses
func filterEmails(words []string) (emails []string, valid int) {
	// emails = make([]string, 1)
	re := regexp.MustCompile(`[A-Z,.,a-z,\d]+(@\S+)(\.\w+)`)
	valid = 0
	for _, x := range words {
		if re.MatchString(x) {
			valid++
			emails = append(emails, re.FindString(x))
		}
	}
	return
}

// Parse file
func Parse(source string, destination string) (err error) {
	words, err := readWords(source)
	if err != nil {
		log.Fatalf("readWords: %s", err)
	}

	emails, valid := filterEmails(words)
	fmt.Printf("Found %d email addresses\n", valid)

	dupes := removeDuplicates(&emails)
	fmt.Printf("Removed %d duplicates\n", dupes)

	if err := writeWords(emails, destination); err != nil {
		log.Fatalf("writeWords: %s", err)
	}

	fmt.Printf("Created document with %d emails", len(emails))

	return nil

}
