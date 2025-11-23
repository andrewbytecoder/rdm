<script setup lang="ts">
import {computed} from "vue";
import {RedisDataTypeEnum, types, validType} from "../../consts/support_redis_type";

interface Props {
  type?: string
  color?: string
  size?: string
}

const props = withDefaults(defineProps<Props>(), {
  type: 'STRING',
  color: '',
  size: undefined,
})

const color: Record<string, string> = {
  [types.STRING]: '#626aef',
  [types.HASH]: '#576bfa',
  [types.LIST]: '#34b285',
  [types.SET]: '#bb7d52',
  [types.ZSET]: '#d053a5',
}

const backgroundColor = computed(() => {
  return color[props.type]
})

</script>

<template>
  <n-tag
    :bordered="false"
    :color="{color:backgroundColor, textColor: 'white'}"
    :size="props.size"
    class="redis-type-tag"
    strong
  >
    {{ props.type }}
  </n-tag>
  <!--  <div class="redis-type-tag flex-box-h" :style="{backgroundColor: backgroundColor}">{{ props.type }}</div>-->
</template>

<style scoped lang="scss">
.redis-type-tag {
  padding: 0 12px;
}


</style>