<template>
  <div id="app">
    <h1>天气查询</h1>
    <input v-model="city" placeholder="请输入城市编码" @keyup.enter="fetchWeatherData" />
    
    <div v-if="weatherData">
      <h2>{{ weatherData.city }} - 实时天气</h2>
      <p>天气现象: {{ weatherData.weather }}</p>
      <p>温度: {{ weatherData.temperature }}°C</p>
      <p>风向: {{ weatherData.winddirection }}</p>
      <p>风力: {{ weatherData.windpower }}</p>
      <p>湿度: {{ weatherData.humidity }}</p>
      <p>更新时间: {{ weatherData.reporttime }}</p>
    </div>

    <div v-if="errorMessage" class="error">{{ errorMessage }}</div>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  name: 'App',
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
      const response = await axios.get(`http://localhost:8080/api/weather?city=${this.city}`);
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

<style>
.error {
  color: red;
}
</style>
