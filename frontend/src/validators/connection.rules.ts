import { helpers, numeric, required } from "@vuelidate/validators";

export const connectionRules = {
  name: { required: helpers.withMessage("Name is required", required) },
  type: { required: helpers.withMessage("Type is required", required) },
  host: { required: helpers.withMessage("Host is required", required) },
  port: {
    required: helpers.withMessage("Port is required", required),
    numeric,
    validatePort: helpers.withMessage("Port must be between 1 and 65535", (value: string) => {
      const num = Number(value);
      return num > 0 && num <= 65535;
    }),
  },
  username: { required: helpers.withMessage("Username is required", required) },
  password: { required: helpers.withMessage("Password is required", required) },
  database: { required: helpers.withMessage("Database name is required", required) },
};
