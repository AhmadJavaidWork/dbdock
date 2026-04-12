import { Component, ref } from "vue";

type ConfirmOptions = {
  title: string;
  message: string;
  icon?: Component;
  iconClass?: string;
};

const defaultOptions: ConfirmOptions = {
  title: "",
  message: "",
};

const isOpen = ref<boolean>(false);
const options = ref<ConfirmOptions>(defaultOptions);
let resolver: (value: boolean) => void;

export function useConfirm() {
  function confirm(suppliedOptions: ConfirmOptions): Promise<boolean> {
    options.value = {
      ...defaultOptions,
      ...suppliedOptions,
    };
    isOpen.value = true;
    return new Promise<boolean>((resolve) => {
      resolver = resolve;
    });
  }

  function onClose(action: boolean) {
    isOpen.value = false;
    resolver(action);
  }

  return { isOpen, options, confirm, onClose };
}
