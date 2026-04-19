<script setup lang="ts">
import { computed } from "vue";

const emit = defineEmits<{}>();

const {
  col = { name: "", rows: [] },
  start = 0,
  end = 0,
  rowHeight = 32,
  width = 300,
} = defineProps<{
  col: { name: string; rows: any[] };
  start: number;
  end: number;
  rowHeight: number;
  width: number;
}>();

const visibleRows = computed(() => col.rows.slice(start, end));

const offsetY = computed(() => start * rowHeight);
const totalHeight = computed(() => col.rows.length * rowHeight);
</script>

<template>
  <div :class="['relative flex-shrink-0 flex flex-col h-full']" :style="{ width: width + 'px' }">
    <div class="relative h-full">
      <div :style="{ height: totalHeight + 'px' }"></div>

      <div class="absolute top-0 left-0 w-full" :style="{ transform: `translateY(${offsetY}px)` }">
        <div
          v-for="(row, i) in visibleRows"
          :key="i"
          class="truncate p-[5px] border-r border-b border-textfield-border-light dark:border-textfield-border-dark"
          :style="{ height: rowHeight + 'px' }"
        >
          {{ row }}
        </div>
      </div>
    </div>
  </div>
</template>
