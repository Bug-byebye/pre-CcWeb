<!--
<template>
  <div id="app">
    <h1>天气查询</h1>
    <input v-model="city" placeholder="请输入城市编码" @keyup.enter="fetchWeatherData" />
    
    <div v-if="weatherData">
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

<style>
.error {
  color: red;
}
</style>
-->


<template>
  <div id="app">
    <router-view></router-view>
  </div>
</template>

<script>
export default {
  name: 'App',
};
</script>

<style>
/* 全局样式 */
.error {
  color: red;
}
/* 基础设置 */
body,h1,h2,h3,h4,h5,h6,hr,p,blockquote,dl,dt,dd,ul,ol,li,pre,form,fieldset,legend,button,input,textarea,th,td{
    margin: 0;
    padding: 0;
}

ul,ol {
    list-style: none;
}

img {
    /* 底部留白 */
    display: block;
	border:0;
}

b,strong {
    font-weight: 400;
}

h1,h2,h3,h4,h5,h6 {
    /* 父元素字号的百分比 */
    font-size: 100%;
    font-weight: normal;
}

i,em {
    /* 不倾斜 */
    font-style: normal;
}

u,ins,s,del {
    /* 去掉中划线和下划线 */
    text-decoration: none;
}

table {
    border: 1px solid #999;
    /* 相当于是cellspacing */
    border-spacing: 0;
    /* 1px边框 */
    border-collapse: collapse;
}

td,th {
    border: 1px solid #999;
}

input,button {
    /* 去掉轮廓线 */
    outline: none;
    border:none;
}

/* 风格设置 */
body {
    font: 12px/1.5 "Microsoft YaHei", Tahoma, Helvetica, Arial, "\5b8b\4f53", sans-serif;
    color: #333;
}

a {
    text-decoration: none;
   color: #666;
}

a:hover {
    color:#DD302D;
    text-decoration: none;
}

.leftfix {
    float: left;
}

.rightfix {
    float: right;
}

.clearfix::after {
    content: "";
    display: block;
    clear: both;
}

</style>
