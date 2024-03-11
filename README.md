This tool was created to help you massively deconstruct urls into files in the fastest way possible.

# Install:
```
▶ go install github.com/al4xs/deconstructurl@latest
```
Or if you prefer manually:

```
▶ git clone github.com/al4xs/deconstructurl
▶ cd deconstructurl
▶ go build deconstructurl.go
▶ ./deconstructurl -h
```
# operation
How does it work?

Imagine the following url:
site.com/directory/uploads/pdf/2023/file.pdf

When using "deconstructorurl", it will deconstruct the url in descending order from right to left, so the output would be:

site.com/directory/uploads/pdf/2023/file.pdf
site.com/directory/uploads/pdf/2023/
site.com/directory/uploads/pdf/
site.com/directory/uploads/
site.com/directory/
site.com/

# How to use
Move your binary to the "/usr/bin" folder

Execution:
```
cat urls-with-paths.txt | deconstructurl
```

you can also combine it with other existing tools
```
echo "domain.com" | waybackurls | deconstructurl | anew
```

I hope this binary written in go helps you find vulnerabilities by stretching your crawler file.

by Al4xs
