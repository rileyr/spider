### spider

  Webcrawler written in Go, to learn Go.

#### General Go Stuffs

  godocs say that most go devs use a single workspace for all repositories. i've
set my gopath to `/Users/riley/golang`, and intend to use the singular workspace
setup. With spider installed, my tree looks like:

```
.
├── bin
│   └── spider
└── src
    └── github.com
        └── rileyr
            └── spider
                ├── README.md
                └── spider.go
```

#### usage
  Run spider, providing the url of the site to be crawled:

```bash
$> bin/spider -url=http://www.google.com
URL to crawl is: http://www.google.com
URL retrieved: http://www.google.com
Found 21 links!
0: http://www.google.com
1: http://www.google.com/imghp?hl=en&tab=wi
2: http://maps.google.com/maps?hl=en&tab=wl
3: https://play.google.com/?hl=en&tab=w8
4: http://www.youtube.com/?tab=w1
5: http://news.google.com/nwshp?hl=en&tab=wn
6: https://mail.google.com/mail/?tab=wm
7: https://drive.google.com/?tab=wo
8: https://www.google.com/intl/en/options/
9: http://www.google.com/history/optout?hl=en
10: /preferences?hl=en
11: https://accounts.google.com/ServiceLogin?hl=en&passive=true&continue=http://www.google.com/
12: /advanced_search?hl=en&authuser=0
13: /language_tools?hl=en&authuser=0
14: https://www.google.com/url?q=https://www.youtube.com/watch%3Fv%3DsmkyorC5qwc&source=hpp&id=5086132&ct=3&usg=AFQjCNGDkcbRPXmyYPvk6xCKx3GFOxvNyw&sa=X&ved=0ahUKEwivmYrAgefPAhXI44MKHbV1Bn8Q8IcBCAU
15: /intl/en/ads/
16: /services/
17: https://plus.google.com/116899029375914044550
18: /intl/en/about.html
19: /intl/en/policies/privacy/
20: /intl/en/policies/terms/
```
