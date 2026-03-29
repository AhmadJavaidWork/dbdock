import type { DirectiveBinding } from "vue";
const handlers: WeakMap<HTMLElement, (event: MouseEvent) => void> = new WeakMap();

const clickOutside = {
  beforeMount(el: HTMLElement, binding: DirectiveBinding) {
    const handleClickOutside = (event: MouseEvent) => {
      if (!(el === event.target || el.contains(event.target as Node))) {
        binding.value(event);
      }
    };

    document.addEventListener("click", handleClickOutside);
    handlers.set(el, handleClickOutside);
  },
  beforeUnmount(el: HTMLElement) {
    const handler = handlers.get(el);
    if (handler) {
      document.removeEventListener("click", handler);
      handlers.delete(el);
    }
  },
};

export default clickOutside;
