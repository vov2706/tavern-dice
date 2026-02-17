<script setup lang="ts">
import { computed, ref } from "vue"
import { onClickOutside } from '@vueuse/core'

export type Option = {
  label: string
  value: string | number
  icon?: string
}

const props = defineProps<{
  modelValue: string | number | null
  options: Option[]
  placeholder?: string
}>()

const emit = defineEmits<{
  (e: "update:modelValue", value: string | number): void
}>()

const isOpen = ref(false)
const root = ref<HTMLElement | null>(null)

const selected = computed(() =>
  props.options.find(o => o.value === props.modelValue)
)

const selectOption = (option: Option) => {
  emit("update:modelValue", option.value)
  isOpen.value = false
}

// клік поза компонентом
onClickOutside(root, () => {
  isOpen.value = false
})
</script>

<template>
  <div ref="root" class="relative w-full">
    <!-- Trigger -->
    <button
      type="button"
      @click="isOpen = !isOpen"
      class="w-full rounded-xl border border-wood-700/35 bg-parchment-50/70 px-4 py-2 text-base text-left flex items-center justify-between transition hover:border-candle-400/60"
    >
      <div class="flex items-center gap-2">
        <img
          v-if="selected?.icon"
          :src="selected.icon"
          class="w-5 h-5"
          :alt="selected.label"
        />
        <span class="truncate">
          {{ selected?.label || placeholder || "Select option" }}
        </span>
      </div>

      <span
        class="transition-transform duration-200"
        :class="{ 'rotate-180': isOpen }"
      >
        ▾
      </span>
    </button>

    <!-- Dropdown -->
    <transition
      enter-active-class="transition duration-150 ease-out"
      enter-from-class="opacity-0 -translate-y-1"
      enter-to-class="opacity-100 translate-y-0"
      leave-active-class="transition duration-100 ease-in"
      leave-from-class="opacity-100 translate-y-0"
      leave-to-class="opacity-0 -translate-y-1"
    >
      <div
        v-if="isOpen"
        class="absolute z-50 mt-2 w-full rounded-xl border border-wood-700/35 bg-parchment-50 shadow-xl backdrop-blur"
      >
        <div
          v-for="option in options"
          :key="option.value"
          @click="selectOption(option)"
          class="px-4 py-2 hover:bg-wood-700/10 text-base cursor-pointer flex items-center gap-2 transition"
        >
          <img
            v-if="option.icon"
            :src="option.icon"
            class="w-5 h-5"
          />
          {{ option.label }}
        </div>
      </div>
    </transition>
  </div>
</template>
