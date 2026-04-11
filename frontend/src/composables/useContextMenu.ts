import { ref } from "vue";

type ContextMenuOptions = {
  position: {
    x: number | null;
    y: number | null;
  };
  options: string[];
};

const defaultOptions: Required<ContextMenuOptions> = {
  position: {
    x: null,
    y: null,
  },
  options: [],
};

const show = ref<boolean>(false);
const options = ref<Required<ContextMenuOptions>>(defaultOptions);
let resolver: (value: string | null) => void;

export function useContextMenu() {
  function contextMenu(suppliedOptions: ContextMenuOptions): Promise<string | null> {
    options.value = {
      ...defaultOptions,
      ...suppliedOptions,
    };
    show.value = true;
    return new Promise<string | null>((resolve) => {
      resolver = resolve;
    });
  }

  function onClose(): void {
    show.value = false;
    resolver(null);
  }

  function onSelect(option: string): void {
    show.value = false;
    resolver(option);
  }

  return { show, options, contextMenu, onClose, onSelect };
}
