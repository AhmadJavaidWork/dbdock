<script setup lang="ts">
import BaseBorderButton from "@/components/buttons/BaseBorderButton.vue";
import BaseBorderedIconButton from "@/components/buttons/BaseBorderedIconButton.vue";
import BasePrimaryButton from "@/components/buttons/BasePrimaryButton.vue";
import ConnectionForm from "@/components/forms/ConnectionForm.vue";
import IconClose from "@/components/icons/IconClose.vue";
import BaseTextField from "@/components/inputs/BaseTextField.vue";
import ConfigModal from "@/components/modals/ConfigModal.vue";
import { useContextMenu } from "@/composables/useContextMenu";
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
import { computed, nextTick, onMounted, onUnmounted, ref } from "vue";

const { prompt } = usePrompt();
const { contextMenu } = useContextMenu();

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
const configRef = ref<InstanceType<typeof ConfigModal> | null>(null);
const newConnectionBtnRef = ref<InstanceType<typeof BasePrimaryButton> | null>(null);
const importConfigBtnRef = ref<InstanceType<typeof BaseBorderButton> | null>(null);

const testing = ref<boolean>(false);
const saving = ref<boolean>(false);
const connecting = ref<boolean>(false);
const search = ref<string>("");
const showConnectionForm = ref<boolean>(false);
const showConfigModal = ref<boolean>(false);

const filteredConnections = computed((): DBConnection[] => {
  if (!search.value) return useConnectionStore().connections;
  return useConnectionStore().connections.filter((c) =>
    c.name.toLowerCase().includes(search.value)
  );
});

async function setConnection(conn: CreateDBConnection): Promise<void> {
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
  if (formRef.value) {
    await nextTick();
    formRef.value.nameTextFieldRef?.inputRef?.focus();
  }
}

async function openContextMenu(e: MouseEvent, conn: DBConnection): Promise<void> {
  const options: string[] = ["Edit"];
  const res = await contextMenu({
    position: {
      x: e.clientX,
      y: e.clientY,
    },
    options,
  });
  switch (res) {
    case "Edit":
      selectConnection(conn);
      break;
    default:
      break;
  }
}

async function selectConnection(conn: DBConnection): Promise<void> {
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

  if (formRef.value && formRef.value.nameTextFieldRef && formRef.value.nameTextFieldRef.inputRef) {
    await nextTick();
    formRef.value.nameTextFieldRef.inputRef.focus();
  }
}

async function addNewConnection(action: "form" | "import"): Promise<void> {
  if (
    action === "form" &&
    formRef.value &&
    formRef.value.nameTextFieldRef &&
    formRef.value.nameTextFieldRef.inputRef
  ) {
    showConnectionForm.value = true;
    await nextTick();
    formRef.value.nameTextFieldRef.inputRef.focus();
  } else if (action === "import" && configRef.value) {
    showConfigModal.value = true;
    await nextTick();
    if (configRef.value.configTextAreaRef && configRef.value.configTextAreaRef.inputRef) {
      configRef.value.configTextAreaRef.inputRef.focus();
    }
  }
}

async function dontAddNewConnection(action: "form" | "import"): Promise<void> {
  if (action === "form" && newConnectionBtnRef.value && newConnectionBtnRef.value.btnRef) {
    hideConnectionForm();
    await nextTick();
    newConnectionBtnRef.value.btnRef.focus();
  } else if (action === "import" && importConfigBtnRef.value && importConfigBtnRef.value.btnRef) {
    showConfigModal.value = false;
    await nextTick();
    importConfigBtnRef.value.btnRef.focus();
  }
}

function hideConnectionForm(): void {
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

function selectDefaultDatabaseDriver(): void {
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

function handleKeydown(e: KeyboardEvent): void {
  if (e.key === "Escape" && showConnectionForm.value) {
    dontAddNewConnection("form");
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

onUnmounted(function (): void {
  window.removeEventListener("keydown", handleKeydown);
});
</script>

<template>
  <ConfigModal
    ref="configRef"
    :tab-index-start="16 + filteredConnections.length"
    :is-open="showConfigModal"
    @close="dontAddNewConnection('import')"
    @submit="setConnection"
  />
  <div class="h-screen flex text-text-light dark:text-text-dark tab">
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
            <BasePrimaryButton
              ref="newConnectionBtnRef"
              type="button"
              tabindex="1"
              @click="addNewConnection('form')"
            >
              New Connection
            </BasePrimaryButton>
            <BaseBorderButton
              ref="importConfigBtnRef"
              type="button"
              tabindex="2"
              @click="addNewConnection('import')"
            >
              Import Config
            </BaseBorderButton>
          </div>
        </div>
        <div
          v-if="useConnectionStore().connections.length > 0"
          class="self-start min-w-[650px] max-w-[650px] flex flex-col gap-[20px]"
        >
          <div class="px-5 flex justify-between items-center">
            <p class="font-normal text-xl">Recent Connections</p>
            <BaseTextField
              tabindex="3"
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
              :tabindex="4 + index"
              :key="`connection-list-name-${index}`"
              :class="[
                'py-2 px-5 flex-wrap cursor-pointer last:border-b-0',
                {
                  'text-select-options-text-light focus-visible:bg-select-options-background-hovered-light hover:bg-select-options-background-hovered-light border-b border-textfield-border-light dark:text-select-options-text-dark dark:focus-visible:bg-select-options-background-hovered-dark dark:hover:bg-select-options-background-hovered-dark dark:border-textfield-border-dark':
                    useConnectionStore().selectedConnection?.id !== connection.id,
                  'text-select-options-text-selected-light bg-select-options-background-selected-light hover:bg-select-options-background-selected-hover-light border-b border-textfield-border-light dark:text-select-options-text-selected-dark dark:bg-select-options-background-selected-dark dark:hover:bg-select-options-background-selected-hover-dark dark:border-textfield-border-dark':
                    useConnectionStore().selectedConnection?.id === connection.id,
                },
              ]"
              @click="selectConnection(connection)"
              @keydown.enter="selectConnection(connection)"
              @keydown.space.prevent="selectConnection(connection)"
              @contextmenu.prevent="openContextMenu($event, connection)"
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
              <BaseBorderedIconButton
                class="absolute top-6 right-0"
                :tabindex="15 + filteredConnections.length"
                @click="dontAddNewConnection('form')"
              >
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
              :tab-index-start="3 + filteredConnections.length"
              @save="handleSave"
              @test="handleTest"
              @connect="handleConnect"
              @cancel="dontAddNewConnection('form')"
            />
          </div>
        </div>
      </Transition>
    </main>
  </div>
</template>
