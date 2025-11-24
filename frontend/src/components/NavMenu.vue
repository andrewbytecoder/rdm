<script setup lang="ts">
import { computed, h, ref, Ref } from 'vue'
import { NIcon, useThemeVars, MenuOption, DropdownOption } from 'naive-ui'
import ToggleDb from './icons/ToggleDb.vue'
import { useI18n } from 'vue-i18n'
import ToggleServer from './icons/ToggleServer.vue'
import IconButton from './common/IconButton.vue'
import Config from './icons/Config.vue'
import useDialogStore from '../stores/dialog'
import Github from './icons/Github.vue'
import { BrowserOpenURL } from '../../wailsjs/runtime'
import Log from './icons/Log.vue'
import useConnectionStore from '../stores/connections.js'
import Update from "./icons/Update.vue";

// 类型定义
type MenuItemKey = 'structure' | 'server' | 'preferences' | 'about' | 'update'

const themeVars = useThemeVars()

const props = defineProps({
  value: {
    type: String,
    default: 'server'
  },
  width: {
    type: Number,
    default: 60
  },
})

const emit = defineEmits(['update:value'])
const iconSize = computed(() => Math.floor(props.width * 0.4))

// 渲染图标函数
const renderIcon = (icon: any) => {
  return () => h(NIcon, null, { default: () => h(icon) })
}
const connectionStore = useConnectionStore()
const i18n = useI18n()

// 主菜单选项, 计算属性，必函数强的点是只有监听的变量改变时才会重新计算
const menuOptions = computed<MenuOption[]>(() => {
  return [
    {
      label: i18n.t('structure'),
      key: 'structure',
      icon: renderIcon(ToggleDb),
      show: connectionStore.anyConnectionOpened,
    },
    {
      label: i18n.t('server'),
      key: 'server',
      icon: renderIcon(ToggleServer),
    },
      {
        label: i18n.t('log'),
        key: 'log',
        icon: renderIcon(Log),
      },
  ]
})

// 偏好设置菜单选项, 计算属性，只有监听的变量改变时才会重新计算
const preferencesOptions = computed<DropdownOption[]>(() => {
  return [
    {
      label: i18n.t('preferences'),
      key: 'preferences',
      icon: renderIcon(Config),
    },
    {
      label: i18n.t('about'),
      key: 'about',
    },
    {
      label: i18n.t('check_update'),
      key: 'update',
      icon: renderIcon(Update),
    },
  ]
})

// 渲染上下文标签, 函数组件，返回一个 VNode
const renderContextLabel = (option: DropdownOption) => {
  const label = typeof option.label === 'string' ? option.label : ''
  return h('div', { class: 'context-menu-item' }, label)
}

// 获取dialog 需要的存储 数据
const dialogStore = useDialogStore()
// 处理偏好设置菜单选择, 函数，根据选择的菜单项执行相应的操作
const onSelectPreferenceMenu = (key: MenuItemKey) => {
  switch (key) {
    case 'preferences':
      //  将偏好设置菜单打开，这里只是打开渲染的标志
      dialogStore.openPreferencesDialog()
      break
    case 'update':
      break
  }
}

// 打开 GitHub 页面
const openGithub = () => {
  //  wails 自带， 调用默认浏览器打开对应网址
  BrowserOpenURL('https://github.com/andrewbytecoder/wrdm')
}
</script>

<template>
  <div id="app-nav-menu"
       :style="{
          width: props.width + 'px',
       }"
       class="flex-box-v">
    <n-menu
        :collapsed-width="props.width"
        :value="props.value"
        :collapsed="true"
        :collapsed-icon-size="iconSize"
        @update:value="(val:  string) => emit('update:value', val)"
        :options="menuOptions"
    ></n-menu>
<!--  用来撑满中间空白区域-->
    <div class="flex-item-expand" ></div>
    <div class="nav-menu-item flex-box-v">
<!--      下拉框，点开是一个拉开的框框-->
<!--      会根据位置进行展开，如果在上面会向下展开，如果在下面会向上展开-->
      <n-dropdown
          :animated="true"
          :keyboard="false"
          :options="preferencesOptions"
          :render-label="renderContextLabel"
          trigger="click"
          @select="onSelectPreferenceMenu"
      >
        <icon-button :icon="Config" :size="iconSize" class="nav-menu-button" />
      </n-dropdown>
      <icon-button :icon="Github" :size="iconSize" class="nav-menu-button" @click="openGithub"></icon-button>
    </div>
  </div>
</template>

<style lang="scss">
#app-nav-menu {
  //width: 60px;
  height: 100vh;
  border-right: var(--border-color) solid 1px;

  .nav-menu-item {
    align-items: center;
    padding: 10px 0;
    gap: 15px;

    .nav-menu-button {
      margin-bottom: 6px;

      :hover {
        color: v-bind('themeVars.primaryColor');
      }
    }
  }
}
</style>
