<script setup lang="ts">
import BasePrimaryButton from "@/components/buttons/BasePrimaryButton.vue";
import BaseTextField from "@/components/inputs/BaseTextField.vue";
import { useConnectionStore } from "@/stores/connection.store";
import { useThemeStore } from "@/stores/theme.store";
import { Table } from "@/types/table.types";
import { storeToRefs } from "pinia";
import { computed, onMounted, ref } from "vue";

const emit = defineEmits<{}>();

const { selectedConnectionTables, selectedConnectionSchemas } = storeToRefs(useConnectionStore());

const sidebarWidth = ref<number>(300);
const isResizing = ref<boolean>(false);

let frameId: number | null = null;
let pendingWidth = 0;

const search = ref<string>("");
const selectedTable = ref<Table>();

const filteredTables = computed(function (): Table[] {
  if (!search.value) return selectedConnectionTables.value;
  return selectedConnectionTables.value.filter((t) => t.name.toLowerCase().includes(search.value));
});

function startResize(e: MouseEvent) {
  e.preventDefault();
  isResizing.value = true;

  document.body.classList.add("resizing");

  window.addEventListener("mousemove", resize);
  window.addEventListener("mouseup", stopResize);
}

function resize(e: MouseEvent) {
  if (!isResizing.value) return;
  pendingWidth = e.clientX;

  if (frameId) return;

  frameId = requestAnimationFrame(() => {
    let newWidth = pendingWidth;

    if (newWidth < 200) newWidth = 200;
    if (newWidth > 600) newWidth = 600;

    sidebarWidth.value = newWidth;

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

onMounted(async function () {
  useConnectionStore().getTables();
});
</script>

<template>
  <aside
    :class="[
      'relative border-r border-textfield-border-light dark:border-textfield-border-dark flex flex-col',
    ]"
    :style="{ width: sidebarWidth + 'px' }"
  >
    <div
      class="absolute top-0 right-0 w-[4px] h-full cursor-col-resize hover:bg-primary/30"
      @mousedown="startResize"
      @dblclick="sidebarWidth = 300"
    ></div>
    <div
      class="flex gap-[15px] p-[10px] border-b border-textfield-border-light dark:border-textfield-border-dark"
    >
      <BaseTextField
        class="w-full"
        v-model="search"
        name="searchTables"
        placeholder="Search tables..."
      />
    </div>

    <nav class="flex-1 overflow-y-auto">
      <div class="font-medium text-xl px-[10px]"><h1>Tables</h1></div>
      <div v-for="(t, index) in filteredTables" :key="`table-list-name-${index}`">
        <div
          :class="[
            'px-5 flex-wrap cursor-pointer',
            {
              'text-select-options-text-light hover:bg-select-options-background-hovered-light dark:text-select-options-text-dark dark:hover:bg-select-options-background-hovered-dark':
                selectedTable?.name !== t.name,
              'text-select-options-text-selected-light bg-select-options-background-selected-light hover:bg-select-options-background-selected-hover-light dark:text-select-options-text-selected-dark dark:bg-select-options-background-selected-dark dark:hover:bg-select-options-background-selected-hover-dark':
                selectedTable?.name === t.name,
            },
          ]"
          @click="selectedTable = t"
        >
          <h3 class="text-lg truncate" :title="t.name">
            {{ t.name }}
          </h3>
        </div>
      </div>
    </nav>

    <div
      class="p-4 border-t border-textfield-border-light dark:border-textfield-border-dark flex items-center justify-center"
    >
      <BasePrimaryButton @click="useThemeStore().toggleTheme"> Toggle Theme </BasePrimaryButton>
    </div>
  </aside>
</template>
