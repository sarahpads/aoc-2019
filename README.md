# GoLang

## Module Resolution
Go uses two environment varialbes for package resolution: GOROOT (where the GO SDK is located; don't change unless you're changing Go versions) and GOPATH (the local working directory, $HOME/go by default). In order to get local packages to resolve, you need to set GOPATH to the local working directory. I updated this in my .zshrc.

- https://medium.com/rungo/everything-you-need-to-know-about-packages-in-go-b8bac62b74cc
- https://blog.learngoprogramming.com/what-are-goroot-and-gopath-1231b084723f