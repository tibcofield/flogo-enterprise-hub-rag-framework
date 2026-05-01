# Flogo Activity Workshop: Build a Weather Checker Activity 🌤️

*A hands-on workshop to create your first Flogo activity in 30 minutes*

## Welcome to Your First Flogo Activity Workshop! 👋

Instead of just reading about Flogo activities, let's build one together! In this workshop, you'll create a "Weather Checker" activity that fetches current weather data for any city. By the end, you'll have a working activity and understand exactly how Flogo activities work.

## What We're Building Today 🎯

**The "Weather Checker" Activity:**
- Takes a city name as input
- Fetches weather data from a free API
- Returns temperature, description, and humidity
- Handles errors gracefully
- Can be configured with different temperature units

**Why this example?**
- It's practical and useful
- Demonstrates real API integration
- Shows error handling
- Easy to test and understand

## Workshop Prerequisites ✅

Before we start, make sure you have:
- [ ] Go installed (version 1.19+) - [Download here](https://golang.org/dl/)
- [ ] A text editor or IDE (VS Code recommended)
- [ ] Internet connection (for API calls)
- [ ] 30 minutes of focused time
- [ ] A curious mind! 🧠

**Optional but helpful:**
- Basic Go programming knowledge
- Understanding of REST APIs
- Familiarity with JSON

## Part 1: Project Setup (5 minutes) 🏗️

### Step 1: Create Your Workshop Space

```bash
# Create the workshop directory
mkdir flogo-weather-workshop
cd flogo-weather-workshop

# Create the activity structure
mkdir -p weatherChecker/{icons,sample,test}
cd weatherChecker

# Initialize Go module
go mod init flogo-weather-workshop/weatherChecker
```

**What just happened?**
- Created a dedicated workshop space
- Set up the standard Flogo activity folder structure
- Initialized a Go module for dependency management

### Step 2: Get Your Free API Key

We'll use OpenWeatherMap's free API:

1. Go to [https://openweathermap.org/api](https://openweathermap.org/api)
2. Sign up for a free account
3. Get your API key from the dashboard
4. Keep it handy - we'll use it soon!

*Don't worry, it's completely free and takes 2 minutes to set up.*

## Part 2: Define the Activity Contract (5 minutes) 📋

### Create activity.json

This file tells Flogo what our activity looks like:

```json
{
  "name": "weather-checker",
  "version": "1.0.0",
  "type": "flogo:activity",
  "title": "Weather Checker",
  "author": "Workshop Participant",
  "description": "Fetches current weather data for any city",
  "homepage": "https://github.com/yourname/weather-checker",
  "display": {
    "category": "Weather",
    "visible": true,
    "description": "Get current weather information for any city worldwide"
  },
  "settings": [
    {
      "name": "apiKey",
      "type": "string",
      "required": true,
      "display": {
        "name": "API Key",
        "description": "OpenWeatherMap API key",
        "type": "password"
      }
    },
    {
      "name": "units",
      "type": "string",
      "required": false,
      "value": "metric",
      "allowed": ["metric", "imperial", "kelvin"],
      "display": {
        "name": "Temperature Units",
        "description": "Units for temperature display"
      }
    },
    {
      "name": "timeout",
      "type": "integer",
      "required": false,
      "value": 10,
      "display": {
        "name": "Timeout (seconds)",
        "description": "API request timeout"
      }
    }
  ],
  "inputs": [
    {
      "name": "city",
      "type": "string",
      "required": true,
      "display": {
        "name": "City Name",
        "description": "Name of the city to get weather for"
      }
    },
    {
      "name": "countryCode",
      "type": "string",
      "required": false,
      "display": {
        "name": "Country Code",
        "description": "ISO 3166 country code (optional, e.g., 'US', 'UK')"
      }
    }
  ],
  "outputs": [
    {
      "name": "temperature",
      "type": "number",
      "display": {
        "name": "Temperature",
        "description": "Current temperature"
      }
    },
    {
      "name": "description",
      "type": "string", 
      "display": {
        "name": "Weather Description",
        "description": "Weather condition description"
      }
    },
    {
      "name": "humidity",
      "type": "integer",
      "display": {
        "name": "Humidity",
        "description": "Humidity percentage"
      }
    },
    {
      "name": "cityFound",
      "type": "string",
      "display": {
        "name": "City Found",
        "description": "Actual city name found by the API"
      }
    },
    {
      "name": "success",
      "type": "boolean",
      "display": {
        "name": "Success",
        "description": "Whether the weather data was retrieved successfully"
      }
    }
  ]
}
```

**💡 Workshop Insight:** Notice how this contract is more detailed than our text processor? That's because this activity will integrate with a real API and needs more configuration options.

## Part 3: Create Data Structures (5 minutes) 📦

### Create metadata.go

```go
package weatherChecker

import (
	"github.com/project-flogo/core/data/coerce"
)

// Field name constants
const (
	// Settings
	sAPIKey  = "apiKey"
	sUnits   = "units"
	sTimeout = "timeout"
	
	// Inputs
	iCity        = "city"
	iCountryCode = "countryCode"
	
	// Outputs
	oTemperature = "temperature"
	oDescription = "description"
	oHumidity    = "humidity"
	oCityFound   = "cityFound"
	oSuccess     = "success"
)

// Settings holds configuration for the weather service
type Settings struct {
	APIKey  string `md:"apiKey,required"`
	Units   string `md:"units"`
	Timeout int    `md:"timeout"`
}

// Input represents the data coming into our activity
type Input struct {
	City        string `md:"city"`
	CountryCode string `md:"countryCode"`
}

// Output represents the weather data we return
type Output struct {
	Temperature float64 `md:"temperature"`
	Description string  `md:"description"`
	Humidity    int     `md:"humidity"`
	CityFound   string  `md:"cityFound"`
	Success     bool    `md:"success"`
}

// WeatherResponse represents the API response structure
type WeatherResponse struct {
	Weather []struct {
		Main        string `json:"main"`
		Description string `json:"description"`
	} `json:"weather"`
	Main struct {
		Temp     float64 `json:"temp"`
		Humidity int     `json:"humidity"`
	} `json:"main"`
	Name string `json:"name"`
	Cod  int    `json:"cod"`
	Message string `json:"message,omitempty"`
}

// FromMap converts map data to Settings struct
func (s *Settings) FromMap(values map[string]interface{}) error {
	if values == nil {
		s.Units = "metric"
		s.Timeout = 10
		return nil
	}

	var err error
	s.APIKey, err = coerce.ToString(values[sAPIKey])
	if err != nil {
		return err
	}

	if units, ok := values[sUnits]; ok && units != nil {
		s.Units, err = coerce.ToString(units)
		if err != nil {
			return err
		}
	} else {
		s.Units = "metric"
	}

	if timeout, ok := values[sTimeout]; ok && timeout != nil {
		s.Timeout, err = coerce.ToInt(timeout)
		if err != nil {
			return err
		}
	} else {
		s.Timeout = 10
	}

	return nil
}

// FromMap converts map data to Input struct
func (i *Input) FromMap(values map[string]interface{}) error {
	if values == nil {
		return nil
	}

	var err error
	i.City, err = coerce.ToString(values[iCity])
	if err != nil {
		return err
	}

	if countryCode, ok := values[iCountryCode]; ok && countryCode != nil {
		i.CountryCode, err = coerce.ToString(countryCode)
		if err != nil {
			return err
		}
	}

	return nil
}

// ToMap converts Input struct to map
func (i *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		iCity:        i.City,
		iCountryCode: i.CountryCode,
	}
}

// ToMap converts Output struct to map
func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		oTemperature: o.Temperature,
		oDescription: o.Description,
		oHumidity:    o.Humidity,
		oCityFound:   o.CityFound,
		oSuccess:     o.Success,
	}
}
```

**💡 Workshop Insight:** We added a `WeatherResponse` struct to handle the JSON response from the weather API. This shows how to work with external data structures.

## Part 4: Implement the Core Logic (10 minutes) ⚡

### Create activity.go

```go
package weatherChecker

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/support/log"
)

// Logger for this activity
var logger = log.ChildLogger(log.RootLogger(), "weather-checker")

// Activity metadata
var activityMd = activity.ToMetadata(&Settings{}, &Input{}, &Output{})

// Activity struct
type Activity struct {
	Settings   *Settings
	httpClient *http.Client
}

// Register the activity
func init() {
	_ = activity.Register(&Activity{}, New)
}

// New creates a new Weather Checker activity instance
func New(ctx activity.InitContext) (activity.Activity, error) {
	s := &Settings{}
	err := s.FromMap(ctx.Settings())
	if err != nil {
		return nil, err
	}

	// Validate API key
	if s.APIKey == "" {
		return nil, fmt.Errorf("API key is required")
	}

	// Create HTTP client with timeout
	httpClient := &http.Client{
		Timeout: time.Duration(s.Timeout) * time.Second,
	}

	logger.Info("Weather Checker activity initialized successfully")

	return &Activity{
		Settings:   s,
		httpClient: httpClient,
	}, nil
}

// Metadata returns the activity metadata
func (a *Activity) Metadata() *activity.Metadata {
	return activityMd
}

// Eval executes the weather checking logic
func (a *Activity) Eval(ctx activity.Context) (done bool, err error) {
	logger.Info("🌤️  Weather Checker starting...")

	// Get input data
	input := &Input{}
	err = ctx.GetInputObject(input)
	if err != nil {
		return false, fmt.Errorf("failed to get input: %w", err)
	}

	// Validate input
	if strings.TrimSpace(input.City) == "" {
		return false, fmt.Errorf("city name cannot be empty")
	}

	logger.Infof("Checking weather for: %s", input.City)

	// Build the API URL
	baseURL := "https://api.openweathermap.org/data/2.5/weather"
	params := url.Values{}
	params.Add("q", buildLocationQuery(input.City, input.CountryCode))
	params.Add("appid", a.Settings.APIKey)
	params.Add("units", a.Settings.Units)

	fullURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())

	// Make the API request
	weatherData, err := a.fetchWeatherData(fullURL)
	if err != nil {
		// Return a "failed" output instead of erroring
		output := &Output{
			Success: false,
			Description: fmt.Sprintf("Failed to fetch weather: %s", err.Error()),
		}
		ctx.SetOutputObject(output)
		logger.Errorf("Weather fetch failed: %v", err)
		return true, nil // Return true but with success=false
	}

	// Create successful output
	output := &Output{
		Temperature: weatherData.Main.Temp,
		Description: weatherData.Weather[0].Description,
		Humidity:    weatherData.Main.Humidity,
		CityFound:   weatherData.Name,
		Success:     true,
	}

	err = ctx.SetOutputObject(output)
	if err != nil {
		return false, fmt.Errorf("failed to set output: %w", err)
	}

	logger.Infof("✅ Weather data retrieved: %.1f°, %s, %d%% humidity", 
		output.Temperature, output.Description, output.Humidity)

	return true, nil
}

// fetchWeatherData calls the OpenWeatherMap API
func (a *Activity) fetchWeatherData(url string) (*WeatherResponse, error) {
	resp, err := a.httpClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned status %d", resp.StatusCode)
	}

	var weatherResp WeatherResponse
	if err := json.NewDecoder(resp.Body).Decode(&weatherResp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	// Check for API errors
	if weatherResp.Cod != 200 {
		return nil, fmt.Errorf("API error: %s", weatherResp.Message)
	}

	if len(weatherResp.Weather) == 0 {
		return nil, fmt.Errorf("no weather data in response")
	}

	return &weatherResp, nil
}

// buildLocationQuery creates the location query string
func buildLocationQuery(city, countryCode string) string {
	if countryCode != "" {
		return fmt.Sprintf("%s,%s", city, countryCode)
	}
	return city
}
```

**💡 Workshop Insight:** This activity demonstrates several important patterns:
- HTTP client initialization with timeout
- Error handling that doesn't crash the flow
- Real API integration
- Input validation
- Structured logging

## Part 5: Test Your Activity (5 minutes) 🧪

### Create activity_test.go

```go
package weatherChecker

import (
	"testing"
	"time"

	"github.com/project-flogo/core/activity"
	"github.com/stretchr/testify/assert"
)

func TestWeatherChecker_Success(t *testing.T) {
	// Skip if no API key in environment
	apiKey := "your-api-key-here" // Replace with your actual API key for testing
	if apiKey == "your-api-key-here" {
		t.Skip("Set your API key to run this test")
	}

	// Create activity settings
	settings := map[string]interface{}{
		"apiKey": apiKey,
		"units":  "metric",
		"timeout": 30,
	}

	// Create test input
	input := map[string]interface{}{
		"city": "London",
		"countryCode": "UK",
	}

	// Set up test context
	ctx := activity.NewContext(activityMd, nil, nil)
	for key, value := range settings {
		ctx.Settings().SetValue(key, value)
	}
	for key, value := range input {
		ctx.SetInput(key, value)
	}

	// Create activity instance
	act, err := New(ctx)
	assert.NoError(t, err)

	// Execute the activity
	done, err := act.Eval(ctx)

	// Verify results
	assert.NoError(t, err)
	assert.True(t, done)

	success := ctx.GetOutput("success").(bool)
	assert.True(t, success)

	if success {
		temp := ctx.GetOutput("temperature")
		desc := ctx.GetOutput("description")
		humidity := ctx.GetOutput("humidity")
		cityFound := ctx.GetOutput("cityFound")

		assert.NotNil(t, temp)
		assert.NotEmpty(t, desc)
		assert.NotNil(t, humidity)
		assert.NotEmpty(t, cityFound)

		t.Logf("Weather for %s: %.1f°C, %s, %d%% humidity", 
			cityFound, temp, desc, humidity)
	}
}

func TestWeatherChecker_InvalidCity(t *testing.T) {
	// Test with invalid city
	settings := map[string]interface{}{
		"apiKey": "test-key", // This will fail, but we test the error handling
		"units":  "metric",
	}

	input := map[string]interface{}{
		"city": "", // Empty city should fail validation
	}

	ctx := activity.NewContext(activityMd, nil, nil)
	for key, value := range settings {
		ctx.Settings().SetValue(key, value)
	}
	for key, value := range input {
		ctx.SetInput(key, value)
	}

	act, err := New(ctx)
	assert.NoError(t, err)

	// This should fail due to empty city
	done, err := act.Eval(ctx)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "city name cannot be empty")
}

func TestWeatherChecker_NoAPIKey(t *testing.T) {
	// Test initialization without API key
	settings := map[string]interface{}{
		"units": "metric",
	}

	ctx := activity.NewContext(activityMd, nil, nil)
	for key, value := range settings {
		ctx.Settings().SetValue(key, value)
	}

	// Should fail during initialization
	_, err := New(ctx)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "API key is required")
}
```

### Run Your Tests

```bash
# Get dependencies
go mod tidy

# Run tests (some will be skipped without API key)
go test -v

# Run with your API key (replace with actual key)
# go test -v -args -apikey=your-actual-api-key
```

**💡 Workshop Insight:** Testing shows different scenarios: success, validation errors, and initialization failures.

## Part 6: Try It Out! (Additional time) 🎮

### Manual Testing

Create a simple test file `manual_test.go`:

```go
package main

import (
	"fmt"
	"log"

	"flogo-weather-workshop/weatherChecker"
	"github.com/project-flogo/core/activity"
)

func main() {
	// Replace with your actual API key
	apiKey := "your-api-key-here"
	
	if apiKey == "your-api-key-here" {
		fmt.Println("Please set your API key in the code first!")
		return
	}

	// Create activity settings
	settings := map[string]interface{}{
		"apiKey": apiKey,
		"units":  "metric",
		"timeout": 30,
	}

	// Test different cities
	cities := []map[string]interface{}{
		{"city": "Tokyo", "countryCode": "JP"},
		{"city": "New York", "countryCode": "US"},
		{"city": "Paris", "countryCode": "FR"},
		{"city": "Sydney", "countryCode": "AU"},
	}

	activityMd := activity.ToMetadata(&weatherChecker.Settings{}, &weatherChecker.Input{}, &weatherChecker.Output{})

	for _, cityData := range cities {
		fmt.Printf("\n🌍 Checking weather for %s, %s...\n", cityData["city"], cityData["countryCode"])
		
		ctx := activity.NewContext(activityMd, nil, nil)
		for key, value := range settings {
			ctx.Settings().SetValue(key, value)
		}
		for key, value := range cityData {
			ctx.SetInput(key, value)
		}

		act, err := weatherChecker.New(ctx)
		if err != nil {
			log.Printf("Failed to create activity: %v", err)
			continue
		}

		done, err := act.Eval(ctx)
		if err != nil {
			log.Printf("Activity failed: %v", err)
			continue
		}

		if done {
			success := ctx.GetOutput("success").(bool)
			if success {
				temp := ctx.GetOutput("temperature")
				desc := ctx.GetOutput("description")
				humidity := ctx.GetOutput("humidity")
				cityFound := ctx.GetOutput("cityFound")

				fmt.Printf("✅ %s: %.1f°C, %s, %d%% humidity\n", 
					cityFound, temp, desc, humidity)
			} else {
				desc := ctx.GetOutput("description")
				fmt.Printf("❌ Failed: %s\n", desc)
			}
		}
	}
}
```

Run it with:
```bash
go run manual_test.go
```

## Workshop Wrap-Up 🎯

### What You've Accomplished

🎉 **Congratulations!** In this workshop, you've:

✅ **Built a real-world Flogo activity** that integrates with an external API  
✅ **Learned proper error handling** for production scenarios  
✅ **Implemented configuration management** with multiple settings  
✅ **Created comprehensive tests** including edge cases  
✅ **Used professional Go patterns** like HTTP client management  
✅ **Handled real JSON data** from an external service  

### Key Concepts Mastered

🔧 **Activity Structure**
- JSON contract definition
- Go struct design for complex data
- Proper initialization and validation

🌐 **API Integration**
- HTTP client configuration
- URL building and parameter handling
- JSON response parsing
- Error handling for network issues

🧪 **Testing Strategies**
- Unit tests with mocked data
- Integration tests with real APIs
- Error scenario testing
- Manual testing approaches

### Differences from the Text Processor

This workshop showed you more advanced patterns:

| Text Processor | Weather Checker |
|---|---|
| Simple string manipulation | External API integration |
| Basic error handling | Production-grade error handling |
| No network calls | HTTP client management |
| Simple configuration | Multiple configuration options |
| Instant execution | Timeout handling |

### Next Steps 🚀

Now that you've mastered the fundamentals, try these challenges:

**🎯 Immediate Challenges**
1. Add support for 5-day forecast
2. Add geolocation-based weather lookup
3. Create a weather alert system
4. Add caching to reduce API calls

**🎯 Advanced Challenges**
1. Build a database connector activity
2. Create a file processing activity
3. Build an email sender activity
4. Design a data transformation activity

**🎯 Real-World Applications**
- Use your weather activity in actual Flogo flows
- Combine it with conditional logic for smart home automation
- Create alerts based on weather conditions
- Build weather-based decision flows

### Resources for Continued Learning

📚 **Documentation**
- [Flogo Official Documentation](https://docs.flogo.io/)
- [Go Programming Language](https://golang.org/doc/)
- [REST API Best Practices](https://restfulapi.net/)

🛠️ **Tools and Libraries**
- [VS Code Go Extension](https://marketplace.visualstudio.com/items?itemName=golang.Go)
- [Postman for API Testing](https://www.postman.com/)
- [Go Testing Framework](https://golang.org/pkg/testing/)

👥 **Community**
- [Flogo GitHub Repository](https://github.com/project-flogo)
- [Go Community Forums](https://forum.golangbridge.org/)
- [Stack Overflow - Flogo Tag](https://stackoverflow.com/questions/tagged/flogo)

### Workshop Feedback

How did this workshop work for you? Consider:
- Was the pacing right?
- Were the examples clear?
- What would you like to see in future workshops?
- What additional topics interest you?

**Remember:** Every expert started as a beginner. You've taken a significant step forward in Flogo activity development. Keep building, keep learning, and don't hesitate to experiment with new ideas!

Happy coding! 🚀🌟