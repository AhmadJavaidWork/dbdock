<script setup lang="ts">
import TableColumn from "@/components/table/TableColumn.vue";
import { DEFAULT_COLUMN_WIDTH } from "@/constants/app";
import { useTabStore } from "@/stores/tab.store";
import { storeToRefs } from "pinia";
import { computed, onMounted, ref } from "vue";

const { activeTab: tab, activeTabIndex } = storeToRefs(useTabStore());

const rowHeight = 32;

const scrollEl = ref<HTMLElement | null>(null);
const scrollTop = ref(0);
const containerHeight = ref(0);

const start = ref(0);
const end = ref(0);
const visibleCount = ref(0);

const columns = computed(function () {
  if (!tab.value) return [];
  return tab.value.columns;
});

function onScroll(e: Event) {
  const el = e.target as HTMLElement;
  scrollTop.value = el.scrollTop;

  start.value = Math.floor(scrollTop.value / rowHeight);
  end.value = start.value + visibleCount.value + 5;
}

function handleKeydown(e: KeyboardEvent): void {
  const key = e.key.toLowerCase();
  const indexes = ["0", "1", "2", "3", "4", "5", "6", "7", "8", "9"];
  if (e.ctrlKey && key === "w") {
    e.preventDefault();
    e.stopPropagation();
    if (!tab.value) return;
    useTabStore().closeTab(activeTabIndex.value);
  } else if (e.ctrlKey && key === "tab") {
    e.preventDefault();
    e.stopPropagation();

    if (!tab.value) return;
    useTabStore().switchTo(activeTabIndex.value + 1);
  } else if (e.ctrlKey && indexes.includes(key)) {
    e.preventDefault();
    e.stopPropagation();

    if (!tab.value) return;
    let index = parseInt(key) - 1;
    if (key === "9") {
      index = -1;
    } else if (key === "0") {
      index = 0;
    }
    useTabStore().switchTo(index);
  }
}

function initViewport(): void {
  if (!scrollEl.value) return;

  containerHeight.value = scrollEl.value.clientHeight;
  visibleCount.value = Math.ceil(containerHeight.value / rowHeight);

  end.value = visibleCount.value;
}

function startResize(e: MouseEvent, index: number) {
  e.preventDefault();
  if (!tab.value) return;
  const col = tab.value.columns[index];

  col.isResizing = true;

  col.startX = e.clientX;
  col.startWidth = col.width;

  document.body.classList.add("resizing");

  const moveHandler = (ev: MouseEvent) => resize(ev, index);
  const upHandler = () => stopResize(index, moveHandler, upHandler);

  window.addEventListener("mousemove", moveHandler);
  window.addEventListener("mouseup", upHandler);
}

function resize(e: MouseEvent, index: number) {
  if (!tab.value) return;
  const col = tab.value.columns[index];
  if (!col.isResizing) return;

  const dx = e.clientX - col.startX;
  col.pendingWidth = col.startWidth + dx;

  if (col.frameId) return;

  col.frameId = requestAnimationFrame(() => {
    let newWidth = col.pendingWidth;

    if (newWidth < 80) newWidth = 80;
    if (newWidth > 600) newWidth = 600;

    col.width = newWidth;

    col.frameId = null;
  });
}

function stopResize(index: number, moveHandler: any, upHandler: any) {
  if (!tab.value) return;
  const col = tab.value.columns[index];
  col.isResizing = false;

  document.body.classList.remove("resizing");

  if (col.frameId) {
    cancelAnimationFrame(col.frameId);
    col.frameId = null;
  }

  window.removeEventListener("mousemove", moveHandler);
  window.removeEventListener("mouseup", upHandler);

  document.body.style.userSelect = "";
  document.body.style.cursor = "";
}

onMounted(function () {
  initViewport();
  window.addEventListener("keydown", handleKeydown, true);
});
</script>

<template>
  <div class="h-full flex flex-col">
    <div
      class="sticky top-0 z-30 flex bg-background-light dark:bg-background-dark border-b border-textfield-border-light dark:border-textfield-border-dark"
    >
      <div
        v-for="(col, index) in columns"
        :key="col.name"
        class="relative flex-shrink-0 text-center p-[5px] font-bold border-r border-textfield-border-light dark:border-textfield-border-dark"
        :style="{ width: col.width + 'px' }"
      >
        {{ col.name }}
        <div
          class="absolute top-0 right-0 w-[4px] h-full cursor-col-resize hover:bg-primary/30 dark:hover:bg-primary-dark z-40"
          @mousedown="startResize($event, index)"
          @dblclick="col.width = DEFAULT_COLUMN_WIDTH"
        ></div>
      </div>
    </div>
    <div ref="scrollEl" class="flex-1 flex overflow-x-auto relative" @scroll="onScroll">
      <TableColumn
        v-for="col in columns"
        :key="col.name"
        :col="col"
        :start="start"
        :end="end"
        :row-height="rowHeight"
        :width="col.width"
      />
    </div>
  </div>
</template>
