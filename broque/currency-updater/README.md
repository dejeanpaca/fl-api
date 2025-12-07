# fl-currency-updater

Fiesta Labs Currency Updater

- Written in Go
- Downloads currency exchange rates from:
  - <https://github.com/fawazahmed0/exchange-api>

# Cron job

Run via cron:
```cron
0 * * * * /root/update-currencies/run
```
