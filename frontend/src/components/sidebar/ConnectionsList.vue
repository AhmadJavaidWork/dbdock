<script setup lang="ts">
import BaseIconButton from "@/components/buttons/BaseIconButton.vue";
import BasePrimaryButton from "@/components/buttons/BasePrimaryButton.vue";
import IconPlus from "@/components/icons/IconPlus.vue";
import FormTextField from "@/components/inputs/FormTextField.vue";
import { useConnectionStore } from "@/stores/connection.store";
import { useThemeStore } from "@/stores/theme.store";
import { DBConnection } from "@/types/connection.type";
import { computed, onMounted, ref } from "vue";

const emit = defineEmits<{
  "select-connection": [connection: DBConnection];
  "add-new-connection": [];
}>();

const search = ref<string>("");
const filteredConnections = computed(() => {
  if (!search.value) return useConnectionStore().connections;
  return useConnectionStore().connections.filter((c) =>
    c.name.toLowerCase().includes(search.value)
  );
});

function selectConnection(conn: DBConnection) {
  emit("select-connection", conn);
}

onMounted(async function () {
  await useConnectionStore().getConnections();
});
</script>

<template>
  <aside
    class="w-[300px] border-r border-textfield-border-light dark:border-textfield-border-dark flex flex-col"
  >
    <div
      class="flex gap-[15px] p-4 border-b border-textfield-border-light dark:border-textfield-border-dark"
    >
      <BaseIconButton
        class="h-[35px]"
        title="Add new connection"
        @click="emit('add-new-connection')"
      >
        <IconPlus class="w-[20px] h-[20px]" />
      </BaseIconButton>
      <FormTextField v-model="search" name="searchConnection" placeholder="Search connections..." />
    </div>

    <nav class="flex-1 overflow-y-auto">
      <div
        v-for="(connection, index) in filteredConnections"
        :key="`connection-list-name-${index}`"
      >
        <div
          :class="[
            'px-5 py-2 flex-wrap cursor-pointer',
            {
              'text-select-options-text-light hover:bg-select-options-background-hovered-light border-b border-textfield-border-light dark:text-select-options-text-dark dark:hover:bg-select-options-background-hovered-dark dark:border-textfield-border-dark':
                useConnectionStore().selectedConnection?.id !== connection.id,
              'text-select-options-text-selected-light bg-select-options-background-selected-light hover:bg-select-options-background-selected-hover-light border-b border-textfield-border-light dark:text-select-options-text-selected-dark dark:bg-select-options-background-selected-dark dark:hover:bg-select-options-background-selected-hover-dark dark:border-textfield-border-dark':
                useConnectionStore().selectedConnection?.id === connection.id,
            },
          ]"
          @click="selectConnection(connection)"
        >
          <h3 class="font-medium text-lg truncate" :title="connection.name">
            {{ connection.name }}
          </h3>
          <div class="flex items-center gap-1 text-md min-w-0">
            <span class="truncate min-w-0" :title="connection.host">
              {{ connection.host }}
            </span>
            <span class="shrink-0">:</span>
            <span class="truncate min-w-0" :title="connection.databaseName">
              {{ connection.databaseName }}
            </span>
          </div>
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
