<script lang="ts" setup>
import BaseBorderButton from "@/components/buttons/BaseBorderButton.vue";
import BasePrimaryButton from "@/components/buttons/BasePrimaryButton.vue";
import FormSelect from "@/components/inputs/FormSelect.vue";
import FormTextField from "@/components/inputs/FormTextField.vue";
import { usePrompt } from "@/composables/usePrompt";
import { connectionErrorToText } from "@/errors/connection.error";
import { testConnection } from "@/services/connection.service";
import { fetchDatabases } from "@/services/database.service";
import { DatabaseDriver } from "@/types/database.types";
import { connectionRules } from "@/validators/connection.rules";
import { useVuelidate } from "@vuelidate/core";
import { computed, onMounted, ref, watch } from "vue";

const { prompt } = usePrompt();

const databases = ref<DatabaseDriver[]>([]);

const name = ref<string>("");
const type = ref<DatabaseDriver | null>(null);
const host = ref<string>("");
const port = ref<number>();
const username = ref<string>("");
const password = ref<string>("");
const database = ref<string>("");

const v$ = useVuelidate(connectionRules, { name, type, host, port, username, password, database });

const nameError = computed((): string => v$.value.name.$errors[0]?.$message.toString() ?? null);
const typeError = computed((): string => v$.value.type.$errors[0]?.$message.toString() ?? null);
const hostError = computed((): string => v$.value.host.$errors[0]?.$message.toString() ?? null);
const portError = computed((): string => v$.value.port.$errors[0]?.$message.toString() ?? null);
const usernameError = computed(
  (): string => v$.value.username.$errors[0]?.$message.toString() ?? null
);
const passwordError = computed(
  (): string => v$.value.password.$errors[0]?.$message.toString() ?? null
);
const databaseError = computed(
  (): string => v$.value.database.$errors[0]?.$message.toString() ?? null
);

const connecting = ref<boolean>(false);
const testing = ref<boolean>(false);

async function connect(): Promise<string | undefined> {
  const valid = await v$.value.$validate();
  if (!valid) return;
  if (!type.value || !port.value) return;
  return testConnection({
    name: name.value,
    type: type.value,
    host: host.value,
    port: port.value,
    username: username.value,
    password: password.value,
    database: database.value,
  });
}

async function handleTest(): Promise<void> {
  testing.value = true;
  if (!type.value || !port.value) return;
  try {
    const res = await connect();
    if (!res) return;
    prompt({
      type: "success",
      title: "Connection status",
      message: "Connection Successful",
    });
  } catch (error: unknown) {
    const message = connectionErrorToText(error, type.value, username.value, database.value);
    prompt({
      type: "error",
      title: "Connection failed",
      message,
    });
    console.log("Connection failed: ", message);
  }
  testing.value = false;
}

async function createConnection(): Promise<void> {
  connecting.value = true;
  if (!type.value || !port.value) return;
  try {
    const res = await connect();
    if (!res) return;
    prompt({
      type: "success",
      title: "Connection status",
      message: "Connection Successful",
    });
  } catch (error: unknown) {
    const message = connectionErrorToText(error, type.value, username.value);
    prompt({
      type: "error",
      title: "Connection failed",
      message,
    });
    console.log("Connection failed: ", message);
  }
  connecting.value = false;
}

onMounted(async function (): Promise<void> {
  databases.value = await fetchDatabases();

  if (databases.value.length > 0) {
    type.value = databases.value[0];
    port.value = databases.value[0].defaultPort;
  }
});

watch(type, function (newType) {
  if (newType) {
    const db = databases.value.find((d) => d.id === newType.id);
    if (db) {
      port.value = db.defaultPort;
    }
  }
});
</script>

<template>
  <div class="grid grid-cols-4 gap-x-[15px] gap-y-[10px]">
    <FormTextField
      class="col-span-4 lg:col-span-4"
      v-model="name"
      label="Name"
      name="name"
      required
      placeholder="Connection Name"
      :error="nameError"
    />
    <FormSelect
      class="col-span-4 lg:col-span-4"
      v-model="type"
      label="Type"
      name="type"
      required
      placeholder="Select Type"
      option-key="id"
      option-label="label"
      :options="databases"
      :error="typeError"
    />
    <FormTextField
      class="col-span-4 lg:col-span-2"
      v-model="host"
      label="Host"
      name="host"
      required
      placeholder="DB Host"
      :error="hostError"
    />
    <FormTextField
      class="col-span-4 lg:col-span-2"
      v-model.number="port"
      label="Port"
      name="port"
      required
      placeholder="DB Port"
      :error="portError"
    />
    <FormTextField
      class="col-span-4 lg:col-span-2"
      v-model="username"
      label="Username"
      name="username"
      required
      placeholder="Username"
      :error="usernameError"
    />
    <FormTextField
      class="col-span-4 lg:col-span-2"
      v-model="password"
      label="Password"
      name="password"
      required
      type="password"
      placeholder="Password"
      :error="passwordError"
    />
    <FormTextField
      class="col-span-4 lg:col-span-4"
      v-model="database"
      label="Database"
      name="databse"
      required
      placeholder="Database name"
      :error="databaseError"
    />
    <div class="col-span-4 flex justify-end gap-[15px]">
      <BaseBorderButton
        type="button"
        @click="handleTest"
        :loading="connecting"
        :disabled="testing || connecting"
      >
        {{ testing ? "Testing Connection..." : "Test Connection" }}
      </BaseBorderButton>
      <BasePrimaryButton
        :disabled="connecting || testing"
        :loading="connecting"
        @click="createConnection"
      >
        {{ connecting ? "Connecting..." : "Connect" }}
      </BasePrimaryButton>
    </div>
  </div>
</template>
