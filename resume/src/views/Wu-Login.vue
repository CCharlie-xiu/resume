<template>
    <div class="wuLogin">
        <input type="text" placeholder="username" v-model="userInfo.username">
        <input type="text" placeholder="password" v-model="userInfo.password">
        <div class="checkout" @click="login2resume">click</div>
    </div>
  </template>
  
  <script>
    export default {
      name: 'wuLogin',
      data() {
        return {
            userInfo: {
                username: '',
                password: ''
            }
        }
      },
      methods: {
        login2resume() {
            this.postRequest("/login",this.userInfo).then(res => {
                if(res) {
                    console.log(res)
                }
            })
        }
      }
             
    }
  </script>
  
  <style lang="less" scoped>
    .wuLogin {

    }
  </style>