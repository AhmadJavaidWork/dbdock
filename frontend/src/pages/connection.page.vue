<script setup lang="ts">
import TableList from "@/components/sidebars/TableList.vue";
import DBTable from "@/components/table/DBTable.vue";
import { Table } from "@/types/table.types";
import { ref } from "vue";

const selectedTable = ref<Table>();

function selectTable(table: Table) {
  selectedTable.value = table;
}
</script>

<template>
  <TableList @select-table="selectTable" />
  <div class="flex-1 overflow-hidden">
    <DBTable :table="selectedTable" />
  </div>
</template>
