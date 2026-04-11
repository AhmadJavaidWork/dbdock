<script setup lang="ts">
import BaseBorderButton from "@/components/buttons/BaseBorderButton.vue";
import BasePrimaryButton from "@/components/buttons/BasePrimaryButton.vue";
import BaseTextArea from "@/components/inputs/BaseTextArea.vue";
import BaseModal from "@/components/modals/BaseModal.vue";
import { fetchDatabaseDrivers } from "@/services/database.service";
import { CreateDBConnection } from "@/types/connection.type";
import { DatabaseDriver } from "@/types/databaseDriver.types";
import { useVuelidate } from "@vuelidate/core";
import { helpers, required } from "@vuelidate/validators";
import { computed, onMounted, Ref, ref } from "vue";

const emit = defineEmits<{
  close: [];
  submit: [conn: CreateDBConnection];
}>();

const { isOpen = false, tabIndexStart = 1 } = defineProps<{
  isOpen: boolean;
  tabIndexStart: number;
}>();

let dragCounter = 0;
const dragging = ref<boolean>(false);
const fileInput = ref<HTMLInputElement | null>(null);
const config = ref<string>("");
const connection = ref<CreateDBConnection>({
  name: "",
  databaseDriverId: null,
  host: "",
  port: null,
  username: "",
  password: "",
  databaseName: "",
});
const databaseDrivers = ref<DatabaseDriver[]>([]);
const configTextAreaRef = ref<InstanceType<typeof BaseTextArea> | null>(null);

const rules = {
  config: {
    required: helpers.withMessage("Please paste a JSON or ENV", required),
    invalid: helpers.withMessage("Invalid JSON or ENV", (val: string): boolean => {
      try {
        return !!parseConfig(val);
      } catch (error) {
        return false;
      }
    }),
  },
};

const v$ = useVuelidate(rules, { config });

const configError = computed((): string => v$.value.config.$errors[0]?.$message.toString() ?? null);

function getField(names: string[], map: any): string {
  for (let i = 0; i < names.length; i++) {
    if (map[names[i]]) return map[names[i]];
  }
  return "";
}

function onDragEnter(e: DragEvent): void {
  e.preventDefault();
  dragCounter++;
  dragging.value = true;
}

function onDragLeave(e: DragEvent): void {
  e.preventDefault();
  dragCounter--;
  if (dragCounter === 0) {
    dragging.value = false;
  }
}

function normalize(input: any): CreateDBConnection {
  const hostFieldNames = ["host", "HOST", "DB_HOST", "DBHOST", "DATABASE_HOST", "DATABASEHOST"];
  const portFieldNames = ["port", "PORT", "DB_PORT", "DBPORT", "DATABASE_PORT", "DATABASEPORT"];
  const usernameFieldNames = [
    "user",
    "username",
    "USER",
    "USERNAME",
    "DB_USER",
    "DB_USERNAME",
    "DBUSER",
    "DBUSERNAME",
    "DATABASE_USER",
    "DATABASE_USERNAME",
    "DATABASEUSER",
    "DATABASEUSERNAME",
  ];
  const passwordFieldNames = [
    "password",
    "PASSWORD",
    "DB_PASSWORD",
    "DBPASSWORD",
    "DATABASE_PASSWORD",
    "DATABASEPASSWORD",
  ];
  const databaseNameFieldNames = [
    "database",
    "dbName",
    "dbname",
    "databaseName",
    "databasename",
    "DATABASE",
    "DBNAME",
    "DATABASENAME",
    "DB_NAME",
    "DATABASE_NAME",
  ];
  const driverFieldNames = [
    "driver",
    "dbDriver",
    "DRIVER",
    "DB_DRIVER",
    "DATABASE_DRIVER",
    "DBDRIVER",
    "DATABASEDRIVER",
  ];

  const driver = getField(driverFieldNames, input);

  let databaseDriverId = null;
  for (let i = 0; i < databaseDrivers.value.length; i++) {
    if (driver.toLowerCase() === databaseDrivers.value[i].name.toLowerCase()) {
      databaseDriverId = databaseDrivers.value[i].id;
    }
  }

  return {
    name: "",
    host: getField(hostFieldNames, input),
    port: Number(getField(portFieldNames, input)),
    username: getField(usernameFieldNames, input),
    password: getField(passwordFieldNames, input),
    databaseName: getField(databaseNameFieldNames, input),
    databaseDriverId,
  };
}

