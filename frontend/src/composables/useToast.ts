import "@/assets/toast.scss";
import IconAlertCircle from "@/components/icons/IconAlertCircle.vue";
import IconCheckCircle from "@/components/icons/IconCheckCircle.vue";
import IconCloseCircle from "@/components/icons/IconCloseCircle.vue";
import IconInfoCircle from "@/components/icons/IconInfoCircle.vue";
import { ONE_SECOND } from "@/constants/essentials";
import { h } from "vue";
import { toast, type ToastContainerOptions } from "vue3-toastify";
import "vue3-toastify/dist/index.css";

export function useToast(message: string, type: "success" | "info" | "warn" | "error") {
  const options: ToastContainerOptions = {
    autoClose: 2 * ONE_SECOND,
    pauseOnHover: true,
    toastClassName:
      "text-text-light bg-background-light dark:bg-background-dark dark:text-text-dark",
    closeButton: false,
  };
  if (type === "success") {
    toast.success(message, {
      ...options,
      progressClassName: "bg-text-success",
      icon: h(IconCheckCircle, {
        class: "text-text-success dark:text-text-success-darker w-[20px] h-[20px]",
      }),
    });
  } else if (type === "info") {
    toast.info(message, {
      ...options,
      icon: h(IconInfoCircle, {
        class: "text-text-info dark:text-text-info-darker w-[20px] h-[20px]",
      }),
      progressClassName: "asdf",
    });
  } else if (type === "warn") {
    toast.warn(message, {
      ...options,
      icon: h(IconAlertCircle, {
        class: "text-text-warning dark:text-text-warning-darker w-[20px] h-[20px]",
      }),
      progressClassName: "asdf",
    });
  } else {
    toast.error(message, {
      ...options,
      icon: h(IconCloseCircle, {
        class: "text-text-danger dark:text-text-danger-darker w-[20px] h-[20px]",
      }),
      progressClassName: "asdf",
    });
  }
}
