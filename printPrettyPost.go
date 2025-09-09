// go
package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/Cheemx/gator/internal/database"
)

func humanTime(t time.Time) string {
	d := time.Since(t)
	if d < time.Minute {
		return "just now"
	}
	if d < time.Hour {
		return fmt.Sprintf("%dm ago", int(d.Minutes()))
	}
	if d < 24*time.Hour {
		return fmt.Sprintf("%dh ago", int(d.Hours()))
	}
	return t.Format("2006-01-02")
}

func truncate(s string, n int) string {
	s = strings.TrimSpace(s)
	if len(s) <= n {
		return s
	}
	// cut at word boundary if possible
	cut := strings.LastIndex(s[:n], " ")
	if cut <= 0 {
		cut = n
	}
	return s[:cut] + "..."
}

func printPost(p database.Post) {
	// Title
	fmt.Printf("• %s\n", p.Title)

	// Source + when
	when := "unknown"
	if p.PublishedAt.Valid {
		when = humanTime(p.PublishedAt.Time)
	} else {
		when = humanTime(p.CreatedAt)
	}
	fmt.Printf("  [%s] %s\n", p.FeedID, when)

	// URL
	fmt.Printf("  %s\n", p.Url)

	// Description
	if p.Description.Valid && strings.TrimSpace(p.Description.String) != "" {
		fmt.Printf("  %s\n", truncate(p.Description.String, 200))
	}

	fmt.Println()
}
