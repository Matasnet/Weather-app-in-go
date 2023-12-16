# Weather-app-in-go
This is an application that allows you to check the weather in various places around the world. The application uses the GO language and API from OpenWeather, which provides weather data in JSON format. The application parses the JSON file and displays the temperature and the name of the city for which the weather was downloaded.

## How use it ?
First, you need to register at https://openweathermap.org/ and obtain your API key.
Your API key you must past to key := "", between quotation marks eg. key := "1234567"

In variable "place" you can write interesting you place. Remember to use + instead of spaces.

When everything is read you can open eg. Visual Studio Code and in terminal write 
go run app.go

## What units I should use

In variable url you can see &units=metric 
You can use: standard (Kelvin), metric(Celsius) or imperial (Fahrenheit) units.
Default are standard, so if you delet +&units=metric you can get temperature in Kelvin. 
