# CurrencyConverter
This is the mini project to fetch the currency rates  

How to Run 
- Run main.go File (This will start the server on 8080 port)
- Hit API from postman 
- FetchCurrency Rates :  http://localhost:8080/currencyconverter/base/{base}/target/{target}/amount/{amount}
      Example : http://localhost:8080/fetchCurrentRate/base/USD/target/EUR
      here base is USD and target is EUR
- Currency Converter with given amount : this API will give us converted amount in target currency  : http://localhost:8080/currencyconverter/base/{base}/target/{target}/amount/{amount}
       Example : http://localhost:8080/currencyconverter/base/USD/target/EUR/amount/100


- Fetched currancy rates will be cached for 3 hours