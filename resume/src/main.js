import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import '../src/api/init.js'

router.beforeEach((to, from, next) => {
    const isAuthenticated = localStorage.getItem('isAuthenticated')
    if (to.name !== 'Login' && !isAuthenticated) {
        next({ name: 'Login' })
    } else {
        next()
    }
})

const app = createApp(App)

import { postRequest, getRequest, putRequest, deleteRequest } from '../src/api/init.js'

app.config.globalProperties.getRequest = getRequest
app.config.globalProperties.postRequest = postRequest
app.config.globalProperties.putRequest = putRequest
app.config.globalProperties.deleteRequest = deleteRequest

createApp(App).use(store).use(router).mount('#app')
