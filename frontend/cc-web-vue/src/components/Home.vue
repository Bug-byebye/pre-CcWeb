<template>
    <div class="container">
      <!-- 顶部导航栏 -->
      <header class="header">
        <div class="logo">
          <img src="@/assets/logo-test.png" alt="网站图标" class="logo-icon"/>
          <div class="site-info">
            <h1 class="site-name">网站名称</h1>
            <p class="site-tagline">这是网站的标语</p>
          </div>
        </div>
        <div class="user-info">
          <img src="@/assets/user-avatar.jpg" alt="用户头像" class="user-avatar"/>
          <span class="username">用户名</span>
        </div>
      </header>
  
      <!-- 中间的可选服务 -->
      <section class="services">
        <button @click="goToWeather">天气查询</button>
        <button @click="goToService(1)">1 + 服务</button>
        <button @click="goToService(2)">2 + 服务</button>
        <button @click="goToService(3)">3 + 服务</button>
        <button @click="goToService(4)">4 + 服务</button>
      </section>
  
      <!-- 底部的相关信息 -->
      <footer class="footer">
        <p>© 2025 示例公司. 版权所有</p>
      </footer>
    </div>
  </template>
  
  <script>
  export default {
    name: 'HomePage0',
    methods: {
      // 跳转到天气查询页面
      goToWeather() {
        this.$router.push({ name: 'weather' });
      },
      // 跳转到其他服务页面（目前是占位用）
      goToService(serviceId) {
        console.log(`跳转到服务 ${serviceId}`);
        // 这里你可以根据不同的 serviceId 跳转到不同的页面
        // 例如：this.$router.push({ name: 'servicePage', params: { id: serviceId } });
      }
    }
  };
  </script>
  
  <style scoped>
  /* 容器整体布局 */
  .container {
    display: flex;
    flex-direction: column;
    height: 100vh;
  }
  
  /* 顶部导航栏样式 */
  .header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    background-color: #4CAF50;
    padding: 15px 30px;
    color: white;
  }
  
  .logo {
    display: flex;
    align-items: center;
  }
  
  .logo-icon {
    width: 40px;
    height: 40px;
    margin-right: 15px;
  }
  
  .site-info {
    display: flex;
    flex-direction: column;
  }
  
  .site-name {
    font-size: 24px;
    font-weight: bold;
    margin: 0;
  }
  
  .site-tagline {
    font-size: 14px;
    margin: 0;
  }
  
  .user-info {
    display: flex;
    align-items: center;
  }
  
  .user-avatar {
    width: 40px;
    height: 40px;
    border-radius: 50%;
    margin-right: 10px;
  }
  
  .username {
    font-size: 16px;
  }
  
  /* 中间的可选服务并排显示 */
  .services {
    display: flex;
    justify-content: space-around;
    padding: 20px 0;
  }
  
  .services button {
    padding: 10px 20px;
    font-size: 16px;
    background-color: #008CBA;
    color: white;
    border: none;
    border-radius: 5px;
    cursor: pointer;
    transition: background-color 0.3s ease;
  }
  
  .services button:hover {
    background-color: #007B9E;
  }
  
  /* 底部信息样式 */
  .footer {
    background-color: #333;
    color: white;
    text-align: center;
    padding: 10px;
  }
  </style>
  