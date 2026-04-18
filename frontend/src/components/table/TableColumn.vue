<script setup lang="ts">
import { computed, ref } from "vue";

const emit = defineEmits<{}>();

const {
  col = { name: "", rows: [] },
  start = 0,
  end = 0,
  rowHeight = 32,
} = defineProps<{
  col: { name: string; rows: any[] };
  start: number;
  end: number;
  rowHeight: number;
}>();

const visibleRows = computed(() => col.rows.slice(start, end));

const offsetY = computed(() => start * rowHeight);
const totalHeight = computed(() => col.rows.length * rowHeight);

const width = ref<number>(300);
const isResizing = ref<boolean>(false);

let frameId: number | null = null;
let pendingWidth = 0;

let startX = 0;
let startWidth = 0;

function startResize(e: MouseEvent) {
  e.preventDefault();
  isResizing.value = true;

  startX = e.clientX;
  startWidth = width.value;

  document.body.classList.add("resizing");

  window.addEventListener("mousemove", resize);
  window.addEventListener("mouseup", stopResize);
}

function resize(e: MouseEvent) {
  if (!isResizing.value) return;

  const dx = e.clientX - startX;
  pendingWidth = startWidth + dx;

  if (frameId) return;

  frameId = requestAnimationFrame(() => {
    let newWidth = pendingWidth;

    if (newWidth < 80) newWidth = 80;
    if (newWidth > 600) newWidth = 600;

    width.value = newWidth;

    frameId = null;
  });
}

function stopResize() {
  isResizing.value = false;

  document.body.classList.remove("resizing");

  if (frameId) {
    cancelAnimationFrame(frameId);
    frameId = null;
  }

  window.removeEventListener("mousemove", resize);
  window.removeEventListener("mouseup", stopResize);

  document.body.style.userSelect = "";
  document.body.style.cursor = "";
}
</script>

<template>
  <div
    :class="[
      'relative flex-shrink-0 border-r border-textfield-border-light dark:border-textfield-border-dark flex flex-col h-full',
    ]"
    :style="{ width: width + 'px' }"
  >
    <div
      class="absolute top-0 right-0 w-[4px] h-full cursor-col-resize hover:bg-primary/30"
      @mousedown="startResize"
      @dblclick="width = 300"
    ></div>
    <div
      class="text-center p-[5px] font-bold border-b border-textfield-border-light dark:border-textfield-border-dark"
    >
      {{ col.name }}
    </div>
    <div class="relative h-full">
      <div :style="{ height: totalHeight + 'px' }"></div>

      <div class="absolute top-0 left-0 w-full" :style="{ transform: `translateY(${offsetY}px)` }">
        <div
          v-for="(row, i) in visibleRows"
          :key="i"
          class="truncate p-[5px] border-b border-textfield-border-light dark:border-textfield-border-dark"
          :style="{ height: rowHeight + 'px' }"
        >
          {{ row }}
        </div>
      </div>
    </div>
  </div>
</template>
