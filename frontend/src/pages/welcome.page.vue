<script setup lang="ts">
import BaseBorderButton from "@/components/buttons/BaseBorderButton.vue";
import BaseBorderedIconButton from "@/components/buttons/BaseBorderedIconButton.vue";
import BasePrimaryButton from "@/components/buttons/BasePrimaryButton.vue";
import ConnectionForm from "@/components/forms/ConnectionForm.vue";
import IconClose from "@/components/icons/IconClose.vue";
import BaseTextField from "@/components/inputs/BaseTextField.vue";
import ConfigModal from "@/components/modals/ConfigModal.vue";
import { usePrompt } from "@/composables/usePrompt";
import { useToast } from "@/composables/useToast";
import { connectionErrorToText } from "@/errors/connection.error";
import { testConnection } from "@/services/connection.service";
import { fetchDatabaseDrivers } from "@/services/database.service";
import { useConnectionStore } from "@/stores/connection.store";
import { useLoaderStore } from "@/stores/loader";
import { CreateDBConnection, DBConnection } from "@/types/connection.type";
import { DatabaseDriver } from "@/types/databaseDriver.types";
import { storeToRefs } from "pinia";
import { computed, onMounted, onUnmounted, ref } from "vue";

const { prompt } = usePrompt();

const { selectedConnection } = storeToRefs(useConnectionStore());

const databaseDrivers = ref<DatabaseDriver[]>([]);
const connection = ref<CreateDBConnection>({
  name: "",
  databaseDriverId: null,
  host: "",
  port: null,
  username: "",
  password: "",
  databaseName: "",
});
const formRef = ref<InstanceType<typeof ConnectionForm> | null>(null);

const testing = ref<boolean>(false);
const saving = ref<boolean>(false);
const connecting = ref<boolean>(false);
const search = ref<string>("");
const showConnectionForm = ref<boolean>(false);
const showConfigModal = ref<boolean>(false);

const filteredConnections = computed(() => {
  if (!search.value) return useConnectionStore().connections;
  return useConnectionStore().connections.filter((c) =>
    c.name.toLowerCase().includes(search.value)
  );
});

function setConnection(conn: CreateDBConnection) {
  connection.value = {
    name: conn.name,
    databaseDriverId: conn.databaseDriverId,
    host: conn.host,
    port: conn.port,
    username: conn.username,
    password: conn.password,
    databaseName: conn.databaseName,
  };

  if (conn.databaseDriverId === null) {
    selectDefaultDatabaseDriver();
  }

  showConnectionForm.value = true;
}

function selectConnection(conn: DBConnection) {
  connection.value = {
    name: conn.name,
    databaseDriverId: conn.databaseDriverId,
    host: conn.host,
    port: conn.port,
    username: conn.username,
    password: conn.password,
    databaseName: conn.databaseName,
  };
  selectedConnection.value = conn;
  showConnectionForm.value = true;
}

function hideConnectionForm() {
  if (formRef.value) {
    formRef.value.v$.$reset();
  }
  selectedConnection.value = null;
  connection.value = {
    name: "",
    databaseDriverId: null,
    host: "",
    port: null,
    username: "",
    password: "",
    databaseName: "",
  };

  selectDefaultDatabaseDriver();
  showConnectionForm.value = false;
}

async function handleTest(): Promise<void> {
  if (!formRef.value) return;
  const valid = await formRef.value.v$.$validate();
  if (!valid) return;
  if (!formRef.value.databaseDriver || !connection.value.port) return;

  useLoaderStore().show();

  testing.value = true;
  try {
    const conn: CreateDBConnection = {
      name: connection.value.name,
      databaseDriverId: connection.value.databaseDriverId,
      host: connection.value.host,
      port: connection.value.port,
      username: connection.value.username,
      password: connection.value.password,
      databaseName: connection.value.databaseName,
    };
    const res = await testConnection(conn, formRef.value.databaseDriver);
    if (!res) return;
    prompt({
      type: "success",
      title: "Connection status",
      message: "Connection Successful",
    });
  } catch (error: unknown) {
    const message = connectionErrorToText(
      error,
      formRef.value.databaseDriver,
      connection.value.username,
      connection.value.databaseName
    );
    prompt({
      type: "error",
      title: "Connection failed",
      message,
    });
    console.log("Connection failed: ", message);
  } finally {
    testing.value = false;
    useLoaderStore().hide();
  }
}

async function createConnection(): Promise<void> {
  if (!formRef.value) return;
  const valid = await formRef.value.v$.$validate();
  if (!valid) return;
  saving.value = true;

  useLoaderStore().show();
  try {
    const conn: CreateDBConnection = {
      name: connection.value.name,
      databaseDriverId: connection.value.databaseDriverId,
      host: connection.value.host,
      port: connection.value.port,
      username: connection.value.username,
      password: connection.value.password,
      databaseName: connection.value.databaseName,
    };
    const message = await useConnectionStore().createConnection(conn);
    useToast(message, "success");
  } catch (error: unknown) {
    useToast(error as string, "error");
    console.log("error", error);
  } finally {
    saving.value = false;
    useLoaderStore().hide();
  }
}