function parseConfig(text: string): CreateDBConnection | undefined {
  if (text.trim().startsWith("{")) {
    const json = JSON.parse(text);
    return normalize(json);
  }

  if (text.includes("=")) {
    const env: Record<string, string> = {};

    text.split("\n").forEach((line) => {
      const [key, value] = line.split("=");
      if (key && value) env[key.trim()] = value.trim();
    });

    return normalize(env);
  }
}

function close(): void {
  v$.value.$reset();
  connection.value = {
    name: "",
    databaseDriverId: null,
    host: "",
    port: null,
    username: "",
    password: "",
    databaseName: "",
  };
  config.value = "";
  dragging.value = false;
  emit("close");
}

async function submit(): Promise<void> {
  const valid = await v$.value.$validate();
  if (!valid) return;
  const conn = parseConfig(config.value);
  if (conn) {
    connection.value = conn;
    emit("submit", connection.value);
    close();
  }
}

async function handleFileUpload(e: Event): Promise<void> {
  const file = (e.target as HTMLInputElement).files?.[0];
  if (!file) return;

  config.value = await file.text();
  submit();
}

async function handleDrop(e: DragEvent): Promise<void> {
  const file = e.dataTransfer?.files?.[0];
  if (!file) return;

  config.value = await file.text();
  submit();
}

function openFileExplorer(): void {
  if (!fileInput.value) return;
  fileInput.value.click();
}

onMounted(async function (): Promise<void> {
  databaseDrivers.value = await fetchDatabaseDrivers();
});

defineExpose<{
  configTextAreaRef: Ref<InstanceType<typeof BaseTextArea> | null>;
}>({ configTextAreaRef });
</script>

<template>
  <BaseModal
    :is-open="isOpen"
    title="Add Database config"
    body-class="p-[30px]"
    footer-class="flex justify-end gap-[15px]"
    @close="close"
  >
    <template #modal-body>
      <div
        class="relative"
        @dragenter.prevent="onDragEnter"
        @dragover.prevent
        @dragleave.prevent="onDragLeave"
        @drop.prevent="handleDrop"
      >
        <div
          :class="[
            'flex flex-col items-center justify-center gap-[15px] transition',
            {
              'pointer-events-none': dragging,
            },
          ]"
        >
          <BaseTextArea
            ref="configTextAreaRef"
            v-model="config"
            name="config"
            placeholder="Paste JSON, or ENV..."
            label="JSON, or ENV..."
            rows="10"
            required
            :tabindex="tabIndexStart + 1"
            :error="configError"
          />
          <p class="font-bold text-xl text-primary">OR</p>
          <BaseBorderButton :tabindex="tabIndexStart + 2" @click="openFileExplorer">
            Drag and Drop or Upload Config
          </BaseBorderButton>
          <input
            type="file"
            accept=".json,.env,.txt"
            ref="fileInput"
            class="hidden"
            @change="handleFileUpload"
          />
        </div>
        <div
          v-show="dragging"
          class="absolute inset-0 z-10 flex items-center justify-center bg-background-light/1 dark:bg-background-dark/1 backdrop-blur-md border-2 border-dashed border-primary rounded-md"
        >
          <p class="text-xl font-semibold text-primary">Drop your config file here</p>
        </div>
      </div>
    </template>
    <template #modal-footer>
      <BasePrimaryButton
        btn-class="px-[30px]"
        type="button"
        :tabindex="tabIndexStart + 3"
        @click="submit"
        >Add</BasePrimaryButton
      >
      <BaseBorderButton
        btn-class="px-[30px]"
        type="button"
        :tabindex="tabIndexStart + 4"
        @click="close"
      >
        Cancel
      </BaseBorderButton>
    </template>
  </BaseModal>
</template>

<style scoped>
.modal-fade-enter-active,
.modal-fade-leave-active {
  transition: opacity 0.2s ease;
}
.modal-fade-enter-from,
.modal-fade-leave-to {
  opacity: 0;
}
</style>
