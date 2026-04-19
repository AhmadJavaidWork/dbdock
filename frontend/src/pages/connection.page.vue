<script setup lang="ts">
import BaseIconButton from "@/components/buttons/BaseIconButton.vue";
import BasePrimaryButton from "@/components/buttons/BasePrimaryButton.vue";
import IconChevronLeft from "@/components/icons/IconChevronLeft.vue";
import IconChevronRight from "@/components/icons/IconChevronRight.vue";
import IconTune from "@/components/icons/IconTune.vue";
import BaseTextField from "@/components/inputs/BaseTextField.vue";
import TableList from "@/components/sidebars/TableList.vue";
import DBTable from "@/components/table/DBTable.vue";
import TabsBar from "@/components/table/TabsBar.vue";
import { MAX_ROWS } from "@/constants/app";
import { useTabStore } from "@/stores/tab.store";
import { storeToRefs } from "pinia";
import { computed, ref } from "vue";

const { activeTab, limitError, offsetError, totalRows } = storeToRefs(useTabStore());

const tuning = ref<boolean>(false);

const appliedOffset = computed(function () {
  if (!activeTab.value) return 0;
  return activeTab.value.appliedOffset;
});

const appliedLimit = computed(function () {
  if (!activeTab.value) return 0;
  return activeTab.value.appliedLimit;
});
</script>

<template>
  <TableList />
  <div class="flex-1 flex flex-col overflow-hidden">
    <TabsBar />
    <div class="flex-1 overflow-hidden">
      <DBTable v-if="activeTab" />
    </div>
    <div
      v-if="activeTab"
      class="shrink-0 px-[20px] flex justify-between items-center h-[60px] border-t border-textfield-border-light dark:border-textfield-border-dark"
    >
      Showing {{ Math.min(appliedOffset + 1, totalRows as number) }} -
      {{ Math.min(appliedOffset + appliedLimit, totalRows as number) }} of
      {{ totalRows }}
      <div class="flex items-center">
        <BaseIconButton
          class="h-[35px]"
          :disabled="appliedOffset - appliedLimit <= -appliedLimit"
          @click="
            useTabStore().loadRecords(appliedLimit, Math.max(appliedOffset - appliedLimit, 0))
          "
        >
          <IconChevronLeft class="w-[25px] h-[25px]" />
        </BaseIconButton>
        <div class="relative inline-block h-[35px]">
          <BaseIconButton class="h-[35px]" @click="tuning = !tuning">
            <IconTune class="w-[25px] h-[25px]" />
          </BaseIconButton>
          <div
            v-show="tuning"
            class="absolute right-0 bottom-full z-50 bg-background-light dark:bg-background-dark border border-textfield-border-light dark:border-textfield-border-dark rounded shadow-lg p-3 flex flex-col gap-[10px]"
          >
            <BaseTextField
              class="w-[200px] flex items-center gap-[17px]"
              label="Limit:"
              name="limit"
              v-model="activeTab.limit"
              type="number"
              :max="MAX_ROWS"
              :error="limitError"
              :should-show-erorr="false"
            ></BaseTextField>
            <BaseTextField
              class="w-[200px] flex items-center gap-[10px]"
              label="Offset:"
              name="offset"
              v-model="activeTab.offset"
              type="number"
              :max="totalRows"
              :error="offsetError"
              :should-show-erorr="false"
            ></BaseTextField>
            <BasePrimaryButton
              :disabled="activeTab.limit === appliedLimit && activeTab.offset === appliedOffset"
              @click="useTabStore().loadRecords(activeTab.limit, activeTab.offset)"
            >
              Apply
            </BasePrimaryButton>
          </div>
        </div>
        <BaseIconButton
          class="h-[35px]"
          :disabled="appliedOffset + appliedLimit >= activeTab.total"
          @click="
            useTabStore().loadRecords(
              appliedLimit,
              Math.min(appliedOffset + appliedLimit, activeTab.total)
            )
          "
        >
          <IconChevronRight class="w-[25px] h-[25px]" />
        </BaseIconButton>
      </div>
    </div>
  </div>
</template>
