<script setup lang="ts">
import IconClose from "@/components/icons/IconClose.vue";
import { useTabStore } from "@/stores/tab.store";
import { storeToRefs } from "pinia";

const { tabs, activeTabIndex } = storeToRefs(useTabStore());

let clickTimeout: number | null = null;

function handleClick(index: number) {
  document.body.classList.add("no-select");
  if (clickTimeout) clearTimeout(clickTimeout);

  clickTimeout = window.setTimeout(() => {
    useTabStore().switchTo(index);
  }, 200);
}

function handleDblClick(index: number) {
  if (clickTimeout) {
    clearTimeout(clickTimeout);
    clickTimeout = null;
  }

  useTabStore().switchTo(index, true);
}
</script>

<template>
  <div class="flex border-b">
    <div
      v-for="(tab, index) in tabs"
      :key="`tab-${tab.table.name}`"
      :class="[
        'px-4 py-2 cursor-pointer flex items-center gap-2 select-none',
        {
          'text-select-options-text-light hover:bg-select-options-background-hovered-light dark:text-select-options-text-dark dark:hover:bg-select-options-background-hovered-dark':
            index !== activeTabIndex,
          'text-select-options-text-selected-light bg-select-options-background-selected-light hover:bg-select-options-background-selected-hover-light dark:text-select-options-text-selected-dark dark:bg-select-options-background-selected-dark dark:hover:bg-select-options-background-selected-hover-dark':
            index === activeTabIndex,
          italic: !tab.attached,
        },
      ]"
      draggable="true"
      :data-tab-index="index"
      @click="handleClick(index)"
      @dblclick="handleDblClick(index)"
    >
      {{ tab.table.name }}
      <button @click.stop="useTabStore().closeTab(index)">
        <IconClose class="w-[16px] h-[16px]"></IconClose>
      </button>
    </div>
  </div>
</template>
