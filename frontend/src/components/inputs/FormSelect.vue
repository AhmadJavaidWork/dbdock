<script setup lang="ts" generic="T extends string | Record<string, any>">
import IconChevronDown from "@/components/icons/IconChevronDown.vue";
import { ref } from "vue";

const {
  label = "",
  name = "",
  options = [],
  placeholder = "",
  required = false,
  error = null,
  optionKey = "",
  optionLabel = "",
} = defineProps<{
  label: string;
  name: string;
  options: T[];
  optionKey?: string;
  optionLabel?: string;
  placeholder?: string;
  required?: boolean;
  error?: string | null;
}>();

const emit = defineEmits<{
  focus: [event: FocusEvent];
}>();

const isDropdownOpen = ref<boolean>(false);

const model = defineModel<T | null>({ required: true });

function closeDropdown(): void {
  isDropdownOpen.value = false;
}

function getLabel(option: T | null): string {
  if (!option) return "";
  if (typeof option === "string") {
    return option;
  }
  if (optionLabel && typeof option === "object") {
    return option[optionLabel] as string;
  }

  return "";
}

function getKey(option: T | null): string | undefined {
  if (!option) return undefined;
  if (typeof option === "string") {
    return option;
  }
  if (optionKey && typeof option === "object") {
    return option[optionKey];
  }

  return "";
}

function selectOption(option: T): void {
  model.value = option;
  isDropdownOpen.value = false;
}
</script>

<template>
  <div v-bind="$attrs">
    <div>
      <label :for="name" class="font-medium">
        {{ label }}
        <span v-if="required" class="text-text-danger dark:text-text-danger-darker">*</span>
      </label>
    </div>
    <div class="relative inline-block text-left w-full" v-click-outside="closeDropdown">
      <div>
        <button
          type="button"
          @click="isDropdownOpen = !isDropdownOpen"
          :class="[
            'inline-flex justify-between items-center w-full rounded-md border px-[10px] py-2 group max-h-[35px]',
            {
              'border-textfield-border-focused-border-light dark:border-textfield-border-focused-border-dark':
                isDropdownOpen && !error,
              'bg-textfield-filled-light border-textfield-border-light dark:bg-textfield-filled-dark dark:border-textfield-border-dark':
                !isDropdownOpen && model && !error,
              'border-textfield-border-light bg-textfield-empty-light dark:border-textfield-border-dark dark:bg-textfield-empty-dark':
                !model && !error && !isDropdownOpen,
              'focus:border-textfield-border-error-light bg-textfield-error-light dark:focus:border-textfield-border-error-dark dark:bg-textfield-error-dark':
                error,
            },
          ]"
          id="menu-button"
          :aria-expanded="isDropdownOpen"
        >
          <span v-if="model">{{ getLabel(model) }}</span>
          <span
            v-else-if="placeholder"
            :class="[
              {
                'text-textfield-placeholder-light dark:text-textfield-placeholder-dark': !error,
                'text-textfield-placeholder-error-light dark:text-textfield-placeholder-error-dark':
                  error,
              },
            ]"
            >{{ placeholder }}</span
          >
          <IconChevronDown
            :class="[
              'w-4 h-4 transition-transform text-gray-placeholder',
              {
                'group-focus:text-textfield-placeholder-error-light dark:group-focus:text-textfield-placeholder-error-light':
                  error,
                'text-textfield-border-focused-border-light dark:text-textfield-border-focused-border-light':
                  !error && (isDropdownOpen || model),
                'text-textfield-placeholder-light dark:text-textfield-placeholder-dark':
                  !error && !isDropdownOpen && !model,
              },
            ]"
            :rotate="isDropdownOpen"
          />
        </button>
      </div>
      <Transition
        enter-active-class="transition ease-out duration-200"
        enter-from-class="opacity-0 -translate-y-2"
        enter-to-class="opacity-100 translate-y-0"
        leave-active-class="transition ease-in duration-150"
        leave-from-class="opacity-100 translate-y-0"
        leave-to-class="opacity-0 -translate-y-2"
      >
        <div
          v-if="isDropdownOpen"
          class="absolute left-0 mt-1 w-full z-50 rounded-md bg-select-options-background-light dark:bg-select-options-background-dark shadow-dropdown text-center"
        >
          <div class="py-1" role="menu" aria-orientation="vertical" aria-labelledby="menu-button">
            <div v-for="(option, index) in options" :key="`name-select-option-${index}`">
              <div
                @click.stop="selectOption(option)"
                :class="[
                  'flex justify-start items-center flex-wrap px-2 py-2 cursor-pointer',
                  {
                    'text-select-options-text-selected-light bg-select-options-background-selected-light hover:bg-select-options-background-selected-hover-light dark:text-select-options-text-selected-dark dark:bg-select-options-background-selected-dark dark:hover:bg-select-options-background-selected-hover-dark':
                      getKey(model) === getKey(option),
                  },
                  {
                    'text-select-options-text-light bg-select-options-background-light hover:bg-select-options-background-hovered-light dark:text-select-options-text-dark dark:bg-select-options-background-dark dark:hover:bg-select-options-background-hovered-dark':
                      getKey(model) !== getKey(option),
                  },
                ]"
              >
                <span>{{ getLabel(option) }}</span>
              </div>
            </div>
          </div>
        </div>
      </Transition>
    </div>
    <p v-if="error" class="text-text-danger dark:text-text-danger-darker text-sm mt-[2px]">
      {{ error }}
    </p>
  </div>
</template>
