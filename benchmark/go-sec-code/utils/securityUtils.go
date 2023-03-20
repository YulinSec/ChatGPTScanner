package utils

import (
	"net/url"
	"regexp"
	"strings"
)

type CommandInjectFilter struct {
}

type CorsFilter struct {
}

type FileUploadFilter struct {
}

type JsonpFilter struct {
}

type PathTraversalFilter struct {
}

type SSRFFilter struct {
}

type XSSFilter struct {
}

func (c *CommandInjectFilter) DoFilter(input string) bool {
	r, _ := regexp.Compile(`^[a-zA-Z0-9_/\.-]+$`)
	return r.MatchString(input)
}

func (c *CorsFilter) DoFilter(input string, whitelists []string) bool {
	for _, v := range whitelists {
		if strings.HasSuffix(input, "."+v) || input == v {
			return true
		}
	}
	return false
}

func (c *FileUploadFilter) DoFilter(input string) bool {
	r, _ := regexp.Compile(`\.\./`)
	return r.MatchString(input)
}

func (c *JsonpFilter) DoFilter(input string, whitelists []string) bool {
	u, err := url.Parse(input)
	if err != nil {
		panic(err)
	}
	for _, v := range whitelists {
		if strings.HasSuffix(u.Host, "."+v) || input == v {
			return true
		}
	}
	return false
}

func (c *PathTraversalFilter) DoFilter(input string) bool {
	r, _ := regexp.Compile(`\.\.`)
	return r.MatchString(input) || strings.HasPrefix(input, "/")
}

func (c *SSRFFilter) DoBlackFilter(input string, blacklists []string) bool {
	u, err := url.Parse(input)
	if err != nil {
		panic(err)
	}
	for _, v := range blacklists {
		if strings.HasSuffix(u.Hostname(), v) {
			return true
		}
	}
	return false
}

func (c *SSRFFilter) DoWhiteFilter(input string, whitelists []string) bool {
	u, err := url.Parse(input)
	if err != nil {
		panic(err)
	}
	//exclude evil-example.com
	if strings.HasPrefix(u.Hostname(), ".") {
		for _, v := range whitelists {
			if strings.HasSuffix(u.Hostname(), v) {
				return false
			}
		}
	} else {
		for _, v := range whitelists {
			if u.Hostname() == v {
				return false
			}
		}
	}
	return true
}

// this filter comes from fix of cve-2022-0870 gogs SSRF
func (c *SSRFFilter) DoGogsFilter(input string) bool {
	return IsLocalHostname(input, nil)
}

func (c *XSSFilter) DoFilter(input string) string {
	mid := strings.ReplaceAll(input, "&", "&amp;")
	mid = strings.ReplaceAll(mid, "<", "&lt;")
	mid = strings.ReplaceAll(mid, ">", "&gt;")
	mid = strings.ReplaceAll(mid, "\"", "&quot;")
	mid = strings.ReplaceAll(mid, "'", "&#x27")
	output := strings.ReplaceAll(mid, "/", "&#x2F")
	return output
}
