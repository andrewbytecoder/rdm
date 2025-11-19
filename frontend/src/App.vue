<script lang="ts" setup>
import HelloWorld from './components/HelloWorld.vue'

import { useThemeVars } from 'naive-ui'
const themeVars = useThemeVars()
</script>

<template>
  <img id="logo" alt="Wails logo" src="./assets/images/logo-universal.png"/>
  <HelloWorld/>
</template>


<style lang="scss">
#app-container {
  height: 100%;
  overflow: hidden;
  border-top: var(--border-color) 1px solid;
  box-sizing: border-box;

  #app-toolbar {
    height: 40px;
    border-bottom: var(--border-color) 1px solid;
  }

  #app-side {
    //overflow: hidden;
    height: 100%;

    .resize-divider {
      //height: 100%;
      width: 2px;
      border-left-width: 5px;
      background-color: var(--border-color);
    }

    .resize-divider-hover {
      width: 5px;
    }

    .resize-divider-drag {
      //background-color: rgb(0, 105, 218);
      width: 5px;
      //background-color: var(--el-color-primary);
      background-color: v-bind('themeVars.primaryColor');
    }
  }
}

.dragging {
  cursor: col-resize !important;
}
</style>