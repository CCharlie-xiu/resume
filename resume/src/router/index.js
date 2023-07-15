import { createRouter, createWebHistory } from 'vue-router'

const routes = [
  {
    path: '/resume',
    name: 'Resume',
    component: () => import('../views/Wu-Resume.vue')
  },
  {
    path: '/',
    name: 'Login',
    component: () => import('../views/Wu-Login.vue')
  }
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router
