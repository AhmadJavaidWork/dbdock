<script setup lang="ts">
import IconAlertCircle from "@/components/icons/IconAlertCircle.vue";
import IconCheckCircle from "@/components/icons/IconCheckCircle.vue";
import IconCloseCircle from "@/components/icons/IconCloseCircle.vue";
import IconInfoCircle from "@/components/icons/IconInfoCircle.vue";
import BaseModal from "@/components/modals/BaseModal.vue";
import { usePrompt } from "@/composables/usePrompt";

const { isOpen, options, onClose } = usePrompt();
</script>

<template>
  <BaseModal
    :is-open="isOpen"
    :title="options.title"
    footer-class="px-0 py-0"
    body-class="flex flex-col items-center gap-[15px]"
    @close="onClose"
  >
    <template #modal-body>
      <IconCheckCircle
        v-if="options.type === 'success'"
        class="text-text-success dark:text-text-success-darker w-[44px] h-[44px]"
      />
      <IconInfoCircle
        v-else-if="options.type === 'info'"
        class="text-text-info dark:text-text-info-darker w-[44px] h-[44px]"
      />
      <IconAlertCircle
        v-else-if="options.type === 'warn'"
        class="text-text-warning dark:text-text-warning-darker w-[44px] h-[44px]"
      />
      <IconCloseCircle
        v-else-if="options.type === 'error'"
        class="text-text-danger dark:text-text-danger-darker w-[44px] h-[44px]"
      />

      <p class="font-medium text-center">{{ options.message }}</p>
    </template>
    <template #modal-footer>
      <button
        class="font-medium px-[20px] py-[8px] w-full bg-background-light dark:bg-background-dark text-text-light dark:text-text-dark rounded-e-prompt rounded-b-prompt hover:bg-select-options-background-hovered-light dark:hover:bg-select-options-background-hovered-dark"
        @click="onClose"
      >
        {{ options.closeBtnTxt }}
      </button>
    </template>
  </BaseModal>
</template>
