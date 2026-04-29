package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"unicode/utf8"
)

const (
	maxTitle       = 60
	maxDescription = 160
	coverWidth     = 1200
	coverHeight    = 630
)

var siteBaseURL = "https://blog.example.com"

type post struct {
	path        string
	title       string
	date        string
	description string
	draft       string
	slug        string
	cover       string
	tags        []string
	topics      []string
	series      []string
	keywords    []string
	aliases     []string
}

func main() {
	siteBaseURL = readBaseURL()
	if rootPosts, _ := filepath.Glob("content/posts/*.md"); len(rootPosts) > 0 {
		fail("posts must use leaf page bundles: content/posts/<slug>/index.md")
	}
	paths, err := filepath.Glob("content/posts/*/index.md")
	if err != nil {
		fail("cannot list posts: %v", err)
	}
	if len(paths) == 0 {
		fail("no posts found in content/posts")
	}

	var failures []string
	for _, path := range paths {
		p, err := readPost(path)
		if err != nil {
			failures = append(failures, fmt.Sprintf("%s: %v", path, err))
			continue
		}
		failures = append(failures, validatePost(p)...)
	}

	if len(failures) > 0 {
		for _, failure := range failures {
			fmt.Fprintln(os.Stderr, failure)
		}
		os.Exit(1)
	}

	fmt.Printf("Validated %d posts\n", len(paths))
}

