import { createRouter, createWebHistory } from 'vue-router';
//import Home from './components/Home.vue';
import Weather from './components/Weather.vue';
import Home1 from './components/Home1.vue';

const routes = [
  {
    path: '/',       // 默认路径，加载 Home 页面
    name: 'home1',
    component: Home1
  },
  {
    path: '/weather',  // 路径为 /weather 显示 Weather 页面
    name: 'weather',
    component: Weather
  }
];

const router = createRouter({
  history: createWebHistory(),
  routes
});

export default router;
