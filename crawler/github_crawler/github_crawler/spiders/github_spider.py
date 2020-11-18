import scrapy


# TODO: move to a config file
members = ['bill0412']


class GithubSpider(scrapy.Spider):
    name = "github"

    def start_requests(self):
        urls = ['https://github.com/{}/mit-test-repo/tree/master/week1'.format(member) for member in members]
        for url in urls:
            yield scrapy.Request(url=url, callback=self.parse)

    def parse(self, response):
        res = response.css('a::attr(title)').getall()
        print(res)
