<script lang="ts" setup>
import BaseBorderButton from "@/components/buttons/BaseBorderButton.vue";
import BasePrimaryButton from "@/components/buttons/BasePrimaryButton.vue";
import BaseSelectField from "@/components/inputs/BaseSelectField.vue";
import BaseTextField from "@/components/inputs/BaseTextField.vue";
import { CreateDBConnection } from "@/types/connection.type";
import { DatabaseDriver } from "@/types/databaseDriver.types";
import { connectionRules } from "@/validators/connection.rules";
import { useVuelidate, Validation } from "@vuelidate/core";
import { computed, Ref, ref, watch } from "vue";

const emit = defineEmits<{
  save: [];
  test: [];
  connect: [];
  cancel: [];
}>();

const {
  databaseDrivers = [],
  testing = false,
  saving = false,
  connecting = false,
  tabIndexStart = 1,
} = defineProps<{
  databaseDrivers: DatabaseDriver[];
  testing: boolean;
  saving: boolean;
  connecting: boolean;
  tabIndexStart: number;
}>();

const form = defineModel<CreateDBConnection>("connection", { required: true });
const databaseDriver = ref<DatabaseDriver | null>(null);
const nameTextFieldRef = ref<InstanceType<typeof BaseTextField> | null>(null);

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
        if (!form.value.port) form.value.port = dd.defaultPort;
      }
    }
  }
);

defineExpose<{
  v$: Ref<Validation>;
  databaseDriver: Ref<DatabaseDriver | null>;
  nameTextFieldRef: Ref<InstanceType<typeof BaseTextField> | null>;
}>({ v$, databaseDriver, nameTextFieldRef });
</script>

<template>
  <div class="flex flex-col gap-[20px]">
    <h1 class="text-xl font-bold">New Connection</h1>
    <div class="grid grid-cols-4 gap-x-[15px] gap-y-[10px]">
      <BaseTextField
        ref="nameTextFieldRef"
        :tabindex="tabIndexStart + 1"
        class="col-span-4 lg:col-span-4"
        v-model="form.name"
        label="Name"
        name="name"
        required
        placeholder="Connection Name"
        :error="nameError"
      />
      <BaseSelectField
        :tabindex="tabIndexStart + 2"
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
      <BaseTextField
        :tabindex="tabIndexStart + 3"
        class="col-span-4 lg:col-span-2"
        v-model="form.host"
        label="Host"
        name="host"
        required
        placeholder="DB Host"
        :error="hostError"
      />
      <BaseTextField
        :tabindex="tabIndexStart + 4"
        class="col-span-4 lg:col-span-2"
        v-model.number="form.port"
        label="Port"
        name="port"
        required
        placeholder="DB Port"
        :error="portError"
      />
      <BaseTextField
        :tabindex="tabIndexStart + 5"
        class="col-span-4 lg:col-span-2"
        v-model="form.username"
        label="Username"
        name="username"
        required
        placeholder="Username"
        :error="usernameError"
      />
      <BaseTextField
        :tabindex="tabIndexStart + 6"
        class="col-span-4 lg:col-span-2"
        v-model="form.password"
        label="Password"
        name="password"
        required
        type="password"
        placeholder="Password"
        :error="passwordError"
      />
      <BaseTextField
        :tabindex="tabIndexStart + 7"
        class="col-span-4 lg:col-span-4"
        v-model="form.databaseName"
        label="Database Name"
        name="databaseName"
        required
        placeholder="Database name"
        :error="databaseError"
      />
    </div>
    <div class="flex justify-between">
      <div class="flex justify-start gap-[15px]">
        <BasePrimaryButton
          :tabindex="tabIndexStart + 8"
          type="button"
          :disabled="saving || testing || connecting"
          :loading="saving"
          @click="emit('save')"
        >
          {{ connecting ? "Saving..." : "Save" }}
        </BasePrimaryButton>
        <BaseBorderButton
          :tabindex="tabIndexStart + 9"
          type="button"
          :disabled="saving || testing || connecting"
          :loading="testing"
          @click="emit('test')"
        >
          {{ testing ? "Testing..." : "Test" }}
        </BaseBorderButton>
      </div>
      <div class="flex justify-start gap-[15px]">
        <BasePrimaryButton
          :tabindex="tabIndexStart + 10"
          type="button"
          :disabled="saving || testing || connecting"
          :loading="connecting"
          @click="emit('connect')"
        >
          {{ connecting ? "Connecting..." : "Connect" }}
        </BasePrimaryButton>
        <BaseBorderButton
          :tabindex="tabIndexStart + 11"
          type="button"
          :disabled="saving || testing || connecting"
          @click="emit('cancel')"
        >
          Cancel
        </BaseBorderButton>
      </div>
    </div>
  </div>
</template>