async function updateConnection(): Promise<void> {
  if (!formRef.value || !selectedConnection.value || !formRef.value.databaseDriver) return;
  const valid = await formRef.value.v$.$validate();
  if (!valid) return;
  saving.value = true;

  useLoaderStore().show();
  try {
    const conn: DBConnection = {
      id: selectedConnection.value.id,
      name: connection.value.name,
      databaseDriverId: connection.value.databaseDriverId,
      host: connection.value.host,
      port: connection.value.port,
      username: connection.value.username,
      password: connection.value.password,
      databaseName: connection.value.databaseName,
      createdAt: selectedConnection.value.createdAt,
      updatedAt: selectedConnection.value.updatedAt,
      databaseDriver: formRef.value.databaseDriver,
    };
    const message = await useConnectionStore().updateConnection(conn);
    useToast(message, "success");
  } catch (error: unknown) {
    useToast(error as string, "error");
    console.log("error", error);
  } finally {
    saving.value = false;
    useLoaderStore().hide();
  }
}

function selectDefaultDatabaseDriver() {
  if (databaseDrivers.value.length > 0) {
    connection.value.databaseDriverId = databaseDrivers.value[0].id;
    connection.value.port = databaseDrivers.value[0].defaultPort;
    if (formRef.value) {
      formRef.value.databaseDriver = databaseDrivers.value[0];
    }
  }
}

async function handleSave(): Promise<void> {
  if (useConnectionStore().selectedConnection) {
    updateConnection();
  } else {
    createConnection();
  }
}

async function handleConnect(): Promise<void> {}

function handleKeydown(e: KeyboardEvent) {
  if (e.key === "Escape" && showConnectionForm.value) {
    hideConnectionForm();
  }
}

window.addEventListener("keydown", handleKeydown);

onMounted(async function (): Promise<void> {
  try {
    useConnectionStore().getConnections();
    databaseDrivers.value = await fetchDatabaseDrivers();
    selectDefaultDatabaseDriver();
  } catch (error) {
    useToast(error as string, "error");
    console.log("error", error);
  }
});

onUnmounted(function () {
  window.removeEventListener("keydown", handleKeydown);
});
</script>

<template>
  <ConfigModal
    :is-open="showConfigModal"
    @close="showConfigModal = false"
    @submit="setConnection"
  />
  <div class="h-screen flex text-text-light dark:text-text-dark">
    <main
      :class="[
        'flex-1 overflow-hidden flex flex-col items-center',
        {
          'py-[100px]': useConnectionStore().connections.length > 0,
          'justify-center': useConnectionStore().connections.length === 0,
        },
      ]"
    >
      <div
        :class="[
          'flex flex-col items-center gap-[50px] justify-center bg-background-light dark:bg-background-dark',
          {
            'max-h-[calc(100vh-200px)]': useConnectionStore().connections.length > 0,
          },
        ]"
      >
        <div class="flex flex-col items-center">
          <div class="text-2xl font-bold mb-[10px]">
            Welcome to
            <span class="text text-secondary">DB</span>
            <span class="text-primary">Dock</span>
          </div>
          <img id="logo" alt="Wails logo" src="@/assets/images/logo.svg" />
        </div>
        <div class="flex flex-col gap-[20px]">
          <p class="font-normal text-xl">Connect to your database to get started</p>
          <div class="flex items-center justify-center gap-[20px]">
            <BasePrimaryButton @click="showConnectionForm = true">New Connection</BasePrimaryButton>
            <BaseBorderButton type="button" @click="showConfigModal = true"
              >Import Config</BaseBorderButton
            >
          </div>
        </div>
        <div
          v-if="useConnectionStore().connections.length > 0"
          class="self-start min-w-[650px] flex flex-col gap-[20px]"
        >
          <div class="flex justify-between items-center">
            <p class="font-normal text-xl">Recent Connections</p>
            <BaseTextField
              v-model="search"
              name="searchConnection"
              placeholder="Search connections..."
            />
          </div>
          <div
            v-if="filteredConnections.length > 0"
            class="overflow-y-scroll overscroll-none max-h-[calc(100vh-200px-158px-100px-83.5px-35px-20px)]"
          >
            <div
              v-for="(connection, index) in filteredConnections"
              :key="`connection-list-name-${index}`"
              :class="[
                'py-2 flex-wrap cursor-pointer',
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
          <div v-else class="text-center max-h-[calc(100vh-200px-158px-100px-83.5px-35px-20px)]">
            No connections found
          </div>
        </div>
      </div>
      <Transition
        enter-active-class="transition-transform duration-300 ease-out"
        enter-from-class="translate-y-full"
        enter-to-class="translate-y-0"
        leave-active-class="transition-transform duration-200 ease-in"
        leave-from-class="translate-y-0"
        leave-to-class="translate-y-full"
      >
        <div
          v-show="showConnectionForm"
          class="fixed inset-0 z-1 max-h-screen bg-background-light dark:bg-background-dark"
        >
          <div
            class="w-full h-full flex flex-col items-center gap-[50px] justify-center bg-background-light dark:bg-background-dark"
          >
            <div class="relative w-[600px] flex flex-col items-center">
              <BaseBorderedIconButton class="absolute top-6 right-0" @click="hideConnectionForm">
                <IconClose class="w-[30px] h-[30px] cursor-pointer text-primary" />
              </BaseBorderedIconButton>
              <div class="text-2xl font-bold mb-[10px]">
                Welcome to
                <span class="text text-secondary">DB</span>
                <span class="text-primary">Dock</span>
              </div>
              <img id="logo" alt="Wails logo" src="@/assets/images/logo.svg" />
            </div>
            <ConnectionForm
              ref="formRef"
              v-model:connection="connection"
              :database-drivers="databaseDrivers"
              :saving="saving"
              :testing="testing"
              :connecting="connecting"
              @save="handleSave"
              @test="handleTest"
              @connect="handleConnect"
              @cancel="hideConnectionForm"
            />
          </div>
        </div>
      </Transition>
    </main>
  </div>
</template>
