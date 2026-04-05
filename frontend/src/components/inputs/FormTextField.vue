<script setup lang="ts" generic="T">
import { Ref, ref } from "vue";

const {
  label = "",
  name = "",
  maxlength = undefined,
  placeholder = "",
  required = false,
  error = null,
  type = "text",
} = defineProps<{
  label?: string;
  name: string;
  maxlength?: number;
  placeholder?: string;
  required?: boolean;
  error?: string | null;
  type?: string;
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
    <input
      ref="inputRef"
      :id="name"
      v-model="model"
      :type="type"
      :placeholder="placeholder"
      :maxlength="maxlength"
      :class="[
        'rounded w-full p-2 placeholder:text-sm min-h-[35px] ring-1 ring-inset outline-none',
        {
          'bg-textfield-empty-light ring-textfield-border-light placeholder:text-textfield-placeholder-light dark:bg-textfield-empty-dark dark:ring-textfield-border-dark dark:placeholder:text-textfield-placeholder-dark':
            !model && !error,
          'bg-textfield-filled-light dark:bg-textfield-filled-dark': model && !error,
          'focus:outline-none focus:ring-textfield-border-focused-border-light focus:bg-textfield-filled-light dark:focus:ring-textfield-border-focused-border-dark dark:focus:bg-textfield-filled-dark':
            !error,
          'ring-textfield-border-light bg-textfield-error-light focus:ring-textfield-border-error-light focus:bg-textfield-error-light placeholder:text-textfield-placeholder-error-light dark:ring-textfield-border-dark dark:bg-textfield-error-dark dark:focus:ring-textfield-border-error-dark dark:focus:bg-textfield-error-dark dark:placeholder:text-textfield-placeholder-error-dark':
            error,
          'ring-textfield-border-light bg-textfield-filled-light dark:ring-textfield-border-dark dark:bg-textfield-filled-dark':
            model,
        },
      ]"
      @focus="emit('focus', $event)"
    />
    <p v-if="error" class="text-text-danger text-sm mt-[3px]">
      {{ error }}
    </p>
  </div>
</template>
