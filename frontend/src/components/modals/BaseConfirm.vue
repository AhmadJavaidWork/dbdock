<script setup lang="ts">
import BaseModal from "@/components/modals/BaseModal.vue";
import { useConfirm } from "@/composables/useConfirm";

const { isOpen, options, onClose } = useConfirm();
</script>

<template>
  <BaseModal
    :is-open="isOpen"
    :title="options.title"
    footer-class="!p-0 flex"
    body-class="flex gap-[15px]"
    @close="onClose(false)"
  >
    <template #modal-body>
      <component :is="options.icon" :class="options.iconClass" />
      <p class="font-medium">
        {{ options.message }}
      </p>
    </template>
    <template #modal-footer>
      <button
        class="font-medium px-[20px] py-[15px] w-full bg-background-light dark:bg-background-dark text-text-light dark:text-text-dark rounded-b-prompt hover:bg-select-options-background-hovered-light dark:hover:bg-select-options-background-hovered-dark border-r border-textfield-border-light dark:border-textfield-border-dark rounded-tr-none rounded-br-none"
        @click="onClose(true)"
      >
        Yes
      </button>
      <button
        class="font-medium px-[20px] py-[15px] w-full bg-background-light dark:bg-background-dark text-text-light dark:text-text-dark rounded-e-prompt hover:bg-select-options-background-hovered-light dark:hover:bg-select-options-background-hovered-dark"
        @click="onClose(false)"
      >
        Cancel
      </button>
    </template>
  </BaseModal>
</template>
