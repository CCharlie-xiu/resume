<template>
  <div class="wuResume">
    <div class="wu-Input">
      <wuInput
       :width="width"
       @html-pdf="handlerHtml2Pdf"></wuInput>
    </div>
    <div class="wu-Show">
      <wuShow
       :title="confrim.title"
       :username="confrim.username"
       :devtitle="confrim.devtitle"
      ></wuShow>
    </div>
  </div>
</template>

<script>
  import WuInput from '@/components/Wu-Input.vue';
  import WuShow from '@/components/Wu-Show.vue';
  export default {
    name: 'wuResume',
    data () {
      return {
        confrim: [],
        width: '300px'
      }
    },
    components: {
      WuInput,WuShow
    },
    methods: {
        handlerHtml2Pdf (confirm) {
          this.confrim = confirm;
          console.log("this is 2",this.confrim)
        }
      }
  }
</script>

<style lang="less" scoped>
  .wuResume {
    display: flex;
    .wu-Input, .wu-Show{
      float: left;
      width: 50%;
    }
  }
</style>