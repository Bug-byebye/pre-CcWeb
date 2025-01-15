<template>
  <div id="app" class="container">
    <h1>天气查询</h1>
    <input 
      v-model="city" 
      placeholder="请输入城市编码" 
      @keyup.enter="fetchWeatherData" 
      class="input-box"
    />
    
    <div v-if="weatherData" class="weather-info">
      <h2>{{ weatherData.city }} - 实时天气</h2>
      <p>天气现象: {{ weatherData.weather }}</p>
      <p>温度: {{ weatherData.temperature }}°C</p>
      <p>风向: {{ weatherData.wind }}</p>
      <p>湿度: {{ weatherData.humidity }}%</p>
      <p>更新时间: {{ weatherData.reporttime }}</p>
    </div>

    <div v-if="errorMessage" class="error">{{ errorMessage }}</div>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  name: 'WeatherPage',
  data() {
    return {
      city: '', // 用户输入的城市编码
      weatherData: null, // 存储天气数据
      errorMessage: '', // 错误信息
    };
  },
  methods: {
    async fetchWeatherData() {
      if (this.city) {
        try {
          // 修改为请求 API Gateway 的接口
          const response = await axios.get(`http://localhost:8080/weather?city=${this.city}`);
          console.log(response.data); // 打印返回的数据，查看其结构

          // 直接获取天气数据并赋值给 weatherData
          if (response.data && response.data.city) {
            this.weatherData = response.data; // 直接把整个返回的数据赋给 weatherData
            this.errorMessage = ''; // 清空错误信息
          } else {
            this.errorMessage = '未能获取有效的天气数据';
            this.weatherData = null;
          }
        } catch (error) {
          console.error(error);
          this.errorMessage = '无法获取天气数据，请检查城市编码或重试。';
        }
      } else {
        this.errorMessage = '请输入有效的城市编码';
      }
    }
  }
};
</script>

<style scoped>
/* 背景图片和布局居中 */
.container {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  height: 100vh;
  background: url('../assets/深蓝背景.png') no-repeat center center;
  background-size: cover;
  text-align: center;
}

/* 标题样式 */
h1 {
  font-size: 6em;
  color: #FFF;
}

/* 输入框样式 */
.input-box {
  width: 300px;
  padding: 10px;
  margin: 20px 0;
  border: 2px solid blue;
  border-radius: 5px;
  outline: none;
  font-size: 16px;
}

/* 天气信息样式 */
.weather-info {
  background-color: rgba(255, 255, 255, 0.8);
  padding: 20px;
  border-radius: 10px;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
}

/* 错误信息样式 */
.error {
  color: red;
  margin-top: 10px;
}
</style>

  