import { ThemeConfig } from "@/types/theme.type";

export function applyThemeConfig(themeConfig: ThemeConfig) {
  const root = document.documentElement;

  for (const [key, value] of Object.entries(themeConfig)) {
    root.style.setProperty(`--color-${key}`, value);
  }
}
