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

When using "deconstructurl", it will deconstruct the url in descending order from right to left, so the output would be:

```
site.com/directory/uploads/pdf/2023/file.pdf
site.com/directory/uploads/pdf/2023/
site.com/directory/uploads/pdf/
site.com/directory/uploads/
site.com/directory/
site.com/
```

# How to use
Move your binary to the "/usr/bin" folder

Execution:
```
cat urls-with-paths.txt | deconstructurl
```

you can also combine it with other existing tools
```
echo "domain.com" | waybackurls | deconstructurl | httpx -title -td
```

I hope this binary written in go helps you find vulnerabilities by stretching your crawler file.

## 🦸 Autor

 <a href="https://github.com/al4xs">
 <img style="border-radius: 50%;" src="https://avatars.githubusercontent.com/u/40411471?v=4" width="100px;" alt=""/>
 <br />
 <sub><b>AL4XS ETH1C4L H4CK3R</b></sub></a> <a href="http://al4xs.github.io/" title="Github Personal Blog"> 🚀</a>
 <br />

[![Instagram Badge](https://img.shields.io/badge/-@michaelferral4xs-1ca0f1?style=flat-square&labelColor=1ca0f1&logo=instagram&logoColor=white&link=https://instagram.com/michaelferral4xs)](https://instagram.com/michaelferral4xs) 
[![Linkedin Badge](https://img.shields.io/badge/-Al4xs-blue?style=flat-square&logo=Linkedin&logoColor=white&link=https://www.linkedin.com/in/michael-al4xs/)](https://www.linkedin.com/in/michael-al4xs/) 
[![protonmail Badge](https://img.shields.io/badge/-@al4xs@protonmail.com-c14438?style=flat-square&logo=protonmail&logoColor=white&link=mailto:al4xs@protonmail.com)](mailto:al4xs@protonmail.com)

---

## 📝 Licença

Feito com ❤️  por Al4xs 👋🏽 [Entre em contato!](https://www.linkedin.com/in/michael-al4xs/)

---
