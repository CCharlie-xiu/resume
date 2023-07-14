<template>
    <div class="wuInput" :style={width}>
        <input type="text" placeholder="title" v-model="confirm.title">
        <input type="text" placeholder="username" v-model="confirm.username">
        <input type="text" placeholder="devtitle" v-model="confirm.devtitle">
        <div class="checkout" @click="html2pdf">click</div>
    </div>
</template>

<script>

  export default {
    name: 'wuInput',
    data() {
        return {
            confirm: {
                title: '',
                username: '',
                devtitle: '',
            }
        }
    },
    props: {
        width: {
            type: String,
            default: '300px'
        },
    },
    // computed() {
    //     if(this.title.trim() == "" || this.username.trim() == "") {
    //         this.resume = false;
    //     } else {
    //         this.confirm.title = this.title.toUpperCase()
    //         this.confirm.username = this.username
    //         this.confirm.devtitle = this.devtitle
    //         localStorage.setItem('confirm',this.confirm)
    //     }
    // },
    methods: {
        html2pdf() {
            const confirm = this.$data.confirm
            console.log(confirm)
            this.$emit('html-pdf', confirm)
        }
    },
  }
</script>

<style lang="less" scoped>
    .wuInput {
        margin: 0;
        padding: 0;
        box-sizing: border-box;

        input {
            width: 100%;
            flex-wrap: wrap;
            flex: 1;
            margin-top: 6px;
        }

        .checkout {
            padding: 10px 5px;
            width: 100px;
            height: 32px;
        }
    }
</style>