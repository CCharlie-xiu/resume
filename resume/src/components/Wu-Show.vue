<template>
    <div class="wuShow" :style={width}>
        <div class="showCar">
            <div class="wuTitle">{{ title }}</div>
            <div class="wusubTitle">
                <div class="wuUsername">{{ username }}</div>
                <div class="wuDevtitle">{{ devtitle }}</div>
            </div>
        </div>
    </div>
</template>

<script>
  export default {
    name: 'wuShow',
    props: {
        title: {
            type: String,
            default: ''
        },
        username: {
            type: String,
            default: ''
        },
        devtitle: {
            type: String,
            default: ''
        }
    },
  }
</script>

<style lang="less" scoped>
    .wuShow {
        margin: 0;
        margin-top: 30px;
        padding: 0;
        box-sizing: border-box;
        width: 60%;
        min-width: 420px;
        height: 96%;
        min-height: 600px;

        .showCar {
            height: 96%;
            border: 2px solid #2c3e50;

            .wuTitle {
                font-size: 48px;
                font-weight: 800;
                line-height: 48px;
                margin: 12px;
                letter-spacing: 2px;
            }

            .wusubTitle {
                display: flex;
                justify-content: space-between;

                .wuUsername {
                    font-size: 24px;
                    font-weight: 800;
                    line-height: 24px;
                    margin: 12px;
                    letter-spacing: 2px;
                }

                .wuDevtitle {
                    font-size: 24px;
                    font-weight: 800;
                    line-height: 24px;
                    margin: 12px;
                    letter-spacing: 2px;
                }
            }
        }

    }
</style>