# faas stonks

This uses serverless technology to get a stock quote and 30 day sparkline from Yahoo Finance.

# Deployment

- Nimbella account
- Namespace with object storage enabled
- Deploy command: `nim project deploy stonks --remote-build`

> Note: the `--remote-build` is required to build this Go function into a binary before deploying the action. 

# Development

- Go 1.15 or later
- Run `dev/run` 
- Running at [localhost:8080](http://localhost:8080/)

> Note: dev/run includes the build tag `dev`, this includes the main function file which conflicts during builds in production

## Special thanks

- [wcharczuk/go-chart](https://github.com/wcharczuk/go-chart)
- [wtfutil/wtf - for Yahoo trend idea](https://github.com/wtfutil/wtf/blob/master/modules/stocks/yfinance/yquote.go#L109)
- [piquette/finance-go](https://github.com/piquette/finance-go)

[![DigitalOcean Referral Badge](https://web-platforms.sfo2.cdn.digitaloceanspaces.com/WWW/Badge%201.svg)](https://www.digitalocean.com/?refcode=cd77e6593231&utm_campaign=Referral_Invite&utm_medium=Referral_Program&utm_source=badge)