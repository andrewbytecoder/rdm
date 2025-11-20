<script setup lang="ts">
import { computed } from 'vue'
import { NIcon } from 'naive-ui'
import type { Component } from 'vue'

interface Props {
  tooltip?: string
  tTooltip?: string
  icon?: string | Component
  size?: number | string
  color?: string
  strokeWidth?: number | string
}

const emit = defineEmits<{
  (e: 'click'): void
  // click: []  vue3
}>()

const props = withDefaults(defineProps<Props>(), {
  size: 20,
  color: 'currentColor',
  strokeWidth: 3,
})

const hasTooltip = computed(() => {
  return props.tooltip || props.tTooltip
})


</script>

<template>
  <n-tooltip v-if="hasTooltip">
    <template #trigger>
      <n-icon :color="props.color" :size="props.size" class="icon-btn" @click="emit('click')">
        <component :is="props.icon" :stroke-width="props.strokeWidth"></component>
      </n-icon>
    </template>
  </n-tooltip>
  <n-icon v-else :color="props.color" :size="props.size" class="icon-btn" @click="emit('click')">
    <component :is="props.icon" :stroke-width="props.strokeWidth"></component>
  </n-icon>
</template>

<style scoped lang="scss">

</style>