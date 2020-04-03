# Subgen - a subdomain concating-utility with some smarts
If you've come to the realisation that you'd like to split up your DNS enumeration methodology by generating combinations and/or pulling passive data from other sources, then subgen is the sorter & concator you're probably looking for


### Installation

If you have a properly configured GOPATH and $GOPATH/bin is in your PATH, then run this command for a one-liner install, thank you golang!
```
go get -u github.com/pry0cc/subgen
```


#### Cat a very large unsorted wordlist.
```
cat wordlist.txt | subgen -d "uber.com" | zdns A | jq -r "select(.data.answers[0].name) | .name" 
```

#### Crawl website with Cewl and generated unresolved subdomain combinations
```
cewl.rb -d 3 -u https://uber.com/ | subgen -d "uber.com"
```

### tldr;
This will take a wordlist, concate with a domain, sort unique in real time (filtering lines that have already been produced) and filtering out with regex any non-DNS safe characters like special characters. 

`wordlist.txt`
```
admin
dashboard
www
helloworld
HELLOWORLD
helLoWorld
autodiscover
aUtOdiscover
*sd&^%$£$%^&*
zabbix
login
```

```
cat wordlist.txt | subgen -d "uber.com"
```

Will produce:

```
admin.uber.com
dashboard.uber.com
www.uber.com
helloworld.uber.com
autodiscover.uber.com
zabbix.uber.com
login.uber.com
```

Which you can pipe to a resolver tool such as MassDNS or ZDNS and print out resolved subdomains.

```
cat ~/lists/sorted-all.txt | subgen -d uber.com | zdns A | jq -r "select(.data.answers[0].name) | .name" 
```


```
cat ~/lists/jhaddix-all.txt | subgen -d uber.com |  massdns -r dns.txt -t A -o S -w results.txt

```


