package main

import (
	"encoding/json"
	"fmt"
	"io"

	//"io/ioutil"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

const apiKey = "835ad9c03e0777cb1095985511eaf078"

func main() {
	fmt.Println("Hello, World!")
	router := gin.Default()
//	router.GET("/api/stock/:symbol", getWeatherData)
	router.Use(cors.Default())
	router.GET("/api/weather", getWeatherData)
	router.Run(":8080")
}

func getWeatherData(c *gin.Context) {
    city := c.DefaultQuery("city", "") // 获取请求中的城市（可选）

    // 如果没有提供城市，返回错误
    if city == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "City is required"})
        return
    }

    // 构造 API 请求的 URL
    url := fmt.Sprintf("https://restapi.amap.com/v3/weather/weatherInfo?key=%s&city=%s", apiKey, city)
    fmt.Println(url)

    // 发送请求获取天气数据
    resp, err := http.Get(url)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch weather data"})
        return
    }
    defer resp.Body.Close()

    // 读取返回的 JSON 数据
    body, err := io.ReadAll(resp.Body)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read weather data"})
        return
    }

    // 将 JSON 数据解析为结构体
    var weatherData map[string]interface{}
    if err := json.Unmarshal(body, &weatherData); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse weather data"})
        return
    }

    // 检查返回的数据是否成功
    if status, exists := weatherData["status"]; exists && status == "1" {
        // 修改返回的数据格式，符合前端预期
        lives := weatherData["lives"].([]interface{})[0].(map[string]interface{})  
        c.JSON(http.StatusOK, gin.H{
            "city":          lives["city"],
            "weather":       lives["weather"],
            "temperature":   lives["temperature"],
            "winddirection": lives["winddirection"],
            "windpower":     lives["windpower"],
            "humidity":      lives["humidity"],
            "reporttime":    lives["reporttime"],
        })
    } else {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid city or failed to fetch data"})
    }
}
