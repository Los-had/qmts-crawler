package utils

var userAgent string = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.54 Safari/537.36"
var proxyList []string = []string{"http://192.155.107.214:1080", "http://213.230.97.10:3128", "http://170.239.255.2:55443"}
//"github.com/gocolly/colly/proxy"

/*
    if py, err := proxy.RoundRobinProxySwitcher(proxyList...); err != nil {
        fmt.Println("Error occurred:", err)
    } else {
        c.SetProxyFunc(py)
    }
*/