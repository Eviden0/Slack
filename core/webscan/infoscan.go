package webscan

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"net/url"
	"runtime"
	"slack-wails/lib/clients"
	"slack-wails/lib/gonmap"
	"slack-wails/lib/util"
	"strconv"
	"strings"
	"sync"
	"time"
)

type TargetINFO struct {
	Protocol      string
	Port          int
	Path          string
	Title         string
	StatusCode    int
	IconHash      string // mmh3
	BodyString    string
	HeadeString   string
	ContentType   string
	Server        string
	ContentLength int
	Banner        string // tcp指纹
	Cert          string // TLS证书
}

// 代理过滤器，防止在走代理扫描时将不存活的网站由于走了代理导致状态码变成200
// var filter = []string{"Burp Suite Professional"}

// DumpResponseHeadersAndRaw returns http headers and response as strings
func DumpResponseHeadersAndRaw(resp *http.Response) (headers, fullresp []byte, err error) {
	// httputil.DumpResponse does not work with websockets
	if resp.StatusCode >= http.StatusContinue && resp.StatusCode <= http.StatusEarlyHints {
		raw := resp.Status + "\n"
		for h, v := range resp.Header {
			raw += fmt.Sprintf("%s: %s\n", h, v)
		}
		return []byte(raw), []byte(raw), nil
	}
	headers, err = httputil.DumpResponse(resp, false)
	if err != nil {
		return
	}
	// logic same as httputil.DumpResponse(resp, true) but handles
	// the edge case when we get both error and data on reading resp.Body
	var buf1, buf2 bytes.Buffer
	b := resp.Body
	if _, err = buf1.ReadFrom(b); err != nil {
		if buf1.Len() <= 0 {
			return
		}
	}
	if err == nil {
		_ = b.Close()
	}

	// rewind the body to allow full dump
	resp.Body = io.NopCloser(bytes.NewReader(buf1.Bytes()))
	err = resp.Write(&buf2)
	fullresp = buf2.Bytes()

	// rewind once more to allow further reuses
	resp.Body = io.NopCloser(bytes.NewReader(buf1.Bytes()))
	return
}

func FingerScan(ti *TargetINFO, targetDB []FingerPEntity) []string {
	var fingerPrintResults []string

	isWeb := ti.Path != "no#web"
	workers := runtime.NumCPU() * 2
	inputChan := make(chan FingerPEntity, len(targetDB))
	defer close(inputChan)
	results := make(chan string, len(targetDB))
	defer close(results)

	var wg sync.WaitGroup

	//接收结果
	go func() {
		for found := range results {
			if found != "" {
				fingerPrintResults = append(fingerPrintResults, found)
			}
			wg.Done()
		}
	}()

	//多线程扫描
	for i := 0; i < workers; i++ {
		go func() {
			for finger := range inputChan {
				rules := finger.Rule
				product := finger.ProductName
				expr := finger.AllString
				for _, singleRule := range rules {
					singleRuleResult := false
					if singleRule.Key == "header" {
						if isWeb && dataCheckString(singleRule.Op, ti.HeadeString, singleRule.Value) {
							singleRuleResult = true
						}
					} else if singleRule.Key == "body" {
						if isWeb && dataCheckString(singleRule.Op, ti.BodyString, singleRule.Value) {
							singleRuleResult = true
						}
					} else if singleRule.Key == "server" {
						if isWeb && dataCheckString(singleRule.Op, ti.Server, singleRule.Value) {
							singleRuleResult = true
						}
					} else if singleRule.Key == "title" {
						if isWeb && dataCheckString(singleRule.Op, ti.Title, singleRule.Value) {
							singleRuleResult = true
						}
					} else if singleRule.Key == "cert" {
						if dataCheckString(singleRule.Op, ti.Cert, singleRule.Value) {
							singleRuleResult = true
						}
					} else if singleRule.Key == "port" {
						value, err := strconv.Atoi(singleRule.Value)
						if err == nil && dataCheckInt(singleRule.Op, ti.Port, value) {
							singleRuleResult = true
						}
					} else if singleRule.Key == "protocol" {
						if singleRule.Op == 0 {
							if ti.Protocol == singleRule.Value {
								singleRuleResult = true
							}
						} else if singleRule.Op == 1 {
							if ti.Protocol != singleRule.Value {
								singleRuleResult = true
							}
						}
					} else if singleRule.Key == "path" {
						if isWeb && dataCheckString(singleRule.Op, ti.Path, singleRule.Value) {
							singleRuleResult = true
						}
					} else if singleRule.Key == "icon_hash" {
						value, err := strconv.Atoi(singleRule.Value)
						hashIcon, errHash := strconv.Atoi(ti.IconHash)
						if isWeb && err == nil && errHash == nil && dataCheckInt(singleRule.Op, hashIcon, value) {
							singleRuleResult = true
						}
					} else if singleRule.Key == "status" {
						value, err := strconv.Atoi(singleRule.Value)
						if isWeb && err == nil && dataCheckInt(singleRule.Op, ti.StatusCode, value) {
							singleRuleResult = true
						}
					} else if singleRule.Key == "content_type" {
						if isWeb && dataCheckString(singleRule.Op, ti.ContentType, singleRule.Value) {
							singleRuleResult = true
						}
					} else if singleRule.Key == "banner" {
						if dataCheckString(singleRule.Op, ti.Banner, singleRule.Value) {
							singleRuleResult = true
						}
					}
					if singleRuleResult {
						expr = expr[:singleRule.Start] + "T" + expr[singleRule.End:]
					} else {
						expr = expr[:singleRule.Start] + "F" + expr[singleRule.End:]
					}
				}

				r := boolEval(expr)
				if r {
					results <- product
				} else {
					results <- ""
				}

			}

		}()
	}
	//添加扫描目标
	for _, input := range targetDB {
		wg.Add(1)
		inputChan <- input
	}
	wg.Wait()
	return util.RemoveDuplicates(fingerPrintResults)
}

// 获取favicon hash值
func FaviconHash(url string, client *http.Client) string {
	resp, body, err := clients.NewRequest("GET", url+"favicon.ico", nil, nil, 10, client)
	if err != nil {
		return ""
	}
	if resp.StatusCode == 200 {
		return util.Mmh3Hash32(util.Base64Encode(body))
	} else {
		return ""
	}
}

type URLINFO struct {
	Protocol string
	Host     string
	Port     int
	Path     string
}

func HostPort(target string) URLINFO {
	var host string
	var port int
	u, _ := url.Parse(target)
	if strings.Contains(u.Host, ":") {
		host = strings.Split(u.Host, ":")[0]
		port, _ = strconv.Atoi(strings.Split(u.Host, ":")[1])
	} else {
		host = u.Host
		port = 80
	}
	return URLINFO{
		Protocol: u.Scheme,
		Host:     host,
		Port:     port,
		Path:     u.Path,
	}
}

func GetBanner(u *URLINFO) string {
	if strings.HasPrefix(u.Protocol, "http") {
		return ""
	}
	scanner := gonmap.New()
	_, response := scanner.Scan(u.Host, u.Port, time.Second*time.Duration(10))
	if response != nil {
		return response.Raw
	}
	return ""
}

func GetHeaderInfo(body []byte, resp *http.Response) (title, server, content_type string) {
	if match := util.RegTitle.FindSubmatch(body); len(match) > 1 {
		title = util.Str2UTF8(string(match[1]))
	}
	for k, v := range resp.Header {
		if k == "Server" {
			server = strings.Join(v, ";")
		}
		if k == "Content-Type" {
			content_type = strings.Join(v, ";")
		}
	}
	return
}
