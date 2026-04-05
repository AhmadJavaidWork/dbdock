<script setup lang="ts">
import ConnectionForm from "@/components/forms/ConnectionForm.vue";
import ConnectionsList from "@/components/sidebar/ConnectionsList.vue";
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
import { onMounted, ref } from "vue";

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
}

async function handleTest(): Promise<void> {
  if (!formRef.value) return;
  const valid = await formRef.value?.v$.$validate();
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

async function upateConnection(): Promise<void> {
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

function addNewConnection() {
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
}

function selectDefaultDatabaseDriver() {
  if (databaseDrivers.value.length > 0) {
    connection.value.databaseDriverId = databaseDrivers.value[0].id;
    connection.value.port = databaseDrivers.value[0].defaultPort;
  }
}

async function handleSave(): Promise<void> {
  if (useConnectionStore().selectedConnection) {
    upateConnection();
  } else {
    createConnection();
  }
}

async function handleConnect(): Promise<void> {}

onMounted(async function (): Promise<void> {
  try {
    databaseDrivers.value = await fetchDatabaseDrivers();

    selectDefaultDatabaseDriver();
  } catch (error) {
    useToast(error as string, "error");
    console.log("error", error);
  }
});
</script>

<template>
  <div
    class="h-screen flex bg-background-light dark:bg-background-dark text-text-light dark:text-text-dark"
  >
    <ConnectionsList @select-connection="selectConnection" @add-new-connection="addNewConnection" />
    <main class="flex-1 overflow-hidden flex flex-col">
      <div class="max-h-screen flex items-center justify-center dark:bg-background-dark">
        <div class="p-6 flex flex-col items-center">
          <div class="text-2xl font-bold mb-[10px]">
            <span class="text text-[#60A5FA]">DB</span>
            <span class="text-primary">Dock</span>
          </div>
          <img id="logo" alt="Wails logo" src="@/assets/images/logo.svg" />
          <ConnectionForm
            ref="formRef"
            :database-drivers="databaseDrivers"
            :connection="connection"
            :saving="saving"
            :testing="testing"
            :connecting="connecting"
            @save="handleSave"
            @test="handleTest"
            @connect="handleConnect"
          />
        </div>
      </div>
    </main>
  </div>
</template>
