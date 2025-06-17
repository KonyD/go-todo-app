import {
  createSystem,
  defaultConfig,
  defineConfig,
} from "@chakra-ui/react";

const customConfig = defineConfig({
  theme: {
    tokens: {
      colors: {
        brand: {
          50: { value: "#e3f8ff" },
          100: { value: "#b3ecff" },
          200: { value: "#81defd" },
          300: { value: "#5ed0fa" },
          400: { value: "#40c3f7" },
          500: { value: "#2bb0ed" },
          600: { value: "#1992d4" },
          700: { value: "#127fbf" },
          800: { value: "#0b69a3" },
          900: { value: "#035388" },
        },
      },
    },
    semanticTokens: {
      colors: {
        "bg": {
          value: { _light: "colors.brand.50", _dark: "#292932"}
        },
        "todo-item": {
          value: {_light: "colors.blue.400", _dark: "colors.gray.700"}
        }
      }
    }
  },
});

export const system = createSystem(defaultConfig, customConfig);
