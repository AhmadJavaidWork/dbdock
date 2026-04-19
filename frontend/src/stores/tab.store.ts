import { useToast } from "@/composables/useToast";
import { DEFAULT_COLUMN_WIDTH, DEFAULT_RECORD_LIMIT, MAX_ROWS } from "@/constants/app";
import { getTableData } from "@/services/connection-manager.service";
import { useConnectionStore } from "@/stores/connection.store";
import { Column, DisplayColumn } from "@/types/column.types";
import { Tab } from "@/types/tab.types";
import { Table } from "@/types/table.types";
import useVuelidate from "@vuelidate/core";
import { helpers, maxValue, minValue } from "@vuelidate/validators";
import { defineStore, storeToRefs } from "pinia";
import { computed, ref } from "vue";

export const useTabStore = defineStore("tab", () => {
  const tabs = ref<Tab[]>([]);
  const activeTabIndex = ref<number>(0);
  const { selectedConnection } = storeToRefs(useConnectionStore());

  const activeTab = computed(() => {
    if (!tabs.value[activeTabIndex.value]) return undefined;
    return tabs.value[activeTabIndex.value];
  });

  const totalRows = computed(() => {
    if (!activeTab.value) return 0;
    return activeTab.value.total;
  });

  const activeTabLimit = computed(() => {
    if (!activeTab.value) return null;
    return activeTab.value.limit;
  });

  const activeTabOffset = computed(() => {
    if (!activeTab.value) return null;
    return activeTab.value.offset;
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
    { limit: activeTabLimit, offset: activeTabOffset }
  );

  const limitError = computed((): string => v$.value.limit.$errors[0]?.$message.toString() ?? null);
  const offsetError = computed(
    (): string => v$.value.offset.$errors[0]?.$message.toString() ?? null
  );

  function openTab(table: Table, attach: boolean) {
    for (let i = 0; i < tabs.value.length; i++) {
      if (tabs.value[i].table.name === table.name) {
        activeTabIndex.value = i;
        tabs.value[i].attached = attach;
        return;
      }
    }

    const newTab: Tab = {
      table,
      columns: [],
      total: 0,
      limit: DEFAULT_RECORD_LIMIT,
      offset: 0,
      appliedLimit: DEFAULT_RECORD_LIMIT,
      appliedOffset: 0,
      attached: attach,
    };

    if (tabs.value.length > 0 && !tabs.value[tabs.value.length - 1].attached) {
      tabs.value[tabs.value.length - 1] = newTab;
    } else {
      tabs.value.push(newTab);
    }

    activeTabIndex.value = tabs.value.length - 1;
    loadRecords(DEFAULT_RECORD_LIMIT, 0);
  }

  function closeTab(index: number) {
    if (!tabs.value[index]) return;

    tabs.value.splice(index, 1);

    if (activeTabIndex.value === index) {
      activeTabIndex.value = Math.min(Math.max(index - 1, 0), index + 1);
    } else if (activeTabIndex.value > index) {
      activeTabIndex.value--;
    }
    if (tabs.value.length === 0) activeTabIndex.value = 0;
  }

  async function loadRecords(curLimit: number, curOffset: number) {
    try {
      if (!activeTab.value || !selectedConnection.value) {
        return;
      }
      const isValid = await v$.value.$validate();
      if (!isValid) return;

      activeTab.value.columns = [];
      const res = await getTableData(
        selectedConnection.value.id,
        selectedConnection.value.databaseDriver.name,
        activeTab.value.table.name,
        curLimit,
        curOffset,
        "",
        "ASC"
      );
      activeTab.value.total = res.total;

      activeTab.value.columns = res.result.map<DisplayColumn>((c: Column) => {
        return {
          name: c.name,
          rows: c.rows,
          width: DEFAULT_COLUMN_WIDTH,
          isResizing: false,
          startX: 0,
          startWidth: 0,
          frameId: null,
          pendingWidth: 0,
        };
      });

      activeTab.value.appliedLimit = curLimit;
      activeTab.value.appliedOffset = curOffset;
      activeTab.value.limit = curLimit;
      activeTab.value.offset = curOffset;
    } catch (error) {
      useToast(error as string, "error");
    }
  }

  function switchTo(index: number, attach?: boolean) {
    if (index === tabs.value.length) {
      index = 0;
    } else if (index > tabs.value.length || index <= -1) {
      index = tabs.value.length - 1;
    }

    activeTabIndex.value = index;
    if (attach) tabs.value[activeTabIndex.value].attached = attach;
  }

  return {
    tabs,
    activeTabIndex,
    activeTab,
    limitError,
    offsetError,
    totalRows,
    openTab,
    closeTab,
    loadRecords,
    switchTo,
  };
});
