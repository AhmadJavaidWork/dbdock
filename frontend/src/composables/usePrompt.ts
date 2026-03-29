import { ref } from "vue";

type PromptOptions = {
  type: "success" | "info" | "warn" | "error";
  title: string;
  message: string;
  closeBtnTxt?: string;
};

const defaultOptions: Required<PromptOptions> = {
  type: "success",
  title: "",
  message: "",
  closeBtnTxt: "Close",
};

const isOpen = ref(false);
const options = ref<Required<PromptOptions>>(defaultOptions);
let resolver: (value: boolean) => void;

export function usePrompt() {
  function prompt(suppliedOptions: PromptOptions): Promise<boolean> {
    options.value = {
      ...defaultOptions,
      ...suppliedOptions,
    };
    isOpen.value = true;
    return new Promise<boolean>((resolve) => {
      resolver = resolve;
    });
  }

  function onClose() {
    isOpen.value = false;
    resolver(true);
  }

  return { isOpen, options, prompt, onClose };
}
