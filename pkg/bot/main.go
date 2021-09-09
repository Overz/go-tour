package bot

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
)

// base url:
// https://iptu.pmf.sc.gov.br/iptu-virtual/main-iptu/

func main() {
	inscricao := "22790120228001282"

	c := colly.NewCollector()

	var rows [][]string
	c.OnHTML("tbody > tr", func(e *colly.HTMLElement) {
		var row []string

		e.DOM.Find("td").Each(func(i int, s *goquery.Selection) {
			if len(row) == 4 {
				return
			}

			text := strings.TrimSpace(s.Contents().Text())
			if text != "" {
				row = append(row, text)
			}
		})

		if len(row) > 0 {
			rows = append(rows, row)
		}
	})

	r, err := regexp.Compile(`^.*sitekey.*'(.*)'.*$`)
	if err != nil {
		log.Fatal(err)
	}

	var siteKey string
	c.OnHTML("script", func(e *colly.HTMLElement) {
		if siteKey != "" {
			return
		}

		lines := strings.Split(e.Text, "\n")
		for _, line := range lines {
			if !strings.Contains(line, "sitekey") {
				continue
			}

			matches := r.FindStringSubmatch(line)
			if len(matches) == 2 {
				siteKey = matches[1]
				break
			}
		}
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Requesting....")
	})

	_ = c.PostMultipart("http://iptu2019.pmf.sc.gov.br/iptu-virtual/main-iptu/index-segunda-via.php",
		map[string][]byte{
			"nu_insc_imbl":           []byte(inscricao),
			"ano-segunda-via-select": []byte("2021"),
		})

	println("Waiting...")
	c.Wait()

	for i, row := range rows {
		fmt.Printf("#%.2d -- %s\n", i, strings.Join(row, ","))
	}

	fmt.Printf("SiteKey: %v\n", siteKey)

	captchaSolved, err := solveCaptcha(siteKey)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Solved captcha: %v\n", captchaSolved)

	fmt.Printf("\n\n")
	var dams []string
	for _, row := range rows {
		dam := strings.ReplaceAll(row[0], "-", "")
		dams = append(dams, dam)
	}

	if err := baixarParcelaUnica(dams, captchaSolved, "TUDOJUNTO.pdf"); err != nil {
		log.Fatal(err)
	}

	println("Done")
}

func solveCaptcha(siteKey string) (string, error) {
	res, err := http.PostForm("http://captchatypers.com/captchaapi/UploadRecaptchaToken.ashx", url.Values{
		"token":         {"xxxxx"},
		"pageurl":       {"https://iptu.pmf.sc.gov.br/iptu-virtual/main-iptu/"},
		"googlekey":     {"6LfatIkUAAAAABA1IFcpmL9KkJbFCDsSPZGeaplW"},
		"recaptchatype": {"2"},
		"action":        {"UPLOADCAPTCHA"},
	})

	if err != nil {
		return "", err
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	captchaID := string(b)
	if _, err := strconv.Atoi(captchaID); err != nil {
		return "", fmt.Errorf("error sending captcha to solve: %s", captchaID)
	}

	timeout := time.After(3 * time.Minute)
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	i := 1
	for {
		select {
		case <-ticker.C:
			res, err := http.PostForm("http://captchatypers.com/captchaapi/GetRecaptchaTextToken.ashx",
				url.Values{
					"token":     {"xxxxx"},
					"captchaID": {captchaID},
					"action":    {"GETTEXT"},
				})
			if err != nil {
				return "", err
			}

			b, err := ioutil.ReadAll(res.Body)
			if err != nil {
				return "", err
			}

			captchaResult := string(b)
			if strings.HasPrefix(captchaResult, "ERROR") {
				fmt.Printf("#%d Captcha not ready (%s)...\n", i, captchaResult)
				i++
			} else {
				return captchaResult, nil
			}

		case <-timeout:
			return "", fmt.Errorf("timeout when waiting for captcha solution")
		}
	}
}

func baixarParcelaUnica(dams []string, captchaResponse string, fileName string) error {
	res, err := http.PostForm("http://iptu2019.pmf.sc.gov.br/iptu-virtual/main-iptu/segunda_via_internet.pdf.php",
		url.Values{
			"nu_dam[]":             dams,
			"controle":             {"ADMIN"},
			"g-recaptcha-response": {captchaResponse},
		})

	if err != nil {
		return err
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(fileName, b, 0755)
}
