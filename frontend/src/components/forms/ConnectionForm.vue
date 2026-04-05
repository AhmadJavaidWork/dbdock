<script lang="ts" setup>
import BaseBorderButton from "@/components/buttons/BaseBorderButton.vue";
import BasePrimaryButton from "@/components/buttons/BasePrimaryButton.vue";
import FormSelect from "@/components/inputs/FormSelect.vue";
import FormTextField from "@/components/inputs/FormTextField.vue";
import { CreateDBConnection } from "@/types/connection.type";
import { DatabaseDriver } from "@/types/databaseDriver.types";
import { connectionRules } from "@/validators/connection.rules";
import { useVuelidate } from "@vuelidate/core";
import { computed, ref, watch } from "vue";

const emit = defineEmits<{
  save: [];
  test: [];
  connect: [];
}>();

const {
  databaseDrivers = [],
  testing = false,
  saving = false,
  connecting = false,
} = defineProps<{
  databaseDrivers: DatabaseDriver[];
  testing: boolean;
  saving: boolean;
  connecting: boolean;
}>();

const form = defineModel<CreateDBConnection>("connection", { required: true });
const databaseDriver = ref<DatabaseDriver | null>(null);

const v$ = useVuelidate(connectionRules, form);

const nameError = computed((): string => v$.value.name.$errors[0]?.$message.toString() ?? null);
const typeError = computed(
  (): string => v$.value.databaseDriverId.$errors[0]?.$message.toString() ?? null
);
const hostError = computed((): string => v$.value.host.$errors[0]?.$message.toString() ?? null);
const portError = computed((): string => v$.value.port.$errors[0]?.$message.toString() ?? null);
const usernameError = computed(
  (): string => v$.value.username.$errors[0]?.$message.toString() ?? null
);
const passwordError = computed(
  (): string => v$.value.password.$errors[0]?.$message.toString() ?? null
);
const databaseError = computed(
  (): string => v$.value.databaseName.$errors[0]?.$message.toString() ?? null
);

watch<number | null>(
  () => form.value.databaseDriverId,
  (ddId: number | null) => {
    if (ddId) {
      const dd = databaseDrivers.find((dd) => dd.id === ddId);
      if (dd) {
        databaseDriver.value = dd;
        form.value.port = dd.defaultPort;
      }
    }
  }
);

defineExpose({ v$, databaseDriver });
</script>

<template>
  <div>
    <h1 class="text-xl font-bold mb-[20px]">New Connection</h1>
    <div class="grid grid-cols-4 gap-x-[15px] gap-y-[10px]">
      <FormTextField
        class="col-span-4 lg:col-span-4"
        v-model="form.name"
        label="Name"
        name="name"
        required
        placeholder="Connection Name"
        :error="nameError"
      />
      <FormSelect
        class="col-span-4 lg:col-span-4"
        v-model="databaseDriver"
        label="Type"
        name="type"
        required
        placeholder="Select Type"
        option-key="id"
        option-label="label"
        :options="databaseDrivers"
        :error="typeError"
      />
      <FormTextField
        class="col-span-4 lg:col-span-2"
        v-model="form.host"
        label="Host"
        name="host"
        required
        placeholder="DB Host"
        :error="hostError"
      />
      <FormTextField
        class="col-span-4 lg:col-span-2"
        v-model.number="form.port"
        label="Port"
        name="port"
        required
        placeholder="DB Port"
        :error="portError"
      />
      <FormTextField
        class="col-span-4 lg:col-span-2"
        v-model="form.username"
        label="Username"
        name="username"
        required
        placeholder="Username"
        :error="usernameError"
      />
      <FormTextField
        class="col-span-4 lg:col-span-2"
        v-model="form.password"
        label="Password"
        name="password"
        required
        type="password"
        placeholder="Password"
        :error="passwordError"
      />
      <FormTextField
        class="col-span-4 lg:col-span-4"
        v-model="form.databaseName"
        label="Database Name"
        name="databaseName"
        required
        placeholder="Database name"
        :error="databaseError"
      />
    </div>
    <div class="flex justify-between mt-[20px]">
      <div class="flex justify-start gap-[15px]">
        <BasePrimaryButton
          type="button"
          :disabled="saving || testing || connecting"
          :loading="saving"
          @click="emit('save')"
        >
          {{ connecting ? "Saving..." : "Save" }}
        </BasePrimaryButton>
        <BaseBorderButton
          type="button"
          :disabled="saving || testing || connecting"
          :loading="testing"
          @click="emit('test')"
        >
          {{ testing ? "Testing..." : "Test" }}
        </BaseBorderButton>
      </div>
      <BasePrimaryButton
        type="button"
        :disabled="saving || testing || connecting"
        :loading="connecting"
        @click="emit('connect')"
      >
        {{ connecting ? "Connecting..." : "Connect" }}
      </BasePrimaryButton>
    </div>
  </div>
</template>
