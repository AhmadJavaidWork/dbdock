<script setup lang="ts">
import TableColumn from "@/components/table/TableColumn.vue";
import { useToast } from "@/composables/useToast";
import { getTableData } from "@/services/connection-manager.service";
import { useConnectionStore } from "@/stores/connection.store";
import { Column } from "@/types/column.types";
import { Table } from "@/types/table.types";
import { storeToRefs } from "pinia";
import { ref, watch } from "vue";

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

function onScroll(e: Event) {
  const el = e.target as HTMLElement;
  scrollTop.value = el.scrollTop;

  start.value = Math.floor(scrollTop.value / rowHeight);
  end.value = start.value + visibleCount.value + 5;
}

watch(
  () => table,
  async () => {
    try {
      if (!selectedConnection.value || !table) return;

      columns.value = [];

      columns.value = await getTableData(
        selectedConnection.value.id,
        selectedConnection.value.databaseDriver.name,
        table.name,
        100,
        0,
        "",
        "ASC"
      );

      initViewport();
    } catch (error) {
      useToast(error as string, "error");
    }
  },
  { immediate: true }
);

function initViewport() {
  if (!scrollEl.value) return;

  containerHeight.value = scrollEl.value.clientHeight;
  visibleCount.value = Math.ceil(containerHeight.value / rowHeight);

  end.value = visibleCount.value;
}
</script>

<template>
  <div class="h-full w-full overflow-x-auto">
    <div ref="scrollEl" class="flex min-w-max h-full overflow-y-auto" @scroll="onScroll">
      <TableColumn
        v-for="col in columns"
        :key="col.name"
        :col="col"
        :start="start"
        :end="end"
        :row-height="rowHeight"
      />
    </div>
  </div>
</template>
