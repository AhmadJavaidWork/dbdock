<script setup lang="ts" generic="T extends string | number | readonly string[] | null | undefined">
import { Ref, ref } from "vue";

const {
  name = "",
  label = "",
  maxlength = undefined,
  placeholder = "",
  required = false,
  error = null,
  type = "text",
  rows = 3,
} = defineProps<{
  name: string;
  label?: string;
  maxlength?: number;
  placeholder?: string;
  required?: boolean;
  error?: string | null;
  type?: string;
  rows?: number | string;
}>();

const emit = defineEmits<{
  focus: [event: FocusEvent];
}>();

const model = defineModel<T | null>({ required: true });

const inputRef = ref<HTMLInputElement | null>(null);

defineExpose<{
  inputRef: Ref<HTMLInputElement | null>;
}>({
  inputRef,
});
</script>

<template>
  <div v-bind="$attrs">
    <label v-if="label" :for="name" class="font-medium">
      {{ label }} <span v-if="required" class="text-danger">*</span>
    </label>
    <textarea
      ref="inputRef"
      :id="name"
      v-model="model"
      :rows="rows"
      :type="type"
      :placeholder="placeholder"
      :maxlength="maxlength"
      :class="[
        'rounded w-full p-2 placeholder:text-sm min-h-[35px] ring-1 ring-inset outline-none',
        {
          'bg-textfield-empty-light ring-textfield-border-light placeholder:text-textfield-placeholder-light dark:bg-textfield-empty-dark dark:ring-textfield-border-dark dark:placeholder:text-textfield-placeholder-dark':
            !model && !error,
          'bg-textfield-filled-light dark:bg-textfield-filled-dark ring-textfield-border-light dark:ring-textfield-border-dark':
            model && !error,
          'focus:outline-none focus:ring-textfield-border-focused-border-light focus:bg-textfield-filled-light dark:focus:ring-textfield-border-focused-border-dark dark:focus:bg-textfield-filled-dark':
            !error,
          'ring-textfield-border-light bg-textfield-error-light focus:ring-textfield-border-error-light focus:bg-textfield-error-light placeholder:text-textfield-placeholder-error-light dark:ring-textfield-border-dark dark:bg-textfield-error-dark dark:focus:ring-textfield-border-error-dark dark:focus:bg-textfield-error-dark dark:placeholder:text-textfield-placeholder-error-dark':
            error,
        },
      ]"
      @focus="emit('focus', $event)"
    />
    <p v-if="error" class="text-text-danger text-sm mt-[-3px]">
      {{ error }}
    </p>
  </div>
</template>