func readPost(path string) (post, error) {
	file, err := os.Open(path)
	if err != nil {
		return post{}, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNo := 0
	inFrontMatter := false
	listField := ""
	p := post{path: path}

	for scanner.Scan() {
		line := scanner.Text()
		lineNo++
		if lineNo == 1 {
			if strings.TrimSpace(line) != "---" {
				return post{}, fmt.Errorf("front matter must start with ---")
			}
			inFrontMatter = true
			continue
		}
		if inFrontMatter && strings.TrimSpace(line) == "---" {
			return p, scanner.Err()
		}
		if !inFrontMatter {
			continue
		}

		trimmed := strings.TrimSpace(line)
		if strings.HasPrefix(trimmed, "- ") && listField != "" {
			value := strings.Trim(strings.TrimSpace(strings.TrimPrefix(trimmed, "- ")), `"`)
			switch listField {
			case "tags":
				p.tags = append(p.tags, value)
			case "topics":
				p.topics = append(p.topics, value)
			case "series":
				p.series = append(p.series, value)
			case "keywords":
				p.keywords = append(p.keywords, value)
			case "aliases":
				p.aliases = append(p.aliases, value)
			}
			continue
		}
		listField = ""
		key, value, ok := strings.Cut(trimmed, ":")
		if !ok {
			continue
		}
		value = strings.Trim(strings.TrimSpace(value), `"`)
		switch key {
		case "title":
			p.title = value
		case "date":
			p.date = value
		case "description":
			p.description = value
		case "draft":
			p.draft = value
		case "slug":
			p.slug = value
		case "cover":
			p.cover = value
		case "tags", "topics", "series", "keywords", "aliases":
			listField = key
		}
	}
	if err := scanner.Err(); err != nil {
		return post{}, err
	}
	return post{}, fmt.Errorf("front matter must end with ---")
}

func validatePost(p post) []string {
	var failures []string
	required := map[string]string{
		"title":       p.title,
		"date":        p.date,
		"description": p.description,
		"draft":       p.draft,
		"cover":       p.cover,
	}
	for field, value := range required {
		if strings.TrimSpace(value) == "" {
			failures = append(failures, fmt.Sprintf("%s: missing %s", p.path, field))
		}
	}
	if utf8.RuneCountInString(p.title) > maxTitle {
		failures = append(failures, fmt.Sprintf("%s: title is %d chars, max %d", p.path, utf8.RuneCountInString(p.title), maxTitle))
	}
	if utf8.RuneCountInString(p.description) > maxDescription {
		failures = append(failures, fmt.Sprintf("%s: description is %d chars, max %d", p.path, utf8.RuneCountInString(p.description), maxDescription))
	}
	if len(p.tags) == 0 {
		failures = append(failures, fmt.Sprintf("%s: missing tags", p.path))
	}
	if len(p.topics) == 0 {
		failures = append(failures, fmt.Sprintf("%s: missing topics", p.path))
	}
	if len(p.series) == 0 {
		failures = append(failures, fmt.Sprintf("%s: missing series", p.path))
	}
	if len(p.keywords) == 0 {
		failures = append(failures, fmt.Sprintf("%s: missing keywords", p.path))
	}
	if len(p.aliases) == 0 {
		failures = append(failures, fmt.Sprintf("%s: missing aliases for redirect generation", p.path))
	}
	if p.draft != "false" {
		failures = append(failures, fmt.Sprintf("%s: draft must be false for publish validation", p.path))
	}
	if p.cover != "" {
		failures = append(failures, validateCover(p)...)
		failures = append(failures, validateGeneratedHTML(p)...)
	}
	return failures
}

func validateCover(p post) []string {
	if !strings.HasPrefix(p.cover, "/") {
		coverPath := filepath.Join(filepath.Dir(p.path), p.cover)
		return validateCoverDimensions(p, coverPath)
	}
	coverPath := filepath.Join("static", strings.TrimPrefix(p.cover, "/"))
	return validateCoverDimensions(p, coverPath)
}

func validateCoverDimensions(p post, coverPath string) []string {
	width, height, err := jpegDimensions(coverPath)
	if err != nil {
		return []string{fmt.Sprintf("%s: cannot read cover dimensions for %s: %v", p.path, coverPath, err)}
	}
	if width != coverWidth || height != coverHeight {
		return []string{fmt.Sprintf("%s: cover must be %dx%d, got %dx%d", p.path, coverWidth, coverHeight, width, height)}
	}
	return nil
}

func jpegDimensions(path string) (int, int, error) {
	if !strings.HasSuffix(strings.ToLower(path), ".jpg") && !strings.HasSuffix(strings.ToLower(path), ".jpeg") {
		return 0, 0, fmt.Errorf("cover must be a JPEG file")
	}
	data, err := os.ReadFile(path)
	if err != nil {
		return 0, 0, err
	}
	if len(data) < 4 || data[0] != 0xff || data[1] != 0xd8 {
		return 0, 0, fmt.Errorf("missing JPEG SOI marker")
	}
	for i := 2; i+9 < len(data); {
		if data[i] != 0xff {
			i++
			continue
		}
		for i < len(data) && data[i] == 0xff {
			i++
		}
		if i >= len(data) {
			break
		}
		marker := data[i]
		i++
		if marker == 0xd9 || marker == 0xda {
			break
		}
		if i+2 > len(data) {
			break
		}
		segmentLength := int(binary.BigEndian.Uint16(data[i : i+2]))
		if segmentLength < 2 || i+segmentLength > len(data) {
			return 0, 0, fmt.Errorf("invalid JPEG segment")
		}
		if isStartOfFrame(marker) {
			if segmentLength < 7 {
				return 0, 0, fmt.Errorf("invalid SOF segment")
			}
			height := int(binary.BigEndian.Uint16(data[i+3 : i+5]))
			width := int(binary.BigEndian.Uint16(data[i+5 : i+7]))
			return width, height, nil
		}
		i += segmentLength
	}
	return 0, 0, fmt.Errorf("JPEG SOF marker not found")
}

func isStartOfFrame(marker byte) bool {
	switch marker {
	case 0xc0, 0xc1, 0xc2, 0xc3, 0xc5, 0xc6, 0xc7, 0xc9, 0xca, 0xcb, 0xcd, 0xce, 0xcf:
		return true
	default:
		return false
	}
}

func validateGeneratedHTML(p post) []string {
	htmlPath := filepath.Join("public", "posts", postSlug(p), "index.html")
	bytes, err := os.ReadFile(htmlPath)
	if err != nil {
		return []string{fmt.Sprintf("%s: generated HTML missing: %s", p.path, htmlPath)}
	}
	html := string(bytes)
	base := strings.TrimRight(siteBaseURL, "/")
	imagePattern := regexp.QuoteMeta(base + p.cover)
	if !strings.HasPrefix(p.cover, "/") {
		imagePattern = regexp.QuoteMeta(base+"/posts/"+postSlug(p)+"/") + `cover[^"\s>]*`
	}
	checks := map[string]*regexp.Regexp{
		"canonical":        regexp.MustCompile(`<link rel="?canonical"? href="?` + regexp.QuoteMeta(base) + `/posts/[^">]+/?`),
		"og:type article":  regexp.MustCompile(`property="?og:type"? content="?article"?`),
		"og:image":         regexp.MustCompile(`property="?og:image"? content="?` + imagePattern),
		"og:image:width":   regexp.MustCompile(`property="?og:image:width"? content="?1200"?`),
		"og:image:height":  regexp.MustCompile(`property="?og:image:height"? content="?630"?`),
		"twitter:image":    regexp.MustCompile(`name="?twitter:image"? content="?` + imagePattern),
		"twitter:card":     regexp.MustCompile(`name="?twitter:card"? content="?summary_large_image"?`),
		"BlogPosting JSON": regexp.MustCompile(`"@type":"BlogPosting"`),
	}
	var failures []string
	for name, pattern := range checks {
		if !pattern.MatchString(html) {
			failures = append(failures, fmt.Sprintf("%s: generated HTML missing %s", p.path, name))
		}
	}
	return failures
}

func postSlug(p post) string {
	if p.slug != "" {
		return p.slug
	}
	base := strings.TrimSuffix(filepath.Base(p.path), filepath.Ext(p.path))
	return urlize(base)
}

func urlize(value string) string {
	value = strings.ToLower(value)
	var b strings.Builder
	dash := false
	for _, r := range value {
		if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') {
			b.WriteRune(r)
			dash = false
			continue
		}
		if !dash {
			b.WriteByte('-')
			dash = true
		}
	}
	return strings.Trim(b.String(), "-")
}

func fail(format string, args ...any) {
	fmt.Fprintf(os.Stderr, format+"\n", args...)
	os.Exit(1)
}

func readBaseURL() string {
	bytes, err := os.ReadFile("config.toml")
	if err != nil {
		return siteBaseURL
	}
	for _, line := range strings.Split(string(bytes), "\n") {
		key, value, ok := strings.Cut(strings.TrimSpace(line), "=")
		if !ok || strings.TrimSpace(key) != "baseURL" {
			continue
		}
		value = strings.Trim(strings.TrimSpace(value), `"`)
		if value != "" {
			return value
		}
	}
	return siteBaseURL
}
