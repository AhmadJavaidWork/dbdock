/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./index.html", "./src/**/*.{vue,js,ts,jsx,tsx}", "./src/components/**/*.{vue,js,ts}"],
  darkMode: "class",
  theme: {
    extend: {
      fontFamily: {
        sans: ["Montserrat", "sans-serif"],
        mono: ["JetBrains Mono", "monospace"],
        roboto: ["Roboto", "sans-serif"],
      },
      colors: {
        primary: "rgb(var(--color-primary) / <alpha-value>)",
        "primary-light": "rgb(var(--color-primary-light) / <alpha-value>)",
        "primary-hover": "rgb(var(--color-primary-hover) / <alpha-value>)",
        "primary-dark": "rgb(var(--color-primary-dark) / <alpha-value>)",
        "primary-dark-hover": "rgb(var(--color-primary-dark-hover) / <alpha-value>)",
        secondary: "rgb(var(--color-secondary) / <alpha-value>)",
        "secondary-dark-hover": "rgb(var(--color-secondary-dark-hover) / <alpha-value>)",
        danger: "rgb(var(--color-danger) / <alpha-value>)",
        "danger-light": "rgb(var(--color-danger-light) / <alpha-value>)",
        "danger-hover": "rgb(var(--color-danger-hover) / <alpha-value>)",
        "danger-darker": "rgb(var(--color-danger-darker) / <alpha-value>)",
        "accent-dark": "rgb(var(--color-accent-dark) / <alpha-value>)",
        surface: {
          light: "rgb(var(--color-surface-light) / <alpha-value>)",
          dark: "rgb(var(--color-surface-dark) / <alpha-value>)",
        },
        background: {
          light: "rgb(var(--color-background-light) / <alpha-value>)",
          dark: "rgb(var(--color-background-dark) / <alpha-value>)",
        },
        text: {
          light: "rgb(var(--color-text-light) / <alpha-value>)",
          dark: "rgb(var(--color-text-dark) / <alpha-value>)",
          "secondary-light": "rgb(var(--color-text-secondary-light) / <alpha-value>)",
          "secondary-dark": "rgb(var(--color-text-secondary-dark) / <alpha-value>)",
          "accent-light": "rgb(var(--color-text-accent-light) / <alpha-value>)",
          "accent-dark": "rgb(var(--color-text-accent-dark) / <alpha-value>)",
          success: "rgb(var(--color-text-success) / <alpha-value>)",
          "success-light": "rgb(var(--color-text-success-light) / <alpha-value>)",
          "success-hover": "rgb(var(--color-text-success-hover) / <alpha-value>)",
          "success-dark": "rgb(var(--color-text-success-dark) / <alpha-value>)",
          "success-darker": "rgb(var(--color-text-success-darker) / <alpha-value>)",
          info: "rgb(var(--color-text-info) / <alpha-value>)",
          "info-light": "rgb(var(--color-text-info-light) / <alpha-value>)",
          "info-hover": "rgb(var(--color-text-info-hover) / <alpha-value>)",
          "info-dark": "rgb(var(--color-text-info-dark) / <alpha-value>)",
          "info-darker": "rgb(var(--color-text-info-darker) / <alpha-value>)",
          danger: "rgb(var(--color-text-danger) / <alpha-value>)",
          "danger-light": "rgb(var(--color-text-danger-light) / <alpha-value>)",
          "danger-hover": "rgb(var(--color-text-danger-hover) / <alpha-value>)",
          "danger-dark": "rgb(var(--color-text-danger-dark) / <alpha-value>)",
          "danger-darker": "rgb(var(--color-text-danger-darker) / <alpha-value>)",
          warning: "rgb(var(--color-text-warning) / <alpha-value>)",
          "warning-light": "rgb(var(--color-text-warning-light) / <alpha-value>)",
          "warning-hover": "rgb(var(--color-text-warning-hover) / <alpha-value>)",
          "warning-dark": "rgb(var(--color-text-warning-dark) / <alpha-value>)",
          "warning-darker": "rgb(var(--color-text-warning-darker) / <alpha-value>)",
        },
        "gray-light": "rgb(var(--color-gray-light) / <alpha-value>)",
        "gray-lighter": "rgb(var(--color-gray-lighter) / <alpha-value>)",
        "gray-placeholder": "rgb(var(--color-gray-placeholder) / <alpha-value>)",
        textfield: {
          empty: {
            light: "rgb(var(--color-textfield-empty-light) / <alpha-value>)",
            dark: "rgb(var(--color-textfield-empty-dark) / <alpha-value>)",
          },
          filled: {
            light: "rgb(var(--color-textfield-filled-light) / <alpha-value>)",
            dark: "rgb(var(--color-textfield-filled-dark) / <alpha-value>)",
          },
          error: {
            light: "rgb(var(--color-textfield-error-light) / <alpha-value>)",
            dark: "rgb(var(--color-textfield-error-dark) / <alpha-value>)",
          },
          border: {
            light: "rgb(var(--color-textfield-border-light) / <alpha-value>)",
            dark: "rgb(var(--color-textfield-border-dark) / <alpha-value>)",
            "error-light": "rgb(var(--color-textfield-border-error-light) / <alpha-value>)",
            "error-dark": "rgb(var(--color-textfield-border-error-dark) / <alpha-value>)",
            "focused-border-light":
              "rgb(var(--color-textfield-border-focused-border-light) / <alpha-value>)",
            "focused-border-dark":
              "rgb(var(--color-textfield-border-focused-border-dark) / <alpha-value>)",
          },
          placeholder: {
            light: "rgb(var(--color-textfield-placeholder-light) / <alpha-value>)",
            dark: "rgb(var(--color-textfield-placeholder-dark) / <alpha-value>)",
            "error-light": "rgb(var(--color-textfield-placeholder-error-light) / <alpha-value>)",
            "error-dark": "rgb(var(--color-textfield-placeholder-error-dark) / <alpha-value>)",
          },
        },
        "select-options": {
          text: {
            light: "rgb(var(--color-select-options-text-light) / <alpha-value>)",
            dark: "rgb(var(--color-select-options-text-dark) / <alpha-value>)",
            "selected-light":
              "rgb(var(--color-select-options-text-selected-light) / <alpha-value>)",
            "selected-dark": "rgb(var(--color-select-options-text-selected-dark) / <alpha-value>)",
          },
          background: {
            light: "rgb(var(--color-select-options-background-light) / <alpha-value>)",
            dark: "rgb(var(--color-select-options-background-dark) / <alpha-value>)",
            "hovered-light":
              "rgb(var(--color-select-options-background-hovered-light) / <alpha-value>)",
            "hovered-dark":
              "rgb(var(--color-select-options-background-hovered-dark) / <alpha-value>)",
            "selected-light":
              "rgb(var(--color-select-options-background-selected-light) / <alpha-value>)",
            "selected-dark":
              "rgb(var(--color-select-options-background-selected-dark) / <alpha-value>)",
            "selected-hover-light":
              "rgb(var(--color-select-options-background-selected-hover-light) / <alpha-value>)",
            "selected-hover-dark":
              "rgb(var(--color-select-options-background-selected-hover-dark) / <alpha-value>)",
          },
        },
      },
      borderRadius: {
        prompt: "10px",
      },
      boxShadow: {
        card: "0 4px 12px rgba(0,0,0,0.1)",
        dropdown: "0 4px 6px -1px rgb(0 0 0 / 0.1), 0 2px 4px -2px rgb(0 0 0 / 0.1)",
      },
    },
  },
  plugins: [],
};
