/*
Exercise8.10
Пакет links предоставляет функцию для извлечения ссылок
*/

package links

import (
	"context"
	"fmt"
	"net/http"

	"golang.org/x/net/html"
)

// Extract выполняет HTTP-запрос GET по определенному URL, выполняет синтаксический анализ HTML и
// возвращает ссылки в HTML-документе
func Extract(url string, cancelled <-chan struct{}) ([]string, error) {
	// req.Cancel - deprecated. Need use context.
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	go func() {
		select {
		case <-cancelled:
			fmt.Println("user stop service")
			cancel()
		case <-ctx.Done():
		}
	}()

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("получение %s: %s", url, resp.Status)
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("анализ %s как HTML: %v", url, err)
	}

	var links []string
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				link, err := resp.Request.URL.Parse(a.Val)
				if err != nil {
					continue // Игнорируем некорректные URL
				}
				links = append(links, link.String())
			}
		}
	}
	forEachNode(doc, visitNode, nil)
	return links, nil
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}

}
