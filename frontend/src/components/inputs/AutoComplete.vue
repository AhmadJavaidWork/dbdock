<script setup lang="ts" generic="T extends string | Record<string, any>">
import IconSpinner from "@/components/icons/IconSpinner.vue";
import BaseTextField from "@/components/inputs/BaseTextField.vue";
import debounce from "lodash.debounce";
import { ComponentPublicInstance, onBeforeUnmount, onMounted, ref, watch } from "vue";

type BaseTextFieldExposed = {
  inputRef: HTMLInputElement | null;
};

const props = defineProps<{
  options?: T[];
  loadOptions?: (
    name: string,
    query: string,
    page: number
  ) => Promise<{ rows: T[]; count: number }>;
  labelKey?: T extends string ? never : keyof T;
  placeholder?: string;
  label?: string;
  required?: boolean;
  error?: string | undefined;
  name: string;
  pageSize?: number;
}>();

const emit = defineEmits<{
  "reach-bottom": [name: string, query: string, page: number];
}>();

const modelValue = defineModel<T | null>("modelValue", { required: true });

const search = ref<string>(modelValue.value ? getLabel(modelValue.value) : "");
const isOpen = ref<boolean>(false);
const filteredOptions = ref<T[]>([]);
const page = ref(0);
const loading = ref(false);
const pages = ref(0);

const containerRef = ref<HTMLElement | null>(null);
const listRef = ref<HTMLElement | null>(null);
const textFieldRef = ref<ComponentPublicInstance<BaseTextFieldExposed> | null>(null);

function getLabel(option: T): string {
  if (typeof option === "string") return option;
  return String(option[props.labelKey as keyof T] ?? "");
}

function filterOptions() {
  if (!props.options) return;
  const query = search.value.toLowerCase();
  filteredOptions.value = props.options.filter((option) =>
    getLabel(option).toLowerCase().includes(query)
  );
}

async function loadAsync(reset = false) {
  if (!props.loadOptions || loading.value) return;

  if (reset) {
    page.value = 0;
    filteredOptions.value = [];
  }

  loading.value = true;
  const result = await props.loadOptions(props.name, search.value, page.value);
  pages.value = Math.ceil(result.count / (props.pageSize || 1));
  filteredOptions.value.push(...result.rows);
  loading.value = false;
}

const debouncedLoad = debounce(() => {
  loadAsync(true);
}, 500);

function selectValue(option: T) {
  modelValue.value = option;
}

function selectOption(option: T) {
  search.value = getLabel(option);
  isOpen.value = false;
}

const dropdownStyle = ref<Record<string, string>>({});

function updateDropdownPosition() {
  const input = textFieldRef.value?.inputRef;
  if (!input) return;

  const rect = input.getBoundingClientRect();

  dropdownStyle.value = {
    top: `${rect.bottom}px`,
    left: `${rect.left}px`,
    width: `${rect.width}px`,
  };
}

function onScroll() {
  const el = listRef.value;
  if (!el) return;

  const atBottom = el.scrollTop + el.clientHeight >= el.scrollHeight - 5;
  if (page.value + 1 === pages.value) return;

  if (atBottom) {
    emit("reach-bottom", props.name, search.value, page.value + 1);
    page.value++;
    loadAsync();
  }
}

watch(
  () => isOpen.value,
  async () => {
    if (isOpen.value) {
      updateDropdownPosition();
      if (props.loadOptions) {
        await loadAsync(true);
      } else {
        filterOptions();
      }
    }
  }
);

watch(
  () => search.value,
  () => {
    if (props.loadOptions) {
      debouncedLoad();
    } else {
      filterOptions();
    }
  }
);

watch(
  () => modelValue.value,
  () => {
    if (modelValue.value) {
      selectOption(modelValue.value);
    }
  }
);

function handleClickOutside(e: MouseEvent) {
  if (!containerRef.value) return;
  if (!containerRef.value.contains(e.target as Node)) {
    isOpen.value = false;
  }
}

onMounted(() => {
  document.addEventListener("click", handleClickOutside);
  window.addEventListener("resize", updateDropdownPosition);
});

onBeforeUnmount(() => {
  debouncedLoad.cancel();
  document.removeEventListener("click", handleClickOutside);
  window.removeEventListener("resize", updateDropdownPosition);
});
</script>

<template>
  <div class="relative w-full" ref="containerRef">
    <BaseTextField
      ref="textFieldRef"
      v-model="search"
      :name="`${name}-autocomplete-search`"
      :label="label"
      :placeholder="placeholder"
      :required="required"
      :error="error"
      @focus="isOpen = true"
    />

    <Teleport to="body">
      <ul
        v-if="isOpen && filteredOptions.length"
        ref="listRef"
        :style="dropdownStyle"
        class="fixed bg-white border rounded-md shadow-md z-[9999] max-h-[200px] overflow-y-auto"
        @scroll="onScroll"
      >
        <li
          v-for="(option, index) in filteredOptions"
          :key="index"
          class="px-4 py-2 text-black text-sm hover:bg-gray-100 cursor-pointer"
          @click="selectValue(option as T)"
        >
          {{ getLabel(option as T) }}
        </li>
        <li v-show="loading" class="flex items-center justify-center py-2">
          <IconSpinner color="[#093EFE]" />
        </li>
      </ul>
    </Teleport>
  </div>
</template>
