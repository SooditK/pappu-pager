
# Pappu Pager - Your Personal News Assistant

![Pappu Pager](https://i.imgur.com/u2zldfb.png)

[![MIT License](https://img.shields.io/badge/License-MIT-green.svg)](https://choosealicense.com/licenses/mit/)
[![GoLang](https://img.shields.io/github/go-mod/go-version/SooditK/pappu-pager/main)]()

Pappu Pager is a custom Slack bot designed to help you stay up-to-date with the latest news. Powered by NewsCatcherApi, Pappu Pager can search for news articles based on your specific queries and send them directly to your Slack channels or DMs.

With Pappu Pager, you no longer have to waste time browsing through various news sources or websites. Simply send a command to the bot, and it will quickly search and deliver the latest news related to your query.

Whether you're interested in business, sports, entertainment, or politics, Pappu Pager can provide you with a tailored news feed that meets your interests.

## Run Locally

Clone the project

```bash
  git clone git@github.com:SooditK/pappu-pager.git
```

Go to the project directory

```bash
  cd pappu-pager
```

Install dependencies

```bash
  go mod tidy
```
Create Environment Variables

```bash
  cp .env.example .env
```

Start the server

```bash
  go run main.go
```


## Authors

- [@SooditK](https://www.github.com/SooditK)


## Acknowledgements

 - [Slacker](https://github.com/shomali11/slacker)
 - [NewsCatcherApi](https://newscatcherapi.com/)
