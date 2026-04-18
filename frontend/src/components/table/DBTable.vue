<script setup lang="ts">
import BaseIconButton from "@/components/buttons/BaseIconButton.vue";
import BasePrimaryButton from "@/components/buttons/BasePrimaryButton.vue";
import IconChevronLeft from "@/components/icons/IconChevronLeft.vue";
import IconChevronRight from "@/components/icons/IconChevronRight.vue";
import IconTune from "@/components/icons/IconTune.vue";
import BaseTextField from "@/components/inputs/BaseTextField.vue";
import TableColumn from "@/components/table/TableColumn.vue";
import { useToast } from "@/composables/useToast";
import { MAX_ROWS } from "@/constants/app";
import { getTableData } from "@/services/connection-manager.service";
import { useConnectionStore } from "@/stores/connection.store";
import { Column } from "@/types/column.types";
import { Result } from "@/types/query.types";
import { Table } from "@/types/table.types";
import { useVuelidate } from "@vuelidate/core";
import { helpers, maxValue, minValue } from "@vuelidate/validators";
import { storeToRefs } from "pinia";
import { computed, ref, watch } from "vue";

const { selectedConnection } = storeToRefs(useConnectionStore());

const { table } = defineProps<{ table: Table | undefined }>();

const rowHeight = 32;

const scrollEl = ref<HTMLElement | null>(null);
const scrollTop = ref(0);
const containerHeight = ref(0);

const start = ref(0);
const end = ref(0);
const visibleCount = ref(0);

const columns = ref<Column[]>([]);
const tableData = ref<Result<Column[]>>();

const tuning = ref<boolean>(false);
const limit = ref<number>(100);
const offset = ref<number>(0);

const appliedLimit = ref<number>(100);
const appliedOffset = ref<number>(0);

const totalRows = computed(function () {
  if (!tableData.value) return 0;
  return tableData.value.total;
});

const v$ = useVuelidate(
  {
    limit: {
      min: minValue(0),
      max: helpers.withMessage(`Limit cannot be more than ${MAX_ROWS}`, maxValue(MAX_ROWS)),
    },
    offset: {
      min: minValue(0),
      max: helpers.withMessage("Offset cannot be more than total", maxValue(totalRows)),
    },
  },
  { limit, offset }
);

const limitError = computed((): string => v$.value.limit.$errors[0]?.$message.toString() ?? null);
const offsetError = computed((): string => v$.value.offset.$errors[0]?.$message.toString() ?? null);

function onScroll(e: Event) {
  const el = e.target as HTMLElement;
  scrollTop.value = el.scrollTop;

  start.value = Math.floor(scrollTop.value / rowHeight);
  end.value = start.value + visibleCount.value + 5;
}

async function goTo(curLimit: number, curOffset: number): Promise<void> {
  try {
    if (!selectedConnection.value || !table) return;

    const isValid = await v$.value.$validate();
    if (!isValid) return;

    columns.value = [];

    tableData.value = await getTableData(
      selectedConnection.value.id,
      selectedConnection.value.databaseDriver.name,
      table.name,
      curLimit,
      curOffset,
      "",
      "ASC"
    );

    columns.value = tableData.value.result;
    appliedLimit.value = curLimit;
    appliedOffset.value = curOffset;
    limit.value = curLimit;
    offset.value = curOffset;

    initViewport();
  } catch (error) {
    useToast(error as string, "error");
  }
}

watch(
  () => table,
  function () {
    goTo(100, 0);
  },
  { immediate: true }
);

function initViewport(): void {
  if (!scrollEl.value) return;

  containerHeight.value = scrollEl.value.clientHeight;
  visibleCount.value = Math.ceil(containerHeight.value / rowHeight);

  end.value = visibleCount.value;
}
</script>

<template>
  <div class="h-full flex flex-col">
    <div ref="scrollEl" class="flex-1 flex overflow-x-auto" @scroll="onScroll">
      <TableColumn
        v-for="col in columns"
        :key="col.name"
        :col="col"
        :start="start"
        :end="end"
        :row-height="rowHeight"
      />
    </div>
    <div
      v-if="table"
      class="px-[20px] flex justify-between items-center h-[60px] shrink-0 border-t border-textfield-border-light dark:border-textfield-border-dark"
    >
      Showing {{ Math.min(appliedOffset + 1, tableData?.total as number) }} -
      {{ Math.min(appliedOffset + appliedLimit, tableData?.total as number) }} of
      {{ tableData?.total }}
      <div class="flex items-center">
        <BaseIconButton
          class="h-[35px]"
          :disabled="appliedOffset - appliedLimit <= -appliedLimit"
          @click="goTo(appliedLimit, Math.max(appliedOffset - appliedLimit, 0))"
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
              v-model="limit"
              type="number"
              :max="MAX_ROWS"
              :error="limitError"
              :should-show-erorr="false"
            ></BaseTextField>
            <BaseTextField
              class="w-[200px] flex items-center gap-[10px]"
              label="Offset:"
              name="offset"
              v-model="offset"
              type="number"
              :max="totalRows"
              :error="offsetError"
              :should-show-erorr="false"
            ></BaseTextField>
            <BasePrimaryButton
              :disabled="limit === appliedLimit && offset === appliedOffset"
              @click="goTo(limit, offset)"
            >
              Apply
            </BasePrimaryButton>
          </div>
        </div>
        <BaseIconButton
          class="h-[35px]"
          :disabled="appliedOffset + appliedLimit >= (tableData?.total || 0)"
          @click="goTo(appliedLimit, Math.min(appliedOffset + appliedLimit, tableData?.total || 0))"
        >
          <IconChevronRight class="w-[25px] h-[25px]" />
        </BaseIconButton>
      </div>
    </div>
  </div>
</template>
