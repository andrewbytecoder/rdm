<script setup lang="ts">
import { ref } from 'vue'
import { isEmpty, reject } from 'lodash'
import Add from '../icons/Add.vue'
import Delete from '../icons/Delete.vue'
import IconButton from '../common/IconButton.vue'
import type { PropType } from 'vue'

interface ZSetItem {
  value: string
  score: number
}

const props = defineProps({
  type: {
    type: Number,
    default: 0
  },
  value: {
    type: Object as PropType<Record<string, number>>,
    default: () => ({})
  }
})

const emit = defineEmits<{
  (e: 'update:value', value: Record<string, number>): void
  (e: 'update:type', type: number): void
}>()

const zset = ref<ZSetItem[]>([{ value: '', score: 0 }])

const onUpdate = (): void => {
  const val = reject(zset.value, (v) => v == null || isEmpty(v.value)) as ZSetItem[]
  const result: Record<string, number> = {}
  for (const elem of val) {
    result[elem.value] = elem.score
  }
  emit('update:value', result)
}


/**
 * @typedef Hash
 * @property {string} key
 * @property {string} [value]
 */
const kvList = ref([{ key: '', value: '' }])


</script>

<template>
    <n-form-item :label="$t('element')" required>
        <n-dynamic-input
            v-model:value="kvList"
            :key-placeholder="$t('enter_field')"
            :value-placeholder="$t('enter_value')"
            preset="pair"
            @update:value="onUpdate"
        >
            <template #action="{ index, create, remove, move }">
                <icon-button v-if="kvList.length > 1" :icon="Delete" size="18" @click="() => remove(index)" />
                <icon-button :icon="Add" size="18" @click="() => create(index)" />
            </template>
        </n-dynamic-input>
    </n-form-item>
</template>

<style lang="scss" scoped></style>
