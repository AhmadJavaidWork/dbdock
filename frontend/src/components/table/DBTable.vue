<script setup lang="ts">
import BasePrimaryButton from "@/components/buttons/BasePrimaryButton.vue";
import BaseTextField from "@/components/inputs/BaseTextField.vue";
import TableColumn from "@/components/table/TableColumn.vue";
import { useToast } from "@/composables/useToast";
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

const limit = ref<number>(100);
const offset = ref<number>(0);
const pageNumber = ref<number>(1);
const curPageNumber = ref<number>(1);

const totalPages = computed(function () {
  if (!tableData.value) return 1;
  return Math.ceil(tableData.value.total / limit.value);
});

const v$ = useVuelidate(
  {
    pageNumber: {
      min: minValue(1),
      max: helpers.withMessage("Page doesn't exist", maxValue(totalPages)),
    },
  },
  { pageNumber }
);

const pageNumError = computed(
  (): string => v$.value.pageNumber.$errors[0]?.$message.toString() ?? null
);

function onScroll(e: Event) {
  const el = e.target as HTMLElement;
  scrollTop.value = el.scrollTop;

  start.value = Math.floor(scrollTop.value / rowHeight);
  end.value = start.value + visibleCount.value + 5;
}

async function goTo(page: number): Promise<void> {
  try {
    if (!selectedConnection.value || !table) return;

    const isValid = await v$.value.$validate();
    if (!isValid) return;

    columns.value = [];
    offset.value = (page - 1) * limit.value;

    tableData.value = await getTableData(
      selectedConnection.value.id,
      selectedConnection.value.databaseDriver.name,
      table.name,
      limit.value,
      offset.value,
      "",
      "ASC"
    );

    columns.value = tableData.value.result;
    curPageNumber.value = page;
    pageNumber.value = page;

    initViewport();
  } catch (error) {
    useToast(error as string, "error");
  }
}

watch(
  () => table,
  function () {
    goTo(1);
    pageNumber.value = 1;
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
      Showing {{ offset + 1 }} - {{ offset + limit }} of {{ tableData?.total }}
      <div class="flex items-center gap-[10px]">
        <BasePrimaryButton :disabled="curPageNumber === 1" @click="goTo(curPageNumber - 1)">
          Previous
        </BasePrimaryButton>
        <BaseTextField
          class="w-[100px]"
          name="pageNumber"
          v-model="pageNumber"
          type="number"
          :max="totalPages"
          :error="pageNumError"
          :should-show-erorr="false"
          @submit="goTo(pageNumber)"
          @blur="pageNumber = curPageNumber"
        ></BaseTextField>
        <BasePrimaryButton
          :disabled="curPageNumber === totalPages"
          @click="goTo(curPageNumber + 1)"
        >
          Next
        </BasePrimaryButton>
      </div>
    </div>
  </div>
</template>
